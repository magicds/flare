package home

import (
	"strings"
	"testing"

	"github.com/soulteary/flare/config/model"
	"github.com/stretchr/testify/assert"
)

func TestRenderBookmarksWithoutCategories_NewTab_ContainsNoreferrer(t *testing.T) {
	var b strings.Builder
	bookmarks := []model.Bookmark{
		{Name: "Test", URL: "https://example.com", Icon: "https://example.com/icon.svg"},
	}

	renderBookmarksWithoutCategories(&b, &bookmarks, true, false, "")

	result := b.String()
	assert.Contains(t, result, `rel="noopener noreferrer"`)
}

func TestRenderBookmarksWithoutCategories_SameTab_ContainsNoreferrer(t *testing.T) {
	var b strings.Builder
	bookmarks := []model.Bookmark{
		{Name: "Test", URL: "https://example.com", Icon: "https://example.com/icon.svg"},
	}

	renderBookmarksWithoutCategories(&b, &bookmarks, false, false, "")

	result := b.String()
	assert.Contains(t, result, `rel="noopener noreferrer"`)
}

func TestRenderBookmarksWithCategories_NewTab_ContainsNoreferrer(t *testing.T) {
	var b strings.Builder
	bookmarks := []model.Bookmark{
		{Name: "Test", URL: "https://example.com", Icon: "https://example.com/icon.svg", Category: "cat1"},
	}
	category := model.Category{ID: "cat1", Name: "Category 1"}
	defaultCategory := model.Category{ID: "cat1", Name: "Category 1"}

	renderBookmarksWithCategories(&b, &bookmarks, &category, &defaultCategory, true, false, "")

	result := b.String()
	assert.Contains(t, result, `rel="noopener noreferrer"`)
}

func TestRenderBookmarksWithCategories_SameTab_ContainsNoreferrer(t *testing.T) {
	var b strings.Builder
	bookmarks := []model.Bookmark{
		{Name: "Test", URL: "https://example.com", Icon: "https://example.com/icon.svg", Category: "cat1"},
	}
	category := model.Category{ID: "cat1", Name: "Category 1"}
	defaultCategory := model.Category{ID: "cat1", Name: "Category 1"}

	renderBookmarksWithCategories(&b, &bookmarks, &category, &defaultCategory, false, false, "")

	result := b.String()
	assert.Contains(t, result, `rel="noopener noreferrer"`)
}
