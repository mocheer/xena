package save

import "github.com/mocheer/xena/pkg/tileset/transform/graph"

type SaveOptions struct {
	OutPath  string
	AsBinary bool
}

// Save
func Save(options *SaveOptions) func(g *graph.Graph) error {
	return func(g *graph.Graph) error {
		if options.AsBinary {
			return g.SaveBinary(options.OutPath)
		} else {
			return g.Save(options.OutPath)
		}
	}
}
