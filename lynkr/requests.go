package lynkr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateRequestBody struct {
	Url string
}

func NewLynkFromRequest(w http.ResponseWriter, r *http.Request) *RequestLynk {
	var body CreateRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	lynk, lynkErr := CreateRequestLynk(body.Url)
	if lynkErr != nil {
		return nil
	}

	fmt.Println("request lynk: ", lynk);

	return lynk
}
