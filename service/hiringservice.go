package service

import "github.com/Jestinepa/hiring/domain"

type FreshersService interface {
	GetAllFresher() ([]domain.Freshers, error)
}

type DefaultFreshersService struct {
	repo domain.ExternalRepository
}

func (s DefaultFreshersService) GetAllFresher() ([]domain.Freshers, error) {
	return s.repo.FindAll()
}

func NewFreshersService(repository domain.ExternalRepository) DefaultFreshersService {
	return DefaultFreshersService{repository}
}
