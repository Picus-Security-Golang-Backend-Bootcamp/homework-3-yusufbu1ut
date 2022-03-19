package helper

import "errors"

var (
	ListErr   = errors.New("Expected List arg > 'list', 'list' 'a' or 'list' 'b'")
	SearchErr = errors.New("Expeted Search arg > 'search' 'some arg/args for search'")
	BuyErr    = errors.New("Expected Buy args > 'buy' 'uint' 'uint'")
	DeleteErr = errors.New("Expected Delete arg> 'delete' 'uint'")
)
