package controllers

import (
	"os"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/xiaosasori/go-fiber/db"
	"github.com/xiaosasori/go-fiber/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Home func
func Home(c *fiber.Ctx) error {
	return c.SendString("hello")
}

var validate = validator.New()

//HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	return err == nil
}

// Register func
func Register(c *fiber.Ctx) error {
	var data models.User
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on register request", "data": err})
	}
	// validate fields
	validationErr := validate.Struct(data)
	if validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on validate register request", "data": validationErr})
	}
	// check user exists
	count, err := db.Repo.Db.Collection("users").CountDocuments(c.Context(), bson.M{"email": data.Email})
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "User already existed"})
	}
	// hash password
	password, err := HashPassword(data.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}
	user := models.User{
		UserName: data.UserName,
		Email:    data.Email,
		Password: password,
	}
	db.Repo.Db.Collection("users").InsertOne(c.Context(), user)
	return c.JSON(fiber.Map{"status": "success", "message": "Success register", "data": user})
}

// Login func
func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	// query user in db
	var foundUser models.User
	query := bson.D{{Key: "email", Value: data["email"]}}
	errNoDoc := db.Repo.Db.Collection("users").FindOne(c.Context(), query).Decode(&foundUser)
	if errNoDoc != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not exist", "data": nil})
	}
	// verify password
	if !VerifyPassword(foundUser.Password, data["password"]) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}
	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = foundUser.UserName
	claims["user_id"] = foundUser.ID
	claims["exp"] = time.Now().Add(time.Second * 60).Unix()
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
