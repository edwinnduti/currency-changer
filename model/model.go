package model

// currency struct to be sent via write function
type Currency struct{
	Type	string	`json:"type"`
	Amount	float32	`json:"amount"`
}

// money struct from r.Body
type Money struct{
	To	string	`json:"to"`
	From	string	`json:"from"`
	Cash	float32	`json:"cash"`
}

// responses
type Response struct{
	Code	int	`json:"code"`
	Message	string	`json:"message"`
}
