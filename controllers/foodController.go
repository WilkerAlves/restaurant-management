package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/WilkerAlves/restaurant-management/database"
	"github.com/WilkerAlves/restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		foodId := c.Param("food_id")
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var food models.Food
		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "error occured while fetching the food item"},
			)
		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var food models.Food
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		var menu models.Menu
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.MenuId}).Decode(&menu)
		if err != nil {
			msg := fmt.Sprintf("menu was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		food.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Id = primitive.NewObjectID()
		food.FoodId = food.Id.Hex()
		food.Price = toFixed(*food.Price, 2)

		result, err := foodCollection.InsertOne(ctx, food)
		if err != nil {
			msg := fmt.Sprintf("Food item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, result)

	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// func round(num float64) int {

// }

func toFixed(num float64, precision int) *float64 {
	temp := 0.0
	return &temp
}
