package handlers_test

import (
	"os"
	"testing"

	"github.com/cpustejovsky/personal-site/handlers"
)

func TestGetResourcesPage(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	wd += "/static/resources.html"
	_, err = handlers.GetResourcesPage(wd)
	if err != nil {
		t.Fatal(err)
	}
}
