package templates

import (
	"fmt"
	"time"
)

func RepliesToCreateNewTask() string {
	return `
💼Format Laporan servis💼🪛
=======================
Nama Pelanggan
Alamat Pelanggan
Merek Mesin 
Type Mesin 
No serie Mesin 
Jenis Kerusakan 
Tindakan Perbaikan
Sparepart yg di ganti
Counter Mesin 
Lokasi Pekerjaan ( kantor,dlm kota,luar kota)
Hasil Pekerjaan (ready,pending,ulang)
Jenis Pekerjaan (Service,Membangun Mesin,Kunjungan Aktif,Install&Training Operator,Lain Lain Office) 
point pekerjaan
    `
}

func RepliesToChannel(user string) string {
	currentTime := time.Now()
	currentHour := currentTime.Hour()
	currenMinute := currentTime.Minute()
	current := fmt.Sprintf("%v.%v", currentHour, currenMinute)
	formatted := currentTime.Format("2006-01-02")
	return fmt.Sprintf(`
	🧑🏾‍🚒On Duty 	: %s
    🚧Start		: %s
    ⧗ Time		:% s
`, user, formatted, current)
}

func RepliesSuccesInsertDataToSheet(spreadSheetId string) string {
	return fmt.Sprint("Berhasil masukkan Data🫡, \n silahkan Cek di https://docs.google.com/spreadsheets/d/", spreadSheetId)
}

func RepliesSuccess() string {
	return `Terimakasih Sudah Menghubungi Bot Insan Mandiri`
}
