package prune

import "github.com/mocheer/xena/pkg/tileset/transform/graph"

// DisposeCounter 用于销毁计数
type DisposeCounter struct {
	disposed map[string]int
}

func NewDisposeCounter() *DisposeCounter {
	return &DisposeCounter{disposed: make(map[string]int)}
}

func (c *DisposeCounter) Empty() bool {
	return len(c.disposed) == 0
}

func (c *DisposeCounter) Entries() []struct {
	Type  string
	Count int
} {
	entries := make([]struct {
		Type  string
		Count int
	}, 0, len(c.disposed))
	for k, v := range c.disposed {
		entries = append(entries, struct {
			Type  string
			Count int
		}{k, v})
	}
	return entries
}

func (c *DisposeCounter) Dispose(prop graph.GraphBase) {
	// c.disposed[prop]++
}
