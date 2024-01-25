package internal

import (
	"embed"
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed www/*.html
var content embed.FS

func staticHandler(c *gin.Context) {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(content))))
}
