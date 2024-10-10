# Modul2 review struktur control
jelaskan secara singkat mengenai modul 2

##code explanation

Soal Latihan Modul 2A
1.Program bertujuan untuk menerima tiga input string dari pengguna, menampilkan urutan awal string tersebut, kemudian menukar nilai dari ketiga string dengan pola tertentu, dan akhirnya menampilkan hasil perubahan urutannya.
2. Program bertujuan untuk menentukan apakah sebuah tahun merupakan tahun kabisat atau bukan. Program meminta input dari pengguna berupa sebuah tahun, kemudian mengevaluasi apakah tahun tersebut memenuhi kriteria tahun kabisat, dan akan menampilkan hasilnya.
3. Program dibuat untuk menentukan apakah sebuah tahun adalah tahun kabisat atau bukan. Dalam kalender Gregorian, tahun kabisat adalah tahun yang memiliki 366 hari, bukan 365 hari seperti tahun biasa. Program akan meminta pengguna memasukkan sebuah tahun, lalu memeriksa apakah tahun tersebut adalah kabisat menggunakan aturan yang telah ditentukan.
4. Program berfungsi untuk melakukan konversi suhu dari satuan Celsius ke tiga satuan suhu lainnya: Fahrenheit, Reamur, dan Kelvin. Pengguna akan diminta untuk memasukkan nilai suhu dalam satuan Celsius dan program akan menampilkan hasil konversi ke ketiga satuan tersebut.
5. Program akan menerima lima input angka dan tiga karakter. Angka-angka akan diperlakukan sebagai karakter berdasarkan kode ASCII-nya. Karakter yang diinputkan akan diubah menjadi karakter selanjutnya dalam urutan ASCII, lalu ditampilkan.

soal latihan 2B
1. Program ini meminta pengguna untuk memasukkan urutan warna untuk 5 eksperimen, di mana setiap eksperimen memiliki 4 warna yang harus dimasukkan. Setelah itu, program akan memeriksa apakah urutan warna yang dimasukkan sesuai dengan urutan yang benar, yaitu: merah, kuning, hijau, dan ungu. Jika semua eksperimen memiliki urutan warna yang benar, program akan mencetak "BERHASIL: true", jika tidak, maka "BERHASIL: false".
2. Program ini bertujuan untuk menerima sejumlah nama bunga dari pengguna, lalu menyusun sebuah pita (string) yang berisi nama-nama bunga tersebut yang dipisahkan oleh tanda " - ". Pada akhirnya, program akan menampilkan hasil pita yang berisi rangkaian nama bunga yang sudah dimasukkan.
3. Program ini berfungsi untuk menghitung berat belanjaan di dua kantong dan mengevaluasi apakah sepeda motor Pak Andi akan oleng berdasarkan berat kantong. Program akan terus meminta input hingga total berat mencapai 150 kg atau salah satu kantong memiliki berat negatif. Program juga akan menentukan apakah ada perbedaan berat yang cukup signifikan antara kedua kantong untuk menyebabkan sepeda motor oleng.
4. Program ini bertujuan untuk menghitung aproksimasi nilai akar dari 2 menggunakan deret konvergensi yang dihasilkan dari rumus tertentu. Pengguna diminta untuk memasukkan jumlah iterasi yang diinginkan, dan program akan menghitung serta menampilkan hasil aproksimasi akar 2 dengan format yang memiliki 10 angka di belakang koma.

soal latihan 2C
1. Program ini digunakan untuk menghitung biaya pengiriman parsel berdasarkan beratnya dalam gram. Program akan menerima input berupa berat parsel dalam gram, kemudian program menghitung biaya pengiriman berdasarkan berat kilogram (kg) dan sisa gram di bawah 1 kg. Biaya pengiriman dihitung dengan aturan tarif tertentu tergantung pada berat keseluruhan dan sisa gram. Aturan penghitungan biaya adalah sebagai berikut: Biaya dasar: Rp 10.000 per kilogram. Jika berat lebih dari atau sama dengan 10 kg, tidak ada biaya tambahan untuk sisa gram. Jika berat kurang dari 10 kg: Jika sisa gram lebih dari atau sama dengan 500 gram, maka biaya tambahan adalah Rp 5 per gram. Jika sisa gram kurang dari 500 gram, maka biaya tambahan adalah Rp 15 per gram. Setelah melakukan perhitungan, program menampilkan rincian berat dalam kilogram dan gram, serta rincian biaya dan total biaya pengiriman parsel.
2. Program ini bertujuan untuk mengkonversi nilai angka dari mata kuliah menjadi nilai huruf berdasarkan skala tertentu. Pertanyaan:

a. Jika nam diberikan nilai 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal? jawab: Program ini tidak dirancang untuk menangani nilai float sebagai input, dan oleh karena itu, program tidak akan dieksekusi sesuai spesifikasi soal.

b. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program yang seharusnya! jawab: 1. Variabel nam dideklarasikan sebagai float64, tetapi dalam blok if, nam diubah menjadi string. Ini akan menyebabkan error tipe data. 2. Nilai huruf seharusnya disimpan dalam variabel nmk, bukan nam. 3. Setiap pernyataan if berdiri sendiri, sehingga nilai yang lebih kecil terus menimpa nilai sebelumnya.

c. Perbaiki Program! jawab: ''' go package main

  import "fmt"

  func main() {
  	var nam float64
  	var nmk string
  
  	fmt.Print("Nilai akhir mata kuliah: ")
  	fmt.Scanln(&nam)
  
  	if nam > 80 {
  		nmk = "A"
  	} else if nam > 72.5 {
  		nmk = "AB"
  	} else if nam > 65 {
  		nmk = "B"
  	} else if nam > 57.5 {
  		nmk = "BC"
  	} else if nam > 50 {
  		nmk = "C"
  	} else if nam > 40 {
  		nmk = "D"
  	} else {
  		nmk = "E"
  	}
  
  	fmt.Println("Nilai mata kuliah: ", nmk)
  }
'''
3. Program ini melakukan dua operasi, operasi pertama menentukan faktor-faktor dari bilangan bulat positif. Program akan meminta pengguna untuk memasukkan bilangan bulat positif lebih dari 1, kemudian mencetak semua faktor dari bilangan tersebut. Operasi kedua program memeriksa apakah bilangan tersebut adalah bilangan prima atau bukan. Sebuah bilangan dikatakan prima jika hanya memiliki dua faktor, yaitu 1 dan dirinya sendiri.
