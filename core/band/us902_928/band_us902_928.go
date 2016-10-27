// Copyright © 2016 The Things Network
// Copyright © 2016 Orne Brocaar
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package us902_928

import (
	"errors"
	"fmt"
	"time"

	. "github.com/edwindvinas/ttn/core/band"
)

func init() {
	// initialize uplink channel 0 - 63
	for i := 0; i < 64; i++ {
		UplinkChannelConfiguration[i] = Channel{
			Frequency: 902300000 + (i * 200000),
			DataRates: []int{0, 1, 2, 3},
		}
	}

	// initialize uplink channel 64 - 71
	for i := 0; i < 8; i++ {
		UplinkChannelConfiguration[i+64] = Channel{
			Frequency: 903000000 + (i * 1600000),
			DataRates: []int{4},
		}
	}

	// initialize downlink channel 0 - 7
	for i := 0; i < 8; i++ {
		DownlinkChannelConfiguration[i] = Channel{
			Frequency: 923300000 + (i * 600000),
			DataRates: []int{10, 11, 12, 13},
		}
	}
}

// Name defines the name of the band
const Name = "US 902-928"

// DataRateConfiguration defines the available data rates
var DataRateConfiguration = [...]DataRate{
	{Modulation: LoRaModulation, SpreadFactor: 10, Bandwidth: 125},
	{Modulation: LoRaModulation, SpreadFactor: 9, Bandwidth: 125},
	{Modulation: LoRaModulation, SpreadFactor: 8, Bandwidth: 125},
	{Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 125},
	{Modulation: LoRaModulation, SpreadFactor: 8, Bandwidth: 500},
	{}, // RFU
	{}, // RFU
	{}, // RFU
	{Modulation: LoRaModulation, SpreadFactor: 12, Bandwidth: 500},
	{Modulation: LoRaModulation, SpreadFactor: 11, Bandwidth: 500},
	{Modulation: LoRaModulation, SpreadFactor: 10, Bandwidth: 500},
	{Modulation: LoRaModulation, SpreadFactor: 9, Bandwidth: 500},
	{Modulation: LoRaModulation, SpreadFactor: 8, Bandwidth: 500},
	{Modulation: LoRaModulation, SpreadFactor: 7, Bandwidth: 500},
	{}, // RFU
	{}, // RFU
}

// DefaultTXPower defines the default TX power in dBm
const DefaultTXPower = 20

// CFListAllowed defines if the optional JoinAccept CFList is allowed for this band
const CFListAllowed = false

// TXPowerConfiguration defines the available TXPower settings in dBm
var TXPowerConfiguration = [...]int{
	30,
	28,
	26,
	24,
	22,
	20,
	18,
	16,
	14,
	12,
	10,
	0,
	0,
	0,
	0,
	0,
}

// MACPayloadSizeConfiguration defines the maximum payload size for each data rate
var MACPayloadSizeConfiguration = [...]MaxPayloadSize{
	{M: 19, N: 11},
	{M: 61, N: 53},
	{M: 137, N: 129},
	{M: 250, N: 242},
	{M: 250, N: 242},
	{}, // Not defined
	{}, // Not defined
	{}, // Not defined
	{M: 41, N: 33},
	{M: 117, N: 109},
	{M: 230, N: 222},
	{M: 230, N: 222},
	{M: 230, N: 222},
	{M: 230, N: 222},
	{}, // Not defined
	{}, // Not defined
}

// RX1DROffsetConfiguration defines the available RX1DROffset configurations
// per data rate.
var RX1DROffsetConfiguration = [...][4]int{
	{10, 9, 8, 8},
	{11, 10, 9, 8},
	{12, 11, 10, 9},
	{13, 12, 11, 10},
	{13, 13, 12, 11},
	{}, // Not defined
	{}, // Not defined
	{}, // Not defined
	{8, 8, 8, 8},
	{9, 8, 8, 8},
	{10, 9, 8, 8},
	{11, 10, 9, 8},
	{12, 11, 10, 9},
	{13, 12, 11, 10},
}

// RX2Frequency defines the RX2 receive window frequency to use (in Hz)
const RX2Frequency = 923300000

// RX2DataRate defines the RX2 receive window data rate to use
const RX2DataRate = 8

// UplinkChannelConfiguration defines the (default) available uplink channels.
var UplinkChannelConfiguration = [72]Channel{}

// DownlinkChannelConfiguration defines the (default) available downlink channels.
var DownlinkChannelConfiguration = [8]Channel{}

// GetDataRate returns the index of the given DataRate.
func GetDataRate(dr DataRate) (int, error) {
	for i, d := range DataRateConfiguration {
		if d == dr {
			return i, nil
		}
	}
	return 0, errors.New("lorawan/band: the given DataRate does not exist")
}

// GetRX1Frequency returns the frequency to be used for RX1 given
// the uplink frequency and data rate.
func GetRX1Frequency(frequency, dataRate int) (int, error) {
	if dataRate > len(DataRateConfiguration) {
		return 0, fmt.Errorf("lorawan/band: given data rate: %d does not exist", dataRate)
	}

	chanNum, err := getChannelNumber(frequency, dataRate)
	if err != nil {
		return 0, err
	}

	return DownlinkChannelConfiguration[chanNum%8].Frequency, nil
}

func getChannelNumber(frequency, dataRate int) (int, error) {
	for chanNum, channel := range UplinkChannelConfiguration {
		if frequency == channel.Frequency {
			for _, dr := range channel.DataRates {
				if dr == dataRate {
					return chanNum, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("lorawan/band: could not get channel number for frequency: %d, data rate: %d", frequency, dataRate)
}

// Default settings for this band
const (
	ReceiveDelay1    time.Duration = time.Second
	ReceiveDelay2    time.Duration = time.Second * 2
	JoinAcceptDelay1 time.Duration = time.Second * 5
	JoinAcceptDelay2 time.Duration = time.Second * 6
	MaxFCntGap       uint32        = 16384
	ADRAckLimit                    = 64
	ADRAckDelay                    = 32
	AckTimeoutMin    time.Duration = time.Second // AckTimeout = 2 +/- 1 (random value between 1 - 3)
	AckTimeoutMax    time.Duration = time.Second * 3
)
