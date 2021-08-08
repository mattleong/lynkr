package synkr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateRequestBody struct {
	Url string
}

type RequestLynk struct {
	Id string
	Url string
}

func CreateRequestLynk(url string) (*RequestLynk, error) {
	id := GenerateId(10)
	lynk := RequestLynk{ Id: id, Url: url }
	return &lynk, nil
}

func NewRequestLynk(w http.ResponseWriter, r *http.Request) *RequestLynk {
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
