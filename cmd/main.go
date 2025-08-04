package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thekrauss/nafyp-gateway/config"
	"github.com/thekrauss/nafyp-gateway/handlers"
	"github.com/thekrauss/nafyp-gateway/internal"
)

func main() {
	// la config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Échec du chargement de la configuration : %v", err)
	}
	config.AppConfig = *cfg

	// init les clients gRPC
	clients := internal.InitGRPCClients()

	// Créer les handlers
	authHandler := handlers.NewAuthHandler(clients)

	// init Gin
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Routes
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.GET("/validate-token", authHandler.ValidateToken)

	// lance le serveur
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.HTTPPort)
	log.Printf("[GATEWAY] API HTTP démarrée sur %s", addr)
	r.Run(addr)
}
