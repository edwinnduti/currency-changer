package model

type Currency struct{
	Type	string	`json:"type"`
	Amount	float32	`json:"amount"`
}

type Money struct{
	To	string	`json:"to"`
	From	string	`json:"from"`
	Cash	float32	`json:"cash"`
}

type Response struct{
	Code	int	`json:"code"`
	Message	string	`json:"message"`
}
