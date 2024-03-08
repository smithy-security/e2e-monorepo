// Copyright (c) 2016 - 2020 Sqreen. All Rights Reserved.
// Please refer to our terms for more information:
// https://www.sqreen.io/terms.html

// Package api provides the base data structures of security signals.
// Higher-level signals can be built from there, such as HTTP traces, metrics,
// events, etc.
package api

import "time"

func NewSumMetric(name, source string, started, ended time.Time, interval time.Duration, values map[string]int64) *Metric {
	return NewMetric(name, source, started, newMetricPayload(started, ended, interval, "sum", values))
}

func NewBinningMetric(name, source string, started, ended time.Time, interval time.Duration, base, unit float64, bins map[string]int64, max float64) *Metric {
	return NewMetric(name, source, started, newBinningMetricPayload(started, ended, interval, "binning", base, unit, bins, max))
}

func newMetricPayload(started, ended time.Time, interval time.Duration, kind string, values map[string]int64) *SignalPayload {
	var header MetricSignalPayloadHeader
	makeMetricSignalPayloadHeader(&header, started, ended, interval, kind)

	kvArray := make([]MetricValueEntry, 0, len(values))
	for k, v := range values {
		kvArray = append(kvArray, MetricValueEntry{Key: k, Value: v})
	}

	return NewPayload(
		"metric/2020-01-01T00:00:00.000Z",
		MetricSignalPayload{
			MetricSignalPayloadHeader: header,
			Values:                    kvArray,
		},
	)
}

func newBinningMetricPayload(started, ended time.Time, interval time.Duration, kind string, base, unit float64, bins map[string]int64, max float64) *SignalPayload {
	var header MetricSignalPayloadHeader
	makeMetricSignalPayloadHeader(&header, started, ended, interval, kind)

	return NewPayload(
		"metric_binning/2020-01-01T00:00:00.000Z",
		BinningMetricsSignalPayload{
			MetricSignalPayloadHeader: header,
			Max:                       max,
			Unit:                      unit,
			Base:                      base,
			Bins:                      bins,
		},
	)
}

func makeMetricSignalPayloadHeader(header *MetricSignalPayloadHeader, started, ended time.Time, interval time.Duration, kind string) {
	captureIntervalSec := int64(interval / time.Second)

	*header = MetricSignalPayloadHeader{
		Type:               "metric",
		CaptureIntervalSec: captureIntervalSec,
		DateStarted:        started,
		DateEnded:          ended,
		Kind:               kind,
	}
}

type (
	MetricSignalPayloadHeader struct {
		Type               string    `json:"type,omitempty"`
		CaptureIntervalSec int64     `json:"capture_interval_s"`
		DateStarted        time.Time `json:"date_started"`
		DateEnded          time.Time `json:"date_ended"`
		Kind               string    `json:"kind"`
	}

	MetricSignalPayload struct {
		MetricSignalPayloadHeader
		Values []MetricValueEntry `json:"values"`
	}

	MetricValueEntry struct {
		Key   string `json:"key"`
		Value int64  `json:"value"`
	}

	BinningMetricsSignalPayload struct {
		MetricSignalPayloadHeader
		Max  float64          `json:"max"`
		Unit float64          `json:"unit"`
		Base float64          `json:"base"`
		Bins map[string]int64 `json:"bins"`
	}
)
