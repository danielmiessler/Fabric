package github

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/go-github/v66/github"
	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2"
)

type Client struct {
	client        *github.Client
	graphqlClient *graphql.Client
	owner         string
	repo          string
	token         string
}

func NewClient(token, owner, repo string) *Client {
	var githubClient *github.Client
	var httpClient *http.Client
	var gqlClient *graphql.Client

	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		httpClient = oauth2.NewClient(context.Background(), ts)
		githubClient = github.NewClient(httpClient)
		gqlClient = graphql.NewClient("https://api.github.com/graphql", httpClient)
	} else {
		httpClient = http.DefaultClient
		githubClient = github.NewClient(nil)
		gqlClient = graphql.NewClient("https://api.github.com/graphql", httpClient)
	}

	return &Client{
		client:        githubClient,
		graphqlClient: gqlClient,
		owner:         owner,
		repo:          repo,
		token:         token,
	}
}

func (c *Client) FetchPRs(prNumbers []int) ([]*PR, error) {
	if len(prNumbers) == 0 {
		return []*PR{}, nil
	}

	ctx := context.Background()
	prs := make([]*PR, 0, len(prNumbers))
	prsChan := make(chan *PR, len(prNumbers))
	errChan := make(chan error, len(prNumbers))

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10)

	for _, prNumber := range prNumbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()

			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			pr, err := c.fetchSinglePR(ctx, num)
			if err != nil {
				errChan <- fmt.Errorf("failed to fetch PR #%d: %w", num, err)
				return
			}
			prsChan <- pr
		}(prNumber)
	}

	go func() {
		wg.Wait()
		close(prsChan)
		close(errChan)
	}()

	var errors []error
	for pr := range prsChan {
		prs = append(prs, pr)
	}
	for err := range errChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return prs, fmt.Errorf("some PRs failed to fetch: %v", errors)
	}

	return prs, nil
}

// GetPRValidationDetails fetches only the data needed for validation (lightweight).
func (c *Client) GetPRValidationDetails(prNumber int) (*PRDetails, error) {
	ctx := context.Background()
	ghPR, _, err := c.client.PullRequests.Get(ctx, c.owner, c.repo, prNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get PR %d: %w", prNumber, err)
	}

	details := &PRDetails{
		PR:        nil,
		State:     getString(ghPR.State),
		Mergeable: ghPR.Mergeable != nil && *ghPR.Mergeable,
	}

	return details, nil
}

// GetPRWithCommits fetches the full PR and its commits.
func (c *Client) GetPRWithCommits(prNumber int) (*PR, error) {
	ctx := context.Background()
	ghPR, _, err := c.client.PullRequests.Get(ctx, c.owner, c.repo, prNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get PR %d: %w", prNumber, err)
	}

	return c.buildPRWithCommits(ctx, ghPR)
}

// GetPRDetails fetches a comprehensive set of details for a single PR.
// Deprecated: Use GetPRValidationDetails + GetPRWithCommits for better performance
func (c *Client) GetPRDetails(prNumber int) (*PRDetails, error) {
	ctx := context.Background()
	ghPR, _, err := c.client.PullRequests.Get(ctx, c.owner, c.repo, prNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get PR %d: %w", prNumber, err)
	}

	pr, err := c.buildPRWithCommits(ctx, ghPR)
	if err != nil {
		return nil, fmt.Errorf("failed to build PR details for %d: %w", prNumber, err)
	}

	details := &PRDetails{
		PR:        pr,
		State:     getString(ghPR.State),
		Mergeable: ghPR.Mergeable != nil && *ghPR.Mergeable,
	}

	return details, nil
}

func (c *Client) buildPRWithCommits(ctx context.Context, ghPR *github.PullRequest) (*PR, error) {
	commits, _, err := c.client.PullRequests.ListCommits(ctx, c.owner, c.repo, *ghPR.Number, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch commits for PR %d: %w", *ghPR.Number, err)
	}

	return c.convertGitHubPR(ghPR, commits), nil
}

