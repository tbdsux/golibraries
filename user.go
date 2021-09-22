package golibraries

import "fmt"

// User() response struct.
type UserResponse struct {
	GithubID     string      `json:"github_id"`
	Login        string      `json:"login"`
	UserType     string      `json:"user_type"`
	CreatedAt    string      `json:"created_at"`
	UpdatedAt    string      `json:"updated_at"`
	Name         string      `json:"name"`
	Company      interface{} `json:"company"`
	Blog         string      `json:"blog"`
	Location     string      `json:"location"`
	Hidden       bool        `json:"hidden"`
	LastSyncedAt string      `json:"last_synced_at"`
	Email        interface{} `json:"email"`
	Bio          string      `json:"bio"`
	UUID         string      `json:"uuid"`
	HostType     string      `json:"host_type"`
}

// UserRepositories() response struct.
type UserRepositoriesResponse []RepositoryResponse

// UserPackages() response struct.
type UserPackagesResponse []RepositoryProjectsResponseElement

// UserPackagesContributions() response struct.
type UserPackagesContributionsResponse []RepositoryProjectsResponseElement

// UserRepositoryContributions() response struct.
type UserRepositoryContributions []RepositoryResponse

// UserDependencies()
type UserDependenciesResponse []RepositoryProjectsResponseElement

// Get information for a givern user or organization.
func (l *LibrariesIO) User(user string) (UserResponse, error) {
	var r = UserResponse{}

	err := l.client.Get(API_URL+fmt.Sprintf("github/%s?api_key=%s", user, l.ApiKey), &r)

	return r, err
}

// Get repositories owned by a user.
func (l *LibrariesIO) UserRepositories(user string, p *PaginateOptions) (UserRepositoriesResponse, error) {
	var r = UserRepositoriesResponse{}

	// query options
	q := QueryOptions{
		ApiKey: l.ApiKey,
	}
	if p != nil {
		q.Page = p.Page
		q.PerPage = p.PerPage
	}

	err := l.client.Get(API_URL+fmt.Sprintf("github/%s/repositories?%s", user, parseQuery(q)), &r)

	return r, err
}

// Get a list of packages referencing the given user's repositories.
func (l *LibrariesIO) UserPackages(user string, p *PaginateOptions) (UserPackagesResponse, error) {
	var r = UserPackagesResponse{}

	// query options
	q := QueryOptions{
		ApiKey: l.ApiKey,
	}
	if p != nil {
		q.Page = p.Page
		q.PerPage = p.PerPage
	}

	err := l.client.Get(API_URL+fmt.Sprintf("github/%s/projects?%s", user, parseQuery(q)), &r)

	return r, err
}

// Get a list of packages that the given user has contributed to.
func (l *LibrariesIO) UserPackagesContributions(user string, p *PaginateOptions) (UserPackagesContributionsResponse, error) {
	var r = UserPackagesContributionsResponse{}

	// query options
	q := QueryOptions{
		ApiKey: l.ApiKey,
	}
	if p != nil {
		q.Page = p.Page
		q.PerPage = p.PerPage
	}

	err := l.client.Get(API_URL+fmt.Sprintf("github/%s/project-contributions?%s", user, parseQuery(q)), &r)

	return r, err
}

// Get a list of repositories that the given user has contributed to.
func (l *LibrariesIO) UserRepositoryContributions(user string, p *PaginateOptions) (UserRepositoriesResponse, error) {
	var r = UserRepositoriesResponse{}

	// query options
	q := QueryOptions{
		ApiKey: l.ApiKey,
	}
	if p != nil {
		q.Page = p.Page
		q.PerPage = p.PerPage
	}

	err := l.client.Get(API_URL+fmt.Sprintf("github/%s/repository-contributions?%s", user, parseQuery(q)), &r)

	return r, err
}

// Get a list of unique packages that the given user's repositories list as a dependency. Ordered by frequency of use in those repositories.
func (l *LibrariesIO) UserDependencies(user string, p *PaginateOptions) (UserDependenciesResponse, error) {
	var r = UserDependenciesResponse{}

	// query options
	q := QueryOptions{
		ApiKey: l.ApiKey,
	}
	if p != nil {
		q.Page = p.Page
		q.PerPage = p.PerPage
	}

	err := l.client.Get(API_URL+fmt.Sprintf("github/%s/dependencies?%s", user, parseQuery(q)), &r)

	return r, err
}
