// Copyright © 2015 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package dutycycle

import (
	"testing"

	testutil "github.com/edwindvinas/ttn/utils/testing"
)

// ----- CHECK utilities
func CheckSubBands(t *testing.T, want subBand, got subBand) {
	testutil.Check(t, want, got, "SubBands")
}

func CheckUsages(t *testing.T, want map[subBand]uint32, got map[subBand]uint32) {
	testutil.Check(t, want, got, "Usages")
}

func CheckConfiguration(t *testing.T, want *Configuration, got *Configuration) {
	testutil.Check(t, want, got, "Configurations")
}

func CheckStates(t *testing.T, want State, got State) {
	testutil.Check(t, want, got, "States")
}
