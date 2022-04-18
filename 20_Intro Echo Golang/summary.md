Echo adalah framework bahasa golang untuk pengembangan aplikasi web.

Salah satu dependensi yang ada di dalamnya adalah router

Dari banyak routing library yang sudah penulis gunakan, hampir seluruhnya mempunyai kemiripan dalam hal penggunaannya, cukup panggil fungsi/method yang dipilih (biasanya namanya sama dengan HTTP Method), lalu sisipkan rute pada parameter pertama dan handler pada parameter kedua.

Berikut contoh sederhana penggunaan echo framework.

r := echo.New()
r.GET("/", handler)
r.Start(":9000")

Sebuah objek router r dicetak lewat echo.New(). Lalu lewat objek router tersebut, dilakukan registrasi rute untuk / dengan method GET dan handler adalah closure handler. Terakhir, dari objek router di-start-lah sebuah web server pada port 9000.
