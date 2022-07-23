package query

type QueryResult struct {
	Result interface{}
	Error  error
}

// ProductQuery interface abstraction
type ProductQuery interface {
	FindByID(id int) <-chan QueryResult
	FindAll() <-chan QueryResult
}
