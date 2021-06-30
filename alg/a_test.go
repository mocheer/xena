package alg_test

import (
	"testing"

	"github.com/mocheer/xena/alg"
)

func TestToWGS84(t *testing.T) {
	lonlats, err := alg.ProjInverse(2633458.580, 470452.902)
	if err != nil {
		t.Log(err)
	}
	t.Log(lonlats)
}
