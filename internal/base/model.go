package base

type Model interface {
	Call() (interface{}, error)
	Validate() error
}
