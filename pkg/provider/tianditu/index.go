package tianditu

import (
	"github.com/mocheer/xena/pkg/provider"
)

type TiandituMap struct {
	provider.Provider
}

func New(typeName string) {

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
