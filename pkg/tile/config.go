package tile

import (
	"math/rand/v2"

	"github.com/mocheer/xena/pkg/gm"
)

type LoadConfig struct {
	URL        string
	DirName    string
	MinZoom    int
	MaxZoom    int
	Bbox       gm.BBox
	Origin     string
	Subdomains []string
	SavePath   string
	FiberCount int
}

func (m LoadConfig) GetRandSubdomains() string {
	s := ""
	if m.Subdomains != nil {
		s = m.Subdomains[rand.IntN(len(m.Subdomains))]
	}
	return s
}
