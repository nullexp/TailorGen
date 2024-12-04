package service

import (
	"context"

	"github.com/nullexp/metagen/internal/port/driver/model"
)

type Generator interface {
	Generate(context.Context, model.GenerateRequest) (model.GenerateResponse, error)
}
