// Code generated by counterfeiter. DO NOT EDIT.
package applychangesfakes

import (
	"sync"

	"github.com/pivotal-cloudops/omen/internal/manifest"
)

type FakeManifestsLoader struct {
	LoadAllDeployedStub        func() (manifest.Manifests, error)
	loadAllDeployedMutex       sync.RWMutex
	loadAllDeployedArgsForCall []struct{}
	loadAllDeployedReturns     struct {
		result1 manifest.Manifests
		result2 error
	}
	loadAllDeployedReturnsOnCall map[int]struct {
		result1 manifest.Manifests
		result2 error
	}
	LoadAllStagedStub        func() (manifest.Manifests, error)
	loadAllStagedMutex       sync.RWMutex
	loadAllStagedArgsForCall []struct{}
	loadAllStagedReturns     struct {
		result1 manifest.Manifests
		result2 error
	}
	loadAllStagedReturnsOnCall map[int]struct {
		result1 manifest.Manifests
		result2 error
	}
	LoadDeployedStub        func(tileGuids []string) (manifest.Manifests, error)
	loadDeployedMutex       sync.RWMutex
	loadDeployedArgsForCall []struct {
		tileGuids []string
	}
	loadDeployedReturns struct {
		result1 manifest.Manifests
		result2 error
	}
	loadDeployedReturnsOnCall map[int]struct {
		result1 manifest.Manifests
		result2 error
	}
	LoadStagedStub        func(tileGuids []string) (manifest.Manifests, error)
	loadStagedMutex       sync.RWMutex
	loadStagedArgsForCall []struct {
		tileGuids []string
	}
	loadStagedReturns struct {
		result1 manifest.Manifests
		result2 error
	}
	loadStagedReturnsOnCall map[int]struct {
		result1 manifest.Manifests
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManifestsLoader) LoadAllDeployed() (manifest.Manifests, error) {
	fake.loadAllDeployedMutex.Lock()
	ret, specificReturn := fake.loadAllDeployedReturnsOnCall[len(fake.loadAllDeployedArgsForCall)]
	fake.loadAllDeployedArgsForCall = append(fake.loadAllDeployedArgsForCall, struct{}{})
	fake.recordInvocation("LoadAllDeployed", []interface{}{})
	fake.loadAllDeployedMutex.Unlock()
	if fake.LoadAllDeployedStub != nil {
		return fake.LoadAllDeployedStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.loadAllDeployedReturns.result1, fake.loadAllDeployedReturns.result2
}

func (fake *FakeManifestsLoader) LoadAllDeployedCallCount() int {
	fake.loadAllDeployedMutex.RLock()
	defer fake.loadAllDeployedMutex.RUnlock()
	return len(fake.loadAllDeployedArgsForCall)
}

func (fake *FakeManifestsLoader) LoadAllDeployedReturns(result1 manifest.Manifests, result2 error) {
	fake.LoadAllDeployedStub = nil
	fake.loadAllDeployedReturns = struct {
		result1 manifest.Manifests
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestsLoader) LoadAllDeployedReturnsOnCall(i int, result1 manifest.Manifests, result2 error) {
	fake.LoadAllDeployedStub = nil
	if fake.loadAllDeployedReturnsOnCall == nil {
		fake.loadAllDeployedReturnsOnCall = make(map[int]struct {
			result1 manifest.Manifests
			result2 error
		})
	}
	fake.loadAllDeployedReturnsOnCall[i] = struct {
		result1 manifest.Manifests
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestsLoader) LoadAllStaged() (manifest.Manifests, error) {
	fake.loadAllStagedMutex.Lock()
	ret, specificReturn := fake.loadAllStagedReturnsOnCall[len(fake.loadAllStagedArgsForCall)]
	fake.loadAllStagedArgsForCall = append(fake.loadAllStagedArgsForCall, struct{}{})
	fake.recordInvocation("LoadAllStaged", []interface{}{})
	fake.loadAllStagedMutex.Unlock()
	if fake.LoadAllStagedStub != nil {
		return fake.LoadAllStagedStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.loadAllStagedReturns.result1, fake.loadAllStagedReturns.result2
}

func (fake *FakeManifestsLoader) LoadAllStagedCallCount() int {
	fake.loadAllStagedMutex.RLock()
	defer fake.loadAllStagedMutex.RUnlock()
	return len(fake.loadAllStagedArgsForCall)
}

func (fake *FakeManifestsLoader) LoadAllStagedReturns(result1 manifest.Manifests, result2 error) {
	fake.LoadAllStagedStub = nil
	fake.loadAllStagedReturns = struct {
		result1 manifest.Manifests
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestsLoader) LoadAllStagedReturnsOnCall(i int, result1 manifest.Manifests, result2 error) {
	fake.LoadAllStagedStub = nil
	if fake.loadAllStagedReturnsOnCall == nil {
		fake.loadAllStagedReturnsOnCall = make(map[int]struct {
			result1 manifest.Manifests
			result2 error
		})
	}
	fake.loadAllStagedReturnsOnCall[i] = struct {
		result1 manifest.Manifests
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestsLoader) LoadDeployed(tileGuids []string) (manifest.Manifests, error) {
	var tileGuidsCopy []string
	if tileGuids != nil {
		tileGuidsCopy = make([]string, len(tileGuids))
		copy(tileGuidsCopy, tileGuids)
	}
	fake.loadDeployedMutex.Lock()
	ret, specificReturn := fake.loadDeployedReturnsOnCall[len(fake.loadDeployedArgsForCall)]
	fake.loadDeployedArgsForCall = append(fake.loadDeployedArgsForCall, struct {
		tileGuids []string
	}{tileGuidsCopy})
	fake.recordInvocation("LoadDeployed", []interface{}{tileGuidsCopy})
	fake.loadDeployedMutex.Unlock()
	if fake.LoadDeployedStub != nil {
		return fake.LoadDeployedStub(tileGuids)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.loadDeployedReturns.result1, fake.loadDeployedReturns.result2
}

func (fake *FakeManifestsLoader) LoadDeployedCallCount() int {
	fake.loadDeployedMutex.RLock()
	defer fake.loadDeployedMutex.RUnlock()
	return len(fake.loadDeployedArgsForCall)
}

func (fake *FakeManifestsLoader) LoadDeployedArgsForCall(i int) []string {
	fake.loadDeployedMutex.RLock()
	defer fake.loadDeployedMutex.RUnlock()
	return fake.loadDeployedArgsForCall[i].tileGuids
}

func (fake *FakeManifestsLoader) LoadDeployedReturns(result1 manifest.Manifests, result2 error) {
	fake.LoadDeployedStub = nil
	fake.loadDeployedReturns = struct {
		result1 manifest.Manifests
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestsLoader) LoadDeployedReturnsOnCall(i int, result1 manifest.Manifests, result2 error) {
	fake.LoadDeployedStub = nil
	if fake.loadDeployedReturnsOnCall == nil {
		fake.loadDeployedReturnsOnCall = make(map[int]struct {
			result1 manifest.Manifests
			result2 error
		})
	}
	fake.loadDeployedReturnsOnCall[i] = struct {
		result1 manifest.Manifests
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestsLoader) LoadStaged(tileGuids []string) (manifest.Manifests, error) {
	var tileGuidsCopy []string
	if tileGuids != nil {
		tileGuidsCopy = make([]string, len(tileGuids))
		copy(tileGuidsCopy, tileGuids)
	}
	fake.loadStagedMutex.Lock()
	ret, specificReturn := fake.loadStagedReturnsOnCall[len(fake.loadStagedArgsForCall)]
	fake.loadStagedArgsForCall = append(fake.loadStagedArgsForCall, struct {
		tileGuids []string
	}{tileGuidsCopy})
	fake.recordInvocation("LoadStaged", []interface{}{tileGuidsCopy})
	fake.loadStagedMutex.Unlock()
	if fake.LoadStagedStub != nil {
		return fake.LoadStagedStub(tileGuids)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.loadStagedReturns.result1, fake.loadStagedReturns.result2
}

func (fake *FakeManifestsLoader) LoadStagedCallCount() int {
	fake.loadStagedMutex.RLock()
	defer fake.loadStagedMutex.RUnlock()
	return len(fake.loadStagedArgsForCall)
}

func (fake *FakeManifestsLoader) LoadStagedArgsForCall(i int) []string {
	fake.loadStagedMutex.RLock()
	defer fake.loadStagedMutex.RUnlock()
	return fake.loadStagedArgsForCall[i].tileGuids
}

func (fake *FakeManifestsLoader) LoadStagedReturns(result1 manifest.Manifests, result2 error) {
	fake.LoadStagedStub = nil
	fake.loadStagedReturns = struct {
		result1 manifest.Manifests
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestsLoader) LoadStagedReturnsOnCall(i int, result1 manifest.Manifests, result2 error) {
	fake.LoadStagedStub = nil
	if fake.loadStagedReturnsOnCall == nil {
		fake.loadStagedReturnsOnCall = make(map[int]struct {
			result1 manifest.Manifests
			result2 error
		})
	}
	fake.loadStagedReturnsOnCall[i] = struct {
		result1 manifest.Manifests
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestsLoader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loadAllDeployedMutex.RLock()
	defer fake.loadAllDeployedMutex.RUnlock()
	fake.loadAllStagedMutex.RLock()
	defer fake.loadAllStagedMutex.RUnlock()
	fake.loadDeployedMutex.RLock()
	defer fake.loadDeployedMutex.RUnlock()
	fake.loadStagedMutex.RLock()
	defer fake.loadStagedMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManifestsLoader) recordInvocation(key string, args []interface{}) {
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
