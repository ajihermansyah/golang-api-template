package response

type APIWithPaginationResponse struct {
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}
