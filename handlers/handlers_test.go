package handlers_test

import (
	"testing"

	"github.com/cpustejovsky/personal-site/handlers"
)

func TestGetResourcesPage(t *testing.T) {
	_, err := handlers.GetResourcesPage()
	if err != nil {
		t.Fatal(err)
	}
}
