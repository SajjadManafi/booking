package helpers

import (
	"booking/internal/config"
	"fmt"
	"net/http"
	"runtime/debug"
)

var app *config.AppConfig

// NewHelpers setup and config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status code of", status)
	http.Error(w, http.StatusText(status), status)

}

func ServerError(w http.ResponseWriter, err error) {
	tracr := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(tracr)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func IsAuthenticated(r *http.Request) bool {
	exist := app.Session.Exists(r.Context(), "user_id")
	return exist
}
