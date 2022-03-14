package main

type Kendaraan struct {
	roda_kendaraan             int
	kecepatan_perjam_kendaraan int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) MobilBerjalan() {
	m.TambahKecepatanMobil(10)
}
func (m *Mobil) TambahKecepatanMobil(KecepatanBaruMobil int) {
	m.kecepatan_perjam_kendaraan = m.kecepatan_perjam_kendaraan + KecepatanBaruMobil
}
func main() {
	MobilCepat := Mobil{}
	MobilCepat.MobilBerjalan()
	MobilCepat.MobilBerjalan()
	MobilCepat.MobilBerjalan()

	MobilLambat := Mobil{}
	MobilLambat.MobilBerjalan()
}
