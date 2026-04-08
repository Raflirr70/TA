Di folder routes, digunakan untuk mendefinisikan rute-rute HTTP yang ada di aplikasi.
Contoh :

- router.POST("/users", handler.CreateUser)
- router.GET("/flights/:id", handler.GetFlightDetails)
- router.PUT("/payments/:id", handler.ProcessPayment)
- router.DELETE("/bookings/:id", handler.CancelBooking)

Yang perlu di folder routes :

- Mendefinisikan rute-rute HTTP beserta metode HTTP-nya (GET, POST, PUT, DELETE, dll)
- Menghubungkan rute dengan handler yang sesuai
- Menyusun middleware yang diperlukan untuk rute tertentu (misal autentikasi, logging, dll)

struktur file bisa seperti ini :

- routes/
  - user_routes.go
  - flight_routes.go
  - payment_routes.go
  - booking_routes.go
  - index.go (opsional, untuk menggabungkan semua rute)
