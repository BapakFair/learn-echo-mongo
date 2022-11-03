package models

type Product struct {
	//ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id, omitempty"`
	Name        string   `json:"product_name" bson:"product_name" validate:"required, max=10"`
	Price       int      `json:"price" bson:"price" validate:"required, max=2000"`
	Currency    string   `json:"currency" bson:"currency" validate:"required, len=3"`
	Discount    int      `json:"discount" bson:"discount"`
	Vendor      string   `json:"vendor" bson:"vendor" validate:"required"`
	Accessories []string `json:"accessories,omitempty" bson:"accessories, omitempty"`
	IsEssential bool     `json:"isEssential" bson:"isEssential"`
}
