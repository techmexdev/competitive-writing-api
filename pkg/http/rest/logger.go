package rest

import (
	"net/http"
	"strings"
	"time"

	"github.com/felixge/httpsnoop"
	log "github.com/sirupsen/logrus"
)

type logger struct {
	*log.Logger
}

func (l logger) Request(r *http.Request, m httpsnoop.Metrics) {
	var reqID string
	if id, ok := r.Context().Value(requestIDKey).(string); ok {
		reqID = id
	}

	var usrID string
	if id, ok := r.Context().Value(userIDKey).(string); ok {
		usrID = id
	}

	l.WithFields(log.Fields{
		"Type":          "Request",
		"Request ID":    reqID,
		"User ID":       usrID,
		"Method":        r.Method,
		"URL":           r.URL.String,
		"Query":         r.URL.Query(),
		"IP":            ipAddress(r),
		"User Agent":    r.UserAgent(),
		"Referer":       r.Header.Get("Referer"),
		"Headers":       r.Header,
		"Time":          m.Duration.Milliseconds,
		"Status Code":   m.Code,
		"Bytes Written": m.Written,
		"Timestamp":     time.Now().Local().String(),
	}).Info()
}

func ipAddress(r *http.Request) string {
	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		// X-Forwarded-For is potentially a list of addresses separated with ","
		ips := strings.Split(forwardedFor, ",")
		for i, ip := range ips {
			ips[i] = strings.TrimSpace(ip)
		}
		// TODO: should return first non-local address
		return ips[0]
	}

	if realIP := r.Header.Get("X-Real-Ip"); realIP != "" {
		return realIP
	}

	return r.RemoteAddr
}

func (l logger) serverError(r *http.Request, err error) {
	serverErrorLog := log.WithFields(log.Fields{
		"Type": "Server Error",
	})
	var reqID string
	if id, ok := r.Context().Value(requestIDKey).(string); ok {
		reqID = id
	}

	var usrID string
	if id, ok := r.Context().Value(userIDKey).(string); ok {
		usrID = id
	}

	serverErrorLog.WithFields(log.Fields{
		"Request ID": reqID,
		"User ID":    usrID,
		"Timestamp":  time.Now().Local().String(),
	}).Error(err)
}
