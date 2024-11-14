package main

import (
	"webclinicalrecords/internal"

	"github.com/gin-gonic/gin"
)

const BASEDIR = "./records"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	storage := internal.NewJSONStorage(BASEDIR)
	routes := internal.NewMyRoutes(storage)

	// Home
	r.GET("/", routes.HomeRoute)

	// GET and POST for inserting a new record
	r.GET("/insert", routes.InsertRecordRoute)
	r.POST("/insert", routes.InsertRecordRoute)

	// Get records of a specified patient
	r.GET("/patient", routes.GetPatientRecordsRoute)

	// Get record providing its ID
	r.GET("/record", routes.GetRecordRoute)

	r.Run("0.0.0.0:8080")
}
