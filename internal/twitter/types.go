package twitter

import (
	"context"
	"time"

	"github.com/Neura-AI-Labs/nova/engine"
	"github.com/Neura-AI-Labs/nova/llm"
	"github.com/Neura-AI-Labs/nova/logger"
	"github.com/Neura-AI-Labs/nova/options"
	"github.com/Neura-AI-Labs/nova/pkg/twitter"

	"gorm.io/gorm"
)

type Twitter struct {
	options.RequiredFields

	ctx       context.Context
	logger    *logger.Logger
	database  *gorm.DB
	llmClient *llm.LLMClient

	assistant *engine.Engine

	twitterClient *twitter.Client
	twitterConfig TwitterConfig

	stopChan chan struct{}
}

type TwitterCredentials struct {
	CT0       string
	AuthToken string
	User      string
}

type IntervalConfig struct {
	Min time.Duration
	Max time.Duration
}

type TwitterConfig struct {
	MonitorInterval IntervalConfig
	Credentials     TwitterCredentials
}