func (c *Client) convertGitHubPR(ghPR *github.PullRequest, commits []*github.RepositoryCommit) *PR {

	result := &PR{
		Number:  *ghPR.Number,
		Title:   getString(ghPR.Title),
		Body:    getString(ghPR.Body),
		URL:     getString(ghPR.HTMLURL),
		Commits: make([]PRCommit, 0, len(commits)),
	}

	if ghPR.MergedAt != nil {
		result.MergedAt = ghPR.MergedAt.Time
	}

	if ghPR.User != nil {
		result.Author = getString(ghPR.User.Login)
		result.AuthorURL = getString(ghPR.User.HTMLURL)
		userType := getString(ghPR.User.Type)

		switch userType {
		case "User":
			result.AuthorType = "user"
		case "Organization":
			result.AuthorType = "organization"
		case "Bot":
			result.AuthorType = "bot"
		default:
			result.AuthorType = "user"
		}
	}

	if ghPR.MergeCommitSHA != nil {
		result.MergeCommit = *ghPR.MergeCommitSHA
	}

	for _, commit := range commits {
		if commit.Commit != nil {
			prCommit := PRCommit{
				SHA:     getString(commit.SHA),
				Message: strings.TrimSpace(getString(commit.Commit.Message)),
			}
			if commit.Commit.Author != nil {
				prCommit.Author = getString(commit.Commit.Author.Name)
				prCommit.Email = getString(commit.Commit.Author.Email)
				if commit.Commit.Author.Date != nil {
					prCommit.Date = commit.Commit.Author.Date.Time
				}
			}
			// The changelog package uses the parent list to detect merge commits.
			if commit.Parents != nil {
				for _, parent := range commit.Parents {
					if parent.SHA != nil {
						prCommit.Parents = append(prCommit.Parents, *parent.SHA)
					}
				}
			}
			result.Commits = append(result.Commits, prCommit)
		}
	}

	return result
}

func (c *Client) fetchSinglePR(ctx context.Context, prNumber int) (*PR, error) {
	ghPR, _, err := c.client.PullRequests.Get(ctx, c.owner, c.repo, prNumber)
	if err != nil {
		return nil, err
	}

	return c.buildPRWithCommits(ctx, ghPR)
}

