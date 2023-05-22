package helper

import "go.uber.org/zap"

type HttpResponse struct {
	logger *zap.Logger
}

func NewHttpResponse(logger *zap.Logger) *HttpResponse {
	return &HttpResponse{
		logger: logger,
	}
}
