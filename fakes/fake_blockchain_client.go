// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	utils "github.com/lgsvl/data-marketplace-stream-delivery/utils"
)

type FakeBlockchainClient struct {
	CheckContractIDStub        func(string, map[string]string) (bool, error)
	checkContractIDMutex       sync.RWMutex
	checkContractIDArgsForCall []struct {
		arg1 string
		arg2 map[string]string
	}
	checkContractIDReturns struct {
		result1 bool
		result2 error
	}
	checkContractIDReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlockchainClient) CheckContractID(arg1 string, arg2 map[string]string) (bool, error) {
	fake.checkContractIDMutex.Lock()
	ret, specificReturn := fake.checkContractIDReturnsOnCall[len(fake.checkContractIDArgsForCall)]
	fake.checkContractIDArgsForCall = append(fake.checkContractIDArgsForCall, struct {
		arg1 string
		arg2 map[string]string
	}{arg1, arg2})
	fake.recordInvocation("CheckContractID", []interface{}{arg1, arg2})
	fake.checkContractIDMutex.Unlock()
	if fake.CheckContractIDStub != nil {
		return fake.CheckContractIDStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.checkContractIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBlockchainClient) CheckContractIDCallCount() int {
	fake.checkContractIDMutex.RLock()
	defer fake.checkContractIDMutex.RUnlock()
	return len(fake.checkContractIDArgsForCall)
}

func (fake *FakeBlockchainClient) CheckContractIDCalls(stub func(string, map[string]string) (bool, error)) {
	fake.checkContractIDMutex.Lock()
	defer fake.checkContractIDMutex.Unlock()
	fake.CheckContractIDStub = stub
}

func (fake *FakeBlockchainClient) CheckContractIDArgsForCall(i int) (string, map[string]string) {
	fake.checkContractIDMutex.RLock()
	defer fake.checkContractIDMutex.RUnlock()
	argsForCall := fake.checkContractIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBlockchainClient) CheckContractIDReturns(result1 bool, result2 error) {
	fake.checkContractIDMutex.Lock()
	defer fake.checkContractIDMutex.Unlock()
	fake.CheckContractIDStub = nil
	fake.checkContractIDReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBlockchainClient) CheckContractIDReturnsOnCall(i int, result1 bool, result2 error) {
	fake.checkContractIDMutex.Lock()
	defer fake.checkContractIDMutex.Unlock()
	fake.CheckContractIDStub = nil
	if fake.checkContractIDReturnsOnCall == nil {
		fake.checkContractIDReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.checkContractIDReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBlockchainClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkContractIDMutex.RLock()
	defer fake.checkContractIDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBlockchainClient) recordInvocation(key string, args []interface{}) {
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

var _ utils.BlockchainClient = new(FakeBlockchainClient)
