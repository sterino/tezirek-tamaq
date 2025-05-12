package order

import (
	"context"
	"errors"
	"order-service/internal/domain/order"
	"order-service/pkg/log"
	"order-service/pkg/store"

	"go.uber.org/zap"
)

// Create создает новый заказ
func (s *Service) Create(ctx context.Context, req order.Request) (order.Response, error) {
	logger := log.LoggerFromContext(ctx).Named("create_order").With(zap.String("user_id", req.UserID))

	entity := order.NewCreate(req)
	id, err := s.orderRepository.Create(ctx, entity)
	if err != nil {
		logger.Error("failed to create order", zap.Error(err))
		return order.Response{}, err
	}
	entity.ID = id

	return order.ParseFromEntity(entity), nil
}

// GetByID возвращает заказ по его ID
func (s *Service) GetByID(ctx context.Context, id string) (order.Response, error) {
	logger := log.LoggerFromContext(ctx).Named("get_order").With(zap.String("id", id))

	entity, err := s.orderRepository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, store.ErrorNotFound) {
			logger.Warn("order not found", zap.Error(err))
			return order.Response{}, err
		}
		logger.Error("failed to get order", zap.Error(err))
		return order.Response{}, err
	}

	return order.ParseFromEntity(entity), nil
}

// Update обновляет заказ
func (s *Service) Update(ctx context.Context, id string, req order.Request) error {
	logger := log.LoggerFromContext(ctx).Named("update_order").With(zap.String("id", id))

	entity := order.NewUpdate(req)
	if err := s.orderRepository.Update(ctx, id, entity); err != nil {
		if errors.Is(err, store.ErrorNotFound) {
			logger.Warn("order not found", zap.Error(err))
			return err
		}
		logger.Error("failed to update order", zap.Error(err))
		return err
	}

	return nil
}

// Delete удаляет заказ
func (s *Service) Delete(ctx context.Context, id string) error {
	logger := log.LoggerFromContext(ctx).Named("delete_order").With(zap.String("id", id))

	if err := s.orderRepository.Delete(ctx, id); err != nil {
		if errors.Is(err, store.ErrorNotFound) {
			logger.Warn("order not found", zap.Error(err))
			return err
		}
		logger.Error("failed to delete order", zap.Error(err))
		return err
	}

	return nil
}

// List возвращает все заказы
func (s *Service) List(ctx context.Context) ([]order.Response, error) {
	logger := log.LoggerFromContext(ctx).Named("list_orders")

	entities, err := s.orderRepository.List(ctx)
	if err != nil {
		logger.Error("failed to list orders", zap.Error(err))
		return nil, err
	}

	return order.ParseFromEntities(entities), nil
}

// ListByUser возвращает заказы пользователя
func (s *Service) ListByUser(ctx context.Context, userID string) ([]order.Response, error) {
	logger := log.LoggerFromContext(ctx).Named("list_orders_by_user").With(zap.String("user_id", userID))

	entities, err := s.orderRepository.ListByUserID(ctx, userID)
	if err != nil {
		logger.Error("failed to list user orders", zap.Error(err))
		return nil, err
	}

	return order.ParseFromEntities(entities), nil
}

// ListByRestaurant возвращает заказы по ресторану
func (s *Service) ListByRestaurant(ctx context.Context, restaurantID string) ([]order.Response, error) {
	logger := log.LoggerFromContext(ctx).Named("list_orders_by_restaurant").With(zap.String("restaurant_id", restaurantID))

	entities, err := s.orderRepository.ListByRestaurantID(ctx, restaurantID)
	if err != nil {
		logger.Error("failed to list restaurant orders", zap.Error(err))
		return nil, err
	}

	return order.ParseFromEntities(entities), nil
}
