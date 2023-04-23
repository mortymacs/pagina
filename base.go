package pagina

// PageIter defines interface of the Pagina instance.
type PageIter[T any] interface {
	// Page actions.
	Next() (*Page[T], error)
	Previous() (*Page[T], error)
	Current() *Page[T]
	HasNext() bool
	HasPrevious() bool

	// Count action.
	Count() uint
	Remained() uint

	// State actions.
	IsLastPage() bool
	IsFirstPage() bool
}

// ItemIter defines interface of the Page instance.
type ItemIter[T any] interface {
	// Item actions.
	Next() (*T, error)
	Previous() (*T, error)
	Current() *T
	HasNext() bool
	HasPrevious() bool

	// Count actions.
	Count() uint
	Remained() uint

	// State actions.
	IsLastItem() bool
	IsFirstItem() bool
}
