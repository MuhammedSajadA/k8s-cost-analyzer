package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/services"
)

type NamespaceHandler struct {
	service *services.ClusterService
}

func NewNamespaceHandler(s *services.ClusterService) *NamespaceHandler {
	return &NamespaceHandler{service: s}
}

func (h *NamespaceHandler) List(c *gin.Context) {
	userID := c.GetString("user_id")
	clusterID := c.Param("id")

	namespaces, err := h.service.ListNamespaces(userID, clusterID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cluster_id": clusterID,
		"namespaces": namespaces,
	})
}
