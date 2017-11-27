package main

type Driver struct {
	Stops       []int
	currentStop int
}

func NewDriver(stops ...int) Driver {

	driver := Driver{currentStop: 1}
	var finalStopSequence []int

	for len(finalStopSequence) < 480 {
		if len(finalStopSequence)+len(stops) > 480 {
			break
		}
		finalStopSequence = append(finalStopSequence, stops...)
	}

	if len(finalStopSequence) < 480 {
		diff := 480 - len(finalStopSequence)
		finalStopSequence = append(finalStopSequence, stops[:diff]...)
	}

	driver.Stops = finalStopSequence

	return driver
}

func (driver Driver) advance() {
	driver.currentStop += 1
}

func (driver Driver) getCurrentStop() int {
	return driver.Stops[driver.currentStop]
}