1. Entities
2. Usecase
3. Controller
4. Framework dan Driver


Models
Layer ini merupakan layer yang menyimpan model yang dipakai pada domain lainnya. Layer ini dapat diakses oleh semua layer dan oleh semua domain.

import "time"

type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

Repository
Layer ini merupakan layer yang menyimpan database handler. Querying, Inserting, Deleting akan dilakukan pada layer ini. Tidak ada business logic disini. Yang ada hanya fungsi standard untuk input-output dari datastore.
Layer ini memiliki tugas utama yakni menentukan datastore apa yang di gunakan. Teman-teman boleh memilih sesuai kepada kebutuhan, mungkin RDBMS (Mysql,PostgreSql, dsb) atau NoSql (Mongodb,CouchDB dsb).
Jika menggunakan arsitektur microservice, maka layer ini akan bertugas sebagai penghubung kepada service lain. Layer ini akan terikat dan bergantung pada datastore yang digunakan.

Usecase
Layer ini merupakan layer yang akan bertugas sebagai pengontrol, yakni menangangi business logic pada setiap domain. Layer ini juga bertugas memilih repository apa yang akan digunakan, dan domain ini bisa memiliki lebih dari satu repository layer.
Tugas utama terbesar dari layer ini, yaitu menjadi penghubung antara datastore (repository layer) dengan delivery layer. Sehingga, layer ini juga bertanggung jawab atas kevalidan data, jika sesuatu terjadi data yang tidak valid pada repository atau delivery, maka layer ini yang pertama kali disalahkan.
Layer ini benar benar harus berisi business logic, contohnya: penjumlahan, total masukan, atau membentuk response yang merupakan gabungan dari beberapa repository/model. Layer ini bergantung pada repository layer. Jadi jika terjadi perubahan di repository secara besar-besaran tentu saja mempengaruhi layer ini. 

Delivery
Layer ini merupakan layer yang akan bertugas sebagai presenter atau menjadi output dari aplikasi. Layer ini bertugas menentukan metode penyampaian yang dipakai, bisa dengan Rest API, HTML, gRPC, File dsb.
Tugas lain dari layer ini, menjadi dinding penghubung antara user dan sistem. Menerima segala input dan validasi input sesuai standar yang digunakan.
Untuk contoh project yang saya gunakan, saya memilih Rest API sebagai delivery layernya. Sehingga, komunikasi antara client/user terhadap sistem dilakukan melalui REST API

Komunikasi Antar Layer
Baiklah, kita sudah tahu ke-4 layer, lalu bagaimana cara memakainya? Bagaimana komunikasi antar layer dilakukan ?
Untuk setiap layer, kecuali Model, akan berkomunikasi menggunakan interface. Interface ini bertindak seperti kontrak/perjanjian antara 2 layer.
Contoh Repository Interface


package repository

import models "github.com/bxcodec/go-clean-arch/article"

type ArticleRepository interface {
	Fetch(cursor string, num int64) ([]*models.Article, error)
	GetByID(id int64) (*models.Article, error)
	GetByTitle(title string) (*models.Article, error)
	Update(article *models.Article) (*models.Article, error)
	Store(a *models.Article) (int64, error)
	Delete(id int64) (bool, error)
} 

Interface ini akan menjadi atribut utama, dan di-inject pada usecase layer. Di harapkan, ketika membuat interface ini untuk lebih berhati-hati dan bertanggung jawab. Pembuatan interface ini harus selesai diawal dan tidak berubah, karena akan dipakai oleh 2 layer yang berbeda bersamaan dan tentu saja jika terjadi perubahan akan menyulitkan nantinya. Usecase hanya mengetahui fungsi di repository melalui interface ini, sehingga repository harus implement interface ini.

Contoh Usecase Interface


package usecase

import (
	"github.com/bxcodec/go-clean-arch/article"
)

type ArticleUsecase interface {
	Fetch(cursor string, num int64) ([]*article.Article, string, error)
	GetByID(id int64) (*article.Article, error)
	Update(ar *article.Article) (*article.Article, error)
	GetByTitle(title string) (*article.Article, error)
	Store(*article.Article) (*article.Article, error)
	Delete(id int64) (bool, error)
}

Testing Setiap Layer
Pada hakikatnya, clean artinya adalah independen.Setiap layer dapat di-test secara independen, meski layer lain masih belum selesai.
Model Layer
Layer ini dapat di-test secara independen karena tidak berketergantungan kepada layer apapun, layer ini dapat di-test, jika dan hanya jika layer ini memiliki fungsi tersendiri
Repository
Untuk membuat test pada layer ini, idealnya adalah dengan melakukan integration testing. Tetapi kita juga dapat menggunakan unit testing. Terimakasih kepada library github.com/DATA-DOG/go-sqlmock yang sudah memberi kemudahan untuk melakukan unit testing mysql pada project saya. Jika anda menggunakan datastore yang berbeda, pastikan anda memiliki unit testing helper-nya atau dengan melakukan integration testing.
Usecase
Karena untuk melakukan test pada layer ini, dibutuhkan sebuah injection dari repository layer, maka kita dapat melakukan mocking pada repository, dan inject mockingnya pada unit test. Untuk mocking berdasarkan interface, saya menggunakan mockery.
Delivery
Seperti usecase, karena layer ini bergantung pada usecase layer, maka dibutuhkan proses inject dari usecase untuk melakukan testing. Karena test bersifat independen, maka kita harus melakukan mocking pada usecase.

