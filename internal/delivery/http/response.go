package http

import "net/http"

var (
	CheckInputData = "Hoops! Something was wrong with input data. Check it and try again!"
	OrderNotFound  = "Hoops! Order with this ID not found"
	OrderFound     = "Order found"
)

type (
	BaseResponse struct {
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
		Data       any    `json:"data"`
	}

	ResponseInput BaseResponse
)

func NewResponse(input *ResponseInput) *BaseResponse {
	return &BaseResponse{
		StatusCode: input.StatusCode,
		Message:    input.Message,
		Data:       input.Data,
	}
}

func OkResponse(input *ResponseInput) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusOK,
		Message:    input.Message,
		Data:       input.Data,
	}
}

func CreatedResponse(input *ResponseInput) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusCreated,
		Message:    input.Message,
		Data:       input.Data,
	}
}

func BadRequestResponse(input *ResponseInput) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusBadRequest,
		Message:    input.Message,
		Data:       input.Data,
	}
}

func NotFoundResponse(input *ResponseInput) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusNotFound,
		Message:    input.Message,
		Data:       input.Data,
	}
}

func InternalServerErrorResponse(input *ResponseInput) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    input.Message,
		Data:       input.Data,
	}
}
