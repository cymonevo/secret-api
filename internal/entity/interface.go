package entity

type BaseModel interface {
	Call() (interface{}, error)
	Validate() error
}
