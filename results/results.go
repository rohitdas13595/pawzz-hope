package results

type APIResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	Error   error  `json:"error"`
	Count   int    `json:"count"`
}

func NewAPIResponse[T any](status string, message string, data T, err error, count int) APIResponse[T] {
	return APIResponse[T]{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
		Count:   count,
	}
}
