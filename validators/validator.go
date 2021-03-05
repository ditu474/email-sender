package validators

//Validator is the interface that validators must implement
type Validator interface {
	Validation(i interface{}) (err error)
}
