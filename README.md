# Nova - An Advanced Conversational AI Framework

<div align="center">
  <img src="./img/neura_banner.png" alt="Nova Banner" width="100%" />
</div>

## Table of Contents
- [Overview](#overview)
- [Core Features](#core-features)
- [Extension Points](#extension-points)
- [Quick Start](#quick-start)
- [Using Nova as a Module](#using-nova-as-a-module)

## Overview
Nova is a next-generation, highly modular AI conversation engine built in Go, designed to provide unparalleled flexibility, extensibility, and platform integration. Nova emphasizes a pluggable architecture and is tailored for diverse conversational use cases through:

- Advanced plugin-based architecture with dynamic component swapping
- Multi-provider LLM support, including OpenAI and custom integrations
- Cross-platform conversation management tools
- Enhanced behavior control via configurable manager systems
- Real-time vector-based semantic search with optimized pgvector
- **New Feature:** Pre-trained personality packs for quick deployment of custom bots
- **New Feature:** Built-in analytics dashboard for monitoring conversation metrics

## Core Features

### Plugin Architecture
- **Enhanced Manager System**: Extend Nova's functionality with modular managers:
  - Insight Manager: Captures and maintains actionable conversation insights.
  - Personality Manager: Configures bot responses, tone, and user preferences.
  - **AI Skill Manager**: Prebuilt skills for FAQs, sentiment analysis, and more.
  - Custom Managers: Develop your own behaviors for unique conversational scenarios.

### State Management
- **Unified Shared State System**:
  - Centralized state storage for seamless data exchange between components.
  - Enhanced support for data caching, improving system performance.
  - New support for state serialization for distributed systems.

### LLM Integration
- **Provider Flexibility**:
  - Multi-provider support, including OpenAI and other third-party LLMs.
  - **Streaming Output**: Handle partial responses for real-time feedback.
  - Configurable provider fallback mechanisms for fail-safe operations.
  - Dynamic model swapping based on runtime performance metrics.

### Platform Support
- **Platform-Neutral Core**:
  - Fully decoupled conversation engine with support for cloud and edge deployments.
  - Out-of-the-box connectors for Slack, Discord, and WhatsApp.
  - Customizable platform manager APIs for seamless integration with any platform.

### Storage Layer
- **Robust Data Storage**:
  - PostgreSQL backend with pgvector for advanced semantic search.
  - Schema customization for unique data storage requirements.
  - **New Feature:** Real-time embeddings monitoring with an interactive CLI interface.

### Toolkit/Function System
- **Pluggable Toolkits**:
  - Seamlessly integrate new tools for extended capabilities.
  - Built-in support for database queries, API calls, and scheduling tasks.
  - **New Feature:** Tool sandboxing to prevent unauthorized tool execution.
  - Context-aware tool execution for state-sensitive interactions.

## Extension Points
Nova allows for significant customization through well-defined extension points:

1. **LLM Providers**: Add support for new AI providers by implementing the LLM interface:
```go
type Provider interface {
    GenerateCompletion(context.Context, CompletionRequest) (string, error)
    GenerateStreamingCompletion(context.Context, StreamingRequest) (<-chan string, error)
    GenerateEmbeddings(context.Context, string) ([]float32, error)
}
```

2. **Managers**: Expand the functionality by creating new manager implementations:
```go
type Manager interface {
    GetID() ManagerID
    GetDependencies() []ManagerID
    Process(state *state.State) error
    PostProcess(state *state.State) error
    Context(state *state.State) ([]state.StateData, error)
    Store(fragment *db.Fragment) error
    StartBackgroundProcesses()
    StopBackgroundProcesses()
    RegisterEventHandler(callback EventCallbackFunc)
    TriggerEvent(eventData EventData)
}
```

## Quick Start
1. Clone the repository:
```bash 
git clone https://github.com/Neura-AI-Labs/nova.git
```   
2. Copy `.env.example` to `.env` and configure your environment variables.
3. Install dependencies:
```bash
go mod download
```
4. Run the CLI example:
```bash
go run examples/cli/main.go
```
5. Run the Slack bot integration:
```bash
go run examples/slack/main.go
```

## Environment Variables
```env
DB_URL=postgresql://user:password@localhost:5432/nova
OPENAI_API_KEY=your_openai_api_key

# Platform-specific credentials
SLACK_API_TOKEN=your_slack_token
DISCORD_BOT_TOKEN=your_discord_token
```

## Architecture
The project follows a clean, modular architecture:

- `engine`: Core conversation engine
- `manager`: Modular plugin manager system
- `managers/*`: Built-in manager implementations
- `state`: Shared state management layer
- `llm`: Interfaces for LLM provider integration
- `stores`: Data storage and retrieval layer
- `tools/*`: Built-in tool implementations
- `examples/`: Sample implementations for various platforms

## Using Nova as a Module

1. Add Nova to your Go project:
```bash
go get github.com/Neura-AI-Labs/nova
```

2. Import Nova into your code:
```go
import (
  "github.com/Neura-AI-Labs/nova/engine"
  "github.com/Neura-AI-Labs/nova/llm"
  "github.com/Neura-AI-Labs/nova/manager"
  "github.com/Neura-AI-Labs/nova/managers/personality"
  "github.com/Neura-AI-Labs/nova/managers/insight"
  "github.com/Neura-AI-Labs/nova/tools/database"
  ... etc
)
```

3. Basic usage example:
```go
// Initialize LLM client
llmClient, err := llm.NewLLMClient(llm.Config{
  ProviderType: llm.ProviderOpenAI,
  APIKey: os.Getenv("OPENAI_API_KEY"),
  ModelConfig: map[llm.ModelType]string{
    llm.ModelTypeDefault: openai.GPT4,
  },
  Logger: logger,
  Context: ctx,
})

// Create engine instance
engine, err := engine.New(
  engine.WithContext(ctx),
  engine.WithLogger(logger),
  engine.WithDB(db),
  engine.WithLLM(llmClient),
)

// Process input
state, err := engine.NewState(actorID, sessionID, "Your input text here")
if err != nil {
  log.Fatal(err)
}

response, err := engine.Process(state)
if err != nil {
  log.Fatal(err)
}
```

4. Explore the available packages:
- `nova/engine`: Core conversation engine
- `nova/llm`: LLM provider interfaces and implementations
- `nova/manager`: Base manager system
- `nova/managers/*`: Built-in manager implementations
- `nova/state`: State management utilities
- `nova/stores`: Data storage implementations

For detailed examples, see the `examples/` directory in the repository.
