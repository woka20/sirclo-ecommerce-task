package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"sirclo-ecommerce-task/auth/src/model"
	"sirclo-ecommerce-task/auth/src/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	// "sirclo-ecommerce-task/auth/src/token"
	// "sirclo-ecommerce-task/auth/src/usecase/mocks"
)

type AccessToken struct {
	Token string `json:"accessToken"`
}
type AuthUseCase struct {
	mock.Mock
}

func (_m *AuthUseCase) AuthUser(email string, password string) <-chan usecase.UseCaseResult {
	ret := _m.Called(email, password)

	var r0 <-chan usecase.UseCaseResult
	if rf, ok := ret.Get(0).(func(string, string) <-chan usecase.UseCaseResult); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan usecase.UseCaseResult)
		}
	}

	return r0
}
func generateIdentityAccessTokenResult() <-chan usecase.UseCaseResult {
	output := make(chan usecase.UseCaseResult)
	go func() {
		var accessToken AccessToken
		accessToken.Token = `Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJQV1QiLCJpYXQiOjE1MTIzMTg0MjMsImlzcyI6InB1cndva2VydG9kZXYuZ2l0aHViLmlvIiwic3ViIjoiMDAxIn0.GeEDHv82F8xp_98QQLiWxZ5aVBzZej0e-Ios8M9l0tYdrrTbdP3zxutSi5H7rxrd43PmlFi0pMMGtbVw64kPkBspCE3Kebbeersa8isn1zBejZO62mIgpRIGRhAJ_rphsxXYqOKKQlgj2ecI39dRR7IRJZdNYoTXXeBktUeUcDU`

		output <- usecase.UseCaseResult{Result: true}
	}()
	return output
}

// func generateIdentityAccessTokenResultError() <-chan usecase.UseCaseResult {
// 	output := make(chan usecase.UseCaseResult)
// 	go func() {
// 		output <- usecase.UseCaseResult{Error: errors.New("Error")}
// 	}()
// 	return output
// }

// func generateIdentityAccessTokenResultErrorNoResult() <-chan usecase.UseCaseResult {
// 	output := make(chan usecase.UseCaseResult)
// 	go func() {
// 		output <- usecase.UseCaseResult{Result: "no result"}
// 	}()
// 	return output
// }

// func generateQueryFindByEmailSuccessIdentityResult() <-chan query.QueryResult {
// 	output := make(chan query.QueryResult)
// 	go func() {
// 		var i model.Identity
// 		i.ID = "M1"
// 		i.Email = "wuriyanto48@yahoo.co.id"
// 		i.Password = "12345"
// 		output <- query.QueryResult{Result: i}
// 	}()
// 	return output
// }

func TestAuthHandler(t *testing.T) {
	_testing = true

	var i model.Identity
	i.ID = "M1"
	i.Email = "wuriyanto48@yahoo.co.id"
	i.Password = "12345"

	// t.Run("Test Password Credentials", func(t *testing.T) {

	mockUseCase := new(AuthUseCase)

	body, _ := json.Marshal(i)

	t.Log(string(body))
	mockUseCase.On("AuthUser", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(generateIdentityAccessTokenResult())
	req := httptest.NewRequest("POST", "/api/login", strings.NewReader(string(body)))
	// req.Header.Set("Content-Type", "application/json")
	// req.BasicAuth()
	rec := httptest.NewRecorder()

	//signKey, _ := getPrivateKey(PrivateKey)

	httpAuthHandler := NewHttpAuthHandler(mockUseCase)

	handler := http.Handler(httpAuthHandler.HandlerLogin())
	t.Log("SUDAH3")
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.

	handler.ServeHTTP(rec, req)
	// t.Log(handler)
	// Check the status code is what we expect.
	assert.Equal(t, http.StatusOK, rec.Code)

	// })

}
