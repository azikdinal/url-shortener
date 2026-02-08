package link

import (
	"strings"
)

func ParseShortCode(code string) int64 {
	var u uint32 = 0
	for i := 0; i < length; i++ {
		ch := code[i]
		idx := strings.IndexByte(alphabet, ch)
		u = u*uint32(base) + uint32(idx)
	}

	id := int64(u ^ uint32(seed))
	return id
}
