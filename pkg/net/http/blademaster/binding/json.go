package binding

import (
	"io/ioutil"
	"net/http"
)


type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (jsonBinding) Bind(req *http.Request, obj interface{}) error {
	content, _ := ioutil.ReadAll(req.Body)
	DefaultWithNoInput(content,obj)
	return validate(obj)
}
