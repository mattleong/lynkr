package lynkr

import (
	"net/http"
	"encoding/json"
)

type Lynk struct {
	Id string `json:"id"`
	Url string `json:"url"`
	GoUrl string `json:"goUrl"`
}

type RequestLynk struct {
	Id string
	Url string
}

type createRequestBody struct {
	Url string
}

func CreateLynk(id string, url string, goUrl string) *Lynk {
	return &Lynk{
		Id: id,
		Url: url,
		GoUrl: goUrl,
	}
}

func newRequestLynk(w http.ResponseWriter, r *http.Request) *RequestLynk {
	var body createRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	id := GenerateId(10)

	return &RequestLynk{ Id: id, Url: body.Url }
}

