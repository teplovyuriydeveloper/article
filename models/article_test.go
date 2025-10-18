package models

import (
	"testing"

	"github.com/teplovyuriydeveloper/article/db"
)

func TestCreateArticle(t *testing.T) {
	db.Connect()

	a := &Article{
		Title:   "Test Article",
		Content: "Testing model.Create()",
		Author:  "Tester",
	}

	err := a.Create()
	if err != nil {
		t.Fatalf("failed to create article: %v", err)
	}

	if a.ID == "" {
		t.Fatal("expected article ID to be set")
	}
}
