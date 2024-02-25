package instatus_go

import (
	"net/url"
	"time"
)

type Translations struct {
	Name        map[string]string `json:"name"`
	Description map[string]string `json:"description"`
}

type Component struct {
	Id                       string  `json:"id"`
	Name                     string  `json:"name"`
	NameTranslationId        *string `json:"nameTranslationId"`
	Description              *string `json:"description"`
	DescriptionTranslationId *string `json:"descriptionTranslationId"`
	// TODO: Enum
	Status                       string       `json:"status"`
	InternalStatus               string       `json:"internalStatus"`
	Order                        int          `json:"order"`
	ShowUptime                   bool         `json:"showUptime"`
	CreatedAt                    time.Time    `json:"createdAt"`
	UpdatedAt                    time.Time    `json:"updatedAt"`
	ArchivedAt                   *time.Time   `json:"archivedAt"`
	SiteId                       string       `json:"siteId"`
	UniqueEmail                  string       `json:"uniqueEmail"`
	OldGroup                     any          `json:"oldGroup"`
	GroupId                      *string      `json:"groupId"`
	IsParent                     bool         `json:"isParent"`
	IsCollapsed                  bool         `json:"isCollapsed"`
	MonitorId                    *string      `json:"monitorId"`
	NameHtml                     *string      `json:"nameHtml"`
	NameHtmlTranslationId        *string      `json:"nameHtmlTranslationId"`
	DescriptionHtml              *string      `json:"descriptionHtml"`
	DescriptionHtmlTranslationId *string      `json:"descriptionHtmlTranslationId"`
	IsThirdParty                 bool         `json:"isThirdParty"`
	ThirdPartyStatus             *string      `json:"thirdPartyStatus"`
	ThirdPartyComponentId        *string      `json:"thirdPartyComponentId"`
	ThirdPartyComponentServiceId *string      `json:"thirdPartyComponentServiceId"`
	ImportedFromStatuspage       bool         `json:"importedFromStatuspage"`
	StartDate                    *time.Time   `json:"startDate"`
	Group                        *Component   `json:"group"`
	Translations                 Translations `json:"translations"`
}

type GetComponentsRequest struct {
	PageDetails

	PageId string
}

func (c *Client) GetComponents(params GetComponentsRequest) ([]Component, error) {
	targetUrl, _ := url.Parse(BaseUrl + "/v1/" + params.PageId + "/components")
	applyPagination(targetUrl, params.PageDetails)

	var components []Component
	err := c.get(targetUrl, &components)

	return components, err
}

type GetComponentRequest struct {
	PageId      string
	ComponentId string
}

func (c *Client) GetComponent(params GetComponentRequest) (Component, error) {
	targetUrl, _ := url.Parse(BaseUrl + "/v1/" + params.PageId + "/components/" + params.ComponentId)

	var component Component
	err := c.get(targetUrl, &component)

	return component, err
}

type UpdateComponentFields struct {
	Status *string `json:"status,omitempty"`

	// TODO: More comprehensive set of fields
}

type UpdateComponentRequest struct {
	PageId      string
	ComponentId string

	UpdatedFields UpdateComponentFields
}

func (c *Client) UpdateComponent(params UpdateComponentRequest) (Component, error) {
	targetUrl, _ := url.Parse(BaseUrl + "/v1/" + params.PageId + "/components/" + params.ComponentId)

	var component Component
	err := c.put(targetUrl, params.UpdatedFields, &component)

	return component, err
}
