Middleware khususnya di dunia web programming adalah suatu software layer yang berada di antara router dengan controller. Karena posisi dari middleware berada di antara router dengan controller, fungsi dari middleware rata — rata memiliki fungsi generik, yaitu :

1.Authentication
2.Authorization
3.Validasi input
4.Sanitasi input
5.Response handler
6.Data logger

tujuan utama dari adanya middleware adalah agar controller fokus dalam menyelesaikan logika alur bisnis dari suatu flow aplikasi tanpa harus untuk melakukan hal — hal di luar itu seperti validasi input untuk setiap flow. Jadi ketika data yang masuk ke controller sudah dalam keadaan bersih dan sudah siap untuk diproses sesuai dengan alur bisnis dari aplikasi tersebut. Dan dengan adanya middleware, kode kita pastinya lebih reusable, maintainable dan readable.

Contohnya seperti ini, misal kita melakukan authentication, jika lokasi kode auth atau modul masih dipasang di controller, pastinya akan repot jika kita cek satu per satu ke setiap controller jika ada kesalahan atau bug. Belum lagi readibility-nya tentu berkurang. Lebih mudah kita menjadikan authentication tersebut menjadi modul middleware. Dan memasangnya pada route agar lebih mudah untuk dicek. Tapi tentu saja aturan ini tidak semuanya berlaku, karena masing — masing framework mempunyai aturan — aturan tersendiri.
