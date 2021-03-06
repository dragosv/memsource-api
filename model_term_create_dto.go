/*
 * Memsource REST API
 *
 * Welcome to Memsource's API documentation. To view our legacy APIs please [visit our documentation](https://wiki.memsource.com/wiki/Memsource_API) and for more information about our new APIs, [visit our blog](https://www.memsource.com/blog/2017/10/24/introducing-rest-apis-qa-with-the-memsource-api-team/).    If you have any questions, please contact [Memsource Support](https://get-help.memsource.com).
 *
 * API version: All
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TermCreateDto struct {
	Text string `json:"text"`
	Lang string `json:"lang,omitempty"`
	// Default: false
	CaseSensitive bool `json:"caseSensitive,omitempty"`
	// Default: false
	ExactMatch bool `json:"exactMatch,omitempty"`
	// Default: false
	Forbidden bool `json:"forbidden,omitempty"`
	// Default: false
	Preferred        bool   `json:"preferred,omitempty"`
	Status           string `json:"status,omitempty"`
	ConceptId        string `json:"conceptId,omitempty"`
	Usage            string `json:"usage,omitempty"`
	Note             string `json:"note,omitempty"`
	ShortTranslation string `json:"shortTranslation,omitempty"`
	TermType         string `json:"termType,omitempty"`
	PartOfSpeech     string `json:"partOfSpeech,omitempty"`
	Gender           string `json:"gender,omitempty"`
	Number           string `json:"number,omitempty"`
}
