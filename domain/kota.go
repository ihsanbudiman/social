package domain

type Kota struct {
	ID      string  `json:"id"`
	Lokasi  string  `json:"lokasi"`
	Daerah  string  `json:"daerah"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Lintang string  `json:"lintang"`
	Bujur   string  `json:"bujur"`
}

type Jadwal struct {
	KotaID  string `json:"kota_id"`
	Tanggal string `json:"tanggal"`
	Imsak   string `json:"imsak"`
	Subuh   string `json:"subuh"`
	Terbit  string `json:"terbit"`
	Dhuha   string `json:"dhuha"`
	Dzuhur  string `json:"dzuhur"`
	Ashar   string `json:"ashar"`
	Maghrib string `json:"maghrib"`
	Isya    string `json:"isya"`
	Date    string `json:"date"`
}
