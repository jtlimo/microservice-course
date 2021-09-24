package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type createMatrixRequest struct {
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=255"`
}

type createMatrixResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewCreateMatrixHandler(s domain.Service, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateMatrixEndpoint(s),
		decodeCreateMatrixRequest,
		encodeCreateMatrixResponse,
		opts...,
	)
}

func makeCreateMatrixEndpoint(s domain.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createMatrixRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		var m domain.Matrix
		data, _ := json.Marshal(req)
		err := json.Unmarshal(data, &m)
		if err != nil {
			return nil, err
		}

		created, err := s.CreateMatrix(ctx, &m)
		if err != nil {
			return nil, err
		}

		return createMatrixResponse{
			UUID:        created.UUID,
			Title:       created.Title,
			Description: created.Description,
			CreatedAt:   created.CreatedAt,
			UpdatedAt:   created.UpdatedAt,
		}, err
	}
}

func decodeCreateMatrixRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createMatrixRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateMatrixResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
