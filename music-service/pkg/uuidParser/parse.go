package uuidParser

import (
	"github.com/e1esm/LyrVibe/music-service/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func ParseUUID(id string) uuid.UUID {
	respID, err := uuid.Parse(id)
	if err != nil {
		logger.GetLogger().Error("Couldn't have extracted ID from string",
			zap.String("err", err.Error()))
		return uuid.UUID{}
	}
	return respID
}
