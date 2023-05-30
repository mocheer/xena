package provider

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/mocheer/pluto/pkg/fn"
	"github.com/mocheer/pluto/pkg/ts/ctp"
	"github.com/mocheer/xena/pkg/gm"
)

type Provider struct {
	URL        string
	Subdomains []string
	Vars       map[string]string
}

func (m Provider) GetTileURL(t *gm.Tile) string {
	data := map[string]any{"z": t.Z, "x": t.X, "y": t.Y, "s": m.GetRandSubdomains()}
	for k, v := range m.Vars {
		data[k] = v
	}
	return fn.FormatByMap(m.URL, data)
}

func (m Provider) LoadTile(t *gm.Tile) ([]byte, error) {
	url := m.GetTileURL(t)
	data, err := ctp.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("reptile '%s' error: %s", url, err))
	}
	return data, nil
}

func (m Provider) GetRandSubdomains() string {
	s := ""
	if m.Subdomains != nil {
		s = m.Subdomains[rand.Intn(len(m.Subdomains))]
	}
	return s
}
