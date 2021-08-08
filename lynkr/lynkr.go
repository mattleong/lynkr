package lynkr

type Lynk struct {
	Id string `json:"id"`
	Url string `json:"url"`
	GoUrl string `json:"goUrl"`
}

func CreateLynk(id string, url string, goUrl string) *Lynk {
	return &Lynk{
		Id: id,
		Url: url,
		GoUrl: goUrl,
	}
}
