package controllers

import (
	"backendnuxt/lib"
	"backendnuxt/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRegister(ctx *gin.Context) {
	var form models.Users
	err := ctx.ShouldBind(&form)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	findUser := models.FindOneUserByEmail(form.Email)
	// if err != nil {
	// ctx.JSON(http.StatusInternalServerError, Response{
	// Success: false,
	// Message: "Error checking user existence",
	// })
	// return
	// }

	// fmt.Println("error = ", findUser)
	if form.Email == findUser.Email {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Email Has Registered",
		})
		return
	}
	hash := lib.CreateHash(form.Password)
	form.Password = hash
	newUser := models.InsertUser(form)
	// if err != nil {
	// ctx.JSON(http.StatusInternalServerError, Response{
	// Success: false,
	// Message: "User Registration Failed",
	// })
	// return
	// }

	if newUser.Id > 0 {
		profile := models.RelationProfile{
			First_Name:   "",
			Last_Name:    "",
			Image:        "",
			Phone_Number: "",
			User_Id:      newUser.Id,
		}

		models.AddProfile(profile)
		// if err != nil {
		// ctx.JSON(http.StatusInternalServerError, Response{
		// Success: false,
		// Message: "Profile Creation Failed",
		// })
		// return
		// }

		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Register Success"})
	} else {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "User Registration Failed",
		})
	}
}

func AuthLogin(ctx *gin.Context) {
	var form models.Users
	ctx.ShouldBind(&form)

	foundUser := models.FindOneUserByEmail(form.Email)
	if form.Email != foundUser.Email {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid Email & Password",
		})
		return
	}

	match := lib.GenerateTokenArgon(form.Password, foundUser.Password)
	if !match {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid Email & Password",
		})
		return
	}

	token := lib.GeneretedToken(struct {
		UserId int `json:"userId"`
	}{
		UserId: foundUser.Id,
	})
	// fmt.Println("data = ", token)
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Login Success",
		Results: token,
	})

}
