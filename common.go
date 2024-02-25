package instatus_go

import (
	"fmt"
	"net/url"
)

const BaseUrl = "https://api.instatus.com"

type Error struct {
	Details struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Details.Code, e.Details.Message)
}

type PageDetails struct {
	Page    int
	PerPage int
}

func applyPagination(targetUrl *url.URL, details PageDetails) {
	page := details.Page
	if page == 0 {
		page = 1
	}

	perPage := details.PerPage
	if perPage == 0 {
		perPage = 50
	}

	q := targetUrl.Query()
	q.Set("page", fmt.Sprintf("%d", page))
	q.Set("per_page", fmt.Sprintf("%d", perPage))
	targetUrl.RawQuery = q.Encode()
}
