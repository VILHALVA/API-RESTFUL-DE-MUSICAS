package repository

import (
    "errors"
    "codigo/models"
)

type MusicRepository interface {
    Save(models.Music) models.Music
    Update(int, models.Music) error
    Delete(int) error
    FindAll() []models.Music
    FindByID(int) (models.Music, error)
    Search(string, string) []models.Music
}

type musicRepository struct {
    musics []models.Music
}

func NewMusicRepository() MusicRepository {
    return &musicRepository{
        musics: []models.Music{
            {ID: 1, Title: "Bohemian Rhapsody", Artist: "Queen", Album: "A Night at the Opera", Year: 1975},
            {ID: 2, Title: "Stairway to Heaven", Artist: "Led Zeppelin", Album: "Led Zeppelin IV", Year: 1971},
        },
    }
}

func (r *musicRepository) Save(music models.Music) models.Music {
    music.ID = len(r.musics) + 1
    r.musics = append(r.musics, music)
    return music
}

func (r *musicRepository) Update(id int, music models.Music) error {
    for i, m := range r.musics {
        if m.ID == id {
            r.musics[i] = music
            r.musics[i].ID = id
            return nil
        }
    }
    return errors.New("music not found")
}

func (r *musicRepository) Delete(id int) error {
    for i, m := range r.musics {
        if m.ID == id {
            r.musics = append(r.musics[:i], r.musics[i+1:]...)
            return nil
        }
    }
    return errors.New("music not found")
}

func (r *musicRepository) FindAll() []models.Music {
    return r.musics
}

func (r *musicRepository) FindByID(id int) (models.Music, error) {
    for _, m := range r.musics {
        if m.ID == id {
            return m, nil
        }
    }
    return models.Music{}, errors.New("music not found")
}

func (r *musicRepository) Search(title, artist string) []models.Music {
    var result []models.Music
    for _, m := range r.musics {
        if (title == "" || m.Title == title) && (artist == "" || m.Artist == artist) {
            result = append(result, m)
        }
    }
    return result
}
