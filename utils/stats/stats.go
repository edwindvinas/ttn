// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

// Package stats supports the collection of metrics from a running component.
package stats

import (
	"time"

	"github.com/rcrowley/go-metrics"
)

// Enabled activates stats collection
var Enabled = true

// Initialize the stats
func Initialize() {
	go func() {
		c := time.Tick(1 * time.Minute)
		for _ = range c {
			Registry.(Ticker).Tick()
		}
	}()
}

// MarkMeter registers an event
func MarkMeter(name string) {
	if Enabled {
		metrics.GetOrRegisterMeter(name, Registry).Mark(1)
		Registry.(Ticker).Renew(name)
	}
}

// UpdateHistogram registers a new value for a histogram
func UpdateHistogram(name string, value int64) {
	if Enabled {
		metrics.GetOrRegisterHistogram(name, Registry, metrics.NewUniformSample(1000)).Update(value)
		Registry.(Ticker).Renew(name)
	}
}

// IncCounter increments a counter by 1
func IncCounter(name string) {
	if Enabled {
		metrics.GetOrRegisterCounter(name, Registry).Inc(1)
		Registry.(Ticker).Renew(name)
	}
}

// DecCounter decrements a counter by 1
func DecCounter(name string) {
	if Enabled {
		metrics.GetOrRegisterCounter(name, Registry).Dec(1)
		Registry.(Ticker).Renew(name)
	}
}

// SetString ...
func SetString(name, tag, value string) {
	if Enabled {
		GetOrRegisterString(name, Registry).Set(tag, value)
		Registry.(Ticker).Renew(name)
	}
}

// SetTimeout ...
func SetTimeout(name string, ttl uint) {
	Registry.(Ticker).SetTTL(name, ttl)
}
