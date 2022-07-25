package usecase

import (
	"errors"

	"github.com/woka20/sirclo-ecommerce-task/auth/src/model"
	"github.com/woka20/sirclo-ecommerce-task/auth/src/query"
)

type authUseCaseImpl struct {
	identityQuery query.IdentityQuery
	// accessTokenGenerator token.AccessTokenGenerator
}

// NewAuthUseCase function for initialise authUseCaseImpl model
func NewAuthUseCase(identityQuery query.IdentityQuery) AuthUseCase {
	return &authUseCaseImpl{
		identityQuery: identityQuery,
		// accessTokenGenerator: accessTokenGenerator,
	}
}

func (a *authUseCaseImpl) AuthUser(email, password string) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		defer close(output)

		identityResult := <-a.identityQuery.FindByEmail(email)

		if identityResult.Error != nil {
			output <- UseCaseResult{Error: identityResult.Error}
			return
		}

		identity, ok := identityResult.Result.(*model.Identity)

		if !ok {
			output <- UseCaseResult{Error: errors.New("Invalid Identity")}
			return
		}

		err := identity.IsValidPassword(password)

		if err != nil {
			output <- UseCaseResult{Error: err}
			return
		} else {
			output <- UseCaseResult{Result: true}

		}

	}()

	return output
}
