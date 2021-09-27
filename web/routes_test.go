package main

import (
	"booking/internal/config"
	"fmt"
	"github.com/go-chi/chi/v5"
	"testing"
)

func TestRoutes(t *testing.T) {

	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {

	case *chi.Mux:
		// do nothing test passe
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, type is %T", v))

	}

}
