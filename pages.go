package instatus_go

import (
	"net/url"
	"time"
)

type Page struct {
	Id                   string     `json:"id"`
	Subdomain            string     `json:"subdomain"`
	Name                 string     `json:"name"`
	WorkspaceId          string     `json:"workspaceId"`
	LogoUrl              *string    `json:"logoUrl"`
	FaviconUrl           *string    `json:"faviconUrl"`
	WebsiteUrl           *string    `json:"websiteUrl"`
	CustomDomain         *string    `json:"customDomain"`
	PublicEmail          *string    `json:"publicEmail"`
	Twitter              *string    `json:"twitter"`
	Status               string     `json:"status"`
	SubscribeBySms       bool       `json:"subscribeBySms"`
	SendSmsNotifications bool       `json:"sendSmsNotifications"`
	Language             string     `json:"language"`
	UseLargeHeader       bool       `json:"useLargeHeader"`
	BrandColor           string     `json:"brandColor"`
	OkColor              string     `json:"okColor"`
	DisruptedColor       string     `json:"disruptedColor"`
	DegradedColor        string     `json:"degradedColor"`
	DownColor            string     `json:"downColor"`
	NoticeColor          string     `json:"noticeColor"`
	UnknownColor         string     `json:"unknownColor"`
	GoogleAnalytics      *string    `json:"googleAnalytics"`
	SmsService           *string    `json:"smsService"`
	HtmlInMeta           *string    `json:"htmlInMeta"`
	HtmlAboveHeader      *string    `json:"htmlAboveHeader"`
	HtmlBelowHeader      *string    `json:"htmlBelowHeader"`
	HtmlAboveFooter      *string    `json:"htmlAboveFooter"`
	HtmlBelowFooter      *string    `json:"htmlBelowFooter"`
	HtmlBelowSummary     *string    `json:"htmlBelowSummary"`
	UptimeDaysDisplay    string     `json:"uptimeDaysDisplay"`
	UptimeOutageDisplay  string     `json:"uptimeOutageDisplay"`
	LaunchDate           *time.Time `json:"launchDate"`
	CssGlobal            *string    `json:"cssGlobal"`
	Onboarded            *bool      `json:"onboarded"`
	CreatedAt            time.Time  `json:"createdAt"`
	UpdatedAt            time.Time  `json:"updatedAt"`
}

type GetPagesRequest struct {
	PageDetails
}

func (c *Client) GetPages(params GetPagesRequest) ([]Page, error) {
	targetUrl, _ := url.Parse(BaseUrl + "/v2/pages")
	applyPagination(targetUrl, params.PageDetails)

	var pages []Page
	err := c.get(targetUrl, &pages)

	return pages, err
}
