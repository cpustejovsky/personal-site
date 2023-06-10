package handlers_test

import (
	"github.com/cpustejovsky/personal-site/handlers"
	"testing"
)

func TestGetResourcesPage(t *testing.T) {
	body, err := handlers.GetResourcesPage()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(body)
}
