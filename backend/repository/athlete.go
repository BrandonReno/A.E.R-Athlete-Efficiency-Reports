package repository

import (
	"context"
	"database/sql"

	"github.com/BrandonReno/A.E.R/models"
	"gorm.io/gorm"
)

type athleteRepository struct {
	db *gorm.DB
}

func NewAthleteRepository(db *gorm.DB) AthleteRepository {
	return &athleteRepository{db: db}
}

func (ar *athleteRepository) Create(ctx context.Context, a *models.Athlete) error {
	if err := ar.db.WithContext(ctx).Create(a).Error; err != nil {
		return err
	}
	return nil
}

func (ar *athleteRepository) GetAll(ctx context.Context) ([]*models.Athlete, error) {
	var athletes []*models.Athlete
	if err := ar.db.WithContext(ctx).Find(athletes).Error; err != nil {
		return nil, err
	}
	return athletes, nil
}

func (ar *athleteRepository) GetByID(ctx context.Context, aid int) (*models.Athlete, error) {
	var athlete *models.Athlete
	if err := ar.db.WithContext(ctx).Where("id = @id", sql.Named("id", aid)).Find(athlete).Error; err != nil {
		return nil, err
	}
	return athlete, nil
}

func (ar *athleteRepository) Update(ctx context.Context, a *models.Athlete) (*models.Athlete, error) {
	if err := ar.db.WithContext(ctx).Save(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}
