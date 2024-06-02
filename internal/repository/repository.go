package repository

import "send-email/internal/domain/model"

type Repository interface {
    Save(campaign *model.Campaign) (string, error)
}

type repository struct {
    // db connection
}

func NewRepository(cfg config.Config) Repository {
    // return new repository with db connection
}

func (r *repository) Save(campaign *model.Campaign) (string, error) {
    // save logic
}
