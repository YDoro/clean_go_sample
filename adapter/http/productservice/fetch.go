package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/ydoro/clean-go/core/dto"
)

func (s service) Fetch(response http.ResponseWriter, request *http.Request) {
	pr, err := dto.FromValuePaginationRequestParams(request)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	products, err := s.usecase.Fetch(pr)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(response).Encode(products)
}
