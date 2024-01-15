package main

import (
	"errors"
	"fmt"

	myerrors "github.com/lucas-code42/custom-errors/my-errors"
)

func main() {
	// Here, a custom error is created using the NewError function from the myerrors package.
	// The error has a status code of 400 and a message of "bad request."
	err := myerrors.NewMyErrors(400, "bad request")

	// This line uses fmt.Errorf to create a new error (errAgain) by formatting a string.
	// The %w verb is used to wrap the original error (err) within the new error.
	// The message includes both the original error's message and a new custom error
	// with a status code of 404 and a message of "not found."
	errAgain := fmt.Errorf("%w: %v", err, myerrors.NewMyErrors(404, "not found"))

	// output: 400: bad request: 404: not found
	fmt.Println(errAgain)

	// The errors.Is function is used to check if the error (errAgain) contains the wrapped error (err).
	// It returns true if errAgain is an instance of err or if err is one of the wrapped errors within errAgain.
	fmt.Println(errors.Is(errAgain, err))

	// The errors.Unwrap function is used to retrieve the original error that was wrapped.
	// In this case, it returns the original error (err), removing the additional wrapping
	fmt.Println(errors.Unwrap(errAgain))
}
