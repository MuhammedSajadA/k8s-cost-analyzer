package repositories

import (
	"errors"

	"gorm.io/gorm"
	"github.com/MuhammedSajadA/k8s-cost-analyzer/internal/models"
)

type ClusterRepository struct {
	db *gorm.DB
}

func NewClusterRepository(db *gorm.DB) *ClusterRepository {
	return &ClusterRepository{db: db}
}

func (r *ClusterRepository) Create(cluster *models.Cluster) error {
	return r.db.Create(cluster).Error
}

func (r *ClusterRepository) FindByIDAndUser(
	clusterID string,
	userID string,
) (*models.Cluster, error) {

	var cluster models.Cluster
	err := r.db.
		Where("id = ? AND user_id = ?", clusterID, userID).
		First(&cluster).
		Error

	if err != nil {
		return nil, errors.New("cluster not found")
	}

	return &cluster, nil
}
