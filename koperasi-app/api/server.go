package api

import (
    // "fmt"

    "github.com/gin-gonic/gin"
    "koperasi-app/blockchain"
)

type Server struct {
    Router *gin.Engine
    Fc     *blockchain.FabricClient
}

func NewServer(fc *blockchain.FabricClient) *Server {
    r := gin.Default()
    s := &Server{Router: r, Fc: fc}

    api := r.Group("/api/v1")
    {
        inv := api.Group("/inventory")
        inv.POST("", s.AddInventory)
        inv.GET(":id", s.GetInventory)
    }

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    return s
}

func (s *Server) Run(addr string) error {
    if addr == "" {
        addr = ":8080"
    }
    return s.Router.Run(addr)
}

// helper used by handlers
func (s *Server) respondError(c *gin.Context, code int, err error) {
    c.JSON(code, gin.H{"error": err.Error()})
}

func (s *Server) respondOK(c *gin.Context, data interface{}) {
    c.JSON(200, gin.H{"data": data})
}
