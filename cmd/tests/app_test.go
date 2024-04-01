package tests

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestApp(t *testing.T) {
	e := httpexpect.Default(t, fmt.Sprintf("http://localhost:%s", os.Getenv("APP_PORT")))

	e.GET("/").
		Expect().
		Status(http.StatusOK).Body().IsEqual("welcome")
}
