package forms

import (
	"net/http"
	"net/url"
	"strings"
)

type Forms struct {
	url.Values
	Errors errors
}

func (f *Forms) Valid() bool {
	return len(f.Errors) == 0
}
func New(data url.Values) *Forms {
	return &Forms{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Forms) Required(r *http.Request, field ...string) {
	for _, v := range field {
		value := f.Get(v)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(v, "This form cannot be Blank")
		}

	}
}
func (f *Forms) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x != "" {
		return true
	}
	f.Errors.Add(field, "This field cannot be blank")
	return false
}
