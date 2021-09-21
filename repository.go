package golibraries

import "fmt"

// Repository() response struct.
type RepositoryResponse struct {
	FullName                 string      `json:"full_name"`
	Description              string      `json:"description"`
	Fork                     bool        `json:"fork"`
	CreatedAt                string      `json:"created_at"`
	UpdatedAt                string      `json:"updated_at"`
	PushedAt                 string      `json:"pushed_at"`
	Homepage                 string      `json:"homepage"`
	Size                     int64       `json:"size"`
	StargazersCount          int64       `json:"stargazers_count"`
	Language                 string      `json:"language"`
	HasIssues                bool        `json:"has_issues"`
	HasWiki                  bool        `json:"has_wiki"`
	HasPages                 bool        `json:"has_pages"`
	ForksCount               int64       `json:"forks_count"`
	MirrorURL                interface{} `json:"mirror_url"`
	OpenIssuesCount          int64       `json:"open_issues_count"`
	DefaultBranch            string      `json:"default_branch"`
	SubscribersCount         int64       `json:"subscribers_count"`
	UUID                     string      `json:"uuid"`
	SourceName               interface{} `json:"source_name"`
	License                  string      `json:"license"`
	Private                  bool        `json:"private"`
	ContributionsCount       int64       `json:"contributions_count"`
	HasReadme                string      `json:"has_readme"`
	HasChangelog             string      `json:"has_changelog"`
	HasContributing          string      `json:"has_contributing"`
	HasLicense               string      `json:"has_license"`
	HasCoc                   string      `json:"has_coc"`
	HasThreatModel           interface{} `json:"has_threat_model"`
	HasAudit                 interface{} `json:"has_audit"`
	Status                   interface{} `json:"status"`
	LastSyncedAt             string      `json:"last_synced_at"`
	Rank                     int64       `json:"rank"`
	HostType                 string      `json:"host_type"`
	HostDomain               interface{} `json:"host_domain"`
	Name                     interface{} `json:"name"`
	SCM                      string      `json:"scm"`
	ForkPolicy               interface{} `json:"fork_policy"`
	PullRequestsEnabled      interface{} `json:"pull_requests_enabled"`
	LogoURL                  interface{} `json:"logo_url"`
	Keywords                 []string    `json:"keywords"`
	GithubContributionsCount int64       `json:"github_contributions_count"`
	GithubID                 string      `json:"github_id"`
}

// RepositoryDependencies() response struct.
type RepositoryDependenciesResponse struct {
	FullName                 string       `json:"full_name"`
	Description              string       `json:"description"`
	Fork                     bool         `json:"fork"`
	CreatedAt                string       `json:"created_at"`
	UpdatedAt                string       `json:"updated_at"`
	PushedAt                 string       `json:"pushed_at"`
	Homepage                 string       `json:"homepage"`
	Size                     int64        `json:"size"`
	StargazersCount          int64        `json:"stargazers_count"`
	Language                 string       `json:"language"`
	HasIssues                bool         `json:"has_issues"`
	HasWiki                  bool         `json:"has_wiki"`
	HasPages                 bool         `json:"has_pages"`
	ForksCount               int64        `json:"forks_count"`
	MirrorURL                interface{}  `json:"mirror_url"`
	OpenIssuesCount          int64        `json:"open_issues_count"`
	DefaultBranch            string       `json:"default_branch"`
	SubscribersCount         int64        `json:"subscribers_count"`
	UUID                     string       `json:"uuid"`
	SourceName               interface{}  `json:"source_name"`
	License                  string       `json:"license"`
	Private                  bool         `json:"private"`
	ContributionsCount       int64        `json:"contributions_count"`
	HasReadme                string       `json:"has_readme"`
	HasChangelog             string       `json:"has_changelog"`
	HasContributing          string       `json:"has_contributing"`
	HasLicense               string       `json:"has_license"`
	HasCoc                   string       `json:"has_coc"`
	HasThreatModel           interface{}  `json:"has_threat_model"`
	HasAudit                 interface{}  `json:"has_audit"`
	Status                   interface{}  `json:"status"`
	LastSyncedAt             string       `json:"last_synced_at"`
	Rank                     int64        `json:"rank"`
	HostType                 string       `json:"host_type"`
	HostDomain               interface{}  `json:"host_domain"`
	Name                     interface{}  `json:"name"`
	SCM                      string       `json:"scm"`
	ForkPolicy               interface{}  `json:"fork_policy"`
	GithubID                 string       `json:"github_id"`
	PullRequestsEnabled      interface{}  `json:"pull_requests_enabled"`
	LogoURL                  interface{}  `json:"logo_url"`
	GithubContributionsCount int64        `json:"github_contributions_count"`
	Keywords                 []string     `json:"keywords"`
	Dependencies             []Dependency `json:"dependencies"`
}

