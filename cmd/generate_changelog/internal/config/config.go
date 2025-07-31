package config

type Config struct {
	RepoPath          string
	OutputFile        string
	Limit             int
	Version           string
	SaveData          bool
	CacheFile         string
	NoCache           bool
	RebuildCache      bool
	GitHubToken       string
	ForcePRSync       bool
	EnableAISummary   bool
	IncomingPR        int
	ProcessPRsVersion string
	IncomingDir       string
	Push              bool
	SyncDB            bool
	Release           string
}
