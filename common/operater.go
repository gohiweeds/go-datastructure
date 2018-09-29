package common

type Operater interface {
	Add(value interface{}) (bool, error)
	Delete(value interface{}) (bool, error)
	Modify(value, update interface{}) (bool, error)
	Get(value interface{}) (bool, error)
	Len() uint64
}
