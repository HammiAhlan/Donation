// donation.go
package donation

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDonasi(nama string, phone_number string, jenis_donasi string, jumlah int) (insertedID interface{}) {
	var donation Donasi
	donation.Nama = nama
	donation.Phone_number = phone_number
	donation.Jenis_donasi = jenis_donasi
	donation.Jumlah = jumlah
	return InsertOneDoc("DonationDB", "donation_data", donation)
}



func InsertBencanaAlam(jenis string, lokasi string, tanggal time.Time, deskripsi string) (insertedID interface{}) {
	var bencana BencanaAlam
	bencana.Jenis = jenis
	bencana.Lokasi = lokasi
	bencana.Tanggal = tanggal
	bencana.Deskripsi = deskripsi
	return InsertOneDoc("Donation", "donation", bencana)
}

func GetDonasiFromPhoneNumber(phone_number string) (donasi Donasi) {
	karyawan := MongoConnect("Donation").Collection("donation")
	filter := bson.M{"phone_number": phone_number}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&donasi)
	if err != nil {
		fmt.Printf("getDonasiFromPhoneNumber: %v\n", err)
	}
	return donasi
}

func GetAllDonasi() (data []Donasi) {
	donation := MongoConnect("DonationDB").Collection("donation_data")
	filter := bson.M{}
	cursor, err := donation.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}