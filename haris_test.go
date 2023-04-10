package haris

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertUser(t *testing.T) {
	dbname := "suratdibai"
	user := UserSurat{
		ID:      primitive.NewObjectID(),
		Nama:    "Haris Riyoni",
		Email:   "harisriyoni@gmail.com",
		Telepon: "081234567890",
	}
	insertedID := InsertUser(dbname, user)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}

func TestInsertSurat(t *testing.T) {
	dbname := "suratdibai"
	surat := Surat{
		ID:       primitive.NewObjectID(),
		Isisurat: "Lorem ipsum dolor, sit amet consectetur adipisicing elit. Eum omnis voluptatem accusantium nemo perspiciatis delectus atque autem!",
		Subject:  "Kamu",
	}
	insertedID := InsertSurat(dbname, surat)
	if insertedID == nil {
		t.Error("Failed to insert surat")
	}
}

func TestInsertKategori(t *testing.T) {
	dbname := "suratdibai"
	surat := Surat{
		ID:       primitive.NewObjectID(),
		Isisurat: "Ini isi surat terserah bebas mau ngisi apa ya gesya",
		Subject:  "Pengumpulan data diri kepada perusahaan haris.com",
	}
	kategori := Kategorisurat{
		ID:           primitive.NewObjectID(),
		NamaKategori: "Kategori pekerjaan",
		Surat:        []Surat{surat},
	}
	insertedID := InsertKategori(dbname, kategori)
	if insertedID == nil {
		t.Error("Failed to insert kategori")
	}
}

func TestInsertLokasi(t *testing.T) {
	dbname := "suratdibai"
	lokasi := Lokasisurat{
		ID:     primitive.NewObjectID(),
		Region: "Jakarta",
	}
	insertedID := InsertLokasi(dbname, lokasi)
	if insertedID == nil {
		t.Error("Failed to insert lokasi")
	}
}

func TestInsertAbout(t *testing.T) {
	dbname := "suratdibai"
	about := Aboutsurat{
		ID:       primitive.NewObjectID(),
		Isi_satu: "Ini isi satu",
		Isi_dua:  "Ini isi dua",
		Image:    "image.jpg",
	}
	insertedID := InsertAbout(dbname, about)
	if insertedID == nil {
		t.Error("Failed to insert about")
	}
}

func TestGetUserData(t *testing.T) {
	nama := "081234567890"
	gege := GetUserData(nama)
	fmt.Println(gege)
}

func TestGetEmailData(t *testing.T) {
	nosubject := "1"
	hiya := GetEmailData(nosubject)
	fmt.Println(hiya)
}

func TestGetNamaUser(t *testing.T) {
	nama := "Haris Riyoni"
	hiya := GetNamaUser(nama)
	fmt.Println(hiya)
}
func TestGetSurat(t *testing.T) {
	nama := "Kamu"
	hiya := GetSurat(nama)
	fmt.Println(hiya)
}

// func TestInsertIsiSurat(t *testing.T) {

// 	result := InsertIsiSurat(MongoConnect())
// 	fmt.Println(result)
// }
