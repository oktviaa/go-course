package main

import (
	"fmt"
	"os"
)

type biodata struct {
	name      string
	alamat    string
	pekerjaan string
	alasan    string
}

var allbiodata = []biodata{
	{name: "Oktaviana", alamat: "Jakarta", pekerjaan: "IT Support", alasan: "Keren"},
	{name: "Via", alamat: "Bandung", pekerjaan: "Backend", alasan: "Susah"},
	{name: "Okta", alamat: "Lampung", pekerjaan: "Frontend", alasan: "Lebih Susah"},
	{name: "Ana", alamat: "Yogya", pekerjaan: "Fullstack Dev", alasan: "Amat teramat sulid"},
}

func main() {
	var argsRaw = os.Args
	coba :=0
	if len(argsRaw) < 2 {
		fmt.Println("Nomor absen tidak terdaftar.\n Contoh: 'go run main.go Fitri' atau 'go run main.go 2")
		// return
	} else {
		input := argsRaw[1]
		// age := argsRaw[2]
		for i, biodata := range allbiodata {
			if biodata.name == input || input == fmt.Sprint(i) {
				println("ID: ", i)
				println("Nama: ", biodata.name)
				println("Alamat: ", biodata.alamat)
				println("Pekerjaan: ", biodata.pekerjaan)
				println("Alasan: ", biodata.alasan)
			} else {
				coba += 1
			}
		}
		if coba == len(allbiodata){
			println("Maaf data yang anda masukkan salah")
		}
	}

}