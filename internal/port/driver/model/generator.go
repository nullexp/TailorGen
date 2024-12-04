package model

type GenerateRequest struct {
	Config string
}

type GenerateResponse struct {
	ExecutionTimeMiliSecond int64
}
