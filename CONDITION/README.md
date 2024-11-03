# Golang Condition Exercises

Proyek ini berisi beberapa latihan pemrograman dasar menggunakan bahasa Go. Setiap latihan dirancang untuk mengasah kemampuan dalam menggunakan kondisi dan logika pemrograman.

## Daftar Latihan

1. **Graduate Student**
   - **Deskripsi**: Menentukan apakah seorang murid lulus atau tidak berdasarkan nilai dan jumlah ketidakhadiran.
   - **Fungsi Utama**: `GraduateStudent(score int, absent int) string`
   - **File Terkait**: 
     - `golang-condition-cp-1-v2/main.go`
     - `golang-condition-cp-1-v2/main_test.go`
     - `golang-condition-cp-1-v2/README.md`

2. **BMI Calculator**
   - **Deskripsi**: Menghitung Body Mass Index (BMI) berdasarkan jenis kelamin dan tinggi badan.
   - **Fungsi Utama**: `BMICalculator(gender string, height int) float64`
   - **File Terkait**: 
     - `golang-condition-cp-2-v2/main.go`
     - `golang-condition-cp-2-v2/main_test.go`
     - `golang-condition-cp-2-v2/README.md`

3. **Predicate Score**
   - **Deskripsi**: Menentukan predikat nilai berdasarkan rata-rata dari beberapa mata pelajaran.
   - **Fungsi Utama**: `GetPredicate(math, science, english, indonesia int) string`
   - **File Terkait**: 
     - `golang-condition-cp-3-v1/main.go`
     - `golang-condition-cp-3-v1/main_test.go`
     - `golang-condition-cp-3-v1/README.md`

4. **Ticket Discount**
   - **Deskripsi**: Menghitung harga tiket bioskop dengan diskon berdasarkan jumlah tiket dan tanggal pembelian.
   - **Fungsi Utama**: `GetTicketPrice(VIP int, regular int, student int, day int) float32`
   - **File Terkait**: 
     - `golang-condition-cp-4-v1/main.go`
     - `golang-condition-cp-4-v1/main_test.go`
     - `golang-condition-cp-4-v1/README.md`

5. **Ticket Playground**
   - **Deskripsi**: Menentukan tarif tiket wahana bermain berdasarkan tinggi dan umur anak.
   - **Fungsi Utama**: `TicketPlayground(height, age int) int`
   - **File Terkait**: 
     - `golang-condition-cp-5-v2/main.go`
     - `golang-condition-cp-5-v2/main_test.go`
     - `golang-condition-cp-5-v2/README.md`

## Cara Menjalankan

1. Pastikan Anda memiliki Go terinstal di sistem Anda.
2. Clone repositori ini.
3. Masuk ke direktori latihan yang ingin Anda jalankan.
4. Gunakan perintah `go run main.go` untuk menjalankan program.
5. Gunakan perintah `go test` untuk menjalankan pengujian.

## Struktur Proyek

- Setiap latihan memiliki direktori sendiri yang berisi file `main.go` untuk implementasi dan `main_test.go` untuk pengujian.
- File `README.md` di setiap direktori memberikan deskripsi dan contoh kasus uji untuk latihan tersebut.