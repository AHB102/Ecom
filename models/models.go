package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      *string            `json:"Name" validate:"required,min=2,max=30"`
	Email     *string            `json:"Email" validate:"required,min=2,max=30"`
	Password  *string            `json:"Password" validate:"required, min=8, max=16"`
	Phone     *string            `json:"Phone" validate:"required"`
	CreatedAT time.Time          `json:"Created At"`
	UserID    string             `json:"userID"`
	UserCart  []ProductUser      `json:"usercart" bson:"usercart"`
	//UserID      string        `json:"userID"`
	//UserCart    []ProductUser `json:"usercart" bson:"usercart"`
	OrderStatus []Order `json:"orders" bson:"orders"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        int                `json:"price"  bson:"price"`
}

type Order struct {
	Order_ID     primitive.ObjectID `bson:"_id"`
	Order_Cart   []ProductUser      `json:"order_list"  bson:"order_list"`
	Orderered_At time.Time          `json:"ordered_on"  bson:"ordered_on"`
	Price        int                `json:"total_price" bson:"total_price"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
}
