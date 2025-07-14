package grpc

import (
	"context"
	"github.com/Sapronovps/RotationBanner/internal/model"
	internalgrpcprotobuf "github.com/Sapronovps/RotationBanner/internal/server/grpc/protobuf"
	"go.uber.org/zap"
	"time"
)

func (s *BannerGrpcServer) AddSlot(_ context.Context, req *internalgrpcprotobuf.RequestAddSlot) (*internalgrpcprotobuf.ResponseSlot, error) {
	var slot model.Slot
	slot.Description = req.GetDescription()

	err := s.app.AddSlot(&slot)
	if err != nil {
		s.logger.Error("failed to add slot", zap.Error(err))
		return nil, err
	}

	response := &internalgrpcprotobuf.ResponseSlot{
		ID:          int64(slot.ID),
		Description: slot.Description,
		CreatedAt:   slot.CreatedAt.Format(time.RFC3339Nano),
	}

	return response, nil
}

func (s *BannerGrpcServer) GetSlot(_ context.Context, req *internalgrpcprotobuf.RequestGetSlot) (*internalgrpcprotobuf.ResponseSlot, error) {
	slot, err := s.app.GetSlot(int(req.GetID()))
	if err != nil {
		s.logger.Error("failed to get slot", zap.Error(err))
		return nil, err
	}

	response := &internalgrpcprotobuf.ResponseSlot{
		ID:          int64(slot.ID),
		Description: slot.Description,
		CreatedAt:   slot.CreatedAt.Format(time.RFC3339Nano),
	}

	return response, nil
}

func (s *BannerGrpcServer) AddBanner(_ context.Context, req *internalgrpcprotobuf.RequestAddBanner) (*internalgrpcprotobuf.ResponseBanner, error) {
	var banner model.Banner
	banner.Title = req.GetTitle()
	banner.Description = req.GetDescription()

	err := s.app.AddBanner(&banner)
	if err != nil {
		s.logger.Error("failed to add banner", zap.Error(err))
		return nil, err
	}

	response := &internalgrpcprotobuf.ResponseBanner{
		ID:          int64(banner.ID),
		Title:       banner.Title,
		Description: banner.Description,
		CreatedAt:   banner.CreatedAt.Format(time.RFC3339Nano),
	}

	return response, nil
}

func (s *BannerGrpcServer) GetBanner(_ context.Context, req *internalgrpcprotobuf.RequestGetBanner) (*internalgrpcprotobuf.ResponseBanner, error) {
	banner, err := s.app.GetBanner(int(req.GetID()))
	if err != nil {
		s.logger.Error("failed to get banner", zap.Error(err))
		return nil, err
	}

	response := &internalgrpcprotobuf.ResponseBanner{
		ID:          int64(banner.ID),
		Title:       banner.Title,
		Description: banner.Description,
		CreatedAt:   banner.CreatedAt.Format(time.RFC3339Nano),
	}

	return response, nil
}

func (s *BannerGrpcServer) AddGroup(_ context.Context, req *internalgrpcprotobuf.RequestAddGroup) (*internalgrpcprotobuf.ResponseGroup, error) {
	var group model.Group
	group.Title = req.GetTitle()
	group.Description = req.GetDescription()

	err := s.app.AddGroup(&group)
	if err != nil {
		s.logger.Error("failed to add group", zap.Error(err))
		return nil, err
	}

	response := &internalgrpcprotobuf.ResponseGroup{
		ID:          int64(group.ID),
		Title:       group.Title,
		Description: group.Description,
		CreatedAt:   group.CreatedAt.Format(time.RFC3339Nano),
	}

	return response, nil
}

func (s *BannerGrpcServer) GetGroup(_ context.Context, req *internalgrpcprotobuf.RequestGetGroup) (*internalgrpcprotobuf.ResponseGroup, error) {
	group, err := s.app.GetGroup(int(req.GetID()))
	if err != nil {
		s.logger.Error("failed to get banner", zap.Error(err))
		return nil, err
	}

	response := &internalgrpcprotobuf.ResponseGroup{
		ID:          int64(group.ID),
		Title:       group.Title,
		Description: group.Description,
		CreatedAt:   group.CreatedAt.Format(time.RFC3339Nano),
	}

	return response, nil
}