func getString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// FetchAllMergedPRs fetches all merged PRs using GitHub's search API
// This is much more efficient than fetching PRs individually
func (c *Client) FetchAllMergedPRs(since time.Time) ([]*PR, error) {
	ctx := context.Background()
	var allPRs []*PR

	query := fmt.Sprintf("repo:%s/%s is:pr is:merged", c.owner, c.repo)
	if !since.IsZero() {
		query += fmt.Sprintf(" merged:>=%s", since.Format("2006-01-02"))
	}

	opts := &github.SearchOptions{
		Sort:  "created",
		Order: "desc",
		ListOptions: github.ListOptions{
			PerPage: 100, // GitHub maximum
		},
	}

	for {
		result, resp, err := c.client.Search.Issues(ctx, query, opts)
		if err != nil {
			return allPRs, fmt.Errorf("failed to search PRs: %w", err)
		}

		prsChan := make(chan *PR, len(result.Issues))
		errChan := make(chan error, len(result.Issues))
		var wg sync.WaitGroup
		semaphore := make(chan struct{}, 10)

		for _, issue := range result.Issues {
			if issue.PullRequestLinks == nil {
				continue // Not a PR
			}

			wg.Add(1)
			go func(prNumber int) {
				defer wg.Done()

				semaphore <- struct{}{}
				defer func() { <-semaphore }()

				pr, err := c.fetchSinglePR(ctx, prNumber)
				if err != nil {
					errChan <- fmt.Errorf("failed to fetch PR #%d: %w", prNumber, err)
					return
				}
				prsChan <- pr
			}(*issue.Number)
		}

		go func() {
			wg.Wait()
			close(prsChan)
			close(errChan)
		}()

		for pr := range prsChan {
			allPRs = append(allPRs, pr)
		}

		for err := range errChan {
			fmt.Fprintf(os.Stderr, "Warning: %v\n", err)
		}

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return allPRs, nil
}

// FetchAllMergedPRsGraphQL fetches all merged PRs with their commits using GraphQL
// This is the ultimate optimization - gets everything in ~5-10 API calls
func (c *Client) FetchAllMergedPRsGraphQL(since time.Time) ([]*PR, error) {
	ctx := context.Background()
	var allPRs []*PR
	var after *string
	totalFetched := 0

	for {
		variables := map[string]any{
			"owner": graphql.String(c.owner),
			"repo":  graphql.String(c.repo),
			"after": (*graphql.String)(after),
		}

		var query PullRequestsQuery
		err := c.graphqlClient.Query(ctx, &query, variables)
		if err != nil {
			return allPRs, fmt.Errorf("GraphQL query failed: %w", err)
		}

		prs := query.Repository.PullRequests.Nodes
		fmt.Fprintf(os.Stderr, "Fetched %d PRs via GraphQL (page %d)\n", len(prs), (totalFetched/100)+1)

		for _, gqlPR := range prs {
			if !since.IsZero() && gqlPR.MergedAt.Before(since) {
				fmt.Fprintf(os.Stderr, "Reached PRs older than %s, stopping\n", since.Format("2006-01-02"))
				return allPRs, nil
			}

			pr := &PR{
				Number:   gqlPR.Number,
				Title:    gqlPR.Title,
				Body:     gqlPR.Body,
				URL:      gqlPR.URL,
				MergedAt: gqlPR.MergedAt,
				Commits:  make([]PRCommit, 0, len(gqlPR.Commits.Nodes)),
			}

			if gqlPR.Author != nil {
				pr.Author = gqlPR.Author.Login
				pr.AuthorURL = gqlPR.Author.URL

				switch gqlPR.Author.Typename {
				case "Bot":
					pr.AuthorType = "bot"
				case "Organization":
					pr.AuthorType = "organization"
				case "User":
					pr.AuthorType = "user"
				default:
					pr.AuthorType = "user"
					if gqlPR.Author.Typename != "" {
						fmt.Fprintf(os.Stderr, "PR #%d: Unknown author typename '%s'\n", gqlPR.Number, gqlPR.Author.Typename)
					}
				}
			} else {
				// GraphQL returns a nil author for a deleted account. REST still returns a login.
				fmt.Fprintf(os.Stderr, "PR #%d: Author is nil in GraphQL response, fetching from REST API\n", gqlPR.Number)

				restPR, err := c.fetchSinglePR(ctx, gqlPR.Number)
				if err == nil && restPR != nil && restPR.Author != "" {
					pr.Author = restPR.Author
					pr.AuthorURL = restPR.AuthorURL
					pr.AuthorType = restPR.AuthorType
				} else {
					pr.Author = "[unknown]"
					pr.AuthorURL = ""
					pr.AuthorType = "user"
				}
			}

			for _, commitNode := range gqlPR.Commits.Nodes {
				commit := PRCommit{
					SHA:     commitNode.Commit.OID,
					Message: strings.TrimSpace(commitNode.Commit.Message),
					Author:  commitNode.Commit.Author.Name,
					Date:    commitNode.Commit.AuthoredDate,
				}
				pr.Commits = append(pr.Commits, commit)
			}

			allPRs = append(allPRs, pr)
		}

		totalFetched += len(prs)

		if !query.Repository.PullRequests.PageInfo.HasNextPage {
			break
		}

		after = &query.Repository.PullRequests.PageInfo.EndCursor
	}

	fmt.Fprintf(os.Stderr, "Total PRs fetched via GraphQL: %d\n", len(allPRs))
	return allPRs, nil
}
