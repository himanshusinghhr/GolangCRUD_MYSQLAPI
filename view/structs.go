package view

//Response
type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

//PostRequest
type PostRequest struct {
	Name string `json:"name"`
	Todo string `jsob:"todo"`
}
