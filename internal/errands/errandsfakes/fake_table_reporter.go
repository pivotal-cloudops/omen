// Code generated by counterfeiter. DO NOT EDIT.
package errandsfakes

import (
	"sync"
)

type FakeTableReporter struct {
	WriteStub        func([]byte) (int, error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 []byte
	}
	writeReturns struct {
		result1 int
		result2 error
	}
	writeReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	FlushStub        func() error
	flushMutex       sync.RWMutex
	flushArgsForCall []struct{}
	flushReturns     struct {
		result1 error
	}
	flushReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTableReporter) Write(arg1 []byte) (int, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("Write", []interface{}{arg1Copy})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.writeReturns.result1, fake.writeReturns.result2
}

func (fake *FakeTableReporter) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeTableReporter) WriteArgsForCall(i int) []byte {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].arg1
}

func (fake *FakeTableReporter) WriteReturns(result1 int, result2 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeTableReporter) WriteReturnsOnCall(i int, result1 int, result2 error) {
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeTableReporter) Flush() error {
	fake.flushMutex.Lock()
	ret, specificReturn := fake.flushReturnsOnCall[len(fake.flushArgsForCall)]
	fake.flushArgsForCall = append(fake.flushArgsForCall, struct{}{})
	fake.recordInvocation("Flush", []interface{}{})
	fake.flushMutex.Unlock()
	if fake.FlushStub != nil {
		return fake.FlushStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.flushReturns.result1
}

func (fake *FakeTableReporter) FlushCallCount() int {
	fake.flushMutex.RLock()
	defer fake.flushMutex.RUnlock()
	return len(fake.flushArgsForCall)
}

func (fake *FakeTableReporter) FlushReturns(result1 error) {
	fake.FlushStub = nil
	fake.flushReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTableReporter) FlushReturnsOnCall(i int, result1 error) {
	fake.FlushStub = nil
	if fake.flushReturnsOnCall == nil {
		fake.flushReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.flushReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTableReporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.flushMutex.RLock()
	defer fake.flushMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTableReporter) recordInvocation(key string, args []interface{}) {
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
