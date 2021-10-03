package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

type deleteMatrixRequest struct {
	UUID string `json:"uuid" validate:"required"`
}

func NewDeleteMatrixHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeDeleteMatrixEndpoint(s),
		decodeDeleteMatrixRequest,
		encodeDeleteMatrixResponse,
		opts...,
	)
}

func makeDeleteMatrixEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(deleteMatrixRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		err := s.DeleteMatrix(ctx, req.UUID)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func decodeDeleteMatrixRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return deleteMatrixRequest{UUID: id}, nil
}

func encodeDeleteMatrixResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}