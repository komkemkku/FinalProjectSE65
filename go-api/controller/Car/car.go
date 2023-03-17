package car

import (
	"net/http"
	"se/jwt-api/orm"
	"github.com/gin-gonic/gin"
)
func CarAll(c *gin.Context) {
	var car []orm.Car
	orm.Db.Find(&car)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Car Read Sucessful","cars": car})
}
// สร้าง structure เพื่อรองรับ json
type CarBody struct {
	Carname string
	Detail string
	Image string
}
func RegisterCar(c *gin.Context) {
	var json CarBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
// เก็บค่าลงใส่ Database
	car := orm.Car{Carname: json.Carname, Detail: json.Detail,Image: json.Image}
orm.Db.Create(&car)
	if car.ID > 0 {
		c.JSON(200, gin.H{"status": "ok", "message": "Car Create Sucessful","Carname": car.ID})
	} else {
		c.JSON(200, gin.H{"status": "error", "message": "Car Register Fail"})
	}
}
