package rehgo

type rehgoError error

// Usage:  defer this function with references to your named return values, and
// it will bubble-up an error caused by a Q() panic (a.k.a. a rehgoError),
// passthrough any other panics, and do nothing if there was no panic.
// This means that any function which contains this deffered at the top can
// wrap any other function call that returns a (T, error) in the Q() function
// to get a Rust-like unwrapped result in successful cases and a bubbled-up
// error when the wrapped function call returns an error.
func Reh[T any](res *T, errout *error) {
	if err := recover(); err != nil {
		if e, ok := err.(rehgoError); ok {
			var myRes T
			*res = myRes
			*errout = e.(error)
		} else {
			panic(err)
		}
	}
}

// Usage: wrap this around any function that returns a (T, error) in order to
// return only the value on success. Failures will cause a panic, which stops
// execution and calls the deferred reh() function, which will set the return
// values for the parent function to the error, bubbling the error up without
// wasting cycles finishing the parent function.
func Q[T any](v T, err error) T {
	if err != nil {
		panic(err.(rehgoError))
	}
	return v
}
