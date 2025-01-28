package engine

import (
	"context"

	"github.com/Neura-AI-Labs/nova/id"
	"github.com/Neura-AI-Labs/nova/llm"
	"github.com/Neura-AI-Labs/nova/logger"
	"github.com/Neura-AI-Labs/nova/manager"
	"github.com/Neura-AI-Labs/nova/options"
	"github.com/Neura-AI-Labs/nova/stores"

	"gorm.io/gorm"
)

type Engine struct {
	options.RequiredFields

	ctx context.Context

	db *gorm.DB

	logger *logger.Logger

	ID   id.ID
	Name string

	// State management
	managers     []manager.Manager
	managerOrder []manager.ManagerID

	// stores
	actorStore   *stores.ActorStore
	sessionStore *stores.SessionStore

	interactionFragmentStore *stores.FragmentStore

	// LLM client
	llmClient *llm.LLMClient
}
