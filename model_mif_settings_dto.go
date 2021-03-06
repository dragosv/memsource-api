/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type MifSettingsDto struct {
	// Default: true
	ExtractBodyPages bool `json:"extractBodyPages,omitempty"`
	// Default: false
	ExtractReferencePages bool `json:"extractReferencePages,omitempty"`
	// Default: true
	ExtractMasterPages bool `json:"extractMasterPages,omitempty"`
	// Default: false
	ExtractHiddenPages bool `json:"extractHiddenPages,omitempty"`
	// Default: false
	ExtractVariables bool `json:"extractVariables,omitempty"`
	// Default: true
	ExtractIndexMarkers bool `json:"extractIndexMarkers,omitempty"`
	// Default: false
	ExtractLinks bool `json:"extractLinks,omitempty"`
	// Default: false
	ExtractXRefDef bool `json:"extractXRefDef,omitempty"`
	// Default: true
	ExtractPgfNumFormat bool `json:"extractPgfNumFormat,omitempty"`
	// Default: true
	ExtractCustomReferencePages bool `json:"extractCustomReferencePages,omitempty"`
	// Default: false
	ExtractDefaultReferencePages bool `json:"extractDefaultReferencePages,omitempty"`
	// Default: true
	ExtractUsedVariables bool `json:"extractUsedVariables,omitempty"`
	// Default: false
	ExtractHiddenCondText bool `json:"extractHiddenCondText,omitempty"`
	// Default: true
	ExtractUsedXRefDef bool `json:"extractUsedXRefDef,omitempty"`
	// Default: true
	ExtractUsedPgfNumFormat bool `json:"extractUsedPgfNumFormat,omitempty"`
}
