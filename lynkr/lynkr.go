package lynkr

type RequestLynk struct {
	Id string
	Url string
}

type Lynk struct {
	Id string `json:"id"`
	Url string `json:"url"`
	GoUrl string `json:"goUrl"`
}

func CreateRequestLynk(url string) (*RequestLynk, error) {
	id := GenerateId(10)
	lynk := RequestLynk{ Id: id, Url: url }
	return &lynk, nil
}

