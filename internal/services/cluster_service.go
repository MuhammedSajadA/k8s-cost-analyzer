package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/models"
	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/repositories"
	"github.com/MuhammedSajadA/k8s-cost-analyzer/pkg/k8s"
)

type ClusterService struct {
	repo *repositories.ClusterRepository
}

func NewClusterService(repo *repositories.ClusterRepository) *ClusterService {
	return &ClusterService{repo: repo}
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

	return s.repo.Create(&cluster)
}
func (s *ClusterService) ListNamespaces(
	userID string,
	clusterID string,
) ([]string, error) {

	// 1️⃣ Verify ownership
	cluster, err := s.repo.FindByIDAndUser(clusterID, userID)
	if err != nil {
		return nil, err
	}

	// 2️⃣ Build K8s client
	clientset, err := k8s.NewClient([]byte(cluster.Kubeconfig))
	if err != nil {
		return nil, err
	}

	// 3️⃣ Fetch namespaces
	return k8s.ListNamespaces(clientset)
}

