package tests

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestApp(t *testing.T) {
	e := httpexpect.Default(t, "http://localhost:3000")

	e.GET("/").
		Expect().
		Status(http.StatusOK).Body().IsEqual("welcome")
}
