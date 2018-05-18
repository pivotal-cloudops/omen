// Code generated by counterfeiter. DO NOT EDIT.
package applychangesfakes

import (
	"sync"
	"time"
)

type FakeOpsmanClient struct {
	PostStub        func(endpoint, data string, timeout time.Duration) ([]byte, error)
	postMutex       sync.RWMutex
	postArgsForCall []struct {
		endpoint string
		data     string
		timeout  time.Duration
	}
	postReturns struct {
		result1 []byte
		result2 error
	}
	postReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOpsmanClient) Post(endpoint string, data string, timeout time.Duration) ([]byte, error) {
	fake.postMutex.Lock()
	ret, specificReturn := fake.postReturnsOnCall[len(fake.postArgsForCall)]
	fake.postArgsForCall = append(fake.postArgsForCall, struct {
		endpoint string
		data     string
		timeout  time.Duration
	}{endpoint, data, timeout})
	fake.recordInvocation("Post", []interface{}{endpoint, data, timeout})
	fake.postMutex.Unlock()
	if fake.PostStub != nil {
		return fake.PostStub(endpoint, data, timeout)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.postReturns.result1, fake.postReturns.result2
}

func (fake *FakeOpsmanClient) PostCallCount() int {
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	return len(fake.postArgsForCall)
}

func (fake *FakeOpsmanClient) PostArgsForCall(i int) (string, string, time.Duration) {
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	return fake.postArgsForCall[i].endpoint, fake.postArgsForCall[i].data, fake.postArgsForCall[i].timeout
}

func (fake *FakeOpsmanClient) PostReturns(result1 []byte, result2 error) {
	fake.PostStub = nil
	fake.postReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsmanClient) PostReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.PostStub = nil
	if fake.postReturnsOnCall == nil {
		fake.postReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.postReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeOpsmanClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.postMutex.RLock()
	defer fake.postMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOpsmanClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
