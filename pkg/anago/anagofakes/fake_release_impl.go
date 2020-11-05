/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package anagofakes

import (
	"sync"
)

type FakeReleaseImpl struct {
	PrepareWorkspaceReleaseStub        func(string, string, string) error
	prepareWorkspaceReleaseMutex       sync.RWMutex
	prepareWorkspaceReleaseArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	prepareWorkspaceReleaseReturns struct {
		result1 error
	}
	prepareWorkspaceReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReleaseImpl) PrepareWorkspaceRelease(arg1 string, arg2 string, arg3 string) error {
	fake.prepareWorkspaceReleaseMutex.Lock()
	ret, specificReturn := fake.prepareWorkspaceReleaseReturnsOnCall[len(fake.prepareWorkspaceReleaseArgsForCall)]
	fake.prepareWorkspaceReleaseArgsForCall = append(fake.prepareWorkspaceReleaseArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.PrepareWorkspaceReleaseStub
	fakeReturns := fake.prepareWorkspaceReleaseReturns
	fake.recordInvocation("PrepareWorkspaceRelease", []interface{}{arg1, arg2, arg3})
	fake.prepareWorkspaceReleaseMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseCallCount() int {
	fake.prepareWorkspaceReleaseMutex.RLock()
	defer fake.prepareWorkspaceReleaseMutex.RUnlock()
	return len(fake.prepareWorkspaceReleaseArgsForCall)
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseCalls(stub func(string, string, string) error) {
	fake.prepareWorkspaceReleaseMutex.Lock()
	defer fake.prepareWorkspaceReleaseMutex.Unlock()
	fake.PrepareWorkspaceReleaseStub = stub
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseArgsForCall(i int) (string, string, string) {
	fake.prepareWorkspaceReleaseMutex.RLock()
	defer fake.prepareWorkspaceReleaseMutex.RUnlock()
	argsForCall := fake.prepareWorkspaceReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseReturns(result1 error) {
	fake.prepareWorkspaceReleaseMutex.Lock()
	defer fake.prepareWorkspaceReleaseMutex.Unlock()
	fake.PrepareWorkspaceReleaseStub = nil
	fake.prepareWorkspaceReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) PrepareWorkspaceReleaseReturnsOnCall(i int, result1 error) {
	fake.prepareWorkspaceReleaseMutex.Lock()
	defer fake.prepareWorkspaceReleaseMutex.Unlock()
	fake.PrepareWorkspaceReleaseStub = nil
	if fake.prepareWorkspaceReleaseReturnsOnCall == nil {
		fake.prepareWorkspaceReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.prepareWorkspaceReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.prepareWorkspaceReleaseMutex.RLock()
	defer fake.prepareWorkspaceReleaseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReleaseImpl) recordInvocation(key string, args []interface{}) {
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