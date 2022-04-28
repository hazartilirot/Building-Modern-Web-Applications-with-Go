package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}

func (f *Form) IsValid() bool {
	return len(f.Errors) == 0
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

/*Required checks each field to be filled in*/
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field is required")
		}
	}
}

/*MinLength checks a field our for a minimum length of characters*/
func (f *Form) MinLength(field string, minCharLength int, r *http.Request) bool {
	if len(r.Form.Get(field)) < minCharLength {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", minCharLength))
		return false
	}
	return true
}

func (f *Form) HasField(field string, r *http.Request) bool {
	res := r.Form.Get(field)
	if res == "" {
		return false
	}
	return true
}
