package donation

import (
	"fmt"
	"testing"
)

func TestInsertDonasi(t *testing.T) {
	nama := "Yoru Radiant"
	phone_number := "081222333222"
	jenis_donasi := "Radianite"
	jumlah := 999999
	hasil := InsertDonasi(nama, phone_number, jenis_donasi, jumlah)
	fmt.Println(hasil)
}

func TestGetDonasiFromPhoneNumber(t *testing.T) {
	phonenumber := "68122221814"
	biodata:=GetDonasiFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}

func TestGetAllDonasi(t *testing.T) {
	data := GetAllDonasi()
	fmt.Println(data)
}
