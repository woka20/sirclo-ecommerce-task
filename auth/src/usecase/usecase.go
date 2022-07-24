package usecase

// UseCaseResult model
type UseCaseResult struct {
	Result interface{}
	Error  error
}

// AuthUseCase interface abstraction
type AuthUseCase interface {
	AuthUser(email, password string) <-chan UseCaseResult
}
