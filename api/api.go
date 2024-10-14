package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"auth/db"
	"auth/db/models"
)

func Login(c echo.Context) error {
	var u models.User
	if err := c.Bind(&u); err != nil{
		return err
	}
	if !db.Login_User(u){
		return c.String(http.StatusConflict, "user not find")
	}
	return nil
}
func Profile(c echo.Context) error {
	jwtplayload, ok := db.JwtPayloadFromRequest(c)
	if !ok {
		return c.String(http.StatusNotFound, "error")
	}
	user := db.Get_Prof(jwtplayload["sub"].(string))

	return c.JSON(http.StatusOK, models.UserRequest{Email: user.Email})
}
