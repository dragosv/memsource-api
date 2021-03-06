/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type QuoteCreateDto struct {
	Name             string             `json:"name"`
	Project          *UidReference      `json:"project"`
	Analyse          *IdReference       `json:"analyse"`
	PriceList        *IdReference       `json:"priceList"`
	NetRateScheme    *IdReference       `json:"netRateScheme,omitempty"`
	WorkflowStepList []IdReference      `json:"workflowStepList,omitempty"`
	Provider         *ProviderReference `json:"provider,omitempty"`
}
