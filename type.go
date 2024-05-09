package donation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Donasi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jenis_donasi string             `bson:"jenis_donasi,omitempty" json:"jenis_donasi,omitempty"`
	Jumlah       int                `bson:"jumlah,omitempty" json:"jumlah,omitempty"`
}


type BencanaAlam struct {
	Jenis     string    `bson:"jenis,omitempty" json:"jenis,omitempty"`
	Lokasi    string    `bson:"lokasi,omitempty" json:"lokasi,omitempty"`
	Tanggal   time.Time `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
	Deskripsi string    `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
}


