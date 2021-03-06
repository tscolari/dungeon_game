// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/tscolari/dungeon_game/dungeon"
	"github.com/tscolari/dungeon_game/solver"
)

type FakeStrategy struct {
	FindBestRouteStub        func(dungeon.Dungeon) (minHP int, bestRoute []string, err error)
	findBestRouteMutex       sync.RWMutex
	findBestRouteArgsForCall []struct {
		arg1 dungeon.Dungeon
	}
	findBestRouteReturns struct {
		result1 int
		result2 []string
		result3 error
	}
}

func (fake *FakeStrategy) FindBestRoute(arg1 dungeon.Dungeon) (minHP int, bestRoute []string, err error) {
	fake.findBestRouteMutex.Lock()
	fake.findBestRouteArgsForCall = append(fake.findBestRouteArgsForCall, struct {
		arg1 dungeon.Dungeon
	}{arg1})
	fake.findBestRouteMutex.Unlock()
	if fake.FindBestRouteStub != nil {
		return fake.FindBestRouteStub(arg1)
	} else {
		return fake.findBestRouteReturns.result1, fake.findBestRouteReturns.result2, fake.findBestRouteReturns.result3
	}
}

func (fake *FakeStrategy) FindBestRouteCallCount() int {
	fake.findBestRouteMutex.RLock()
	defer fake.findBestRouteMutex.RUnlock()
	return len(fake.findBestRouteArgsForCall)
}

func (fake *FakeStrategy) FindBestRouteArgsForCall(i int) dungeon.Dungeon {
	fake.findBestRouteMutex.RLock()
	defer fake.findBestRouteMutex.RUnlock()
	return fake.findBestRouteArgsForCall[i].arg1
}

func (fake *FakeStrategy) FindBestRouteReturns(result1 int, result2 []string, result3 error) {
	fake.FindBestRouteStub = nil
	fake.findBestRouteReturns = struct {
		result1 int
		result2 []string
		result3 error
	}{result1, result2, result3}
}

var _ solver.Strategy = new(FakeStrategy)
