package response

type Position struct {
	X       float32 `json:"x"`
	Y       float32 `json:"y"`
	Message string  `json:"message"`
}

type data struct {
	Result string
}

type Response struct {
	Position any    `json:"position,omitempty" mask:"struct"`
	Result   Result `json:"result"`
}

type ResponseWithList struct {
	Data   []interface{} `json:"data,omitempty" mask:"struct"`
	Result Result        `json:"result"`
}

type Result struct {
	Details []Detail `json:"details"`
	Source  string   `json:"source"`
}

type Detail struct {
	InternalCode string `json:"internalCode"`
	Message      string `json:"message"`
	Detail       string `json:"detail"`
}
