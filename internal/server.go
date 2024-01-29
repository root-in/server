package internal

import (
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
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("get user failed with %v", err))
		return
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("can't read file for user %v", *user))
		return
	}

	f, err := fileHeader.Open()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("can't open file for user %v", *user))
		return
	}

	fileName := user.getId()

	if err := storeGCS(f, bucket, fileName); err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to store file for user %v with %v", *user, err))
		return
	}

	if err := saveUser(user, fileName); err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to save user %v with %v", *user, err))
		return
	}

	fmt.Printf("Registered: %v\n", *user)
}

func getUser(c *gin.Context) (*User, error) {
	var user User
	var err error

	if user.Name, err = getMandatoryField(c, "name"); err != nil {
		return nil, err
	}

	if user.Surname, err = getMandatoryField(c, "surname"); err != nil {
		return nil, err
	}

	user.Email = getOptionalField(c, "email")

	if user.Phone, err = getMandatoryField(c, "phone"); err != nil {
		return nil, err
	}

	return &user, nil
}

func getMandatoryField(c *gin.Context, name string) (string, error) {
	result := c.PostForm(name)
	if result == "" {
		return "", fmt.Errorf("field %q is missing", name)
	}
	return result, nil
}

func getOptionalField(c *gin.Context, name string) string {
	return c.PostForm(name)
}
