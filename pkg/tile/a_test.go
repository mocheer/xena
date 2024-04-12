package tile_test

import (
	"fmt"
	"testing"

	"github.com/mocheer/xena/pkg/gm"
	"github.com/mocheer/xena/pkg/tile"
)

var china = []float64{109, 20.2, 122, 40}

func TestXxx(t *testing.T) {
	z := 8
	startX, startY, endX, endY := tile.GetTileLimitByBbox(4326, z, gm.BBox(china))
	fmt.Println(startX, startY, endX, endY)
	fmt.Println((endX - startX) * (endY - startY))
}
