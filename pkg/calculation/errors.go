package calculation

import "errors"

var (
	ErrInvalidExpression  = errors.New("invalid expression")
	ErrDivisionByZero     = errors.New("division by zero")
	ErrFailedToReadInput  = errors.New("failed to read expression from console")
	ErrInvalidRequestBody = errors.New("invalid request body")
)
