package rest

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"

	"github.com/go-resty/resty/v2"

	"github.com/spacetab-io/solar-staff-sdk-go/config"
)

type Repo struct {
	resty *resty.Client
}

// NewRepo returns a new rest api Repo
func NewRepo(cfg *config.SolarStaffConfig) *Repo {
	s := new(Repo)
	s.resty = resty.New()
	s.resty.SetDebug(true) //configuration.Config.Log.Debug)
	//s.resty.SetLogger(logs.Log.Writer())
	// nolint:gomnd
	s.resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(5))
	// Registering Request Middleware
	s.resty.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		req.SetHeader("Content-Type", "application/json")
		req.URL = cfg.URL + req.URL
		bodyParams := req.Body.(map[string]interface{})
		bodyParams["client_id"] = cfg.ClientID
		bodyParams["signature"] = makeSignature(cfg.Salt, bodyParams)
		req.SetBody(bodyParams)

		return nil // if its success otherwise return error
	})

	return s
}

//func makeURL(service string, path string) string {
//	return strings.TrimRight(configuration.Config.Services[service].URL, "/") + "/" + strings.TrimLeft(path, "/")
//}

func (r *Repo) sendPOST(uri string, bodyParams map[string]interface{}) (*resty.Response, error) {
	return r.resty.R().SetBody(bodyParams).Post(uri)
}

func makeSignature(salt string, params map[string]interface{}) string {
	paramsValues := make([]string, 0)
	for p, v := range params {
		paramValue := p + ":" + fmt.Sprint(v)

		if v == "" {
			paramValue = p
		}

		paramsValues = append(paramsValues, paramValue)
	}
	sort.Slice(paramsValues, func(i, j int) bool {
		return paramsValues[i] < paramsValues[j]
	})

	str := strings.Join(paramsValues, ";") + ";" + salt
	//return str
	h := sha1.New()
	h.Write([]byte(str))

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}
