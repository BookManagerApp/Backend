package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var secretKey = []byte("your_secret_key") // Ganti dengan kunci rahasia yang kuat

// GenerateToken menghasilkan JWT token untuk pengguna
func GenerateToken(userID uint, email string) (string, error) {
    claims := jwt.MapClaims{
        "id_user": userID,
        "email":   email,
        "exp":     time.Now().Add(time.Hour * 72).Unix(), // Token berlaku selama 72 jam
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

// ParseToken mem-parse JWT token dan mengembalikan klaim
func ParseToken(tokenStr string) (*jwt.Token, *jwt.MapClaims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err != nil {
        return nil, nil, err
    }

    claims, ok := token.Claims.(*jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, nil, err
    }

    return token, claims, nil
}
