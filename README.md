# rehgo

A simple implementation of error-wrapping similar to Rust that allows you to handle errors without having to do the `if err != nil` check all the damned time.

The name stands for Rust-like Error Handling in GO.

The Reh() function must be deferred at the top of any function utilizing this error handling, and the utilizing function's return values must be (T, error). The Q() function wraps any function which also returns (T,error) and returns the T if there is no error. If there is an error, execution stops due to a panic, but the Reh() function recovers the panic and returns the error for the utilizing function. It does this by being passed references to the named return values of the utilizing function.

I'm sure there are more robust ways of doing this, but this is my quick-and-dirty method of reducing my error-handling pain in the most common case.

Example:

```
import (
    "strconv"
    r "github.com/maxweisspoker/rehgo"
)

func makeANumberNoError() (int, error) {
    return 4, nil
}

func makeANumberWithError() (int, error) {
    return 6, errors.New("An error")
}

func myFunction() (myStr string, err error) {
    defer r.Reh(&myStr, &err)

    myStr = "My string"

    var intOne int = r.Q(makeANumberNoError())
    // intOne now equals 4

    var intTwo int = r.Q(makeANumberWithError())
    // intTwo is not assigned. myFunction() returns with myStr equaling the
    // default initialization value of "" and err being errors.New("An error")

    // This code is never reached
    var intThree = 3
    return "Success " + strconv.Itoa(intOne + intTwo + intThree), nil
}
```
