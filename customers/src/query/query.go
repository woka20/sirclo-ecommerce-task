package query

// QueryResult model
type QueryResult struct {
	Result interface{}
	Error  error
}

type CustomerQuery interface {
	FindByEmail(email string) <-chan QueryResult
}
