package internal

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(content fs.FS) {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.StaticFS("/", http.FS(content))
	r.POST("/register", registerHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func registerHandler(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("can't read file"))
		return
	}

	f, err := fileHeader.Open()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("can't open file"))
		return
	}

	if err := storeGCS(f, "rootin-web", user.getId()); err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to store file with %v", err))
	}

	if err := saveUser(user); err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to save user with %v", err))
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' registered", user))
}

func getUser(c *gin.Context) (*User, error) {
	var user User
	var err error

	if user.Name, err = getField(c, "name"); err != nil {
		return nil, c.AbortWithError(http.StatusBadRequest, err)
	}

	if user.Surname, err = getField(c, "surname"); err != nil {
		return nil, c.AbortWithError(http.StatusBadRequest, err)
	}

	if user.Email, err = getField(c, "email"); err != nil {
		return nil, c.AbortWithError(http.StatusBadRequest, err)
	}

	if user.Phone, err = getField(c, "phone"); err != nil {
		return nil, c.AbortWithError(http.StatusBadRequest, err)
	}

	return &user, nil
}

func getField(c *gin.Context, name string) (string, error) {
	result := c.PostForm(name)
	if result == "" {
		return "", fmt.Errorf("field %q is missing", name)
	}
	return result, nil
}
