package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sazzo/sazztv/backend/database"
	"github.com/sazzo/sazztv/backend/model"
	"github.com/sazzo/sazztv/backend/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

type ChangeUsernameDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}


func changeUsername(c echo.Context) (err error) {
	authUser := c.Get("user").(model.User)

	changeUsernameData := new(ChangeUsernameDTO)
	if err = c.Bind(changeUsernameData); err != nil {
		return c.JSON(http.StatusBadRequest, &util.APIError{
			Message: "Invalid request body",
		})
	}
	if err = c.Validate(changeUsernameData); err != nil {
		return err 
	}
	
	db := database.GetConnection()

	comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(changeUsernameData.Password))
	if comparePasswordErr != nil {
		return c.JSON(http.StatusBadRequest, &util.APIError{
			Message: "Invalid user credentials",
		})
	}

	foreignUser := new(struct {
		ID string `json:"id"`
		Username string `json:"username"`
		IsAdmin bool `json:"is_admin"`
		IsLive bool `json:"is_live"`
		LastStreamAt string `json:"last_stream_at"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	})
	
	updateUserQuery := db.Model(&model.User{}).Clauses(clause.Returning{}).Where("id = ?", authUser.ID).Update("username", changeUsernameData.Username).Scan(&foreignUser)
	if util.IsUniqueConstraintError(updateUserQuery.Error) {
		return c.JSON(http.StatusConflict, &util.APIError{
			Message: "Username already exists",
		})
	}


	return c.JSON(http.StatusOK, foreignUser)
}