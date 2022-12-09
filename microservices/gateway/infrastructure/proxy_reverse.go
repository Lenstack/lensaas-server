package infrastructure

import (
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ProxyReverse struct {
	logger          *zap.Logger
	proxy           *httputil.ReverseProxy
	remoteServerURL *url.URL
}

func NewProxyReverse(target string, logger *zap.Logger) *ProxyReverse {
	remoteServerURL, err := url.Parse(target)
	if err != nil {
		logger.Sugar().Fatal("Failed to parse origin server URL: ", err)
		return nil
	}
	return &ProxyReverse{proxy: httputil.NewSingleHostReverseProxy(remoteServerURL), remoteServerURL: remoteServerURL, logger: logger}
}

func (p *ProxyReverse) ProxyReverseHandler(proxy http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		p.logger.Sugar().Info("ProxyReverse: ", req.URL.Path)

		req.URL.Host = p.remoteServerURL.Host
		req.URL.Scheme = p.remoteServerURL.Scheme
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = p.remoteServerURL.Host

		proxy.ServeHTTP(res, req)
	}
}
