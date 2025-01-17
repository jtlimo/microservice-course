package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/pkg/validator"
)

type updateSubjectRequest struct {
	UUID      uuid.UUID `json:"uuid" validate:"required"`
	Code      string    `json:"code" validate:"required,max=45"`
	Name      string    `json:"name" validate:"required,max=100"`
	Objective string    `json:"objective" validate:"max=245"`
	Credit    float32   `json:"credit"`
	Workload  float32   `json:"workload"`
}

type updateSubjectResponse struct {
	UUID      uuid.UUID `json:"uuid"`
	Code      string    `json:"code" validate:"required,max=45"`
	Name      string    `json:"name" validate:"required,max=100"`
	Objective string    `json:"objective,omitempty" validate:"max=245"`
	Credit    float32   `json:"credit,omitempty"`
	Workload  float32   `json:"workload,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUpdateSubjectHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeUpdateSubjectEndpoint(s),
		decodeUpdateSubjectRequest,
		encodeUpdateSubjectResponse,
		opts...,
	)
}

//nolint:dupl
func makeUpdateSubjectEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(updateSubjectRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		c := domain.Subject{}
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, err
		}

		if err := s.UpdateSubject(ctx, &c); err != nil {
			return nil, err
		}

		return updateSubjectResponse{
			UUID:      c.UUID,
			Code:      c.Code,
			Name:      c.Name,
			Objective: c.Objective,
			Credit:    c.Credit,
			Workload:  c.Workload,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		}, nil
	}
}

func decodeUpdateSubjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	var req updateSubjectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.UUID = uuid.MustParse(id)

	return req, nil
}

func encodeUpdateSubjectResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
