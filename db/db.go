package db

import (
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"auth/db/models"
)

var Key string = "fjb13jb2hgjbh35bhbjbjfcj1klkj4tr932490gvj9290vj1o3finhb13iulrfh89i42ghfvn9023ruj"


const contextKeyUser = "user"

// jwtPayloadFromRequest извлекает JWT токен из контекста и возвращает его указанные заявленные, если это возможно.
func JwtPayloadFromRequest(c echo.Context) (jwt.MapClaims, bool) {
    jwtToken, ok := c.Get(contextKeyUser).(*jwt.Token)
    if !ok {
        logrus.WithFields(logrus.Fields{
            "jwt_token_context_value": c.Get(contextKeyUser),
        }).Error("wrong type of JWT token in context")
        return nil, false
    }

    payload, ok := jwtToken.Claims.(jwt.MapClaims)
    if !ok {
        logrus.WithFields(logrus.Fields{
            "jwt_token_claims": jwtToken.Claims,
        }).Error("wrong type of JWT token claims")
        return nil, false
    }

    return payload, true
}
func Login_User(u models.User) bool {
	dsn := "host=localhost user=osami password= dbname=users port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn))
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err != nil {
		loger.Error(err.Error())
	}

	if errs := db.Where("Email = ?", u.Email).Where("Password = ?", u.Password).Update("Token", GenerateToken(u)).Error; errs != nil {
		loger.Error(errs.Error())
		return false
	}
	return true
}
func GenerateToken(u models.User) string {
	payload := jwt.MapClaims{
		"sub": u.Email,
		"exp" : time.Now().Add(time.Hour * 72).Unix(),
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := tkn.SignedString(Key)
	if err != nil{
		slog.New(slog.NewJSONHandler(os.Stdout, nil)).Error(err.Error())
	}
	return t
}
func Get_Prof(email string) models.User {
	var u models.User
	
	dsn := "host=localhost user=osami password= dbname=users port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil{
		slog.New(slog.NewJSONHandler(os.Stdout, nil)).Error(err.Error())
	}

	if err := db.Where("Email = ?", email).Find(&u); err != nil{
		slog.New(slog.NewJSONHandler(os.Stdout, nil)).Error(err.Error.Error())
	}

	return u
}