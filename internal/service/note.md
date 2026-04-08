Folder service digunakan untuk menyimpan logic bisnis aplikasi.
Contoh :

- func (s \*Service) CreateUser(params CreateUserParams) (models.User, error) {}
- func (s \*Service) GetFlightDetails(flightID int) (models.Flight, error) {}

Yang perlu di folder service :

- Logic bisnis aplikasi, contoh membuat user, mencari penerbangan, memproses pembayaran, dll (Perfile sesuai dengan fitur aplikasi)
- Koordinasi antara repository dan komponen lainnya, contoh memanggil beberapa repository untuk menyelesaikan satu operasi bisnis
- Validasi bisnis, contoh memastikan user memiliki saldo cukup sebelum melakukan pemesanan
- Pengolahan data sebelum disimpan atau dikembalikan ke handler, contoh mapping harga, menghitung diskon, dll

Yang tidak perlu di folder service :

- Handler HTTP
- Query database langsung (pake repository)
- Validasi input request (letakkan di handler)
