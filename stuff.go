package golibraries

type SearchSort string

const (
	SortRank                     SearchSort = "rank"
	SortStars                    SearchSort = "stars"
	SortDependentsCount          SearchSort = "dependents_count"
	SortDependentReposCount      SearchSort = "dependent_repos_count"
	SortLatestReleasePublishedAt SearchSort = "latest_release_published_at"
	SortContributorsCount        SearchSort = "contributions_count"
	SortCreatedAt                SearchSort = "created_at"
)
