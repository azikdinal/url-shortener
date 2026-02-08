package link

// Для хеширования ссылок используется LinkID и seed

const (
	alphabet        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	base            = int64(len(alphabet))
	length          = 10
	seed     uint32 = 2549583949
)

func GenerateShortCode(id int64) string {
	u := uint32(id) ^ seed

	buf := make([]byte, length)
	for i := length - 1; i >= 0; i-- {
		idx := int(u % uint32(base))
		buf[i] = alphabet[idx]
		u /= uint32(base)
	}

	return string(buf)
}
