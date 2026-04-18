package api

import (
    "net/http"

    "github.com/gin-gonic/gin"
    // "koperasi-app/blockchain" // Removed unused import
)

type InventoryRequest struct {
    ID              string `json:"id" binding:"required"`
    Code            string `json:"code" binding:"required"`
    Name            string `json:"name" binding:"required"`
    Description     string `json:"description"`
    Symbol          string `json:"symbol"`
    ConversionFactor int   `json:"conversionFactor"`
    BaseUnit        bool   `json:"baseUnit"`
    Category        string `json:"category"`
    Status          bool   `json:"status"`
}

// AddInventory handles POST /api/v1/inventory
func (s *Server) AddInventory(c *gin.Context) {
    var req InventoryRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // call blockchain wrapper
    tx, err := s.Fc.AddInventoryItem(req.ID, req.Code, req.Name, req.Description, req.Symbol, req.ConversionFactor, req.BaseUnit, req.Category, req.Status)
    if err != nil {
        s.respondError(c, http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusCreated, gin.H{"tx": tx})
}

// GetInventory handles GET /api/v1/inventory/:id
func (s *Server) GetInventory(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
        return
    }

    item, err := s.Fc.ReadItem(id)
    if err != nil {
        s.respondError(c, http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusOK, gin.H{"item": item})
}