type Dependency struct {
	ProjectName        string              `json:"project_name"`
	Name               string              `json:"name"`
	Platform           Platform            `json:"platform"`
	Requirements       string              `json:"requirements"`
	LatestStable       *string             `json:"latest_stable"`
	Latest             *string             `json:"latest"`
	Deprecated         *bool               `json:"deprecated"`
	Outdated           *bool               `json:"outdated"`
	Filepath           Filepath            `json:"filepath"`
	Kind               Kind                `json:"kind"`
	NormalizedLicenses []NormalizedLicense `json:"normalized_licenses"`
}

type Filepath string
type Kind string

type NormalizedLicense string

// Generated by https://quicktype.io

type RepositoryProjectsResponse []RepositoryProjectsResponseElement

type RepositoryProjectsResponseElement struct {
	DependentReposCount            int64       `json:"dependent_repos_count"`
	DependentsCount                int64       `json:"dependents_count"`
	DeprecationReason              interface{} `json:"deprecation_reason"`
	Description                    string      `json:"description"`
	Forks                          int64       `json:"forks"`
	Homepage                       string      `json:"homepage"`
	Keywords                       []string    `json:"keywords"`
	Language                       string      `json:"language"`
	LatestDownloadURL              string      `json:"latest_download_url"`
	LatestReleaseNumber            string      `json:"latest_release_number"`
	LatestReleasePublishedAt       string      `json:"latest_release_published_at"`
	LatestStableReleaseNumber      string      `json:"latest_stable_release_number"`
	LatestStableReleasePublishedAt string      `json:"latest_stable_release_published_at"`
	LicenseNormalized              bool        `json:"license_normalized"`
	Licenses                       Licenses    `json:"licenses"`
	Name                           string      `json:"name"`
	NormalizedLicenses             []string    `json:"normalized_licenses"`
	PackageManagerURL              string      `json:"package_manager_url"`
	Platform                       string      `json:"platform"`
	Rank                           int64       `json:"rank"`
	RepositoryLicense              string      `json:"repository_license"`
	RepositoryURL                  string      `json:"repository_url"`
	Stars                          int64       `json:"stars"`
	Status                         interface{} `json:"status"`
	Versions                       []Version   `json:"versions"`
}

// Get info for a repository. Currently only works for open source repositories.
func (l *LibrariesIO) Repository(owner, name string) (RepositoryResponse, error) {
	var r = RepositoryResponse{}

	err := l.client.Get(API_URL+fmt.Sprintf("github/%s/%s?api_key=%s", owner, name, l.ApiKey), &r)

	return r, err
}

// Get a list of dependencies for a repositories. Currently only works for open source repositories.
func (l *LibrariesIO) RepositoryDependencies(owner, name string) (RepositoryDependenciesResponse, error) {
	var r = RepositoryDependenciesResponse{}

	err := l.client.Get(API_URL+fmt.Sprintf("github/%s/%s/dependencies?api_key=%s", owner, name, l.ApiKey), &r)

	return r, err
}

// Get a list of packages referencing the given repository.
func (l *LibrariesIO) RepositoryProjects(owner, name string, p PaginateOptions) (RepositoryProjectsResponse, error) {
	var r = RepositoryProjectsResponse{}

	q := QueryOptions{
		ApiKey: l.ApiKey, Page: p.Page, PerPage: p.PerPage,
	}

	err := l.client.Get(API_URL+fmt.Sprintf("github/%s/%s/projects?%s", owner, name, parseQuery(q)), &r)

	return r, err
}
