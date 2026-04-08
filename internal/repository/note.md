Folder ini digunakan untuk menyimpan query dengan GORM.

Contoh :

- type SearchFlights struct {}
- func (r \*Repository) SearchFlights(params SearchFlights) ([]models.Flight, error) {}

Yang tidak perlu di folder repository :

- Logic bisnis
- Handler HTTP
- Validasi input
