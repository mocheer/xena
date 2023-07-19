package tianditu

import (
	"github.com/mocheer/pluto/pkg/ts/ctp"
	"github.com/mocheer/xena/pkg/gm"
	"github.com/mocheer/xena/pkg/provider"
)

type TiandituMap struct {
	provider.Provider
}

func New(typeName string) {

}

func (m *TiandituMap) LoadTile(t *gm.Tile) ([]byte, error) {
	if m.Loader == nil {
		m.Loader = ctp.New()
		m.Loader.SetProxies(
			"http://1.15.156.141:7890",
		)
	}
	return m.Provider.LoadTile(t)
}

func NewNormal() *TiandituMap {
	return &TiandituMap{
		Provider: provider.Provider{
			URL:        URL,
			Subdomains: subdomains,
			Vars: map[string]string{
				"v": EPSG3857_TYPES["Normal"],
			},
		},
	}
}

func NewSatellite() *TiandituMap {
	return &TiandituMap{
		Provider: provider.Provider{
			URL:        URL,
			Subdomains: subdomains,
			Vars: map[string]string{
				"v": EPSG3857_TYPES["Satellite"],
			},
		},
	}
}
func NewSatelliteA() *TiandituMap {
	return &TiandituMap{
		Provider: provider.Provider{
			URL:        URL,
			Subdomains: subdomains,
			Vars: map[string]string{
				"v": EPSG3857_TYPES["Satellite_A"],
			},
		},
	}
}
