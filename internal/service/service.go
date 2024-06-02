package service

import (
    "send-email/internal/domain/model"
    "send-email/internal/repository"
)

type Service struct {
    repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
    return &Service{repo: repo}
}

func (s *Service) CreateCampaign(name, content string, emails []string) (string, error) {
    campaign := &model.Campaign{Name: name, Content: content, Emails: emails}
    return s.repo.Save(campaign)
}
