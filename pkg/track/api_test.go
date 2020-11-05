/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package track

import (
	"context"
	"testing"
	"time"

	"mosn.io/mosn/pkg/buffer"
)

func TestTrackFromContext(t *testing.T) {
	ctx := buffer.NewBufferPoolContext(context.Background())
	defer func() {
		if c := buffer.PoolContext(ctx); c != nil {
			c.Give()
		}
	}()
	for _, ph := range []TrackPhase{
		ProtocolDecode, StreamFilterBeforeRoute, MatchRoute,
	} {
		StartTrack(ctx, ph)
		EndTrack(ctx, ph)
	}
	RangeCosts(ctx, func(p TrackPhase, tk TrackTime) bool {
		switch p {
		case ProtocolDecode, StreamFilterBeforeRoute, MatchRoute:
			if len(tk.Costs) != 1 {
				t.Fatalf("%d phase is not setted", p)
			}
		default:
		}
		return true
	})
	s := GetTrackCosts(ctx)
	if !outexp.MatchString(s) {
		t.Fatalf("unexpected output: %s", s)
	}
	t.Logf("output is %s", s)
}

func TestTrackTransmit(t *testing.T) {
	dstCtx := buffer.NewBufferPoolContext(context.Background())
	srcCtx := buffer.NewBufferPoolContext(context.Background())
	defer func() {
		for _, ctx := range []context.Context{
			dstCtx, srcCtx,
		} {
			if c := buffer.PoolContext(ctx); c != nil {
				c.Give()
			}
		}
	}()
	// set value
	AddDataReceived(dstCtx)
	for _, ph := range []TrackPhase{
		ProtocolDecode, StreamFilterBeforeRoute, MatchRoute,
	} {
		StartTrack(dstCtx, ph)
		EndTrack(dstCtx, ph)
	}
	time.Sleep(100 * time.Millisecond)
	StartTrack(srcCtx, ProtocolDecode)
	EndTrack(srcCtx, ProtocolDecode)
	AddDataReceived(srcCtx)
	// Transmit
	TransmitBufferByContext(dstCtx, srcCtx)
	// Verify
	RangeCosts(dstCtx, func(p TrackPhase, tk TrackTime) bool {
		switch p {
		case ProtocolDecode:
			if len(tk.Costs) != 2 {
				t.Fatalf("%d phase is not setted", p)
			}
		case StreamFilterBeforeRoute, MatchRoute:
			if len(tk.Costs) != 1 {
				t.Fatalf("%d phase is not setted", p)
			}
		default:
		}
		return true
	})
	ts := GetDataReceived(dstCtx)
	if len(ts) != 2 {
		t.Fatalf("no data received")
	}
}
