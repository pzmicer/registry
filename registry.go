package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/jmoiron/sqlx"

	"github.com/joho/godotenv"
)

type Service struct {
	ID       int    `db:"id" json:"id,omitempty"`
	Name     string `db:"name" json:"name"`
	Cost     string `db:"cost" json:"cost"`
	Duration int    `db:"duration" json:"duration"`
	Currency string `db:"currency" json:"currency"`
	Url      string `db:"url" json:"url"`
	Key      string `db:"key" json:"key"`
}

type Method struct {
	ID        int `db:"id" json:"id,omitempty"`
	Name      int `db:"method_name" json:"method_name"`
	ServiceId int `db:"service_id" json:"service_id"`
}

func getConnection() *sqlx.DB {
	godotenv.Load()
	db, err := sqlx.Connect("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func checkService(c *gin.Context) {
	serviceName := c.Query("name")

	db := getConnection()

	rows, err := db.Queryx("SELECT * FROM services")
	if err != nil {
		fmt.Println(err)
		return
	}

	service := Service{}
	for rows.Next() {
		err := rows.StructScan(&service)
		if err != nil {
			log.Fatalln(err)
		}
		if service.Name == serviceName {
			c.JSON(http.StatusOK, gin.H{
				"result": true,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"result:": false,
	})
}

func getServiceInfo(c *gin.Context) {
	serviceName := c.Query("name")

	db := getConnection()

	service := Service{}
	err := db.Get(&service, "SELECT * FROM services WHERE name = $1", serviceName)
	if err != nil {
		log.Fatalln(err)
	}

	service.ID = 0
	c.JSON(http.StatusOK, service)
}

func getServiceList(c *gin.Context) {
	db := getConnection()

	services := []Service{}
	err := db.Select(&services, "SELECT (name, cost, duration, currency, url, key) FROM services")
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, services)
}

func addService(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var newService Service
	err := json.Unmarshal(body, &newService)
	if err != nil {
		log.Fatalln(err)
	}

	db := getConnection()

	db.NamedExec(`INSERT INTO services (name, cost, duration, currency, url, key) 
		VALUES (:name, :cost, :duration, :currency, :url, :key)`, newService)

	c.Status(http.StatusOK)
}

func main() {
	router := gin.Default()
	router.GET("/checkService", checkService)
	router.GET("/getServiceInfo", getServiceInfo)
	router.GET("/getServiceList", getServiceList)
	router.POST("/addService", addService)

	router.Run(":8080")
}
