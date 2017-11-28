package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdvanceTwoPositions(t *testing.T) {
	assert := assert.New(t);
	driver := NewDriver(3, 1, 2, 3)
	driver.advance()
	driver.advance()

	assert.Equal(driver.getCurrentStop(), 2)
}

func TestAdvanceFourPositions(t *testing.T) {
	assert := assert.New(t);
	driver := NewDriver(3, 1, 2, 3)
	driver.advance()
	driver.advance()
	driver.advance()
	driver.advance()

	assert.Equal(driver.getCurrentStop(), 3)
}

func TestTwoDriversInDifferentStopNotShareGossip(t *testing.T) {
	assert := assert.New(t);
	driver1 := NewDriver(1)
	driver2 := NewDriver(2)

	gossipsShared := driver1.ShareGossips(driver2, 1)

	assert.False(gossipsShared)
}

func TestTwoDriversInSameStopShareGossip(t *testing.T) {
	assert := assert.New(t);
	driver1 := NewDriver(1)
	driver2 := NewDriver(1)

	gossipsShared := driver1.ShareGossips(driver2, 1)

	assert.True(gossipsShared)
}

func TestDriverHave480StopsInitialy(t *testing.T) {
	assert := assert.New(t);
	driver1 := NewDriver(1, 2)

	assert.Equal(len(driver1.Stops), 480)
}
