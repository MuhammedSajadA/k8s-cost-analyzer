package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    
	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/models"
	"github.com/MuhammedSajadA/k8s-cost-analyzer/pkg/k8s"
)

type ClusterService struct {
	db *gorm.DB
}

func NewClusterService(db *gorm.DB) *ClusterService {
	return &ClusterService{db: db}
}

func (s *ClusterService) AddCluster(
	userID string,
	name string,
	kubeconfig []byte,
) error {

	// 1️⃣ Validate kubeconfig by connecting to cluster
	clientset, err := k8s.NewClient(kubeconfig)
	if err != nil {
		return errors.New("invalid kubeconfig")
	}

	// 2️⃣ Test access (simple API call)
	_, err = clientset.CoreV1().
		Namespaces().
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return errors.New("cannot access cluster")
	}

	// 3️⃣ Save cluster
	cluster := models.Cluster{
		ID:         uuid.New().String(),
		UserID:     userID,
		Name:       name,
		Kubeconfig: string(kubeconfig),
	}

	return s.db.Create(&cluster).Error
}
