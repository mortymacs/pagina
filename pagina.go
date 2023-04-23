package pagina

import (
	"errors"
	"math"
)

type Pagina[T any] struct {
	Items []T
	Size  uint

	// internal attributes.
	totalPages  uint
	currentPage uint
	resetFlag   bool
}

func New[T any](items []T, size uint) *Pagina[T] { // Iter
	return &Pagina[T]{
		Items:       items,
		Size:        size,
		resetFlag:   true,
		totalPages:  uint(math.Ceil(float64(len(items)) / float64(size))),
		currentPage: 0,
	}
}

// getRange returns range numbers.
func (p *Pagina[T]) getRange() (start, end uint) {
	if p.resetFlag {
		start = p.Size
	} else {
		start = (p.currentPage - 1) * p.Size
	}
	end = p.currentPage * p.Size
	return
}

// selectRange returns a slice of items.
func (p *Pagina[T]) selectRange(start, end uint) ([]*T, error) {
	// Validate range.
	if end > uint(len(p.Items)) {
		return nil, errors.New("out of range")
	}

	var selectedRange []*T
	for index, item := range p.Items {
		if uint(index) >= start && uint(index) <= end {
			selectedRange = append(selectedRange, &item)
		}
	}
	return selectedRange, nil
}

// Next returns next page.
func (p *Pagina[T]) Next() (*Page[T], error) {
	if p.Remained() == 0 {
		return nil, errors.New("next item not found")
	}

	// Move to next page.
	if p.resetFlag {
		p.resetFlag = false
	} else {
		p.currentPage = p.currentPage + 1
	}

	// Select range of items.
	items, err := p.selectRange(p.getRange())
	if err != nil {
		return nil, err
	}

	return newPage(items), nil
}

// Previous returns next page.
func (p *Pagina[T]) Previous() (*Page[T], error) {
	if p.IsFirstPage() {
		return nil, errors.New("previous item not found")
	}

	// Move to previous page.
	p.currentPage = p.currentPage - 1
	if p.currentPage == 0 {
		p.resetFlag = true
	}

	// Select range of items.
	items, err := p.selectRange(p.getRange())
	if err != nil {
		return nil, err
	}

	return newPage(items), nil
}

// Current returns the current page.
func (p *Pagina[T]) Current() *Page[T] {
	items, _ := p.selectRange(p.getRange())
	return newPage(items)
}

// HasNext defines you have a next page or not.
func (p *Pagina[T]) HasNext() bool {
	return !p.IsLastPage()
}

// HasPrevious defines you have a previous page or not.
func (p *Pagina[T]) HasPrevious() bool {
	return !p.IsFirstPage()
}

// Count returns count of pages.
func (p *Pagina[T]) Count() uint {
	return p.totalPages
}

// Remained returns remained count of pages.
func (p *Pagina[T]) Remained() uint {
	return p.totalPages - p.currentPage
}

// IsLastPage defines you're in the last page or not.
func (p *Pagina[T]) IsLastPage() bool {
	return p.Remained() == 0
}

// IsFirstPage defines you're in the first page or not.
func (p *Pagina[T]) IsFirstPage() bool {
	return p.currentPage == 0
}
