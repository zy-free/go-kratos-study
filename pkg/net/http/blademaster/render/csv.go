package render

import (
	//"fmt"
	"github.com/pkg/errors"
	"net/http"
)

var csvContentType = []string{"application/csv; charset=utf-8"}

// CSV common bytes struct.
type CSV struct {
	Content []byte
	Title   string
}

// Render (Data) writes data with custom ContentType.
func (r CSV) Render(w http.ResponseWriter) (err error) {
	r.WriteContentType(w)

	if _, err = w.Write(r.Content); err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// WriteContentType writes data with custom ContentType.
func (r CSV) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, csvContentType)
}
