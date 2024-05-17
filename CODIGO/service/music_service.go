package service

import (
    "codigo/models"
    "codigo/repository"
)

type MusicService interface {
    Save(models.Music) models.Music
    Update(int, models.Music) error
    Delete(int) error
    FindAll() []models.Music
    FindByID(int) (models.Music, error)
    Search(string, string) []models.Music
}

type musicService struct {
    repository repository.MusicRepository
}

func NewMusicService() MusicService {
    return &musicService{
        repository: repository.NewMusicRepository(),
    }
}

func (s *musicService) Save(music models.Music) models.Music {
    return s.repository.Save(music)
}

func (s *musicService) Update(id int, music models.Music) error {
    return s.repository.Update(id, music)
}

func (s *musicService) Delete(id int) error {
    return s.repository.Delete(id)
}

func (s *musicService) FindAll() []models.Music {
    return s.repository.FindAll()
}

func (s *musicService) FindByID(id int) (models.Music, error) {
    return s.repository.FindByID(id)
}

func (s *musicService) Search(title, artist string) []models.Music {
    return s.repository.Search(title, artist)
}
