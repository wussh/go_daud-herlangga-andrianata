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

Repository Test
Seperti yang saya tulis sebelumnya di atas, saya menggunakan go-sqlmock untuk melakukan mocking connection ke database, karena kebetulan untuk project yang saya buat menggunakan mysql.
Menggunakan Mock

Usecase Test
Contoh usecase test
package usecase_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/bxcodec/faker"
	models "github.com/bxcodec/go-clean-arch/article"
	"github.com/bxcodec/go-clean-arch/article/repository/mocks"
	ucase "github.com/bxcodec/go-clean-arch/article/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockArticleRepo := new(mocks.ArticleRepository)
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	mockListArtilce := make([]*models.Article, 0)
	mockListArtilce = append(mockListArtilce, &mockArticle)
	mockArticleRepo.On("Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockListArtilce, nil)
	u := ucase.NewArticleUsecase(mockArticleRepo)
	num := int64(1)
	cursor := "12"
	list, nextCursor, err := u.Fetch(cursor, num)
	cursorExpected := strconv.Itoa(int(mockArticle.ID))
	assert.Equal(t, cursorExpected, nextCursor)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListArtilce))

	mockArticleRepo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64"))

}

Delivery Test
Untuk melakukan test pada delivery layer, bergantung pada bagaimana data anda akan di representasikan. Untuk contoh yang saya buat, saya menggunakan HTTP Rest API sebagai delivery methodnya.
Untuk melakukan test pada HTTP, kita dapat melakukan package built-in dari golang: httptest untuk membantu dalam melakukan testing. Selain itu, karena bergantung pada usecase, maka saya akan menggunakan usecase yang sudah di-mocking pada delivery layer.

Final Output and The Merging
Tahap terakhir yaitu menggabungkan semua layer, agar aplikasi dapat berjalan. Penggabungan di lakukan di root folder, di package main.
