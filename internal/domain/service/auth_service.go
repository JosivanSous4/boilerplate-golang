// package service

// import (
// 	"boilerplate-go/internal/domain/model"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// )

// var jwtKey = []byte("your_secret_key") // Change this in real applications

// // Claims defines the structure of the JWT claims
// type Claims struct {
//     Username string `json:"username"`
//     jwt.StandardClaims
// }

// // GenerateToken generates a JWT for the given user
// func GenerateToken(user model.User) (string, error) {
//     expirationTime := time.Now().Add(5 * time.Minute) // Token expires in 5 minutes

//     claims := &Claims{
//         Username: user.Username,
//         StandardClaims: jwt.StandardClaims{
//             ExpiresAt: expirationTime.Unix(),
//         },
//     }

//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//     signedToken, err := token.SignedString(jwtKey)
//     if err != nil {
//         return "", err
//     }

//     return signedToken, nil
// }

// // VerifyToken verifies the validity of the given token
// func VerifyToken(tokenString string) (*jwt.Token, error) {
//     token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//         return jwtKey, nil
//     })
//     if err != nil {
//         return nil, err
//     }
//     return token, nil
// }

package service

import (
	"boilerplate-go/internal/domain/repository"
	"boilerplate-go/internal/infrastructure/security"
	"context"
	"errors"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type AuthService interface {
    Login(ctx context.Context, login LoginRequest) (string, error)
}

type authService struct {
    repo     repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
    return &authService{
        repo:     repo,
    }
}

func (s *authService) Login(ctx context.Context, login LoginRequest) (string, error) {

    user, err := s.repo.GetUserByEmail(ctx, login.Username)

    if err != nil {
        return "", err
    }

    checkPass := security.CheckPasswordHash(login.Password, user.Password)

    if !checkPass {
        return "", errors.New("usuário ou senha incorreta")
    }

    token, _ := security.GenerateJWT(user.Username)

    return token, nil
}
