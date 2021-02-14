package controllers

import (
	"fmt"

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

// Register func
func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		UserName: data["username"],
		Email:    data["email"],
		Password: password,
	}
	db.Repo.Db.Collection("users").InsertOne(c.Context(), user)
	return c.JSON(user)
}

// Login func
func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.User
	query := bson.D{{Key: "email", Value: data["email"]}}
	db.Repo.Db.Collection("users").FindOne(c.Context(), query).Decode(&user)
	fmt.Print(user)
	return nil
}
