package provider

import (
	"Nexus/internal/domain/model"
	"context"
)

type Provider interface {
	Chat(ctx context.Context, req *model.ChatRequest) (*model.ChatResponse, error)
}
