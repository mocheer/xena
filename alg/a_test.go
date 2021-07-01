package alg_test

import (
	"testing"

	"github.com/mocheer/xena/alg/proj4"
)

func TestToWGS84(t *testing.T) {
	lonlats, err := proj4.Inverse(2633458.580, 470452.902)
	if err != nil {
		t.Log(err)
	}
	t.Log(lonlats)
}
