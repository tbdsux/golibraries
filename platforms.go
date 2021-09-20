package golibraries

type PlatformsResponse []PlatformsResponseElement

type PlatformsResponseElement struct {
	Name            string  `json:"name"`
	ProjectCount    int64   `json:"project_count"`
	Homepage        string  `json:"homepage"`
	Color           string  `json:"color"`
	DefaultLanguage *string `json:"default_language"`
}

// Get a list of supported package managers.
func (l *LibrariesIO) Platforms(p PaginateOptions) (PlatformsResponse, error) {
	var r = PlatformsResponse{}

	// query options
	q := QueryOptions{
		ApiKey: l.ApiKey, Page: p.Page, PerPage: p.PerPage,
	}

	l.client.Get(API_URL+"/platforms?"+parseQuery(q), &r)

	return r, nil
}
