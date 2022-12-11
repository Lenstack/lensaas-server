package infrastructure

import (
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type IProxy interface {
	ServeHTTP(proxy *httputil.ReverseProxy) http.HandlerFunc
}

type Proxy struct {
	logger *zap.Logger
}

func NewProxy(targetUrl string, protocol string, logger *zap.Logger) (*httputil.ReverseProxy, error) {
	target, err := url.Parse(targetUrl)
	if err != nil {
		logger.Sugar().Errorf("Error parsing target url: %s", err)
		return nil, err
	}

	target.Scheme = "https"
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Transport = &http2.Transport{}

	proxy.ModifyResponse = func(response *http.Response) error {
		dumpResponse, err := httputil.DumpResponse(response, false)
		if err != nil {
			logger.Sugar().Errorf("Error dumping response: %v", err)
			return err
		}
		logger.Sugar().Infof("Response: \r\n%v", string(dumpResponse))
		return nil
	}
	return proxy, nil
}

func (p *Proxy) ServeHTTP(proxy *httputil.ReverseProxy) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		p.logger.Sugar().Infof("Request: %v", request)
		request.Header.Set("X-Forwarded-Host", request.Header.Get("Host"))
		proxy.ServeHTTP(writer, request)
	}
}
