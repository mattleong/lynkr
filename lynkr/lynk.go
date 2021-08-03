package lynkr

type Lynk struct {
	Id string
	Url string
}

func CreateLynk(url string) *Lynk {
	id := RandomString(10)
	lynk := Lynk{ Id: id, Url: url }
	return &lynk
}

