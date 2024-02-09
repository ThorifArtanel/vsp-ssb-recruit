package apis

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/marcboeker/go-duckdb"
)

type SignUpForm struct {
	Name           string `form:"name"`
	Email          string `form:"email"`
	Phone          string `form:"phone"`
	Age            string `form:"age"`
	Religion       string `form:"religion"`
	Gender         string `form:"gender"`
	Degree         string `form:"degree"`
	CampusName     string `form:"campus_name"`
	CityBirthplace string `form:"city_birthplace"`
	CityResidency  string `form:"city_residency"`
}

func FormHandler(c *gin.Context) {
	log.Print("INFO: Parsing Payload")
	var req SignUpForm

	if err := c.Bind(&req); err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "Bad Request",
			"msg":    "The request could not be understood by the server due to incorrect syntax."})
		return
	}

	db, err := sql.Open("duckdb", "")
	if err != nil {
		fmt.Printf("db unable to connect: %+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "failure",
		})
		return
	}
	defer db.Close()

	_, err = db.Exec(`
		SET autoinstall_known_extensions=1;
		SET autoload_known_extensions=1;
		SET s3_endpoint = 'storage.googleapis.com';
		SET s3_access_key_id = '` + os.Getenv("KEYID") + `';
		SET s3_secret_access_key = '` + os.Getenv("KEYSECRET") + `';`)
	if err != nil {
		fmt.Printf("failed setting credentials: %+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "failure",
		})
		return
	}

	_, err = db.Exec(`
		CREATE TABLE datum(
			name varchar(100),
			email varchar(100),
			phone varchar(100),
			age varchar(100),
			religion varchar(100),
			gender varchar(100),
			degree varchar(100),
			campus_name varchar(100),
			city_birthplace varchar(100),
			city_residency varchar(100)
		)`)
	if err != nil {
		fmt.Printf("failed creating table: %+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "failure",
		})
		return
	}

	_, err = db.Exec(`
		INSERT INTO datum(
			name,
			email,
			phone,
			age,
			religion,
			gender,
			degree,
			campus_name,
			city_birthplace,
			city_residency
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`, req.Name, req.Email, req.Phone, req.Age, req.Religion, req.Gender, req.Degree, req.CampusName, req.CityBirthplace, req.CityResidency)
	if err != nil {
		fmt.Printf("failed inserting data to table: %+v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "failure",
		})
		return
	}

	id := uuid.New()
	_, err = db.Exec(`
		COPY datum TO 's3://` + os.Getenv("BUCKET_NAME") + `/` + id.String() + `.parquet'`)
	if err != nil {
		fmt.Printf("failed uploading to S3: %+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "failure",
		})
		return
	}
}
