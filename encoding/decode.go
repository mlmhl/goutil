package encoding

type Decoder interface{
	Int64()
	Int()
}

var (
	DefaultDecoder = &defaultDecoder{}
)

type defaultDecoder struct {}

func (defaultDecoder *defaultDecoder) Int64(buffer []byte) (int64, int) {
	num := int64(0)
	for i := 0; i < 8; i++ {
		num <<= 8;
		num |= int64(buffer[i])
	}
	return num, 8
}

// return the decoded value and remained buffer
func (defaultEncoder *defaultDecoder) Int(buffer []byte) (int, int) {
	 val, size := defaultEncoder.Int64(buffer)
	return int64(val), size
}