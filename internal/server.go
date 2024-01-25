package internal

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.GET("/", staticHandler)
	r.POST("/register", registerHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func registerHandler(c *gin.Context) {
	var name, surname, email, phone string
	var err error

	if name, err = getField(c, "name"); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if surname, err = getField(c, "surname"); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if email, err = getField(c, "email"); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if phone, err = getField(c, "phone"); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("can't read file"))
		return
	}

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, "/tmp/file.Filename")

	fmt.Printf("name: %s; surname: %s; email: %s; phone: %s; file: %s", name, surname, email, phone, file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func getField(c *gin.Context, name string) (string, error) {
	result := c.PostForm(name)
	if result == "" {
		return "", fmt.Errorf("field %q is missing", name)
	}
	return result, nil
}
