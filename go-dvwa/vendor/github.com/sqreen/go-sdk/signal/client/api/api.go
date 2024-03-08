// Copyright (c) 2016 - 2020 Sqreen. All Rights Reserved.
// Please refer to our terms for more information:
// https://www.sqreen.io/terms.html

package api

import "time"

type Point Signal

func NewPoint(name, source string, t time.Time, actor, trigger, infra, location interface{}, context *SignalContext, payload *SignalPayload) *Point {
	return (*Point)(newSignal("point", name, source, t, actor, trigger, infra, location, context, payload))
}

type Metric Signal

func NewMetric(name, source string, t time.Time, payload *SignalPayload) *Metric {
	return (*Metric)(newSignal("metric", name, source, t, nil, nil, nil, nil, nil, payload))
}

func newSignal(typ, name, source string, t time.Time, actor, trigger, infra, location interface{}, context *SignalContext, payload *SignalPayload) *Signal {
	return &Signal{
		SignalPayload: payload,
		Type:          typ,
		Name:          name,
		Source:        source,
		Time:          t,
		Actor:         actor,
		SignalContext: context,
		Trigger:       trigger,
		LocationInfra: infra,
		Location:      location,
	}
}

type (
	Signal struct {
		Type          string      `json:"type"`
		Name          string      `json:"signal_name,omitempty"`
		Source        string      `json:"source,omitempty"`
		Time          time.Time   `json:"time,omitempty"`
		Actor         interface{} `json:"actor,omitempty"`
		Trigger       interface{} `json:"trigger,omitempty"`
		LocationInfra interface{} `json:"location_infra,omitempty"`
		Location      interface{} `json:"location,omitempty"`
		*SignalPayload
		*SignalContext
	}

	SignalPayload struct {
		Schema  string      `json:"payload_schema,omitempty"`
		Payload interface{} `json:"payload,omitempty"`
	}

	SignalContext struct {
		Schema  string      `json:"context_schema,omitempty"`
		Context interface{} `json:"context,omitempty"`
	}
)

// Trace is a set of signals. Common signal fields can be factored in the trace
// root fields.
type Trace struct {
	Signal
	Data []*Signal `json:"data"`
}

func NewTrace(name, source string, t time.Time, actor, trigger, infra, location interface{}, context *SignalContext, payload *SignalPayload, d []*Signal) *Trace {
	return &Trace{
		Signal: *newSignal("trace", name, source, t, actor, trigger, infra, location, context, payload),
		Data:   d,
	}
}

func NewContext(schema string, context interface{}) *SignalContext {
	return &SignalContext{
		Schema:  schema,
		Context: context,
	}
}

func NewPayload(schema string, payload interface{}) *SignalPayload {
	return &SignalPayload{
		Schema:  schema,
		Payload: payload,
	}
}

type (
	Batch []SignalFace

	// SignalFace is a simple helper to make sure only a given set of types defined
	// in this package can be added to the batch array (private interface method
	// can indeed only be implemented in the same package).
	SignalFace interface {
		isSignal()
	}
)

func (Signal) isSignal() {}
func (Point) isSignal()  {}
func (Metric) isSignal() {}

// Static assert that SignalFace is correctly implemented.
var (
	_ SignalFace = &Trace{}
	_ SignalFace = &Point{}
	_ SignalFace = &Metric{}
	_ SignalFace = Trace{}
	_ SignalFace = Point{}
	_ SignalFace = Metric{}
)
