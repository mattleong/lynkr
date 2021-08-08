package lynkr

type RequestLynk struct {
	Id string
	Url string
}

type Lynk struct {
	Url string
}

func CreateRequestLynk(url string) (*RequestLynk, error) {
	id := RandomString(10)
	lynk := RequestLynk{ Id: id, Url: url }
	return &lynk, nil
}

