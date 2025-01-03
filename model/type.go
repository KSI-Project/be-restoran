package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama      string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Harga     float64            `bson:"harga,omitempty" json:"harga,omitempty"`
	Deskripsi string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Gambar    string             `bson:"gambar,omitempty" json:"gambar,omitempty"`
	Kategori  Kategori           `bson:"kategori,omitempty" json:"kategori,omitempty"`
	BahanBaku BahanBaku          `bson:"bahan_baku,omitempty" json:"bahan_baku,omitempty"`
}

type Kategori struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kategori string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
}

type BahanBaku struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	BahanBaku string             `bson:"bahan_baku,omitempty" json:"bahan_baku,omitempty"`
	Jumlah    string             `bson:"jumlah,omitempty" json:"jumlah,omitempty"`
}

type Customer struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	BahanBaku string             `bson:"bahan_baku,omitempty" json:"bahan_baku,omitempty"`
	Jumlah    string             `bson:"jumlah,omitempty" json:"jumlah,omitempty"`
}
