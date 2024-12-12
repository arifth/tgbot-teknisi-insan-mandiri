package templates

func RepliesToCreateNewTask() string {
	return `
Format Laporan servis 

Tanggal
Jam
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

func RepliesSuccesInsertDataToSheet() string {
	return `
 Berhasil masukkan Data
    `
}

func RepliesSuccess() string {
	return `Terimakasih Sudah Menghubungi Bot Insan Mandiri`
}
