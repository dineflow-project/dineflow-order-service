package services

import (
	"context"
	"errors"
	"time"

	"dineflow-order-service/models"
	"dineflow-order-service/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderServiceImpl struct {
	orderCollection *mongo.Collection
	ctx             context.Context
}

func NewOrderService(orderCollection *mongo.Collection, ctx context.Context) OrderService {
	return &OrderServiceImpl{orderCollection, ctx}
}

func (p *OrderServiceImpl) CreateOrder(order *models.CreateOrderRequest) (*models.DBOrder, error) {
	order.CreateAt = time.Now()
	order.UpdatedAt = order.CreateAt
	res, err := p.orderCollection.InsertOne(p.ctx, order)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("order with that title already exists")
		}
		return nil, err
	}

	opt := options.Index()
	opt.SetUnique(true)

	// index := mongo.IndexModel{Keys: bson.M{"title": 1}, Options: opt}

	// if _, err := p.orderCollection.Indexes().CreateOne(p.ctx, index); err != nil {
	// 	return nil, errors.New("could not create index for title")
	// }

	var newOrder *models.DBOrder
	query := bson.M{"_id": res.InsertedID}
	if err = p.orderCollection.FindOne(p.ctx, query).Decode(&newOrder); err != nil {
		return nil, err
	}

	return newOrder, nil
}

func (p *OrderServiceImpl) UpdateOrder(id string, data *models.UpdateOrder) (*models.DBOrder, error) {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.orderCollection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedOrder *models.DBOrder
	if err := res.Decode(&updatedOrder); err != nil {
		return nil, errors.New("no order with that Id exists")
	}

	return updatedOrder, nil
}

func (p *OrderServiceImpl) FindOrderById(id string) (*models.DBOrder, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var order *models.DBOrder

	if err := p.orderCollection.FindOne(p.ctx, query).Decode(&order); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return order, nil
}

func (p *OrderServiceImpl) FindOrders(page int, limit int) ([]*models.DBOrder, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	skip := (page - 1) * limit

	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))
	opt.SetSort(bson.M{"created_at": -1})

	query := bson.M{}

	cursor, err := p.orderCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(p.ctx)

	var orders []*models.DBOrder

	for cursor.Next(p.ctx) {
		order := &models.DBOrder{}
		err := cursor.Decode(order)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return []*models.DBOrder{}, nil
	}

	return orders, nil
}

func (p *OrderServiceImpl) DeleteOrder(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := p.orderCollection.DeleteOne(p.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}

func (p *OrderServiceImpl) FindOrdersByUserId(UserId string, page int, limit int) ([]*models.DBOrder, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	skip := (page - 1) * limit

	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))
	opt.SetSort(bson.M{"created_at": -1})

	query := bson.M{"user_id": UserId} // Filter by UserId

	cursor, err := p.orderCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(p.ctx)

	var orders []*models.DBOrder

	for cursor.Next(p.ctx) {
		order := &models.DBOrder{}
		err := cursor.Decode(order)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return []*models.DBOrder{}, nil
	}

	return orders, nil
}
