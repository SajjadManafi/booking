package handler

import (
	"booking/internal/models"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
	key, value string
}

var theTest = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
}{
	{"home", "/", "GET", http.StatusOK},
	{"about", "/about", "GET", http.StatusOK},
	{"gq", "/generals-quarters", "GET", http.StatusOK},
	{"ms", "/majors-suite", "GET", http.StatusOK},
	{"sa", "/search-availability", "GET", http.StatusOK},
	{"contact", "/contact", "GET", http.StatusOK},
	{"rs", "/reservation-summery", "GET", http.StatusOK},

	//{"post-search-avail", "/search-availability", "POST", []postData{
	//	{key: "start-date", value: "2020-01-01"},
	//	{key: "end-date", value: "2020-01-02"},
	//}, http.StatusOK},
	//{"post-search-avail-json", "/search-availability-json", "POST", []postData{
	//	{key: "start-date", value: "2020-01-01"},
	//	{key: "end-date", value: "2020-01-02"},
	//}, http.StatusOK},
	//{"post-make-reservation", "/make-reservation", "POST", []postData{
	//	{key: "first_name", value: "Harry"},
	//	{key: "last_name", value: "Potter"},
	//	{key: "phone_number", value: "50505-66-42"},
	//	{key: "email", value: "Harry@magic.com"},
	//}, http.StatusOK},
}

func TestHandlers(t *testing.T) {

	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTest {

		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}

	}
}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.Reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Reservation Handler returned wrong code: got %d, want %d", rr.Code, http.StatusOK)
	}

	// test case where reservation is not in session (rest everything)
	req, _ = http.NewRequest("GET", "make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation Handler returned wrong code: got %d, want %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test with none-exist room
	req, _ = http.NewRequest("GET", "make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	reservation.RoomID = 100
	session.Put(ctx, "reservation", reservation)
}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
