package auth

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"open-btm.com/common"
	"open-btm.com/models"
	"open-btm.com/observe"
)

type LoginResponse struct {
	AccessToken  string `json:"AccessToken"`
	RefreshToken string `json:"RefreshToken"`
}
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login
// @Summary LogIn User
// @Description Login User
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body UserLogin true "Login User"
// @Success 200 {object} common.ResponseHTTP{data=LoginResponse}
// @Failure 400 {object} common.ResponseHTTP{}
// @Failure 500 {object} common.ResponseHTTP{}
// @Router /login [post]
func LoginUser(contx echo.Context) error {
	//  Geting tracer
	tracer := contx.Get("tracer").(*observe.RouteTracer)

	// validator initialization
	validate := validator.New()

	//validating post data
	posted_user := new(UserLogin)

	//first parse request data
	if err := contx.Bind(&posted_user); err != nil {
		return contx.JSON(http.StatusBadRequest, common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// then validate structure
	if err := validate.Struct(posted_user); err != nil {
		return contx.JSON(http.StatusBadRequest, common.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	//  creating user via Api call to blue Admin
	resp, err := models.LoginUserBlueAdmin(tracer.Tracer, posted_user.Email, posted_user.Password)
	if err != nil {
		return err
	}

	result := LoginResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}

	// return data if transaction is sucessfull
	return contx.JSON(http.StatusOK, common.ResponseHTTP{
		Success: true,
		Message: "User Logged in successfully",
		Data:    result,
	})
}
