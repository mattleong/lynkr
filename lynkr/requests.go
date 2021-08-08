package lynkr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateRequestBody struct {
	Url string
}

func NewLynkFromRequest(w http.ResponseWriter, r *http.Request) *Lynk {
	var body CreateRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	lynk, lynkErr := CreateLynk(body.Url)
	if lynkErr != nil {
		return nil
	}

	fmt.Println("lynk: ", lynk);

	return lynk
}
