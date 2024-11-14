package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyRoutes struct {
	storage Storage
}

func NewMyRoutes(storage Storage) *MyRoutes {
	return &MyRoutes{storage}
}

func (r *MyRoutes) HomeRoute(c *gin.Context) {

	c.HTML(200, "home.html", gin.H{
		"message": "This is the Home",
		"title":   "Home",
	})
}

func (r *MyRoutes) InsertRecordRoute(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{
		"message": "Inserting",
		"title":   "Insert Record",
	})
}

func (r *MyRoutes) GetPatientRecordsRoute(c *gin.Context) {
	records := r.storage.GetAllPatientRecords(c.Query("id"))

	if records == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.HTML(200, "home.html", gin.H{
		"message": records,
		"title":   "All Patient Records",
	})
}

func (r *MyRoutes) GetRecordRoute(c *gin.Context) {
	record := r.storage.GetRecord(c.Query("id"))

	if record == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.HTML(200, "home.html", gin.H{
		"message": record,
		"title":   "Single Record",
	})
}
