package validators

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//RequiredField is the validator used for validate a specific field in the request
type RequiredField struct {
	FieldName string
	data      map[string]interface{}
}

//Validation verifies if the field required is present
func (rf *RequiredField) Validation(i interface{}) (err error) {
	req := i.(*http.Request)
	err = json.NewDecoder(req.Body).Decode(&rf.data)
	if err != nil && err != io.EOF {
		return
	}
	if rf.data[rf.FieldName] == nil {
		return fmt.Errorf("Missing param: %s", rf.FieldName)
	}
	return nil
}
