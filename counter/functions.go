package counter

func IncrementCounter(start int) func() int {
	counter := start
	return func() int {
		counter++
		return counter
	}
}
