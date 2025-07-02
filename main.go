package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type entity struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Sales int    `json:"sales"`
}

var entityList = []entity{
	{Id: "1", Name: "Bloodborne", Sales: 2000},
	{Id: "2", Name: "Dark Souls 3", Sales: 28000},
	{Id: "3", Name: "Elden Ring", Sales: 15000},
}

func getEntityList(context *gin.Context) {
	context.JSON(http.StatusOK, entityList)
}

func addEntity(context *gin.Context) {
	var newEnt entity

	if err := context.BindJSON(&newEnt); err != nil {
		return
	}

	entityList = append(entityList, newEnt)
	context.JSON(http.StatusCreated, newEnt)

}

func getEntityById(id string) (*entity, error) {
	for i, j := range entityList {
		if j.Id == id {
			return &entityList[i], nil
		}
	}

	return nil, errors.New("entity not found")
}

func getEntity(context *gin.Context) {
	id := context.Param("id")
	entity, err := getEntityById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "entitiy not found"})
		return
	}

	context.JSON(http.StatusOK, entity)
}

/*
func toggleEntityStatus(context *gin.Context){
	id := context.Param("id")
	entity, err := getEntityById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "entitiy not found"})
		return
	}

	entity.
}
*/

func main() {
	router := gin.Default()
	router.GET("/entityList", getEntityList)
	router.GET("/entityList/:id", getEntity)
	router.PATCH("/entityList/:id", getEntity)
	router.POST("/entityList", addEntity)
	router.Run("localhost:8080")
}
