package web

type ErrorStruct struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Url     string         `json:"url"`
	Method  string         `json:"method"`
	Details *[]interface{} `json:"details,omitempty"`
}
