package link

// Для хеширования ссылок используется LinkID и seed

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	base     = int64(len(alphabet))
	length   = 10
)

func encodeIDWithSeed(id int64, seedInt int64) string {
	value := id ^ seedInt

	buf := make([]byte, length)
	for i := length - 1; i >= 0; i-- {
		buf[i] = alphabet[value%base]
		value /= base
	}

	return string(buf)
}
