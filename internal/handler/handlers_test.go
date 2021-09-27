package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
	key, value string
}

var theTest =[]struct{
	name string
	url string
	method string
	params []postData
	expectedStatusCode int
}{
	{"home", "/","GET",[]postData{}, http.StatusOK},
	{"about", "/about","GET",[]postData{}, http.StatusOK},
	{"gq", "/generals-quarters","GET",[]postData{}, http.StatusOK},
	{"ms", "/majors-suite","GET",[]postData{}, http.StatusOK},
	{"sa", "/search-availability","GET",[]postData{}, http.StatusOK},
	{"contact", "/contact","GET",[]postData{}, http.StatusOK},
	{"mr", "/make-reservation","GET",[]postData{}, http.StatusOK},
	{"rs", "/reservation-summery","GET",[]postData{}, http.StatusOK},
}


func TestHandlers(t *testing.T) {

	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _,e := range theTest{
		if e.method == "GET" {
			resp, err:= ts.Client().Get(ts.URL+ e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {}
	}
}


