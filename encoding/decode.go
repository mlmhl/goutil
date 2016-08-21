package encoding

type Decoder interface{
	Int64()
}

var (
	DefaultDecoder = &defaultDecoder{}
)

type defaultDecoder struct {}

func (defaultDecoder *defaultDecoder) Int64(buffer []byte) int64 {
	num := int64(0)
	for _, b := range(buffer) {
		num <<= 8;
		num |= int64(b)
	}
	return num
}