package container

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
