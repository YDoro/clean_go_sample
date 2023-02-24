package productservice

import (
	"encoding/json"
	"net/http"

	"github.com/ydoro/clean-go/core/dto"
)

func (s service) Create(response http.ResponseWriter, request *http.Request) {
	pr, err := dto.FromJSONCreateProductRequest(request.Body)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	product, err := s.usecase.Create(pr)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(response).Encode(product)
}
