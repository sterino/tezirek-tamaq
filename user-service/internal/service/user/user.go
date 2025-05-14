package user

import (
	"context"
	"errors"
	"user-service/internal/domain/user"
	"user-service/pkg/log"
	"user-service/pkg/store"

	"go.uber.org/zap"
)

func (s *Service) GetByID(ctx context.Context, id string) (user.Response, error) {
	logger := log.LoggerFromContext(ctx).Named("get_user").With(zap.String("id", id))

	entity, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, store.ErrorNotFound) {
			logger.Warn("order not found", zap.Error(err))
			return user.Response{}, err
		}
		logger.Error("failed to get order", zap.Error(err))
		return user.Response{}, err
	}

	return user.ParseFromEntity(entity), nil
}

func (s *Service) Update(ctx context.Context, id string, req user.Request) error {
	logger := log.LoggerFromContext(ctx).Named("update_user").With(zap.String("id", id), zap.Any("req", req))

	entity := user.NewUpdate(req)
	if err := s.userRepository.Update(ctx, id, entity); err != nil {
		if errors.Is(err, store.ErrorNotFound) {
			logger.Warn("order not found", zap.Error(err))
			return err
		}
		logger.Error("failed to update order", zap.Error(err))
		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	logger := log.LoggerFromContext(ctx).Named("delete_user").With(zap.String("id", id))

	if err := s.userRepository.Delete(ctx, id); err != nil {
		if errors.Is(err, store.ErrorNotFound) {
			logger.Warn("order not found", zap.Error(err))
			return err
		}
		logger.Error("failed to delete order", zap.Error(err))
		return err
	}

	return nil
}

func (s *Service) List(ctx context.Context) ([]user.Response, error) {
	logger := log.LoggerFromContext(ctx).Named("list_users")

	entities, err := s.userRepository.List(ctx)
	if err != nil {
		logger.Error("failed to list users", zap.Error(err))
		return nil, err
	}

	return user.ParseFromEntities(entities), nil
}
