package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"webkodes.com/admin/models"
	"webkodes.com/admin/utils"
)

var jwtKey = []byte("my_secret_key")

type claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

//CreateToken utility function to create jwt tokens
func CreateToken(u models.User) models.User {

	//Creating Access Token String
	accessTokenExpirationTime := time.Now().Add(30 * time.Minute)
	accessClaims := &claims{
		Username: u.Username,
		Role:     u.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpirationTime.Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, _ := accessToken.SignedString(jwtKey)

	u.Tokens.AccessToken = accessTokenString
	u.Tokens.Validity = accessTokenExpirationTime
	return u
}

// Login controller to handel login request
func Login(c *gin.Context) {
	var u models.User
	err := c.BindJSON(&u)
	if err != nil {
		fmt.Println(err)
	}
	if u.Username != "" && u.Password != "" {
		user := models.GetSingleUser(bson.M{"username": u.Username})
		if user.Username == u.Username {
			if utils.CheckPasswordHash(u.Password, user.Password) {
				u = CreateToken(u)
				models.UpdateUser(bson.M{"tokens": u.Tokens}, bson.M{"_id": user.ID})
				user := models.GetSingleUser(bson.M{"_id": user.ID})
				if user.Tokens.AccessToken == u.Tokens.AccessToken {
					c.SetCookie("token", u.Tokens.AccessToken, 1.577e+7, "", "webkodes.com", true, true)
					c.SetCookie("isLoggedIn", "true", 1.577e+7, "", "webkodes.com", true, false)
					c.JSON(http.StatusOK, gin.H{
						"status":  "success",
						"message": "User logged in",
						"tokens": gin.H{
							"access_token": u.Tokens.AccessToken,
							"validity":     u.Tokens.Validity,
						},
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"status":  "failed",
						"type":    "not_updated_tokens",
						"message": "Unable to Update Token",
					})
				}
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  "failed",
					"type":    "password_incorrect",
					"message": "Password Incorrect",
				})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"type":    "username_incorrect",
				"message": "User Not Registered",
			})
		}
	} else {

		c.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"type":    "data_not_received",
			"message": "Data not received",
		})
	}
}

// helper function to add user to database
func addUserAfterVerification(c *gin.Context, userFromRequest models.User) {
	checkUser := models.GetSingleUser(bson.M{"username": userFromRequest.Username})
	if checkUser.Username == userFromRequest.Username {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"type":    "username_exists",
			"message": "Username Already Exists",
		})
	} else {
		models.InsertUser(userFromRequest)
		user := models.GetSingleUser(bson.M{"username": userFromRequest.Username})
		if user.Username == userFromRequest.Username {
			c.SetCookie("token", user.Tokens.AccessToken, 1.577e+7, "", "webkodes.com", true, true)
			c.SetCookie("isLoggedIn", "true", 1.577e+7, "", "webkodes.com", true, false)
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "Account created successful",
				"user_id": user.ID,
				"tokens": gin.H{
					"access_token": userFromRequest.Tokens.AccessToken,
					"validity":     userFromRequest.Tokens.Validity,
				},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"type":    "no_record_after_signup",
				"message": "Internal Server Error",
			})
		}
	}
}

