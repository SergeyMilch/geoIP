package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeoLocation struct {
	IP           string  `json:"query"`
	Country      string  `json:"country"`
	Region       string  `json:"regionName"`
	City         string  `json:"city"`
	Zip          string  `json:"zip"`
	Latitude     float64 `json:"lat"`
	Longitude    float64 `json:"lon"`
	ISP          string  `json:"isp"`
	Organization string  `json:"org"`
}

func GetGeoLocation(c *gin.Context) {
	ipAddress := c.Param("ip")

	url := fmt.Sprintf("http://ip-api.com/json/%s", ipAddress)
	response, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve location information"})
		return
	}
	defer response.Body.Close()

	var geoInfo GeoLocation

	err = json.NewDecoder(response.Body).Decode(&geoInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
		return
	}

	c.JSON(http.StatusOK, geoInfo)
}

func main() {
	r := gin.Default()

	r.GET("/geolocation/:ip", GetGeoLocation)

	r.Run(":8000")
}
