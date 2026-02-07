package link

import (
	"hash/fnv"
)

type ShortCodeGenerator struct {
	seed string
}

func NewShortCodeGenerator(seed string) *ShortCodeGenerator {
	return &ShortCodeGenerator{seed: seed}
}

func (g *ShortCodeGenerator) Generate(id int64) string {
	var seedInt = seedToInt64(g.seed)
	return string(encodeIDWithSeed(id, seedInt))
}

func seedToInt64(seed string) int64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(seed))
	return int64(h.Sum64())
}
