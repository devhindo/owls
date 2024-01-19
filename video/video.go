package video

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleVideo(c *gin.Context) {

	// Return an immediate response to the client if necessary
	c.JSON(http.StatusOK, gin.H{"message": "Video request received"})
}