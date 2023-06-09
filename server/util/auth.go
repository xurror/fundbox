package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strconv"
	"time"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func CheckPasswordHash(password string, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(userPassword))

	if err != nil {
		log.Print("E-Mail or Password is incorrect")
		return false
	}
	return true
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func ValidateToken(token string) (uuid.UUID, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{},
		func(parsedToken *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := parsedToken.Claims.(*jwt.StandardClaims)
	if !ok {
		//msg = fmt.Sprintf("the token is invalid")
		//msg = err.Error()
		return uuid.Nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		//msg = fmt.Sprintf("token is expired")
		//msg = err.Error()
		return uuid.Nil, err
	}
	return uuid.MustParse(claims.Id), nil
}

func GenerateJWT(userID uuid.UUID) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString("privateKey")
}

func GenerateAllTokens(id uuid.UUID, email string) (signedToken string, signedRefreshToken string, err error) {
	claims := &jwt.StandardClaims{
		Audience:  "USER",
		Subject:   email,
		Id:        id.String(),
		Issuer:    "API-GRAPHQL",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
	}

	refreshClaims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}
