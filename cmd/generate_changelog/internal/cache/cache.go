package cache

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/danielmiessler/fabric/cmd/generate_changelog/internal/git"
	"github.com/danielmiessler/fabric/cmd/generate_changelog/internal/github"
	_ "github.com/mattn/go-sqlite3"
)

type Cache struct {
	db *sql.DB
}

func New(dbPath string) (*Cache, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	cache := &Cache{db: db}
	if err := cache.createTables(); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return cache, nil
}

func (c *Cache) Close() error {
	return c.db.Close()
}

func (c *Cache) createTables() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS metadata (
			key TEXT PRIMARY KEY,
			value TEXT NOT NULL,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS versions (
			name TEXT PRIMARY KEY,
			date DATETIME,
			commit_sha TEXT,
			pr_numbers TEXT,
			ai_summary TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS commits (
			sha TEXT PRIMARY KEY,
			version TEXT NOT NULL,
			message TEXT,
			author TEXT,
			email TEXT,
			date DATETIME,
			is_merge BOOLEAN,
			pr_number INTEGER,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (version) REFERENCES versions(name)
		)`,
		`CREATE TABLE IF NOT EXISTS pull_requests (
			number INTEGER PRIMARY KEY,
			title TEXT,
			body TEXT,
			author TEXT,
			author_url TEXT,
			author_type TEXT DEFAULT 'user',
			url TEXT,
			merged_at DATETIME,
			merge_commit TEXT,
			commits TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE INDEX IF NOT EXISTS idx_commits_version ON commits(version)`,
		`CREATE INDEX IF NOT EXISTS idx_commits_pr_number ON commits(pr_number)`,
		`CREATE TABLE IF NOT EXISTS commit_pr_mapping (
			commit_sha TEXT PRIMARY KEY,
			pr_number INTEGER NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (pr_number) REFERENCES pull_requests(number)
		)`,
		`CREATE INDEX IF NOT EXISTS idx_commit_pr_mapping_sha ON commit_pr_mapping(commit_sha)`,
	}

	for _, query := range queries {
		if _, err := c.db.Exec(query); err != nil {
			return fmt.Errorf("failed to execute query: %w", err)
		}
	}

	return nil
}

func (c *Cache) GetLastProcessedTag() (string, error) {
	var tag string
	err := c.db.QueryRow("SELECT value FROM metadata WHERE key = 'last_processed_tag'").Scan(&tag)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return tag, err
}

func (c *Cache) SetLastProcessedTag(tag string) error {
	_, err := c.db.Exec(`
		INSERT OR REPLACE INTO metadata (key, value, updated_at)
		VALUES ('last_processed_tag', ?, CURRENT_TIMESTAMP)
	`, tag)
	return err
}

func (c *Cache) SaveVersion(v *git.Version) error {
	prNumbers, _ := json.Marshal(v.PRNumbers)

	_, err := c.db.Exec(`
		INSERT OR REPLACE INTO versions (name, date, commit_sha, pr_numbers, ai_summary)
		VALUES (?, ?, ?, ?, ?)
	`, v.Name, v.Date, v.CommitSHA, string(prNumbers), v.AISummary)

	return err
}

// UpdateVersionAISummary updates only the AI summary for a specific version
func (c *Cache) UpdateVersionAISummary(versionName, aiSummary string) error {
	_, err := c.db.Exec(`
		UPDATE versions SET ai_summary = ? WHERE name = ?
	`, aiSummary, versionName)
	return err
}

func (c *Cache) SaveCommit(commit *git.Commit, version string) error {
	_, err := c.db.Exec(`
		INSERT OR REPLACE INTO commits
		(sha, version, message, author, email, date, is_merge, pr_number)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, commit.SHA, version, commit.Message, commit.Author, commit.Email,
		commit.Date, commit.IsMerge, commit.PRNumber)

	return err
}

func (c *Cache) SavePR(pr *github.PR) error {
	commits, _ := json.Marshal(pr.Commits)

	_, err := c.db.Exec(`
		INSERT OR REPLACE INTO pull_requests
		(number, title, body, author, author_url, author_type, url, merged_at, merge_commit, commits)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, pr.Number, pr.Title, pr.Body, pr.Author, pr.AuthorURL, pr.AuthorType,
		pr.URL, pr.MergedAt, pr.MergeCommit, string(commits))

	return err
}

func (c *Cache) GetPR(number int) (*github.PR, error) {
	var pr github.PR
	var commitsJSON string

	err := c.db.QueryRow(`
		SELECT number, title, body, author, author_url, COALESCE(author_type, 'user'), url, merged_at, merge_commit, commits
		FROM pull_requests WHERE number = ?
	`, number).Scan(
		&pr.Number, &pr.Title, &pr.Body, &pr.Author, &pr.AuthorURL, &pr.AuthorType,
		&pr.URL, &pr.MergedAt, &pr.MergeCommit, &commitsJSON,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(commitsJSON), &pr.Commits); err != nil {
		return nil, fmt.Errorf("failed to unmarshal commits: %w", err)
	}

	return &pr, nil
}

func (c *Cache) GetVersions() (map[string]*git.Version, error) {
	rows, err := c.db.Query(`
		SELECT name, date, commit_sha, pr_numbers, ai_summary FROM versions
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	versions := make(map[string]*git.Version)

	for rows.Next() {
		var v git.Version
		var dateStr sql.NullString
		var prNumbersJSON string
		var aiSummary sql.NullString

		if err := rows.Scan(&v.Name, &dateStr, &v.CommitSHA, &prNumbersJSON, &aiSummary); err != nil {
			return nil, err
		}

		if dateStr.Valid {
			// Try RFC3339Nano first (for nanosecond precision), then fall back to RFC3339
			v.Date, err = time.Parse(time.RFC3339Nano, dateStr.String)
			if err != nil {
				v.Date, err = time.Parse(time.RFC3339, dateStr.String)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing date '%s' for version '%s': %v. Expected format: RFC3339 or RFC3339Nano.\n", dateStr.String, v.Name, err)
				}
			}
		}

		if prNumbersJSON != "" {
			json.Unmarshal([]byte(prNumbersJSON), &v.PRNumbers)
		}

		if aiSummary.Valid {
			v.AISummary = aiSummary.String
		}

		v.Commits, err = c.getCommitsForVersion(v.Name)
		if err != nil {
			return nil, err
		}

		versions[v.Name] = &v
	}

	return versions, rows.Err()
}

func (c *Cache) getCommitsForVersion(version string) ([]*git.Commit, error) {
	rows, err := c.db.Query(`
		SELECT sha, message, author, email, date, is_merge, pr_number
		FROM commits WHERE version = ?
		ORDER BY date DESC
	`, version)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commits []*git.Commit

	for rows.Next() {
		var commit git.Commit
		if err := rows.Scan(
			&commit.SHA, &commit.Message, &commit.Author, &commit.Email,
			&commit.Date, &commit.IsMerge, &commit.PRNumber,
		); err != nil {
			return nil, err
		}
		commits = append(commits, &commit)
	}

	return commits, rows.Err()
}

func (c *Cache) Clear() error {
	tables := []string{"metadata", "versions", "commits", "pull_requests"}
	for _, table := range tables {
		if _, err := c.db.Exec("DELETE FROM " + table); err != nil {
			return err
		}
	}
	return nil
}

// VersionExists checks if a version already exists in the cache
func (c *Cache) VersionExists(version string) (bool, error) {
	var count int
	err := c.db.QueryRow("SELECT COUNT(*) FROM versions WHERE name = ?", version).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CommitExists checks if a commit already exists in the cache
func (c *Cache) CommitExists(hash string) (bool, error) {
	var count int
	err := c.db.QueryRow("SELECT COUNT(*) FROM commits WHERE sha = ?", hash).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetLastPRSync returns the timestamp of the last PR sync
func (c *Cache) GetLastPRSync() (time.Time, error) {
	var timestamp string
	err := c.db.QueryRow("SELECT value FROM metadata WHERE key = 'last_pr_sync'").Scan(&timestamp)
	if err == sql.ErrNoRows {
		return time.Time{}, nil
	}
	if err != nil {
		return time.Time{}, err
	}

	return time.Parse(time.RFC3339, timestamp)
}

// SetLastPRSync updates the timestamp of the last PR sync
func (c *Cache) SetLastPRSync(timestamp time.Time) error {
	_, err := c.db.Exec(`
		INSERT OR REPLACE INTO metadata (key, value, updated_at)
		VALUES ('last_pr_sync', ?, CURRENT_TIMESTAMP)
	`, timestamp.Format(time.RFC3339))
	return err
}

// SavePRBatch saves multiple PRs in a single transaction for better performance
func (c *Cache) SavePRBatch(prs []*github.PR) error {
	tx, err := c.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT OR REPLACE INTO pull_requests
		(number, title, body, author, author_url, author_type, url, merged_at, merge_commit, commits)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, pr := range prs {
		commits, _ := json.Marshal(pr.Commits)
		_, err := stmt.Exec(
			pr.Number, pr.Title, pr.Body, pr.Author, pr.AuthorURL, pr.AuthorType,
			pr.URL, pr.MergedAt, pr.MergeCommit, string(commits),
		)
		if err != nil {
			return fmt.Errorf("failed to save PR #%d: %w", pr.Number, err)
		}
	}

	return tx.Commit()
}

// GetAllPRs returns all cached PRs
func (c *Cache) GetAllPRs() (map[int]*github.PR, error) {
	rows, err := c.db.Query(`
		SELECT number, title, body, author, author_url, COALESCE(author_type, 'user'), url, merged_at, merge_commit, commits
		FROM pull_requests
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prs := make(map[int]*github.PR)

	for rows.Next() {
		var pr github.PR
		var commitsJSON string

		if err := rows.Scan(
			&pr.Number, &pr.Title, &pr.Body, &pr.Author, &pr.AuthorURL, &pr.AuthorType,
			&pr.URL, &pr.MergedAt, &pr.MergeCommit, &commitsJSON,
		); err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(commitsJSON), &pr.Commits); err != nil {
			return nil, fmt.Errorf("failed to unmarshal commits for PR #%d: %w", pr.Number, err)
		}

		prs[pr.Number] = &pr
	}

	return prs, rows.Err()
}

// MarkPRAsNonExistent marks a PR number as non-existent to avoid future fetches
func (c *Cache) MarkPRAsNonExistent(prNumber int) error {
	_, err := c.db.Exec(`
		INSERT OR REPLACE INTO metadata (key, value, updated_at)
		VALUES (?, 'non_existent', CURRENT_TIMESTAMP)
	`, fmt.Sprintf("pr_non_existent_%d", prNumber))
	return err
}

// IsPRMarkedAsNonExistent checks if a PR is marked as non-existent
func (c *Cache) IsPRMarkedAsNonExistent(prNumber int) bool {
	var value string
	err := c.db.QueryRow("SELECT value FROM metadata WHERE key = ?",
		fmt.Sprintf("pr_non_existent_%d", prNumber)).Scan(&value)
	return err == nil && value == "non_existent"
}

// SaveCommitPRMappings saves SHA→PR mappings for all commits in PRs
func (c *Cache) SaveCommitPRMappings(prs []*github.PR) error {
	tx, err := c.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT OR REPLACE INTO commit_pr_mapping (commit_sha, pr_number)
		VALUES (?, ?)
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, pr := range prs {
		for _, commit := range pr.Commits {
			_, err := stmt.Exec(commit.SHA, pr.Number)
			if err != nil {
				return fmt.Errorf("failed to save commit mapping %s→%d: %w", commit.SHA, pr.Number, err)
			}
		}
	}

	return tx.Commit()
}

// GetPRNumberBySHA returns the PR number for a given commit SHA
func (c *Cache) GetPRNumberBySHA(sha string) (int, bool) {
	var prNumber int
	err := c.db.QueryRow("SELECT pr_number FROM commit_pr_mapping WHERE commit_sha = ?", sha).Scan(&prNumber)
	if err == sql.ErrNoRows {
		return 0, false
	}
	if err != nil {
		return 0, false
	}
	return prNumber, true
}

// GetCommitSHAsForPR returns all commit SHAs for a given PR number
func (c *Cache) GetCommitSHAsForPR(prNumber int) ([]string, error) {
	rows, err := c.db.Query("SELECT commit_sha FROM commit_pr_mapping WHERE pr_number = ?", prNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shas []string
	for rows.Next() {
		var sha string
		if err := rows.Scan(&sha); err != nil {
			return nil, err
		}
		shas = append(shas, sha)
	}

	return shas, rows.Err()
}

// GetUnreleasedContentHash returns the cached content hash for Unreleased
func (c *Cache) GetUnreleasedContentHash() (string, error) {
	var hash string
	err := c.db.QueryRow("SELECT value FROM metadata WHERE key = 'unreleased_content_hash'").Scan(&hash)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("no content hash found")
	}
	return hash, err
}

// SetUnreleasedContentHash stores the content hash for Unreleased
func (c *Cache) SetUnreleasedContentHash(hash string) error {
	_, err := c.db.Exec(`
		INSERT OR REPLACE INTO metadata (key, value, updated_at)
		VALUES ('unreleased_content_hash', ?, CURRENT_TIMESTAMP)
	`, hash)
	return err
}
