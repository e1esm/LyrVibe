package service

import (
	"context"
	"github.com/e1esm/LyrVibe/artist-service/pkg/logger"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"go.uber.org/zap"
)

type RolesProvider interface {
	UpdateRole(context.Context, string, string) (*proto.UpdatingRoleResponse, error)
}

type RoleService struct {
	conn proto.AuthServiceClient
}

func NewRolesService(conn proto.AuthServiceClient) RolesProvider {
	return &RoleService{conn: conn}
}

func (rs *RoleService) UpdateRole(ctx context.Context, userID string, role string) (*proto.UpdatingRoleResponse, error) {
	resp, err := rs.conn.UpdateRole(ctx, &proto.UpdatingRoleRequest{
		UserId:        userID,
		RequestedRole: role,
	})
	if err != nil {
		logger.GetLogger().Error("RoleService:UpdateRole", zap.String("", err.Error()))
		return nil, err
	}
	return resp, nil
}
