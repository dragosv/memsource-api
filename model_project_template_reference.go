/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ProjectTemplateReference struct {
	TemplateName string                 `json:"templateName,omitempty"`
	SourceLang   string                 `json:"sourceLang,omitempty"`
	TargetLangs  []string               `json:"targetLangs,omitempty"`
	Id           string                 `json:"id,omitempty"`
	Owner        *UserReference         `json:"owner,omitempty"`
	Domain       *DomainReference       `json:"domain,omitempty"`
	SubDomain    *SubDomainReference    `json:"subDomain,omitempty"`
	CostCenter   *CostCenterReference   `json:"costCenter,omitempty"`
	BusinessUnit *BusinessUnitReference `json:"businessUnit,omitempty"`
	Note         string                 `json:"note,omitempty"`
	Client       *ClientReference       `json:"client,omitempty"`
}
