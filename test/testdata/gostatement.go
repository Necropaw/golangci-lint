package testdata

func foogofunc() {
	go func() { // ERROR "go statement found"

	}()

	go bar() // ERROR "go statement found"
}
