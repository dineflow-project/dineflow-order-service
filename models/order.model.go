package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateOrderRequest struct {
	Status    string    `json:"status" bson:"status" binding:"required"`
	MenuId    string    `json:"menu_id" bson:"menu_id" binding:"required"`
	VendorId  string    `json:"vendor_id" bson:"vendor_id" binding:"required"`
	Price     float32   `json:"price" bson:"price" binding:"required"`
	Request   string    `json:"request,omitempty" bson:"request,omitempty"` // Use a pointer for optional field
	UserId    string    `json:"user_id" bson:"user_id" binding:"required"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBOrder struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Status    string             `json:"status,omitempty" bson:"status,omitempty"`
	MenuId    string             `json:"menu_id,omitempty" bson:"menu_id,omitempty"`
	VendorId  string             `json:"vendor_id,omitempty" bson:"vendor_id,omitempty"`
	Price     float32            `json:"price,omitempty" bson:"price,omitempty"`
	Request   string             `json:"request,omitempty" bson:"request,omitempty"`
	UserId    string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	CreateAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateOrder struct {
	Status    string    `json:"status,omitempty" bson:"status,omitempty"`
	MenuId    string    `json:"menu_id,omitempty" bson:"menu_id,omitempty"`
	VendorId  string    `json:"vendor_id,omitempty" bson:"vendor_id,omitempty"`
	Price     float32   `json:"price,omitempty" bson:"price,omitempty"`
	Request   string    `json:"request,omitempty" bson:"request,omitempty"` // Use a pointer for optional field
	UserId    string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