// Signup controller to handel sign up request
func Signup(c *gin.Context) {
	var userFromRequest models.User
	err := c.BindJSON(&userFromRequest)
	if err != nil {
		fmt.Println(err)
	}
	if userFromRequest.Username != "" && userFromRequest.Password != "" && userFromRequest.Email != "" && userFromRequest.Name != "" && userFromRequest.Role != "" {
		userFromRequest.Password, _ = utils.HashPassword(userFromRequest.Password)
		userFromRequest.CreatedAt = time.Now()
		userFromRequest = CreateToken(userFromRequest)
		if userFromRequest.Role != "default" {
			tokenFromHeader := c.Request.Header.Get("Authorization")
			if tokenFromHeader != "" {
				splitToken := strings.Split(tokenFromHeader, "Bearer")
				tokenFromHeader = strings.TrimSpace(splitToken[1])
				claims := jwt.MapClaims{}
				token, err := jwt.ParseWithClaims(tokenFromHeader, claims, func(token *jwt.Token) (interface{}, error) {
					return jwtKey, nil
				})
				if err != nil || !token.Valid {
					c.JSON(http.StatusOK, gin.H{
						"status":  "failed",
						"type":    "not_authorized",
						"message": "You are not authorized to create such user",
					})
				} else {
					userForVerification := models.GetSingleUser(bson.M{"username": claims["username"]})
					if userForVerification.Role == "admin" {
						addUserAfterVerification(c, userFromRequest)
					} else {
						c.JSON(http.StatusOK, gin.H{
							"status":  "failed",
							"type":    "not_authorized",
							"message": "You are not authorized to create such user",
						})
					}
				}
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  "failed",
					"type":    "not_authorized",
					"message": "No Authorization token found",
				})
			}
		} else {
			addUserAfterVerification(c, userFromRequest)
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"type":    "data_not_received",
			"message": "Data not received",
		})
	}
}

// Users controller to handle get all users request
func Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"users":  models.GetUsers(bson.M{}),
	})
}

// User controller to handle operation of getting a single user
func User(c *gin.Context) {
	var u models.User
	err := c.BindJSON(&u)
	if err != nil {
		fmt.Println(err)
	}
	user := models.GetSingleUser(bson.M{"username": u.Username})
	if user.Username == u.Username {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"user":   user,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"type":    "no_record",
			"message": "No users found",
		})
	}
}

// UpdateUser controller to handle update request details of a user
func UpdateUser(c *gin.Context) {
	var u models.User
	err := c.BindJSON(&u)
	if err != nil {
		fmt.Println(err)
	}
	if u.Username != "" && u.Email != "" && u.Name != "" {
		models.UpdateUser(bson.M{"name": u.Name, "email": u.Email, "username": u.Username}, bson.M{"_id": u.ID})
		user := models.GetSingleUser(bson.M{"_id": u.ID})
		if user.Name == u.Name && user.Email == u.Email && user.Username == u.Username {
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"user":   user,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"type":    "not_updated",
				"message": "No user updated",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"type":    "data_not_received",
			"message": "Data not received",
		})
	}
}

// DeleteUser controller to handle delete request for a user
func DeleteUser(c *gin.Context) {
	var u models.User
	err := c.BindJSON(&u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u.ID)
	models.DeleteUser(bson.M{"_id": u.ID})
	user := models.GetSingleUser(bson.M{"_id": u.ID})
	if user.ID != u.ID {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"type":    "not_deleted",
			"message": "No user deleted",
		})
	}

}

// RefreshToken controller to handle refreshing of the access token
func RefreshToken(c *gin.Context) {
	var u models.User
	err := c.BindJSON(&u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
	if u.ID.String() != "" {
		user := models.GetSingleUser(bson.M{"_id": u.ID})
		fmt.Println(user)
		if u.ID == user.ID {
			u = CreateToken(u)
			models.UpdateUser(bson.M{"tokens": u.Tokens}, bson.M{"_id": user.ID})
			user := models.GetSingleUser(bson.M{"_id": user.ID})
			fmt.Println(user)
			if user.Tokens.AccessToken == u.Tokens.AccessToken {
				c.SetCookie("token", u.Tokens.AccessToken, 1.577e+7, "", "webkodes.com", true, true)
				c.SetCookie("isLoggedIn", "true", 1.577e+7, "", "webkodes.com", true, false)
				c.JSON(http.StatusOK, gin.H{
					"status":  "success",
					"message": "User logged in",
					"tokens": gin.H{
						"access_token": u.Tokens.AccessToken,
						"validity":     u.Tokens.Validity,
					},
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  "failed",
					"type":    "not_updated_tokens",
					"message": "Unable to update tokens",
				})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"type":    "no_record",
				"message": "No users found with the id",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"type":    "data_not_received",
			"message": "Data not received",
		})
	}

}
