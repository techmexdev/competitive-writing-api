package rest

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/felixge/httpsnoop"
)

type key string

const (
	requestIDKey key = "REQUEST_ID_KEY"
	userIDKey        = "USER_ID_KEY"
)

func requestID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func trace(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := requestID()
		ctx := context.WithValue(r.Context(), requestIDKey, reqID)
		w.Header().Set("X-Request-Id", reqID)
		hf(w, r.WithContext(ctx))
	}
}

func cors(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Conetent-Length, Accept-Encoding, X-CSRF-Token, Authentication")

		if r.Method == "OPTIONS" {
			return
		}

		hf(w, r)
	}
}

type hdl struct{ fn http.HandlerFunc }

func (h hdl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.fn(w, r)
}

func logRequest(lgr logger, hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		buf, _ := ioutil.ReadAll(r.Body)
		rBody := ioutil.NopCloser(bytes.NewBuffer(buf))
		rBodyCopy := ioutil.NopCloser(bytes.NewBuffer(buf))

		r.Body = rBody
		m := httpsnoop.CaptureMetrics(hdl{hf}, w, r)

		r.Body = rBodyCopy
		lgr.Request(r, m)
	}
}
