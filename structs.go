package main

//payRequest request object
type payRequest struct {
	APIKey       string
	Amount       int64
	Redirect     string
	FactorNumber string
}

//payResponse object response
type payResponse struct {
	Status       int64  `json:"status"`
	TransID      int64  `json:"transId"`
	ErrorCode    int64  `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

//callbackResponse handles pay.ir callback post data
type callbackResponse struct {
	Status       int64
	TransID      int64
	FactorNumber string
	Message      string
}

//verifyRequest request object
type verifyRequest struct {
	APIKEY  string
	TransID int64
}

//verifyResponse object response
type verifyResponse struct {
	Status       int64  `json:"status"`
	Amount       string `json:"amount"`
	ErrorCode    int64  `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
