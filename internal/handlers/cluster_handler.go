package handlers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/services"
)

type ClusterHandler struct {
	service *services.ClusterService
}

func NewClusterHandler(s *services.ClusterService) *ClusterHandler {
	return &ClusterHandler{service: s}
}

func (h *ClusterHandler) AddCluster(c *gin.Context) {
	userID := c.GetString("user_id")
	name := c.PostForm("name")

	file, _, err := c.Request.FormFile("kubeconfig")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "kubeconfig required"})
		return
	}
	defer file.Close()

	data, _ := io.ReadAll(file)

	if err := h.service.AddCluster(userID, name, data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "cluster added"})
}
