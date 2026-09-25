package gm_alg_test

import (
	"testing"

	"github.com/mocheer/xena/pkg/gm"
	gm_alg "github.com/mocheer/xena/pkg/metry_alg"
	"github.com/stretchr/testify/assert"
)

func TestGetZoom(t *testing.T) {
	z := gm_alg.GetZoom(gm.BBox{
		115.49377441406251,
		24.422143781858985,
		120.45410156250001,
		28.01380137638074,
	}, [2]float64{903, 729})
	assert.Equal(t, 8.0, z)
}

func TestGetZoom2(t *testing.T) {
	z := gm_alg.GetZoom(gm.BBox{
		117.47955322265626,
		25.456914906486638,
		118.71963500976564,
		26.35742006833118,
	}, [2]float64{903, 729})
	assert.Equal(t, 10.0, z)
}
