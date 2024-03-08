// Copyright (c) 2016 - 2020 Sqreen. All Rights Reserved.
// Please refer to our terms for more information:
// https://www.sqreen.io/terms.html

package client

import (
	"context"
	"errors"

	"github.com/sqreen/go-sdk/signal/client/api"
)

type SignalService Client

func (s *SignalService) unwrap() *Client { return (*Client)(s) }

func (s *SignalService) SendBatch(ctx context.Context, b api.Batch) error {
	if len(b) == 0 {
		return errors.New("unexpected empty batch")
	}
	c := s.unwrap()
	r, err := c.newRequest("POST", "batches", b)
	if err != nil {
		return err
	}
	return c.do(ctx, r, nil)
}

func (s *SignalService) SendTrace(ctx context.Context, trace *api.Trace) error {
	if trace == nil {
		return errors.New("unexpected trace argument value `nil`")
	}
	if len(trace.Data) == 0 {
		return errors.New("unexpected empty trace data array")
	}
	c := s.unwrap()
	r, err := c.newRequest("POST", "traces", trace)
	if err != nil {
		return err
	}
	return c.do(ctx, r, nil)
}

func (s *SignalService) SendSignal(ctx context.Context, signal *api.Signal) error {
	if signal == nil {
		return errors.New("unexpected signal argument value `nil`")
	}
	c := s.unwrap()
	r, err := c.newRequest("POST", "signals", signal)
	if err != nil {
		return err
	}
	return c.do(ctx, r, nil)
}
