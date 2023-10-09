package server

import (
	"fmt"
	"github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	badRequest          = "Bad request: %s"
	verificationFailed  = "Verification failed for user: %s"
	verificationSucceed = "Verification succeed for user: %s"
	releaseError        = "Track cannot be released: %s"
	deleteError         = "Track cannot be deleted: %s"
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

	accessCookie := &http.Cookie{
		Name:     accessTokenName,
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Unix(0, 0),
	}
	http.SetCookie(c.Writer, accessCookie)

	refreshCookie := &http.Cookie{
		Name:     refreshTokenName,
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Unix(0, 0),
	}
	http.SetCookie(c.Writer, refreshCookie)
	c.JSON(http.StatusOK, gin.H{
		"status":       resp.RequestStatus.String(),
		"verification": fmt.Sprintf(verificationSucceed, verificationRequest.Username),
	})
}

func (ps *ProxyServer) ReleaseTrack(c *gin.Context) {
	releaseRequest := proto.NewTrackRequest{}
	releaseRequest.ArtistId = c.GetString("id")
	if err := c.BindJSON(&releaseRequest); err != nil {
		logger.GetLogger().Error("Bad request",
			zap.String("err", err.Error()))
		c.JSON(http.StatusBadRequest, fmt.Sprintf(badRequest, err.Error()))
		return
	}
	resp, err := ps.Services.ArtistService.ReleaseTrack(&releaseRequest)
	if err != nil {
		logger.GetLogger().Error("Track cannot be released", zap.String("err", err.Error()))
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(releaseError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title":  resp.Title,
		"status": resp.RequestStatus.String(),
	})
}

func (ps *ProxyServer) DeleteTrack(c *gin.Context) {
	deleteRequest := proto.DeleteTrackRequest{}
	if err := c.BindJSON(&deleteRequest); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf(badRequest, err.Error()))
		return
	}
	deleteRequest.AuthorId = c.GetString("id")
	resp, err := ps.Services.ArtistService.DeleteTrack(&deleteRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf(deleteError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title":  resp.Title,
		"status": resp.RequestStatus.String(),
	})
}

func (ps *ProxyServer) ReleaseAlbum(c *gin.Context) {

}
