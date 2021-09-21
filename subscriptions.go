package golibraries

import (
	"fmt"
	"net/http"
)

// UserSubscriptions() response struct.
type UserSubscriptionsResponse []UserSubscriptionsResponseElement

// SubscribeToProject() response struct.
type SubscribeToProjectResponse UserSubscriptionsResponseElement

// CheckSubscriptionToProject() response struct.
type CheckSubscriptionToProjectResponse UserSubscriptionsResponseElement

// UpdateSubscription() response struct.
type UpdateSubscriptionResponse UserSubscriptionsResponseElement

type UserSubscriptionsResponseElement struct {
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
	IncludePrerelease bool    `json:"include_prerelease"`
	Project           Project `json:"project"`
}

type Project struct {
	ID                             int64                 `json:"id"`
	Name                           string                `json:"name"`
	Platform                       string                `json:"platform"`
	CreatedAt                      string                `json:"created_at"`
	UpdatedAt                      string                `json:"updated_at"`
	Description                    string                `json:"description"`
	Keywords                       []string              `json:"keywords"`
	Homepage                       string                `json:"homepage"`
	Licenses                       string                `json:"licenses"`
	RepositoryURL                  string                `json:"repository_url"`
	RepositoryID                   int64                 `json:"repository_id"`
	NormalizedLicenses             []string              `json:"normalized_licenses"`
	VersionsCount                  int64                 `json:"versions_count"`
	Rank                           int64                 `json:"rank"`
	LatestReleasePublishedAt       string                `json:"latest_release_published_at"`
	LatestReleaseNumber            string                `json:"latest_release_number"`
	PmID                           interface{}           `json:"pm_id"`
	KeywordsArray                  []string              `json:"keywords_array"`
	DependentsCount                int64                 `json:"dependents_count"`
	Language                       string                `json:"language"`
	Status                         interface{}           `json:"status"`
	LastSyncedAt                   string                `json:"last_synced_at"`
	DependentReposCount            int64                 `json:"dependent_repos_count"`
	RuntimeDependenciesCount       int64                 `json:"runtime_dependencies_count"`
	Score                          int64                 `json:"score"`
	ScoreLastCalculated            string                `json:"score_last_calculated"`
	LatestStableReleaseNumber      string                `json:"latest_stable_release_number"`
	LatestStableReleasePublishedAt string                `json:"latest_stable_release_published_at"`
	LicenseSetByAdmin              bool                  `json:"license_set_by_admin"`
	LicenseNormalized              bool                  `json:"license_normalized"`
	DeprecationReason              interface{}           `json:"deprecation_reason"`
	PackageManagerURL              string                `json:"package_manager_url"`
	Stars                          int64                 `json:"stars"`
	Forks                          int64                 `json:"forks"`
	LatestStableRelease            string                `json:"latest_stable_release"`
	Versions                       []SubscriptionVersion `json:"versions"`
}

type SubscriptionVersion struct {
	Number      string `json:"number"`
	PublishedAt string `json:"published_at"`
}

// List packages that a user is subscribed to receive notifications about new releases.
func (l *LibrariesIO) UserSubscriptions(p PaginateOptions) (UserSubscriptionsResponse, error) {
	var r = UserSubscriptionsResponse{}

	q := QueryOptions{
		ApiKey: l.ApiKey, Page: p.Page, PerPage: p.PerPage,
	}

	err := l.client.Get(API_URL+fmt.Sprintf("subscriptions?%s", parseQuery(q)), &r)

	return r, err
}

// Subscribe to receive notifications about new releases of a project.
func (l *LibrariesIO) SubscribeToProject(platform string, name string, include_prerelease bool) (SubscribeToProjectResponse, error) {
	var r = SubscribeToProjectResponse{}

	err := l.client.Post(API_URL+fmt.Sprintf("subscriptions/%s/%s?include_prerelease=%v&api_key=%s", platform, name, include_prerelease, l.ApiKey), nil, &r)

	return r, err
}

// Check if a users is subscribed to receive notifications about new releases of a project.
func (l *LibrariesIO) CheckSubscriptionToProject(platform string, name string) (CheckSubscriptionToProjectResponse, error) {
	var r = CheckSubscriptionToProjectResponse{}

	err := l.client.Get(API_URL+fmt.Sprintf("subscriptions/%s/%s?api_key=%s", platform, name, l.ApiKey), &r)

	return r, err
}

// Update the options for a subscription.
func (l *LibrariesIO) UpdateSubscription(platform string, name string, include_prerelease bool) (UpdateSubscriptionResponse, error) {
	var r = UpdateSubscriptionResponse{}

	req, err := http.NewRequest("PUT", API_URL+fmt.Sprintf("subscriptions/%s/%s?include_prerelease=%v&api_key=%s", platform, name, include_prerelease, l.ApiKey), nil)
	if err != nil {
		return r, err
	}

	err = l.client.Request(req, &r)

	return r, err
}

// Stop receiving release notifications from a project.
func (l *LibrariesIO) UnsubscribeFromProject(platform string, name string) error {
	req, err := http.NewRequest("DELETE", API_URL+fmt.Sprintf("subscriptions/%s/%s?api_key=%s", platform, name, l.ApiKey), nil)
	if err != nil {
		return err
	}

	return l.client.Request(req, nil)
}
