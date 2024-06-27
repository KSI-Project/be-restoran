package module

import (
	"context"
	"errors"
	"fmt"
	"github.com/ghaidafasya24/be-tubes/model"
	// "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

// INSERT MENU
func InsertMenu(db *mongo.Database, col string, nama string, harga float64, deskripsi string) (insertedID primitive.ObjectID, err error) {
	menu := bson.M{
		"nama":      nama,
		"harga":     harga,
		"deskripsi": deskripsi,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), menu)
	if err != nil {
		fmt.Printf("InsertMenu: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetMenuByID(_id primitive.ObjectID, db *mongo.Database, col string) (menu Menu, err error) {
    koleksi := db.Collection(col)
    filter := bson.M{"_id": _id}
    err = koleksi.FindOne(context.TODO(), filter).Decode(&menu)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return menu, fmt.Errorf("data dengan ID %s tidak ditemukan", _id)
        }
        return menu, fmt.Errorf("kesalahan dalam mengambil data dengan ID %s: %s", _id, err.Error())
    }
    return menu, nil
}





// func GetKaryawanFromPhoneNumber(phone_number string, db *mongo.Database, col string) (staf model.Presensi, errs error) {
// 	karyawan := db.Collection(col)
// 	filter := bson.M{"phone_number": phone_number}
// 	err := karyawan.FindOne(context.TODO(), filter).Decode(&staf)
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return staf, fmt.Errorf("no data found for phone number %s", phone_number)
// 		}
// 		return staf, fmt.Errorf("error retrieving data for phone number %s: %s", phone_number, err.Error())
// 	}
// 	return staf, nil
// }

// func GetAllPresensi(db *mongo.Database, col string) (data []model.Presensi) {
// 	karyawan := db.Collection(col)
// 	filter := bson.M{}
// 	cursor, err := karyawan.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("GetALLData :", err)
// 	}
// 	err = cursor.All(context.TODO(), &data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }

// func GetPresensiFromID(_id primitive.ObjectID, db *mongo.Database, col string) (presensi model.Presensi, errs error) {
// 	karyawan := db.Collection(col)
// 	filter := bson.M{"_id": _id}
// 	err := karyawan.FindOne(context.TODO(), filter).Decode(&presensi)
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return presensi, fmt.Errorf("no data found for ID %s", _id)
// 		}
// 		return presensi, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
// 	}
// 	return presensi, nil
// }

// func UpdatePresensi(db *mongo.Database, col string, id primitive.ObjectID, long float64, lat float64, lokasi string, phonenumber string, checkin string, biodata model.Karyawan) (err error) {
// 	filter := bson.M{"_id": id}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"longitude":    long,
// 			"latitude":     lat,
// 			"location":     lokasi,
// 			"phone_number": phonenumber,
// 			"checkin":      checkin,
// 			"biodata":      biodata,
// 		},
// 	}
// 	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		fmt.Printf("UpdatePresensi: %v\n", err)
// 		return
// 	}
// 	if result.ModifiedCount == 0 {
// 		err = errors.New("no data has been changed with the specified ID")
// 		return
// 	}
// 	return nil
// }

// func DeletePresensiByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
// 	karyawan := db.Collection(col)
// 	filter := bson.M{"_id": _id}

// 	result, err := karyawan.DeleteOne(context.TODO(), filter)
// 	if err != nil {
// 		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
// 	}

// 	if result.DeletedCount == 0 {
// 		return fmt.Errorf("data with ID %s not found", _id)
// 	}

// 	return nil
// }