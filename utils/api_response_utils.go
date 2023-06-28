package utils

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ResponseWithPaginate struct {
	Meta      Meta        `json:"meta"`
	Data      interface{} `json:"data"`
	TotalList *int64      `json:"total_list"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// handle api response
func ApiResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

// handle api response with pagination
func ApiResponseWithPaginate(message string, code int, status string, data interface{}, total *int64) ResponseWithPaginate {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := ResponseWithPaginate{
		Meta:      meta,
		Data:      data,
		TotalList: total,
	}

	return jsonResponse
}
