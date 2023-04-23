package pagina

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Post struct {
	Title string
}

var posts = []*Post{
	{Title: "Test1"},
	{Title: "Test2"},
	{Title: "Test3"},
	{Title: "Test4"},
	{Title: "Test5"},
}
var pageSize uint = 2

func TestPaginaNew(t *testing.T) {
	p := New(posts, 2)

	assert.Equal(t, p.totalPages, uint(3))
	assert.Equal(t, p.currentPage, uint(0))
}

func TestPaginaNext(t *testing.T) {
	p := New(posts, 2)

	// Next page: Test1, Test2
	pl, err := p.Next()
	assert.NoError(t, err)
	t.Log(p.Items)
	assert.Equal(t, uint(2), pl.Count())
}
