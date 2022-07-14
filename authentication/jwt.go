package authentication

import (
    "crypto/rand"
    "fmt"
    "github.com/golang-jwt/jwt"
    "github.com/google/uuid"
    "log"
)

var secret = make([]byte, 16)

func init() {
    _, err := rand.Read(secret)
    if err != nil {
        log.Fatalln(err)
    }
}

func Sign(userId uuid.UUID) (string, error) {
    return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"id": userId.String()}).SignedString(secret)
}

func Id(tokenString string) (uuid.UUID, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return secret, nil
    })
    if err != nil {
        return uuid.Nil, err
    }
    if !token.Valid {
        return uuid.Nil, fmt.Errorf("token is not valid: %s", tokenString)
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return uuid.Nil, fmt.Errorf("cannot parse claims: %v", token.Claims)
    }
    idValue, ok := claims["id"]
    if !ok {
        return uuid.Nil, fmt.Errorf("cannot find 'id' field: %v", claims)
    }
    idString, ok := idValue.(string)
    if !ok {
        return uuid.Nil, fmt.Errorf("cannot cast id to string: %v", idValue)
    }
    id, err := uuid.Parse(idString)
    if err != nil {
        return uuid.Nil, err
    }
    return id, nil
}
