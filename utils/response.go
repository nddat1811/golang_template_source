package utils

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PageInfo struct {
	Total        int         `json:"total"`
	CurrentTotal int         `json:"current_total"`
	CurrentPage  int         `json:"current_page"`
	Data         interface{} `json:"data"`
}

type PagingResponse struct {
	Message string   `json:"message"`
	Data    PageInfo `json:"data"`
}

func NewResponse(Message string, Data interface{}) *Response {
	return &Response{Message: Message, Data: Data}
}
