package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Specs struct {
	Description    string `json:"description"`
	MaxVelocity    string `json:"maxVelocity"`
	NumberOfPeople string `json:"numberOfPeople"`
	FuelType       string `json:"fueltype"`
	ZeroToHundred  string `json:"zeroToHundred"`
	HorsePower     string `json:"horsePower"`
	Engine         string `json:"engine"`
}

type Car struct {
	ID       uuid.UUID `json:"id"`
	Brand    string    `json:"brand"`
	Model    string    `json:"model"`
	Price    string    `json:"price"`
	Specs    Specs     `json:"specs"`
}

var cars []Car

func main() {
	r := gin.Default()

	r.POST("/api/v1/register", RegisterCars)

	r.GET("/api/v1/cars", ListCars)

	r.PUT("/api/v1/cars/:id", UpdateCar)

	r.DELETE("/api/v1/cars/:id", DeleteCar)

	r.Run(":3030")
}

func RegisterCars(c *gin.Context) {
	var newCar Car

	if err := c.BindJSON(&newCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newID, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newCar.ID = newID

	cars = append(cars, newCar)

	c.JSON(http.StatusCreated, newCar)
}

func ListCars(c *gin.Context) {
	c.JSON(http.StatusOK, cars)
}

func UpdateCar(c *gin.Context) {
	var updatedCar Car

	carID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&updatedCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, car := range cars {
		if car.ID == carID {
			updatedCar.ID = carID
			cars[i] = updatedCar
			c.JSON(http.StatusOK, updatedCar)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "car not found"})
}

func DeleteCar(c *gin.Context) {
	carID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, car := range cars {
		if car.ID == carID {
			cars = append(cars[:i], cars[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"status": "success"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "car not found"})
}