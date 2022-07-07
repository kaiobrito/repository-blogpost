package repository

type NotFound struct {

}

func (NotFound) Error() string {
	return "Object not found"
}