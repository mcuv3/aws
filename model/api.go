package model

type ErrorResponse struct {
	Status  int
	Message string
}

type SuccessResponse struct { 
	Message string
}
