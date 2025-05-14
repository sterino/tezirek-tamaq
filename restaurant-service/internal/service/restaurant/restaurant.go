package restaurant

import (
	"context"
	"errors"
	"restaurant-service/internal/domain/restaurant"
	"restaurant-service/pkg/log"
	"restaurant-service/pkg/store"

	"go.uber.org/zap"
)

func (s *Service) Create(ctx context.Context, req restaurant.Request) (restaurant.Response, error) {
	logger := log.LoggerFromContext(ctx).Named("create_restaurant").With(zap.String("name", req.Name))

	entity := restaurant.NewCreate(req)
	id, err := s.restaurantRepository.Create(ctx, entity)
	if err != nil {
		logger.Error("failed to create restaurant", zap.Error(err))
		return restaurant.Response{}, err
	}
	entity.ID = id

	return restaurant.ParseFromEntity(entity), nil
}

func (s *Service) GetByID(ctx context.Context, id string) (restaurant.Response, error) {
	logger := log.LoggerFromContext(ctx).Named("get_restaurant").With(zap.String("id", id))

	entity, err := s.restaurantRepository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, store.ErrorNotFound) {
			logger.Warn("restaurant not found", zap.Error(err))
			return restaurant.Response{}, err
		}
		logger.Error("failed to get restaurant", zap.Error(err))
		return restaurant.Response{}, err
	}

	return restaurant.ParseFromEntity(entity), nil
}

func (s *Service) Update(ctx context.Context, id string, req restaurant.Request) error {
	logger := log.LoggerFromContext(ctx).Named("update_restaurant").With(zap.String("id", id))

	entity := restaurant.NewUpdate(req)
	if err := s.restaurantRepository.Update(ctx, id, entity); err != nil {
		if errors.Is(err, store.ErrorNotFound) {
			logger.Warn("restaurant not found", zap.Error(err))
			return err
		}
		logger.Error("failed to update restaurant", zap.Error(err))
		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	logger := log.LoggerFromContext(ctx).Named("delete_restaurant").With(zap.String("id", id))

	if err := s.restaurantRepository.Delete(ctx, id); err != nil {
		if errors.Is(err, store.ErrorNotFound) {
			logger.Warn("restaurant not found", zap.Error(err))
			return err
		}
		logger.Error("failed to delete restaurant", zap.Error(err))
		return err
	}

	return nil
}

func (s *Service) List(ctx context.Context) ([]restaurant.Response, error) {
	logger := log.LoggerFromContext(ctx).Named("list_restaurants")

	entities, err := s.restaurantRepository.List(ctx)
	if err != nil {
		logger.Error("failed to list restaurants", zap.Error(err))
		return nil, err
	}

	return restaurant.ParseFromEntities(entities), nil
}
