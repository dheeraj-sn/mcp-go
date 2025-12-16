package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"log"
	"net/http"
	"internal/tools"
	"github.com/dheeraj-sn/mcp-go/internal/mcp"
	"github.com/dheeraj-sn/mcp-go/internal/middleware"
	"github.com/dheeraj-sn/mcp-go/internal/domain/limiter"	
)

func main() {
	// Root context and shutdown signal handling
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Domain services
	limiterService := limiter.NewService()

	// Tool registry
	registry := tools.NewRegistry()

	registry.Register(
		limit.NewCreateTool(limiterService),
	)
	registry.Register(
		limit.NewGetTool(limiterService),
	)
	registry.Register(
		limit.NewDeleteTool(limiterService),
	)

	// Middleware chain, the order of the middleware is important
	toolsMiddleware := middleware.Chain(
		middleware.WithAuth(),
		middleware.WithRateLimit(100),
		middleware.WithTimeout(2*time.Second),
		middleware.WithMetrics(),
	)

	// mcp server
	server := mcp.NewServer(
		mcp.WithToolRegistry(registry),
		mcp.WithToolMiddleware(toolMiddleware),
	)

	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           server.Handler(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func(){
		log.Println("Starting MCP server on :8080")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}

	<-ctx.Done()
	log.Println("Shutting down MCP server...")
	shutdownCtx, cancel := context.WithTimeout(
        context.Background(),
        10*time.Second,
    )
    defer cancel()
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
        log.Printf("graceful shutdown failed: %v", err)
    }
	log.Println("server stopped")
}
