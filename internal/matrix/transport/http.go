package transport

import (
	"net/http"

	"github.com/sumelms/microservice-course/internal/matrix/endpoints"
	"github.com/sumelms/microservice-course/pkg/errors"

	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"

	"github.com/sumelms/microservice-course/internal/matrix/domain"
)

func NewHTTPHandler(r *mux.Router, s domain.ServiceInterface, logger log.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	listMatrixHandler := endpoints.NewListMatrixHandler(s, opts...)
	createMatrixHandler := endpoints.NewCreateMatrixHandler(s, opts...)
	findMatrixHandler := endpoints.NewFindMatrixHandler(s, opts...)
	updateMatrixHandler := endpoints.NewUpdateMatrixHandler(s, opts...)
	deleteMatrixHandler := endpoints.NewDeleteMatrixHandler(s, opts...)

	r.Handle("/matrices", listMatrixHandler).Methods(http.MethodGet)
	r.Handle("/matrices", createMatrixHandler).Methods(http.MethodPost)
	r.Handle("/matrices/{uuid}", findMatrixHandler).Methods(http.MethodGet)
	r.Handle("/matrices/{uuid}", updateMatrixHandler).Methods(http.MethodPut)
	r.Handle("/matrices/{uuid}", deleteMatrixHandler).Methods(http.MethodDelete)

	listSubjectHandler := endpoints.NewListSubjectHandler(s, opts...)
	createSubjectHandler := endpoints.NewCreateSubjectHandler(s, opts...)
	findSubjectHandler := endpoints.NewFindSubjectHandler(s, opts...)
	updateSubjectHandler := endpoints.NewUpdateSubjectHandler(s, opts...)
	deleteSubjectHandler := endpoints.NewDeleteSubjectHandler(s, opts...)

	r.Handle("/subjects", createSubjectHandler).Methods(http.MethodPost)
	r.Handle("/subjects", listSubjectHandler).Methods(http.MethodGet)
	r.Handle("/subjects/{uuid}", findSubjectHandler).Methods(http.MethodGet)
	r.Handle("/subjects/{uuid}", updateSubjectHandler).Methods(http.MethodPut)
	r.Handle("/subjects/{uuid}", deleteSubjectHandler).Methods(http.MethodDelete)
}
