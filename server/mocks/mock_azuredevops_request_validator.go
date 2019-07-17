// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server (interfaces: AzureDevopsRequestValidator)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	http "net/http"
	"reflect"
	"time"
)

type MockAzureDevopsRequestValidator struct {
	fail func(message string, callerSkip ...int)
}

func NewMockAzureDevopsRequestValidator(options ...pegomock.Option) *MockAzureDevopsRequestValidator {
	mock := &MockAzureDevopsRequestValidator{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockAzureDevopsRequestValidator) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockAzureDevopsRequestValidator) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockAzureDevopsRequestValidator) Validate(r *http.Request, user []byte, pass []byte) ([]byte, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockAzureDevopsRequestValidator().")
	}
	params := []pegomock.Param{r, user, pass}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Validate", params, []reflect.Type{reflect.TypeOf((*[]byte)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []byte
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]byte)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockAzureDevopsRequestValidator) VerifyWasCalledOnce() *VerifierMockAzureDevopsRequestValidator {
	return &VerifierMockAzureDevopsRequestValidator{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockAzureDevopsRequestValidator) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockAzureDevopsRequestValidator {
	return &VerifierMockAzureDevopsRequestValidator{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockAzureDevopsRequestValidator) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockAzureDevopsRequestValidator {
	return &VerifierMockAzureDevopsRequestValidator{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockAzureDevopsRequestValidator) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockAzureDevopsRequestValidator {
	return &VerifierMockAzureDevopsRequestValidator{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockAzureDevopsRequestValidator struct {
	mock                   *MockAzureDevopsRequestValidator
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockAzureDevopsRequestValidator) Validate(r *http.Request, user []byte, pass []byte) *MockAzureDevopsRequestValidator_Validate_OngoingVerification {
	params := []pegomock.Param{r, user, pass}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Validate", params, verifier.timeout)
	return &MockAzureDevopsRequestValidator_Validate_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockAzureDevopsRequestValidator_Validate_OngoingVerification struct {
	mock              *MockAzureDevopsRequestValidator
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockAzureDevopsRequestValidator_Validate_OngoingVerification) GetCapturedArguments() (*http.Request, []byte, []byte) {
	r, user, pass := c.GetAllCapturedArguments()
	return r[len(r)-1], user[len(user)-1], pass[len(pass)-1]
}

func (c *MockAzureDevopsRequestValidator_Validate_OngoingVerification) GetAllCapturedArguments() (_param0 []*http.Request, _param1 [][]byte, _param2 [][]byte) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*http.Request, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*http.Request)
		}
		_param1 = make([][]byte, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.([]byte)
		}
		_param2 = make([][]byte, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.([]byte)
		}
	}
	return
}