func (s *BannerGrpcServer) AddBannerGroupStats(_ context.Context, req *internalgrpcprotobuf.RequestAddBannerGroupStats) (*internalgrpcprotobuf.ResponseBannerGroupStats, error) {
	var bannerGroupStats model.BannerGroupStats
	bannerGroupStats.SlotID = int(req.GetSlotID())
	bannerGroupStats.BannerID = int(req.GetBannerID())
	bannerGroupStats.GroupID = int(req.GetGroupID())

	err := s.app.AddBannerGroupStats(&bannerGroupStats)
	if err != nil {
		s.logger.Error("failed to add banner group stats", zap.Error(err))
		return nil, err
	}

	response := &internalgrpcprotobuf.ResponseBannerGroupStats{
		ID:        int64(bannerGroupStats.ID),
		SlotID:    int64(bannerGroupStats.SlotID),
		BannerID:  int64(bannerGroupStats.BannerID),
		GroupID:   int64(bannerGroupStats.GroupID),
		Shows:     int64(bannerGroupStats.Shows),
		Clicks:    int64(bannerGroupStats.Clicks),
		CreatedAt: bannerGroupStats.CreatedAt.Format(time.RFC3339Nano),
		UpdatedAt: bannerGroupStats.UpdatedAt.Format(time.RFC3339Nano),
	}

	return response, nil
}

func (s *BannerGrpcServer) GetBannerGroupStats(_ context.Context, req *internalgrpcprotobuf.RequestGetBannerGroupStats) (*internalgrpcprotobuf.ResponseBannerGroupStats, error) {
	bannerGroupStats, err := s.app.GetBannerGroupStats(int(req.GetSlotID()), int(req.GetBannerID()), int(req.GetGroupID()))
	if err != nil {
		s.logger.Error("failed to get banner group stats", zap.Error(err))
		return nil, err
	}

	response := &internalgrpcprotobuf.ResponseBannerGroupStats{
		ID:        int64(bannerGroupStats.ID),
		SlotID:    int64(bannerGroupStats.SlotID),
		BannerID:  int64(bannerGroupStats.BannerID),
		GroupID:   int64(bannerGroupStats.GroupID),
		Shows:     int64(bannerGroupStats.Shows),
		Clicks:    int64(bannerGroupStats.Clicks),
		CreatedAt: bannerGroupStats.CreatedAt.Format(time.RFC3339Nano),
		UpdatedAt: bannerGroupStats.UpdatedAt.Format(time.RFC3339Nano),
	}

	return response, nil
}

func (s *BannerGrpcServer) RegisterClick(_ context.Context, req *internalgrpcprotobuf.RequestRegisterClick) (*internalgrpcprotobuf.ResponseRegisterClick, error) {
	err := s.app.RegisterClick(int(req.GetSlotID()), int(req.GetBannerID()), int(req.GetGroupID()))
	if err != nil {
		s.logger.Error("failed to register click", zap.Error(err))
		return nil, err
	}

	response := &internalgrpcprotobuf.ResponseRegisterClick{
		BannerID: req.BannerID,
	}

	return response, nil
}

func (s *BannerGrpcServer) GetBannerByMultiArmBandit(_ context.Context, req *internalgrpcprotobuf.RequestGetBannerByMultiArmBandit) (*internalgrpcprotobuf.ResponseBanner, error) {
	banner, err := s.app.GetBannerByMultiArmBandit(int(req.GetSlotID()), int(req.GetGroupID()))
	if err != nil {
		s.logger.Error("failed to get banner by arm bandit", zap.Error(err))
		return nil, err
	}

	response := &internalgrpcprotobuf.ResponseBanner{
		ID:          int64(banner.ID),
		Title:       banner.Title,
		Description: banner.Description,
		CreatedAt:   banner.CreatedAt.Format(time.RFC3339Nano),
	}

	return response, nil
}
