package controllers

import (
	"fmt"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/xiaosasori/go-fiber/db"
	"github.com/xiaosasori/go-fiber/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetUser get user with id
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	query := bson.D{{Key: "_id", Value: id}}
	db.Repo.Db.Collection("users").FindOne(c.Context(), query).Decode(&user)
	fmt.Print(user)
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

// UpdateUser update an user
func UpdateUser(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on update user request", "data": err})
	}
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["user_id"].(string)

	// Upsert to insert a new document if a document matching the filter isn't found
	opts := options.FindOneAndUpdate().SetUpsert(true)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid Objectid", "data": nil})
	}
	filter := bson.D{{Key: "_id", Value: objectID}}
	update := bson.D{{"$set", bson.D{{Key: "email", Value: data["email"]}}}}
	var updatedUser models.User
	db.Repo.Db.Collection("users").FindOneAndUpdate(c.Context(), filter, update, opts).Decode(&updatedUser)
	return c.JSON(fiber.Map{"status": "success", "message": "User successfully updated", "data": updatedUser})
}
