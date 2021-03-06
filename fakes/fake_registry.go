//
// Copyright (c) 2019 LG Electronics Inc.
// SPDX-License-Identifier: Apache-2.0
//

// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	context "context"
	sync "sync"

	resources "github.com/lgsvl/data-marketplace-stream-delivery/resources"
)

type FakeRegistry struct {
	AddChannelStub        func(context.Context, string, chan string) error
	addChannelMutex       sync.RWMutex
	addChannelArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 chan string
	}
	addChannelReturns struct {
		result1 error
	}
	addChannelReturnsOnCall map[int]struct {
		result1 error
	}
	AddStreamStub        func(context.Context, resources.DataStream) error
	addStreamMutex       sync.RWMutex
	addStreamArgsForCall []struct {
		arg1 context.Context
		arg2 resources.DataStream
	}
	addStreamReturns struct {
		result1 error
	}
	addStreamReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteChannelStub        func(context.Context, string) error
	deleteChannelMutex       sync.RWMutex
	deleteChannelArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	deleteChannelReturns struct {
		result1 error
	}
	deleteChannelReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStreamStub        func(context.Context, string) error
	deleteStreamMutex       sync.RWMutex
	deleteStreamArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	deleteStreamReturns struct {
		result1 error
	}
	deleteStreamReturnsOnCall map[int]struct {
		result1 error
	}
	GetChannelStub        func(context.Context, string) (chan string, error)
	getChannelMutex       sync.RWMutex
	getChannelArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getChannelReturns struct {
		result1 chan string
		result2 error
	}
	getChannelReturnsOnCall map[int]struct {
		result1 chan string
		result2 error
	}
	GetStreamStub        func(context.Context, string) (resources.DataStream, error)
	getStreamMutex       sync.RWMutex
	getStreamArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getStreamReturns struct {
		result1 resources.DataStream
		result2 error
	}
	getStreamReturnsOnCall map[int]struct {
		result1 resources.DataStream
		result2 error
	}
	HasStreamStub        func(context.Context, string) bool
	hasStreamMutex       sync.RWMutex
	hasStreamArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	hasStreamReturns struct {
		result1 bool
	}
	hasStreamReturnsOnCall map[int]struct {
		result1 bool
	}
	ListStreamsStub        func() []resources.DataStream
	listStreamsMutex       sync.RWMutex
	listStreamsArgsForCall []struct {
	}
	listStreamsReturns struct {
		result1 []resources.DataStream
	}
	listStreamsReturnsOnCall map[int]struct {
		result1 []resources.DataStream
	}
	UpdateStreamStub        func(context.Context, resources.DataStream) error
	updateStreamMutex       sync.RWMutex
	updateStreamArgsForCall []struct {
		arg1 context.Context
		arg2 resources.DataStream
	}
	updateStreamReturns struct {
		result1 error
	}
	updateStreamReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRegistry) AddChannel(arg1 context.Context, arg2 string, arg3 chan string) error {
	fake.addChannelMutex.Lock()
	ret, specificReturn := fake.addChannelReturnsOnCall[len(fake.addChannelArgsForCall)]
	fake.addChannelArgsForCall = append(fake.addChannelArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 chan string
	}{arg1, arg2, arg3})
	fake.recordInvocation("AddChannel", []interface{}{arg1, arg2, arg3})
	fake.addChannelMutex.Unlock()
	if fake.AddChannelStub != nil {
		return fake.AddChannelStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addChannelReturns
	return fakeReturns.result1
}

func (fake *FakeRegistry) AddChannelCallCount() int {
	fake.addChannelMutex.RLock()
	defer fake.addChannelMutex.RUnlock()
	return len(fake.addChannelArgsForCall)
}

func (fake *FakeRegistry) AddChannelCalls(stub func(context.Context, string, chan string) error) {
	fake.addChannelMutex.Lock()
	defer fake.addChannelMutex.Unlock()
	fake.AddChannelStub = stub
}

