package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdvanceOnePosition(t *testing.T) {
	assert := assert.New(t);
	driver := []int{3, 1, 2, 3}
	nextStop := advance(driver)

	assert.Equal(nextStop, 1)
}

func TestTwoDriversInDifferentStopDontShareGossip(t *testing.T) {
	assert := assert.New(t);
	driver1 := NewDriver(1)
	driver2 := NewDriver(2)

	gossipsShared := shareGossips(driver1, driver2, 1)

	assert.False(gossipsShared)
}

func TestTwoDriversInSameStopShareGossip(t *testing.T) {
	assert := assert.New(t);
	driver1 := NewDriver(1)
	driver2 := NewDriver(1)

	gossipsShared := shareGossips(driver1, driver2, 1)

	assert.True(gossipsShared)
}

func TestDriverHave480StopsInitialy(t *testing.T) {
	assert := assert.New(t);
	driver1 := NewDriver(1, 2)

	assert.True(len(driver1.Stops) == 480)

}

func shareGossips(firstDriver *Driver, secondDriver *Driver, stop int) bool {

	return firstDriver.Stops[stop] == secondDriver.Stops[stop]
}

func advance(ints []int) int {
	return 1
}
