package main

import "errors"

var (
	ErrNotNumber        = errors.New("Data is not numeric")
	ErrInvalidColumn    = errors.New("Invalid column number")
	ErrNoFiles          = errors.New("No input files provided")
	ErrInvalidOperation = errors.New("Invalid operation")
)
