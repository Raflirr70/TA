Folder handler digunakan untuk menangani request dan response HTTP mirip controller di framework lain.
Contoh :

- func (h \*Handler) CreateUser(w http.ResponseWriter, r \*http.Request) {}
- func (h \*Handler) GetFlightDetails(w http.ResponseWriter, r \*http.Request) {}

Yang perlu di folder handler :

- Menangani request dan response HTTP, contoh parsing request body, query parameters, headers, dll
- Validasi input request, contoh memastikan data yang diterima sesuai format yang diharapkan
- Memanggil service untuk menjalankan logic bisnis aplikasi
- Menyusun response yang akan dikirimkan ke client, contoh mapping response (Sesuaikan dengan Figma)
- Menangani error yang terjadi selama proses request, misal mengembalikan status code dan pesan error yang sesuai, contoh 404 NOT FOUND penerbangan tidak ditemukan

Yang tidak perlu di folder handler :

- Logic bisnis aplikasi (letakkan di service)
- Query database langsung (pake repository)
- Validasi bisnis (letakkan di service)
- Pengolahan data yang kompleks (letakkan di service)
