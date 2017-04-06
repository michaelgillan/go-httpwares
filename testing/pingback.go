package httpwares_testing

import (
	"encoding/json"
	"net/http"
)

const (
	DefaultPingBackStatusCode = http.StatusCreated
)

// PingBackResponse is a JSON struct that encodes the inbound request.
type PingBackResponse struct {
	ProtoMajor int               `json:"protoMajor"`
	Method     string            `json:"method"`
	UrlHost    string            `json:"urlHost"`
	UrlPath    string            `json:"urlPath"`
	HdHost     string            `json:"hdHost"`
	Headers    map[string]string `json:"headers"`
}

// PingBackHandler is an http.Handler that pings back the request info as JSON.
func PingBackHandler(retCode int) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		respJs := &PingBackResponse{
			ProtoMajor: req.ProtoMajor,
			Method:     req.Method,
			UrlHost:    req.URL.Host,
			UrlPath:    req.URL.Path,
			HdHost:     req.Host,
			Headers:    make(map[string]string),
		}
		for k, _ := range req.Header {
			respJs.Headers[k] = req.Header.Get(k)
		}

		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(retCode)
		json.NewEncoder(resp).Encode(respJs)
	}
}
