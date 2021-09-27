package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever", r)
	if has {
		t.Error("form shows has field when it should not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)

	t.Log(`form.Has("a", r) returns`, form.Has("a", r))

	if !form.Has("a", r) {
		t.Error("shows form does not have field when it does")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10, r)
	if form.Valid() {
		t.Error("for shows min length exists for nonexistent field")
	}

	postedData := url.Values{}
	postedData.Add("some_field", "some value")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.MinLength("some_field", 100, r)
	if form.Valid() {
		t.Error("shows minlength of 100 when data is shorter")
	}

	postedData = url.Values{}
	postedData.Add("some_field", "some value")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.MinLength("some_field", 1, r)
	if !form.Valid() {
		t.Error("shows minlength is not at least 1 when it is")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.IsEmail("somefield")

	if form.Valid() {
		t.Error("form shows valid email for non existent field")
	}

	x := form.Errors.Get("somefield")
	if x == "" {
		t.Error("should have got an error from form")
	}

	postedData := url.Values{}
	postedData.Add("email", "me@here.com")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("shows invalid email when should not")
	}
}

func TestErrors_Get(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	x := form.Errors.Get("something")
	if x != "" {
		t.Error("got an error for form when we should not have")
	}

	form.Errors.Add("somefield", "some error")
	x = form.Errors.Get("somefield")
	if x == "" {
		t.Error("did not get an error when we should have")
	}
}
