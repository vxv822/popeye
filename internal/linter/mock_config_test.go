// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/derailed/popeye/internal/linter (interfaces: Config)

package linter

import (
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	"time"
)

type MockConfig struct {
	fail func(message string, callerSkip ...int)
}

func NewMockConfig() *MockConfig {
	return &MockConfig{fail: pegomock.GlobalFailHandler}
}

func (mock *MockConfig) ActiveNamespace() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConfig().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ActiveNamespace", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockConfig) NodeCPULimit() float64 {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConfig().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("NodeCPULimit", params, []reflect.Type{reflect.TypeOf((*float64)(nil)).Elem()})
	var ret0 float64
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(float64)
		}
	}
	return ret0
}

func (mock *MockConfig) NodeMEMLimit() float64 {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConfig().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("NodeMEMLimit", params, []reflect.Type{reflect.TypeOf((*float64)(nil)).Elem()})
	var ret0 float64
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(float64)
		}
	}
	return ret0
}

func (mock *MockConfig) PodCPULimit() float64 {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConfig().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PodCPULimit", params, []reflect.Type{reflect.TypeOf((*float64)(nil)).Elem()})
	var ret0 float64
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(float64)
		}
	}
	return ret0
}

func (mock *MockConfig) PodMEMLimit() float64 {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConfig().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PodMEMLimit", params, []reflect.Type{reflect.TypeOf((*float64)(nil)).Elem()})
	var ret0 float64
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(float64)
		}
	}
	return ret0
}

func (mock *MockConfig) RestartsLimit() int {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockConfig().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("RestartsLimit", params, []reflect.Type{reflect.TypeOf((*int)(nil)).Elem()})
	var ret0 int
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(int)
		}
	}
	return ret0
}

func (mock *MockConfig) VerifyWasCalledOnce() *VerifierConfig {
	return &VerifierConfig{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockConfig) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierConfig {
	return &VerifierConfig{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockConfig) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierConfig {
	return &VerifierConfig{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockConfig) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierConfig {
	return &VerifierConfig{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierConfig struct {
	mock                   *MockConfig
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierConfig) ActiveNamespace() *Config_ActiveNamespace_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ActiveNamespace", params, verifier.timeout)
	return &Config_ActiveNamespace_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Config_ActiveNamespace_OngoingVerification struct {
	mock              *MockConfig
	methodInvocations []pegomock.MethodInvocation
}

func (c *Config_ActiveNamespace_OngoingVerification) GetCapturedArguments() {
}

func (c *Config_ActiveNamespace_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierConfig) NodeCPULimit() *Config_NodeCPULimit_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "NodeCPULimit", params, verifier.timeout)
	return &Config_NodeCPULimit_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Config_NodeCPULimit_OngoingVerification struct {
	mock              *MockConfig
	methodInvocations []pegomock.MethodInvocation
}

func (c *Config_NodeCPULimit_OngoingVerification) GetCapturedArguments() {
}

func (c *Config_NodeCPULimit_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierConfig) NodeMEMLimit() *Config_NodeMEMLimit_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "NodeMEMLimit", params, verifier.timeout)
	return &Config_NodeMEMLimit_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Config_NodeMEMLimit_OngoingVerification struct {
	mock              *MockConfig
	methodInvocations []pegomock.MethodInvocation
}

func (c *Config_NodeMEMLimit_OngoingVerification) GetCapturedArguments() {
}

func (c *Config_NodeMEMLimit_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierConfig) PodCPULimit() *Config_PodCPULimit_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PodCPULimit", params, verifier.timeout)
	return &Config_PodCPULimit_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Config_PodCPULimit_OngoingVerification struct {
	mock              *MockConfig
	methodInvocations []pegomock.MethodInvocation
}

func (c *Config_PodCPULimit_OngoingVerification) GetCapturedArguments() {
}

func (c *Config_PodCPULimit_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierConfig) PodMEMLimit() *Config_PodMEMLimit_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PodMEMLimit", params, verifier.timeout)
	return &Config_PodMEMLimit_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Config_PodMEMLimit_OngoingVerification struct {
	mock              *MockConfig
	methodInvocations []pegomock.MethodInvocation
}

func (c *Config_PodMEMLimit_OngoingVerification) GetCapturedArguments() {
}

func (c *Config_PodMEMLimit_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierConfig) RestartsLimit() *Config_RestartsLimit_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "RestartsLimit", params, verifier.timeout)
	return &Config_RestartsLimit_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Config_RestartsLimit_OngoingVerification struct {
	mock              *MockConfig
	methodInvocations []pegomock.MethodInvocation
}

func (c *Config_RestartsLimit_OngoingVerification) GetCapturedArguments() {
}

func (c *Config_RestartsLimit_OngoingVerification) GetAllCapturedArguments() {
}