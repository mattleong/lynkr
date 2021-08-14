package lynkr

type Lynk struct {
	Id string `json:"id"`
	Url string `json:"url"`
	GoUrl string `json:"goUrl"`
}

func CreateLynk(goUrl string) *Lynk {
	id := generateLynkId(8)
	url := "/z/" + id
	return &Lynk{
		Id: id,
		Url: url,
		GoUrl: goUrl,
	}
}

