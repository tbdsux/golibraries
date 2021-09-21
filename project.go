package golibraries

import "fmt"

// Project() response struct.
type ProjectResponse struct {
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
	Licenses                       string      `json:"licenses"`
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

type Version struct {
	Number            string      `json:"number"`
	PublishedAt       string      `json:"published_at"`
	SpdxExpression    string      `json:"spdx_expression"`
	OriginalLicense   string      `json:"original_license"`
	ResearchedAt      interface{} `json:"researched_at"`
	RepositorySources []string    `json:"repository_sources"`
}

// ProjectDependencies() response struct.
type ProjectDependenciesResponse struct {
	ID                             int64       `json:"id"`
	Name                           string      `json:"name"`
	Platform                       string      `json:"platform"`
	CreatedAt                      string      `json:"created_at"`
	UpdatedAt                      string      `json:"updated_at"`
	Description                    string      `json:"description"`
	Keywords                       []string    `json:"keywords"`
	Homepage                       string      `json:"homepage"`
	Licenses                       string      `json:"licenses"`
	RepositoryURL                  string      `json:"repository_url"`
	RepositoryID                   int64       `json:"repository_id"`
	NormalizedLicenses             []string    `json:"normalized_licenses"`
	VersionsCount                  int64       `json:"versions_count"`
	Rank                           int64       `json:"rank"`
	LatestReleasePublishedAt       string      `json:"latest_release_published_at"`
	LatestReleaseNumber            string      `json:"latest_release_number"`
	PmID                           interface{} `json:"pm_id"`
	KeywordsArray                  []string    `json:"keywords_array"`
	DependentsCount                int64       `json:"dependents_count"`
	Language                       string      `json:"language"`
	Status                         interface{} `json:"status"`
	LastSyncedAt                   string      `json:"last_synced_at"`
	DependentReposCount            int64       `json:"dependent_repos_count"`
	RuntimeDependenciesCount       int64       `json:"runtime_dependencies_count"`
	Score                          int64       `json:"score"`
	ScoreLastCalculated            string      `json:"score_last_calculated"`
	LatestStableReleaseNumber      string      `json:"latest_stable_release_number"`
	LatestStableReleasePublishedAt string      `json:"latest_stable_release_published_at"`
	LicenseSetByAdmin              bool        `json:"license_set_by_admin"`
	LicenseNormalized              bool        `json:"license_normalized"`
	DeprecationReason              interface{} `json:"deprecation_reason"`
	Dependencies                   []string    `json:"dependencies"`
}

// ProjectDependents() response struct.
type ProjectDependentsResponse []ProjectResponse

// ProjectDependentRepositories() response struct.
type ProjectDependentRepositoriesResponse []ProjectDependentRepositoriesResponseElement

type ProjectDependentRepositoriesResponseElement struct {
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
	Status                   string      `json:"status"`
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

// ProjectContributors() response struct.
type ProjectContributorsResponse []ProjectContributorsResponseElement

type ProjectContributorsResponseElement struct {
	GithubID     string  `json:"github_id"`
	Login        string  `json:"login"`
	UserType     string  `json:"user_type"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	Name         *string `json:"name"`
	Company      *string `json:"company"`
	Blog         *string `json:"blog"`
	Location     *string `json:"location"`
	Hidden       bool    `json:"hidden"`
	LastSyncedAt string  `json:"last_synced_at"`
	Email        *string `json:"email"`
	Bio          *string `json:"bio"`
	UUID         string  `json:"uuid"`
	HostType     string  `json:"host_type"`
}

// ProjectSourceRank() response struct.
type ProjectSourceRankResponse struct {
	BasicInfoPresent        int64 `json:"basic_info_present"`
	RepositoryPresent       int64 `json:"repository_present"`
	ReadmePresent           int64 `json:"readme_present"`
	LicensePresent          int64 `json:"license_present"`
	VersionsPresent         int64 `json:"versions_present"`
	FollowsSemver           int64 `json:"follows_semver"`
	RecentRelease           int64 `json:"recent_release"`
	NotBrandNew             int64 `json:"not_brand_new"`
	OnePointOh              int64 `json:"one_point_oh"`
	DependentProjects       int64 `json:"dependent_projects"`
	DependentRepositories   int64 `json:"dependent_repositories"`
	Stars                   int64 `json:"stars"`
	Contributors            int64 `json:"contributors"`
	Subscribers             int64 `json:"subscribers"`
	AllPrereleases          int64 `json:"all_prereleases"`
	AnyOutdatedDependencies int64 `json:"any_outdated_dependencies"`
	IsDeprecated            int64 `json:"is_deprecated"`
	IsUnmaintained          int64 `json:"is_unmaintained"`
	IsRemoved               int64 `json:"is_removed"`
}

// ProjectSearch() response struct.
type ProjectSearchResponse []ProjectResponse

// Get information about a package and its versions.
func (l *LibrariesIO) Project(platform, name string) (ProjectResponse, error) {
	var r = ProjectResponse{}

	err := l.client.Get(API_URL+fmt.Sprintf("%s/%s?api_key=%s", platform, name, l.ApiKey), &r)

	return r, err
}

// Get a list of dependecies for a version of a project, pass `latest` to get dependecy info for the latest avaialable version.
func (l *LibrariesIO) ProjectDependencies(platform, name, version string) (ProjectDependenciesResponse, error) {
	var r = ProjectDependenciesResponse{}

	err := l.client.Get(API_URL+fmt.Sprintf("%s/%s/%s/dependencies?api_key=%s", platform, name, version, l.ApiKey), &r)

	return r, err
}

// Get packages that have at least one version that depends on a given project.
func (l *LibrariesIO) ProjectDependents(platform, name string, p PaginateOptions) (ProjectDependentsResponse, error) {
	var r = ProjectDependentsResponse{}

	// query options
	q := QueryOptions{
		ApiKey: l.ApiKey, Page: p.Page, PerPage: p.PerPage,
	}

	err := l.client.Get(API_URL+fmt.Sprintf("%s/%s/dependents?%s", platform, name, parseQuery(q)), &r)

	return r, err
}

// Get repositories that depend on a given project.
func (l *LibrariesIO) ProjectDependentRepositories(platform, name string, p PaginateOptions) (ProjectDependentRepositoriesResponse, error) {
	var r = ProjectDependentRepositoriesResponse{}

	// query options
	q := QueryOptions{
		ApiKey: l.ApiKey, Page: p.Page, PerPage: p.PerPage,
	}

	err := l.client.Get(API_URL+fmt.Sprintf("%s/%s/dependent_repositories?%s", platform, name, parseQuery(q)), &r)

	return r, err
}

// Get users that have contributed to a given project.
func (l *LibrariesIO) ProjectContributors(platform, name string, p PaginateOptions) (ProjectContributorsResponse, error) {
	var r = ProjectContributorsResponse{}

	q := QueryOptions{
		ApiKey: l.ApiKey, Page: p.Page, PerPage: p.PerPage,
	}

	err := l.client.Get(API_URL+fmt.Sprintf("%s/%s/contributors?%s", platform, name, parseQuery(q)), &r)

	return r, err
}

// Get breakdown of SourceRank score for a given project.
func (l *LibrariesIO) ProjectSourceRank(platform, name string) (ProjectSourceRankResponse, error) {
	var r = ProjectSourceRankResponse{}

	err := l.client.Get(API_URL+fmt.Sprintf("%s/%s/sourcerank?api_key=%s", platform, name, l.ApiKey), &r)
	return r, err
}

type ProjectSearchOptions struct {
	ApiKey    string     `url:"api_key,omitempty"`
	Page      int        `url:"page,omitempty"`
	PerPage   int        `url:"per_page,omitempty"`
	Sort      SearchSort `url:"sort,omitempty"`
	Languages []string   `url:"languages,omitempty" del:","`
	Licenses  []string   `url:"licenses,omitempty" del:","`
	Keywords  []string   `url:"keywords,omitempty" del:","`
	Platforms []string   `url:"platforms,omitempty" del:","`
}

// Search for projects.
func (l *LibrariesIO) ProjectSearch(search string, p ProjectSearchOptions) (ProjectSearchResponse, error) {
	var r = ProjectSearchResponse{}

	err := l.client.Get(API_URL+fmt.Sprintf("search?q=%s&", search)+parseQuery(p)+fmt.Sprintf("&api_key=%s", l.ApiKey), &r)

	return r, err
}
