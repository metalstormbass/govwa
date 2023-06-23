package setup

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/govwa/util/database"
)

const (
	DropUsersTable = `DROP TABLE IF EXISTS Users`

	CreateUsersTable = `CREATE TABLE Users (
		id int(10) NOT NULL AUTO_INCREMENT,
		uname varchar(100) NOT NULL,
		pass varchar(100) NOT NULL,
		PRIMARY KEY (id)
	  ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1`

	InsertUsers = `INSERT INTO Users VALUES (1,'admin','9f3b6fa4703a5ba96fda0dee48ec76fc'),(2,'user1','ff1d5c0015a535b01a5d03a373bf06f6')`

	DropProfilesTable = `DROP TABLE IF EXISTS Profile`

	CreateProfilesTable = `CREATE TABLE Profile (
		profile_id int(10) NOT NULL AUTO_INCREMENT,
		user_id int(10) NOT NULL,
		full_name varchar(100) NOT NULL,
		city varchar(100) NOT NULL,
		phone_number varchar(15) NOT NULL,
		PRIMARY KEY (profile_id)
	  ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1`

	InsertProfile = `INSERT INTO Profile VALUES (1,1,'Andro','Jakarta','08882112345'),(2,2,'Rocky','Bandung','08882112345')`
)

var DB *sql.DB
var err error

/*func init() {
	DB, err = database.Connect()
	if err != nil {
		log.Println(err.Error())
	}
}*/

func createUsersTable() error {

	DB, err = database.Connect()

	_, err = DB.Exec(DropUsersTable)
	if err != nil {
		return err
	}
	_, err = DB.Exec(CreateUsersTable)
	if err != nil {
		return err
	}
	_, err = DB.Exec(InsertUsers)
	if err != nil {
		return err
	}
	return nil
}

func createProfileTable() error {
	_, err = DB.Exec(DropProfilesTable)
	if err != nil {
		return err
	}
	_, err = DB.Exec(CreateProfilesTable)
	if err != nil {
		return err
	}
	_, err = DB.Exec(InsertProfile)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	router := gin.Default()

	// Endpoint for uploading a file
	router.POST("/upload", func(c *gin.Context) {
		//file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save the uploaded file to a desired location
		//err = c.SaveUploadedFile(file, "uploads/"+file.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully!"})
	})

	router.Run(":8080")
}






func bad() {
	r := gin.Default()
	r.POST("/upload", handleUpload)
	r.Run(":8080")
}

func handleUpload(c *gin.Context) {
	// Dummy file
	file := &multipart.FileHeader{
		Filename: "example.txt",
		Size:     1024,
		//Header:   make(http.Header),
	}

	err := saveFile(file)
	if err != nil {
		log.Println("Error saving file: ", err)
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.String(http.StatusOK, "File uploaded successfully")
}

func saveFile(file *multipart.FileHeader) error {
	// Dummy source file
	src, err := os.Open("path/to/source/file.txt")
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer src.Close()

	// Dummy destination file
	dst, err := os.Create("path/to/destination/file.txt")
	if err != nil {
		return fmt.Errorf("error creating destination file: %w", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	return nil
}










