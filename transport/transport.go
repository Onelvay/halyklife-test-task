package transport

import (
	"github.com/Onelvay/halyklife-test-task/pkg/service"
	"net/http"
	"net/http/httputil"
)

type Transport struct {
	Audit *service.AuditServer
}

func (t *Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	response, err := http.DefaultTransport.RoundTrip(request)

	body, err := httputil.DumpResponse(response, true)
	if err != nil {
		return nil, err
	}
	t.Audit.LogResponse(string(body))
	return response, err
}
