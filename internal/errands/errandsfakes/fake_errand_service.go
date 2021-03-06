// Code generated by counterfeiter. DO NOT EDIT.
package errandsfakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type FakeErrandService struct {
	ListStagedProductErrandsStub        func(productID string) (api.ErrandsListOutput, error)
	listStagedProductErrandsMutex       sync.RWMutex
	listStagedProductErrandsArgsForCall []struct {
		productID string
	}
	listStagedProductErrandsReturns struct {
		result1 api.ErrandsListOutput
		result2 error
	}
	listStagedProductErrandsReturnsOnCall map[int]struct {
		result1 api.ErrandsListOutput
		result2 error
	}
	UpdateStagedProductErrandsStub        func(productID string, errandName string, postDeployState interface{}, preDeleteState interface{}) error
	updateStagedProductErrandsMutex       sync.RWMutex
	updateStagedProductErrandsArgsForCall []struct {
		productID       string
		errandName      string
		postDeployState interface{}
		preDeleteState  interface{}
	}
	updateStagedProductErrandsReturns struct {
		result1 error
	}
	updateStagedProductErrandsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeErrandService) ListStagedProductErrands(productID string) (api.ErrandsListOutput, error) {
	fake.listStagedProductErrandsMutex.Lock()
	ret, specificReturn := fake.listStagedProductErrandsReturnsOnCall[len(fake.listStagedProductErrandsArgsForCall)]
	fake.listStagedProductErrandsArgsForCall = append(fake.listStagedProductErrandsArgsForCall, struct {
		productID string
	}{productID})
	fake.recordInvocation("ListStagedProductErrands", []interface{}{productID})
	fake.listStagedProductErrandsMutex.Unlock()
	if fake.ListStagedProductErrandsStub != nil {
		return fake.ListStagedProductErrandsStub(productID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listStagedProductErrandsReturns.result1, fake.listStagedProductErrandsReturns.result2
}

func (fake *FakeErrandService) ListStagedProductErrandsCallCount() int {
	fake.listStagedProductErrandsMutex.RLock()
	defer fake.listStagedProductErrandsMutex.RUnlock()
	return len(fake.listStagedProductErrandsArgsForCall)
}

func (fake *FakeErrandService) ListStagedProductErrandsArgsForCall(i int) string {
	fake.listStagedProductErrandsMutex.RLock()
	defer fake.listStagedProductErrandsMutex.RUnlock()
	return fake.listStagedProductErrandsArgsForCall[i].productID
}

func (fake *FakeErrandService) ListStagedProductErrandsReturns(result1 api.ErrandsListOutput, result2 error) {
	fake.ListStagedProductErrandsStub = nil
	fake.listStagedProductErrandsReturns = struct {
		result1 api.ErrandsListOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeErrandService) ListStagedProductErrandsReturnsOnCall(i int, result1 api.ErrandsListOutput, result2 error) {
	fake.ListStagedProductErrandsStub = nil
	if fake.listStagedProductErrandsReturnsOnCall == nil {
		fake.listStagedProductErrandsReturnsOnCall = make(map[int]struct {
			result1 api.ErrandsListOutput
			result2 error
		})
	}
	fake.listStagedProductErrandsReturnsOnCall[i] = struct {
		result1 api.ErrandsListOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeErrandService) UpdateStagedProductErrands(productID string, errandName string, postDeployState interface{}, preDeleteState interface{}) error {
	fake.updateStagedProductErrandsMutex.Lock()
	ret, specificReturn := fake.updateStagedProductErrandsReturnsOnCall[len(fake.updateStagedProductErrandsArgsForCall)]
	fake.updateStagedProductErrandsArgsForCall = append(fake.updateStagedProductErrandsArgsForCall, struct {
		productID       string
		errandName      string
		postDeployState interface{}
		preDeleteState  interface{}
	}{productID, errandName, postDeployState, preDeleteState})
	fake.recordInvocation("UpdateStagedProductErrands", []interface{}{productID, errandName, postDeployState, preDeleteState})
	fake.updateStagedProductErrandsMutex.Unlock()
	if fake.UpdateStagedProductErrandsStub != nil {
		return fake.UpdateStagedProductErrandsStub(productID, errandName, postDeployState, preDeleteState)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateStagedProductErrandsReturns.result1
}

func (fake *FakeErrandService) UpdateStagedProductErrandsCallCount() int {
	fake.updateStagedProductErrandsMutex.RLock()
	defer fake.updateStagedProductErrandsMutex.RUnlock()
	return len(fake.updateStagedProductErrandsArgsForCall)
}

func (fake *FakeErrandService) UpdateStagedProductErrandsArgsForCall(i int) (string, string, interface{}, interface{}) {
	fake.updateStagedProductErrandsMutex.RLock()
	defer fake.updateStagedProductErrandsMutex.RUnlock()
	return fake.updateStagedProductErrandsArgsForCall[i].productID, fake.updateStagedProductErrandsArgsForCall[i].errandName, fake.updateStagedProductErrandsArgsForCall[i].postDeployState, fake.updateStagedProductErrandsArgsForCall[i].preDeleteState
}

func (fake *FakeErrandService) UpdateStagedProductErrandsReturns(result1 error) {
	fake.UpdateStagedProductErrandsStub = nil
	fake.updateStagedProductErrandsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeErrandService) UpdateStagedProductErrandsReturnsOnCall(i int, result1 error) {
	fake.UpdateStagedProductErrandsStub = nil
	if fake.updateStagedProductErrandsReturnsOnCall == nil {
		fake.updateStagedProductErrandsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStagedProductErrandsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeErrandService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listStagedProductErrandsMutex.RLock()
	defer fake.listStagedProductErrandsMutex.RUnlock()
	fake.updateStagedProductErrandsMutex.RLock()
	defer fake.updateStagedProductErrandsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeErrandService) recordInvocation(key string, args []interface{}) {
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
