package synkr

import (
	"github.com/mattleong/lynkr/lynkr"
)

// @TODO save link in db
func SaveLynk(requestLynk *lynkr.RequestLynk) (*lynkr.Lynk, error) {
	url := "/z/" + requestLynk.Id
	lynk := lynkr.Lynk{ Url: url }
	return &lynk, nil
}
