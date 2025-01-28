package stores

import (
	"context"
	"time"

	"github.com/Neura-AI-Labs/nova/cache"
	"github.com/Neura-AI-Labs/nova/db"
	"github.com/Neura-AI-Labs/nova/id"

	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

type Store struct {
	db  *gorm.DB
	ctx context.Context
}

type SessionStore struct {
	Store
}

type FragmentStore struct {
	Store
	fragmentTable db.FragmentTable
	cache         *cache.Cache
}

type FragmentFilter struct {
	ActorID   *id.ID
	SessionID *id.ID
	Metadata  []MetadataCondition
	StartTime *time.Time
	EndTime   *time.Time
	Embedding *pgvector.Vector
	Limit     int
}

type ActorStore struct {
	Store
}

type MetadataCondition struct {
	Key      string
	Value    interface{}
	Operator MetadataOperator
}

type MetadataOperator string

const (
	MetadataOpEquals    MetadataOperator = "="
	MetadataOpNotEquals MetadataOperator = "!="
	MetadataOpContains  MetadataOperator = "?"
	MetadataOpIn        MetadataOperator = "IN"
)
