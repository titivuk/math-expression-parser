package object

const (
	RESULT_OBJ = "RESULT"
	ERROR_OBJ  = "ERROR"
)

type Object interface {
	Type() string
}

type Result struct {
	Value float64
}

func (r *Result) Type() string {
	return RESULT_OBJ
}

type Error struct {
	Message string
}

func (e *Error) Type() string {
	return ERROR_OBJ
}
