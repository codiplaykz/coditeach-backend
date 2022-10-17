package controllers

import (
	"coditeach/dao"
	"coditeach/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/mborders/logmatic"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const (
	SecretKey  = "Secret"
	authHeader = "Authorization"
	//userCtx    = "userId"
)

var userDAO = dao.UserDAO{Logger: logmatic.NewLogger()}
var logger = logmatic.NewLogger()

func SignUp(c *fiber.Ctx) error {
	//Get data of user
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	userToCheck := &models.User{
		Email: data["email"],
	}

	err = userDAO.GetByEmail(userToCheck)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to sign up.",
		})
	}

	if userToCheck.Id != 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User with current email already exist",
		})
	}

	//Crypting password
	pass, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	//Setting user data
	user := models.User{
		Name:     data["name"],
		Surname:  data["surname"],
		Login:    data["login"],
		Role:     data["role"],
		Email:    data["email"],
		Password: pass,
	}

	////Creating user row in DB
	err = userDAO.Create(&user)
	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to sign up.",
		})
	}

	result := fiber.Map{
		"userData": fiber.Map{
			"Id":      user.Id,
			"Login":   user.Login,
			"Role":    user.Role,
			"Name":    user.Name,
			"Surname": user.Surname,
			"Email":   user.Email,
		},
		"tokens": generateTokenPair(int(user.Id)),
	}

	//Return user
	c.Status(fiber.StatusCreated)
	return c.JSON(result)
}

func SignIn(c *fiber.Ctx) error {
	//Get data of user
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	var user models.User
	user.Email = data["email"]

	//Get user by email
	err = userDAO.GetByEmail(&user)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to sign in.",
		})
	}

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	//Check password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		logger.Error("ERROR while comparing passwords: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	result := fiber.Map{
		"userData": fiber.Map{
			"Id":      user.Id,
			"Login":   user.Login,
			"Role":    user.Role,
			"Name":    user.Name,
			"Surname": user.Surname,
			"Email":   user.Email,
		},
		"tokens": generateTokenPair(int(user.Id)),
	}

	//Return user
	c.Status(fiber.StatusOK)
	return c.JSON(result)
}

func Refresh(c *fiber.Ctx) error {
	//Get token
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	refreshToken := data["token"]

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, err := strconv.Atoi(claims["iss"].(string))
		if err != nil {
			return nil
		}
		c.Status(fiber.StatusOK)
		return c.JSON(generateTokenPair(id))
	}

	return c.SendStatus(fiber.StatusUnauthorized)
}

func UserIdentifyMiddleware(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	token := headers[authHeader]

	if token == "" {
		c.Status(401)
		return c.JSON(fiber.Map{"error": "empty auth header"})
	}

	isAuthorized, userId := CheckToken(token)

	if isAuthorized {
		c.Set("userId", string(userId))
	} else {
		c.Status(401)
		return c.JSON(fiber.Map{"error": "invalid auth header"})
	}

	return c.Next()
}

func AdminIdentifyMiddleware(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	token := headers[authHeader]

	if token == "" {
		c.Status(401)
		return c.JSON(fiber.Map{"error": "empty auth header"})
	}

	isAuthorized, userId := CheckToken(token)

	if isAuthorized {
		user := models.User{Id: uint(userId)}

		err := userDAO.GetById(&user)

		if err != nil {
			logger.Error("ERROR: %s", err)
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "Invalid user.",
			})
		}

		if user.Role != "admin" {
			c.Status(401)
			return c.JSON(fiber.Map{"error": "invalid auth user"})
		}
	} else {
		c.Status(401)
		return c.JSON(fiber.Map{"error": "invalid auth header"})
	}

	return c.Next()
}

func CheckToken(userToken string) (bool, int) {
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SecretKey), nil
	})

	if err != nil {
		return false, -1
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if _, ok := claims["iss"]; ok {
			id, err := strconv.Atoi(claims["iss"].(string))
			if err != nil {
				return false, -1
			}
			return true, id
		}
	}
	return false, -1
}

func generateTokenPair(userId int) fiber.Map {
	//Creating jwt token
	accessClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(userId),
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(), //30 min
	})

	refreshClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(userId),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	accessToken, err := accessClaims.SignedString([]byte(SecretKey))
	refreshToken, err := refreshClaims.SignedString([]byte(SecretKey))

	if err != nil {
		return fiber.Map{
			"Error": err,
		}
	}

	result := fiber.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}

	return result
}
