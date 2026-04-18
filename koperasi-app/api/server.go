package api

import (
    "github.com/gin-gonic/gin"
    "koperasi-app/blockchain"
)

type Server struct {
    Router *gin.Engine
    Fc     blockchain.ChaincodeAPI
}

func NewServer(fc blockchain.ChaincodeAPI) *Server {
    r := gin.Default()
    s := &Server{Router: r, Fc: fc}

    api := r.Group("/api/v1")
    {
        inv := api.Group("/inventory")
        inv.POST("", s.AddInventory)
        inv.GET(":id", s.GetInventory)
        inv.PUT(":id", s.UpdateInventory)
        inv.DELETE(":id", s.DeleteInventory)
        inv.GET(":id/history", s.GetInventoryHistory)
    }

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    return s
}

func (s *Server) Run(addr string) error {
    return s.Router.Run(addr)
}

func (s *Server) respondError(c *gin.Context, code int, err error) {
    c.JSON(code, gin.H{"error": err.Error()})
}