func (fake *FakeRegistry) AddChannelArgsForCall(i int) (context.Context, string, chan string) {
	fake.addChannelMutex.RLock()
	defer fake.addChannelMutex.RUnlock()
	argsForCall := fake.addChannelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRegistry) AddChannelReturns(result1 error) {
	fake.addChannelMutex.Lock()
	defer fake.addChannelMutex.Unlock()
	fake.AddChannelStub = nil
	fake.addChannelReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry) AddChannelReturnsOnCall(i int, result1 error) {
	fake.addChannelMutex.Lock()
	defer fake.addChannelMutex.Unlock()
	fake.AddChannelStub = nil
	if fake.addChannelReturnsOnCall == nil {
		fake.addChannelReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addChannelReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry) AddStream(arg1 context.Context, arg2 resources.DataStream) error {
	fake.addStreamMutex.Lock()
	ret, specificReturn := fake.addStreamReturnsOnCall[len(fake.addStreamArgsForCall)]
	fake.addStreamArgsForCall = append(fake.addStreamArgsForCall, struct {
		arg1 context.Context
		arg2 resources.DataStream
	}{arg1, arg2})
	fake.recordInvocation("AddStream", []interface{}{arg1, arg2})
	fake.addStreamMutex.Unlock()
	if fake.AddStreamStub != nil {
		return fake.AddStreamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addStreamReturns
	return fakeReturns.result1
}

func (fake *FakeRegistry) AddStreamCallCount() int {
	fake.addStreamMutex.RLock()
	defer fake.addStreamMutex.RUnlock()
	return len(fake.addStreamArgsForCall)
}

func (fake *FakeRegistry) AddStreamCalls(stub func(context.Context, resources.DataStream) error) {
	fake.addStreamMutex.Lock()
	defer fake.addStreamMutex.Unlock()
	fake.AddStreamStub = stub
}

func (fake *FakeRegistry) AddStreamArgsForCall(i int) (context.Context, resources.DataStream) {
	fake.addStreamMutex.RLock()
	defer fake.addStreamMutex.RUnlock()
	argsForCall := fake.addStreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRegistry) AddStreamReturns(result1 error) {
	fake.addStreamMutex.Lock()
	defer fake.addStreamMutex.Unlock()
	fake.AddStreamStub = nil
	fake.addStreamReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry) AddStreamReturnsOnCall(i int, result1 error) {
	fake.addStreamMutex.Lock()
	defer fake.addStreamMutex.Unlock()
	fake.AddStreamStub = nil
	if fake.addStreamReturnsOnCall == nil {
		fake.addStreamReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addStreamReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry) DeleteChannel(arg1 context.Context, arg2 string) error {
	fake.deleteChannelMutex.Lock()
	ret, specificReturn := fake.deleteChannelReturnsOnCall[len(fake.deleteChannelArgsForCall)]
	fake.deleteChannelArgsForCall = append(fake.deleteChannelArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DeleteChannel", []interface{}{arg1, arg2})
	fake.deleteChannelMutex.Unlock()
	if fake.DeleteChannelStub != nil {
		return fake.DeleteChannelStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteChannelReturns
	return fakeReturns.result1
}

func (fake *FakeRegistry) DeleteChannelCallCount() int {
	fake.deleteChannelMutex.RLock()
	defer fake.deleteChannelMutex.RUnlock()
	return len(fake.deleteChannelArgsForCall)
}

func (fake *FakeRegistry) DeleteChannelCalls(stub func(context.Context, string) error) {
	fake.deleteChannelMutex.Lock()
	defer fake.deleteChannelMutex.Unlock()
	fake.DeleteChannelStub = stub
}

func (fake *FakeRegistry) DeleteChannelArgsForCall(i int) (context.Context, string) {
	fake.deleteChannelMutex.RLock()
	defer fake.deleteChannelMutex.RUnlock()
	argsForCall := fake.deleteChannelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRegistry) DeleteChannelReturns(result1 error) {
	fake.deleteChannelMutex.Lock()
	defer fake.deleteChannelMutex.Unlock()
	fake.DeleteChannelStub = nil
	fake.deleteChannelReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry) DeleteChannelReturnsOnCall(i int, result1 error) {
	fake.deleteChannelMutex.Lock()
	defer fake.deleteChannelMutex.Unlock()
	fake.DeleteChannelStub = nil
	if fake.deleteChannelReturnsOnCall == nil {
		fake.deleteChannelReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteChannelReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry) DeleteStream(arg1 context.Context, arg2 string) error {
	fake.deleteStreamMutex.Lock()
	ret, specificReturn := fake.deleteStreamReturnsOnCall[len(fake.deleteStreamArgsForCall)]
	fake.deleteStreamArgsForCall = append(fake.deleteStreamArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DeleteStream", []interface{}{arg1, arg2})
	fake.deleteStreamMutex.Unlock()
	if fake.DeleteStreamStub != nil {
		return fake.DeleteStreamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteStreamReturns
	return fakeReturns.result1
}

func (fake *FakeRegistry) DeleteStreamCallCount() int {
	fake.deleteStreamMutex.RLock()
	defer fake.deleteStreamMutex.RUnlock()
	return len(fake.deleteStreamArgsForCall)
}

func (fake *FakeRegistry) DeleteStreamCalls(stub func(context.Context, string) error) {
	fake.deleteStreamMutex.Lock()
	defer fake.deleteStreamMutex.Unlock()
	fake.DeleteStreamStub = stub
}

func (fake *FakeRegistry) DeleteStreamArgsForCall(i int) (context.Context, string) {
	fake.deleteStreamMutex.RLock()
	defer fake.deleteStreamMutex.RUnlock()
	argsForCall := fake.deleteStreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRegistry) DeleteStreamReturns(result1 error) {
	fake.deleteStreamMutex.Lock()
	defer fake.deleteStreamMutex.Unlock()
	fake.DeleteStreamStub = nil
	fake.deleteStreamReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry) DeleteStreamReturnsOnCall(i int, result1 error) {
	fake.deleteStreamMutex.Lock()
	defer fake.deleteStreamMutex.Unlock()
	fake.DeleteStreamStub = nil
	if fake.deleteStreamReturnsOnCall == nil {
		fake.deleteStreamReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteStreamReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry) GetChannel(arg1 context.Context, arg2 string) (chan string, error) {
	fake.getChannelMutex.Lock()
	ret, specificReturn := fake.getChannelReturnsOnCall[len(fake.getChannelArgsForCall)]
	fake.getChannelArgsForCall = append(fake.getChannelArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetChannel", []interface{}{arg1, arg2})
	fake.getChannelMutex.Unlock()
	if fake.GetChannelStub != nil {
		return fake.GetChannelStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getChannelReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRegistry) GetChannelCallCount() int {
	fake.getChannelMutex.RLock()
	defer fake.getChannelMutex.RUnlock()
	return len(fake.getChannelArgsForCall)
}

func (fake *FakeRegistry) GetChannelCalls(stub func(context.Context, string) (chan string, error)) {
	fake.getChannelMutex.Lock()
	defer fake.getChannelMutex.Unlock()
	fake.GetChannelStub = stub
}

func (fake *FakeRegistry) GetChannelArgsForCall(i int) (context.Context, string) {
	fake.getChannelMutex.RLock()
	defer fake.getChannelMutex.RUnlock()
	argsForCall := fake.getChannelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRegistry) GetChannelReturns(result1 chan string, result2 error) {
	fake.getChannelMutex.Lock()
	defer fake.getChannelMutex.Unlock()
	fake.GetChannelStub = nil
	fake.getChannelReturns = struct {
		result1 chan string
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistry) GetChannelReturnsOnCall(i int, result1 chan string, result2 error) {
	fake.getChannelMutex.Lock()
	defer fake.getChannelMutex.Unlock()
	fake.GetChannelStub = nil
	if fake.getChannelReturnsOnCall == nil {
		fake.getChannelReturnsOnCall = make(map[int]struct {
			result1 chan string
			result2 error
		})
	}
	fake.getChannelReturnsOnCall[i] = struct {
		result1 chan string
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistry) GetStream(arg1 context.Context, arg2 string) (resources.DataStream, error) {
	fake.getStreamMutex.Lock()
	ret, specificReturn := fake.getStreamReturnsOnCall[len(fake.getStreamArgsForCall)]
	fake.getStreamArgsForCall = append(fake.getStreamArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetStream", []interface{}{arg1, arg2})
	fake.getStreamMutex.Unlock()
	if fake.GetStreamStub != nil {
		return fake.GetStreamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStreamReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRegistry) GetStreamCallCount() int {
	fake.getStreamMutex.RLock()
	defer fake.getStreamMutex.RUnlock()
	return len(fake.getStreamArgsForCall)
}

func (fake *FakeRegistry) GetStreamCalls(stub func(context.Context, string) (resources.DataStream, error)) {
	fake.getStreamMutex.Lock()
	defer fake.getStreamMutex.Unlock()
	fake.GetStreamStub = stub
}

func (fake *FakeRegistry) GetStreamArgsForCall(i int) (context.Context, string) {
	fake.getStreamMutex.RLock()
	defer fake.getStreamMutex.RUnlock()
	argsForCall := fake.getStreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRegistry) GetStreamReturns(result1 resources.DataStream, result2 error) {
	fake.getStreamMutex.Lock()
	defer fake.getStreamMutex.Unlock()
	fake.GetStreamStub = nil
	fake.getStreamReturns = struct {
		result1 resources.DataStream
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistry) GetStreamReturnsOnCall(i int, result1 resources.DataStream, result2 error) {
	fake.getStreamMutex.Lock()
	defer fake.getStreamMutex.Unlock()
	fake.GetStreamStub = nil
	if fake.getStreamReturnsOnCall == nil {
		fake.getStreamReturnsOnCall = make(map[int]struct {
			result1 resources.DataStream
			result2 error
		})
	}
	fake.getStreamReturnsOnCall[i] = struct {
		result1 resources.DataStream
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistry) HasStream(arg1 context.Context, arg2 string) bool {
	fake.hasStreamMutex.Lock()
	ret, specificReturn := fake.hasStreamReturnsOnCall[len(fake.hasStreamArgsForCall)]
	fake.hasStreamArgsForCall = append(fake.hasStreamArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("HasStream", []interface{}{arg1, arg2})
	fake.hasStreamMutex.Unlock()
	if fake.HasStreamStub != nil {
		return fake.HasStreamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.hasStreamReturns
	return fakeReturns.result1
}

func (fake *FakeRegistry) HasStreamCallCount() int {
	fake.hasStreamMutex.RLock()
	defer fake.hasStreamMutex.RUnlock()
	return len(fake.hasStreamArgsForCall)
}

func (fake *FakeRegistry) HasStreamCalls(stub func(context.Context, string) bool) {
	fake.hasStreamMutex.Lock()
	defer fake.hasStreamMutex.Unlock()
	fake.HasStreamStub = stub
}

func (fake *FakeRegistry) HasStreamArgsForCall(i int) (context.Context, string) {
	fake.hasStreamMutex.RLock()
	defer fake.hasStreamMutex.RUnlock()
	argsForCall := fake.hasStreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRegistry) HasStreamReturns(result1 bool) {
	fake.hasStreamMutex.Lock()
	defer fake.hasStreamMutex.Unlock()
	fake.HasStreamStub = nil
	fake.hasStreamReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRegistry) HasStreamReturnsOnCall(i int, result1 bool) {
	fake.hasStreamMutex.Lock()
	defer fake.hasStreamMutex.Unlock()
	fake.HasStreamStub = nil
	if fake.hasStreamReturnsOnCall == nil {
		fake.hasStreamReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.hasStreamReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRegistry) ListStreams() []resources.DataStream {
	fake.listStreamsMutex.Lock()
	ret, specificReturn := fake.listStreamsReturnsOnCall[len(fake.listStreamsArgsForCall)]
	fake.listStreamsArgsForCall = append(fake.listStreamsArgsForCall, struct {
	}{})
	fake.recordInvocation("ListStreams", []interface{}{})
	fake.listStreamsMutex.Unlock()
	if fake.ListStreamsStub != nil {
		return fake.ListStreamsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.listStreamsReturns
	return fakeReturns.result1
}

func (fake *FakeRegistry) ListStreamsCallCount() int {
	fake.listStreamsMutex.RLock()
	defer fake.listStreamsMutex.RUnlock()
	return len(fake.listStreamsArgsForCall)
}

func (fake *FakeRegistry) ListStreamsCalls(stub func() []resources.DataStream) {
	fake.listStreamsMutex.Lock()
	defer fake.listStreamsMutex.Unlock()
	fake.ListStreamsStub = stub
}

func (fake *FakeRegistry) ListStreamsReturns(result1 []resources.DataStream) {
	fake.listStreamsMutex.Lock()
	defer fake.listStreamsMutex.Unlock()
	fake.ListStreamsStub = nil
	fake.listStreamsReturns = struct {
		result1 []resources.DataStream
	}{result1}
}

func (fake *FakeRegistry) ListStreamsReturnsOnCall(i int, result1 []resources.DataStream) {
	fake.listStreamsMutex.Lock()
	defer fake.listStreamsMutex.Unlock()
	fake.ListStreamsStub = nil
	if fake.listStreamsReturnsOnCall == nil {
		fake.listStreamsReturnsOnCall = make(map[int]struct {
			result1 []resources.DataStream
		})
	}
	fake.listStreamsReturnsOnCall[i] = struct {
		result1 []resources.DataStream
	}{result1}
}

func (fake *FakeRegistry) UpdateStream(arg1 context.Context, arg2 resources.DataStream) error {
	fake.updateStreamMutex.Lock()
	ret, specificReturn := fake.updateStreamReturnsOnCall[len(fake.updateStreamArgsForCall)]
	fake.updateStreamArgsForCall = append(fake.updateStreamArgsForCall, struct {
		arg1 context.Context
		arg2 resources.DataStream
	}{arg1, arg2})
	fake.recordInvocation("UpdateStream", []interface{}{arg1, arg2})
	fake.updateStreamMutex.Unlock()
	if fake.UpdateStreamStub != nil {
		return fake.UpdateStreamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateStreamReturns
	return fakeReturns.result1
}

func (fake *FakeRegistry) UpdateStreamCallCount() int {
	fake.updateStreamMutex.RLock()
	defer fake.updateStreamMutex.RUnlock()
	return len(fake.updateStreamArgsForCall)
}

func (fake *FakeRegistry) UpdateStreamCalls(stub func(context.Context, resources.DataStream) error) {
	fake.updateStreamMutex.Lock()
	defer fake.updateStreamMutex.Unlock()
	fake.UpdateStreamStub = stub
}

func (fake *FakeRegistry) UpdateStreamArgsForCall(i int) (context.Context, resources.DataStream) {
	fake.updateStreamMutex.RLock()
	defer fake.updateStreamMutex.RUnlock()
	argsForCall := fake.updateStreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRegistry) UpdateStreamReturns(result1 error) {
	fake.updateStreamMutex.Lock()
	defer fake.updateStreamMutex.Unlock()
	fake.UpdateStreamStub = nil
	fake.updateStreamReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry) UpdateStreamReturnsOnCall(i int, result1 error) {
	fake.updateStreamMutex.Lock()
	defer fake.updateStreamMutex.Unlock()
	fake.UpdateStreamStub = nil
	if fake.updateStreamReturnsOnCall == nil {
		fake.updateStreamReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStreamReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addChannelMutex.RLock()
	defer fake.addChannelMutex.RUnlock()
	fake.addStreamMutex.RLock()
	defer fake.addStreamMutex.RUnlock()
	fake.deleteChannelMutex.RLock()
	defer fake.deleteChannelMutex.RUnlock()
	fake.deleteStreamMutex.RLock()
	defer fake.deleteStreamMutex.RUnlock()
	fake.getChannelMutex.RLock()
	defer fake.getChannelMutex.RUnlock()
	fake.getStreamMutex.RLock()
	defer fake.getStreamMutex.RUnlock()
	fake.hasStreamMutex.RLock()
	defer fake.hasStreamMutex.RUnlock()
	fake.listStreamsMutex.RLock()
	defer fake.listStreamsMutex.RUnlock()
	fake.updateStreamMutex.RLock()
	defer fake.updateStreamMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRegistry) recordInvocation(key string, args []interface{}) {
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

var _ resources.Registry = new(FakeRegistry)
