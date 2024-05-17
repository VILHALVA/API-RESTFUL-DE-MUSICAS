package service

import (
    "errors"
    "time"
    "github.com/dgrijalva/jwt-go"
)

type AuthService interface {
    Login(username, password string) (string, error)
    ValidateToken(token string) (*jwt.Token, error)
}

type authService struct {
    secretKey string
    users     map[string]string // username:password
}

func NewAuthService() AuthService {
    return &authService{
        secretKey: "mysecretkey",
        users: map[string]string{
            "user1": "password1",
            "user2": "password2",
        },
    }
}

func (s *authService) Login(username, password string) (string, error) {
    if pass, ok := s.users[username]; ok && pass == password {
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "username": username,
            "exp":      time.Now().Add(time.Hour * 24).Unix(),
        })
        return token.SignedString([]byte(s.secretKey))
    }
    return "", errors.New("invalid credentials")
}

func (s *authService) ValidateToken(token string) (*jwt.Token, error) {
    return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid signing method")
        }
        return []byte(s.secretKey), nil
    })
}
