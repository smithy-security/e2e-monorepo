// Copyright (c) 2016 - 2020 Sqreen. All Rights Reserved.
// Please refer to our terms for more information:
// https://www.sqreen.io/terms.html

// Package http provides security signals describing HTTP traces. An HTTP trace
// describes a HTTP request and response, having a set of security signals such
// as events, attacks, errors, etc.
package http

import (
	"time"

	"github.com/sqreen/go-sdk/signal/client/api"
)

type Trace api.Trace

type Actor struct {
	IPAddresses []string          `json:"ip_addresses"`
	UserAgent   string            `json:"user_agent"`
	Identifiers map[string]string `json:"identifiers,omitempty"`
}

type (
	Context struct {
		Request  RequestContext  `json:"request"`
		Response ResponseContext `json:"response"`
	}

	RequestContext struct {
		Start      time.Time   `json:"start_processing_time"`
		End        time.Time   `json:"end_processing_time"`
		Headers    [][]string  `json:"headers"`
		UserAgent  string      `json:"user_agent"`
		Scheme     string      `json:"scheme"`
		Verb       string      `json:"verb"`
		Host       string      `json:"host"`
		Port       uint64      `json:"port"`
		RemoteIP   string      `json:"remote_ip"`
		RemotePort uint64      `json:"remote_port"`
		Path       string      `json:"path"`
		Referer    string      `json:"referer"`
		Parameters interface{} `json:"parameters"`
		Rid        string      `json:"rid"`
	}

	ResponseContext struct {
		Status        int    `json:"status"`
		ContentType   string `json:"content_type,omitempty"`
		ContentLength int64  `json:"content_length,omitempty"`
	}
)

func NewTrace(source string, t time.Time, a *Actor, infra interface{}, c *Context, d []*api.Signal) *Trace {
	return (*Trace)(api.NewTrace("", source, t, a, nil, infra, nil, newContext(c), nil, d))
}

func NewActor(ipAddresses []string, userAgent string, uid map[string]string) *Actor {
	return &Actor{
		IPAddresses: ipAddresses,
		UserAgent:   userAgent,
		Identifiers: uid,
	}
}

func newContext(context *Context) *api.SignalContext {
	return api.NewContext("http/2020-01-01T00:00:00.000Z", context)
}

func NewContext(req *RequestContext, resp *ResponseContext) *Context {
	return &Context{
		Request:  *req,
		Response: *resp,
	}
}

func NewRequestContext(start, end time.Time, requestID string, headers [][]string, userAgent, scheme, verb, host, remoteIP, path, referer string, port, remotePort uint64, parameters interface{}) *RequestContext {
	return &RequestContext{
		Start:      start,
		End:        end,
		Rid:        requestID,
		Headers:    headers,
		UserAgent:  userAgent,
		Scheme:     scheme,
		Verb:       verb,
		Host:       host,
		Port:       port,
		RemoteIP:   remoteIP,
		RemotePort: remotePort,
		Path:       path,
		Referer:    referer,
		Parameters: parameters,
	}
}

func NewResponseContext(status int, contentType string, contentLength int64) *ResponseContext {
	return &ResponseContext{
		Status:        status,
		ContentType:   contentType,
		ContentLength: contentLength,
	}
}
