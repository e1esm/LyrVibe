package server

import (
	"fmt"
	"github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	badRequest          = "Bad request: %s"
	verificationFailed  = "Verification failed for user: %s"
	verificationSucceed = "Verification succeed for user: %s"
)

func (ps *ProxyServer) NewArtist(c *gin.Context) {
	verificationRequest := proto.VerificationRequest{}
	verificationRequest.Username = c.GetString("username")
	if err := c.BindJSON(&verificationRequest); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf(badRequest, err.Error()))
		return
	}
	verificationRequest.Id = c.GetString("id")
	resp, err := ps.Services.ArtistService.New(&verificationRequest)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(verificationFailed, verificationRequest.Username))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":       resp.RequestStatus,
		"verification": fmt.Sprintf(verificationSucceed, verificationRequest.Username),
	})
}

func (ps *ProxyServer) ReleaseTrack(c *gin.Context) {
	releaseRequest := proto.NewTrackRequest{}

}
