package haris

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
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
func InsertUser(db string, user User) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Printf("InsertUser: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertSurat(db string, surat Surat) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("surat").InsertOne(context.TODO(), surat)
	if err != nil {
		fmt.Printf("InsertSurat: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertKategori(db string, kategori Kategori) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("kategori").InsertOne(context.TODO(), kategori)
	if err != nil {
		fmt.Printf("InsertKategori: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertLokasi(db string, lokasi Lokasi) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("lokasi").InsertOne(context.TODO(), lokasi)
	if err != nil {
		fmt.Printf("InsertLokasi: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertAbout(db string, about About) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("about").InsertOne(context.TODO(), about)
	if err != nil {
		fmt.Printf("InsertAbout: %v\n", err)
	}
	return insertResult.InsertedID
}

// get
//
//	func GetUserData(nama string) (data User) {
//		user := MongoConnect("suratdb").Collection("users")
//		filter := bson.M{"nama": nama}
//		err := user.FindOne(context.TODO(), filter).Decode(&data)
//		if err != nil {
//			fmt.Printf("GetUserData: %v\n", err)
//		}
//		return data
//	}
func GetDataSurat(isisurat string) (data Surat) {
	surat := MongoConnect("suratdb").Collection("surat")
	filter := bson.M{"isisurat": isisurat}
	err := surat.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("GetDataSurat: %v\n", err)
	}
	return data
}

func GetUserData(nama string) (data []User) {
	user := MongoConnect("suratdb").Collection("users")
	filter := bson.M{"nama": nama}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetUserData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
