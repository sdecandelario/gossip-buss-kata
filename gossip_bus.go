package main

type Driver struct {
	Stops []int
}

func NewDriver(stops ...int) *Driver {

	driver := new(Driver)
	var finalStopSequence []int

	for len(finalStopSequence) < 480 {
		if len(finalStopSequence) + len(stops) > 480{
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
