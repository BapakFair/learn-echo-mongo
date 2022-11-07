package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	NickName     string             `json:"nickName" bson:"nickName" validate:"required, max=10"`
	FullName     string             `json:"fullName" bson:"fullName" validate:"required, max=50"`
	NomorAnggota string             `json:"nomorAnggota" bson:"nomorAnggota" validate:"required, max=50"`
	TotalPetak   int                `json:"totalPetak" bson:"totalPetak" validate:"required, max=2000"`
	TotalAndil   int                `json:"totalAndil" bson:"totalAndil" validate:"required, max=2000"`
	Email        string             `json:"email" bson:"email" validate:"required"`
	Password     string             `json:"password" bson:"password" validate:"required"`
	PhoneNumber  int                `json:"phoneNumber" bson:"phoneNumber" validate:"required"`
}
