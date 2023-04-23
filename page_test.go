package pagina

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	Name string
}

var users = []*User{
	{Name: "Test1"},
	{Name: "Test2"},
	{Name: "Test3"},
}

func TestPageNext(t *testing.T) {
	p := newPage(users)

	// First item
	i, err := p.Next()
	assert.NoError(t, err)
	assert.Equal(t, "Test1", i.Name)

	// Second
	i, err = p.Next()
	assert.NoError(t, err)
	assert.Equal(t, "Test2", i.Name)

	// Third
	i, err = p.Next()
	assert.NoError(t, err)
	assert.Equal(t, "Test3", i.Name)

	// Invalid
	i, err = p.Next()
	assert.Error(t, err)
	assert.Nil(t, i)
}

func TestPagePrevious(t *testing.T) {
	p := newPage(users)

	// First item
	i, err := p.Next()
	assert.NoError(t, err)
	assert.Equal(t, "Test1", i.Name)

	// Second
	i, err = p.Next()
	assert.NoError(t, err)
	assert.Equal(t, "Test2", i.Name)

	// First item.
	i, err = p.Previous()
	assert.NoError(t, err)
	assert.Equal(t, "Test1", i.Name)

	// Invalid
	i, err = p.Previous()
	assert.Error(t, err)
	assert.Nil(t, i)
}

func TestPageCurrent(t *testing.T) {
	p := newPage(users)

	// First item
	p.Next()

	// Second
	p.Next()

	// Current
	i := p.Current()
	assert.Equal(t, "Test2", i.Name)
}

func TestPageHasNext(t *testing.T) {
	p := newPage(users)

	assert.True(t, p.HasNext())

	// First item ("Test1")
	p.Next()
	assert.True(t, p.HasNext())

	// Next item ("Test2")
	p.Next()
	assert.True(t, p.HasNext())

	// Next item ("Test3")
	p.Next()
	assert.False(t, p.HasNext())
}

func TestPageHasPrevious(t *testing.T) {
	p := newPage(users)

	assert.False(t, p.HasPrevious())

	// First item ("Test1")
	p.Next()
	assert.False(t, p.HasPrevious())

	// Next item ("Test2")
	p.Next()
	assert.True(t, p.HasPrevious())

	// Previous item ("Test1")
	p.Previous()
	assert.False(t, p.HasPrevious())
}

func TestPageCount(t *testing.T) {
	p := newPage(users)
	assert.Equal(t, p.Count(), uint(3))
}

func TestRemained(t *testing.T) {
	p := newPage(users)
	assert.Equal(t, p.Remained(), uint(3))

	// Next item ("Test1")
	p.Next()
	assert.Equal(t, p.Remained(), uint(2))

	// Next item ("Test2")
	p.Next()
	assert.Equal(t, p.Remained(), uint(1))

	// Next item ("Test3")
	p.Next()
	assert.Equal(t, p.Remained(), uint(0))
}

func TestPageIsLastItem(t *testing.T) {
	p := newPage(users)
	assert.False(t, p.IsLastItem())

	// Next item ("Test1")
	p.Next()
	assert.False(t, p.IsLastItem())

	// Next item ("Test2")
	p.Next()
	assert.False(t, p.IsLastItem())

	// Next item ("Test3")
	p.Next()
	assert.True(t, p.IsLastItem())

	// Previous item ("Test2")
	p.Previous()
	assert.False(t, p.IsLastItem())

	// Previous item ("Test3")
	p.Next()
	assert.True(t, p.IsLastItem())
}

func TestPageIsFirstItem(t *testing.T) {
	p := newPage(users)
	assert.True(t, p.IsFirstItem())

	// Next item ("Test1")
	p.Next()
	assert.True(t, p.IsFirstItem())

	// Next item ("Test2")
	p.Next()
	assert.False(t, p.IsFirstItem())

	// Next item ("Test3")
	p.Next()
	assert.False(t, p.IsFirstItem())

	// Previous items ("Test2", "Test1")
	p.Previous()
	p.Previous()
	assert.True(t, p.IsFirstItem())
}
