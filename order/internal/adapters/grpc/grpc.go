package grpc

import (
	"context"

	"github.com/garren/microservices-proto/golang/order"
	"github.com/garren/microservices/order/internal/application/core/domain"
)

// defines handlers

// accepts an order request from the client, converts it to an order domain model, places order via api
func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
  var orderItems []domain.OrderItem
  for _, it := range request.OrderItems {
    orderItems = append(orderItems, domain.OrderItem{
      ProductCode: it.ProductCode,
      UnitPrice: it.UnitPrice,
      Quantity: it.Quantity,
    })
  }
  newOrder := domain.NewOrder(request.UserId, orderItems)
  result, err := a.api.PlaceOrder(newOrder)
  if err != nil {
    return nil, err
  }
  return &order.CreateOrderResponse{OrderId: result.ID}, nil
}
