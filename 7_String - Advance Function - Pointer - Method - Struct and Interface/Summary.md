Advance function
-variadic func
-anonymous func==literal func
closure
defer func

string
-len
-compare
-contains
-substring
-replace
-insert

pointer. pointer adalah sebuah variabel yang mempunyai isi berupa sebuah alamat atau lokasi memori.

struct. Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu. Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti map, hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda.

Dari sebuah struct, kita bisa buat variabel baru, yang memiliki atribut sesuai skema struct tersebut. Kita sepakati dalam buku ini, variabel tersebut dipanggil dengan istilah object atau object struct.

method. Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya). Method bisa diakses lewat variabel objek.

Keunggulan method dibanding fungsi biasa adalah memiliki akses ke property struct hingga level private (level akses nantinya akan dibahas lebih detail pada chapter selanjutnya). Dan juga, dengan menggunakan method sebuah proses bisa di-enkapsulasi dengan baik.

interface. Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan nama tertentu.

Interface merupakan tipe data. Nilai objek bertipe interface zero value-nya adalah nil. Interface mulai bisa digunakan jika sudah ada isinya, yaitu objek konkret yang memiliki definisi method minimal sama dengan yang ada di interface-nya.

package. Package adalah koleksi dari func dan data.

error handling.Error merupakan topik yang sangat penting dalam pemrograman Go. Di bagian ini kita akan belajar mengenai pemanfaatan error dan cara membuat custom error sendiri. Selain itu, kita juga akan belajar tentang penggunaan panic untuk memunculkan panic error, dan recover untuk mengatasinya.
error merupakan sebuah tipe. Error memiliki 1 buah property berupa method Error(), method ini mengembalikan detail pesan error dalam string. Error termasuk tipe yang isinya bisa nil.

Di Go, banyak sekali fungsi yang mengembalikan nilai balik lebih dari satu. Biasanya, salah satu kembalian adalah bertipe error. Contohnya seperti pada fungsi strconv.Atoi(). Fungsi tersebut digunakan untuk konversi data string menjadi numerik. Fungsi ini mengembalikan 2 nilai balik. Nilai balik pertama adalah hasil konversi, dan nilai balik kedua adalah error. Ketika konversi berjalan mulus, nilai balik kedua akan bernilai nil. Sedangkan ketika konversi gagal, penyebabnya bisa langsung diketahui dari error yang dikembalikan.
