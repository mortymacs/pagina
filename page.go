package pagina

import (
	"errors"
)

// Page defines the page data structure.
type Page[T any] struct {
	// Internal attributes.
	items       []*T
	currentItem uint
	resetFlag   bool
}

// newPage returns a new instance of Page.
func newPage[T any](items []*T) *Page[T] {
	return &Page[T]{
		items:       items,
		currentItem: 0,
		resetFlag:   true,
	}
}

// Next returns next item of the page.
func (p *Page[T]) Next() (*T, error) {
	if p.Remained() == 0 {
		return nil, errors.New("next item not found")
	}

	if p.resetFlag {
		p.resetFlag = false
	} else {
		p.currentItem = p.currentItem + 1
	}
	item := p.items[p.currentItem]
	return item, nil
}

// Previous returns previous item of the page.
func (p *Page[T]) Previous() (*T, error) {
	if p.currentItem == 0 {
		return nil, errors.New("previous item not found")
	}

	p.currentItem = p.currentItem - 1
	if p.currentItem == 0 {
		p.resetFlag = true
	}
	return p.items[p.currentItem], nil
}

// Current returns current item of the page.
func (p *Page[T]) Current() *T {
	return p.items[p.currentItem]
}

// HasNext defines you have a next item or not..
func (p *Page[T]) HasNext() bool {
	return !p.IsLastItem()
}

// HasPrevious defines you have a previous item or not.
func (p *Page[T]) HasPrevious() bool {
	return !p.IsFirstItem()
}

// Count returns count of items.
func (p *Page[T]) Count() uint {
	return uint(len(p.items))
}

// Count returns count of items.
func (p *Page[T]) Remained() uint {
	if p.resetFlag {
		return p.Count()
	}
	return p.Count() - p.currentItem - 1
}

// IsLastItem defines you're in the last page or not.
func (p *Page[T]) IsLastItem() bool {
	return p.currentItem == (p.Count() - 1)
}

// IsFirstPage defines you're in the first page or not.
func (p *Page[T]) IsFirstItem() bool {
	return p.currentItem == 0
}
