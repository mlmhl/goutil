package encoding

type Encoder interface {
	Int64() []byte
	Int() []byte
}

var (
	DefaultEncoder = &defaultEncoder{}
)

type defaultEncoder struct {}

func (defaultEncoder *defaultEncoder) Int64(number int64) []byte {
	mask := int64((1 << 8) - 1)
	buffer := make([]byte, 8)

	for i := 7; i >= 0; i-- {
		buffer[i] = byte(number & mask)
		number >>= 8
	}

	return buffer
}

func (defaultEncoder *defaultEncoder) Int(number int) []byte {
	return defaultEncoder.Int64(int64(number))
}