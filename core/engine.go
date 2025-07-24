package core

type Engine interface {
	Name() string
	Query(query string, page int, pageSize int) ([]Result, error)
	Check() bool
}
