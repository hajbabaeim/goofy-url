package usecase

import "github.com/hajbabaeim/goofy-url/url/domain"

type ShortenerService struct {
	repo  repository.URLRepository // Interface dependency
	cache cache.Cache
}

func (s *ShortenerService) Create(original string) (*domain.ShortURL, error) {
	if err := domain.ValidateURL(original); err != nil {
		return nil, err
	}
}
