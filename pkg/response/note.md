Tempat nyimpan template response yang akan dikirim ke client.
Contoh :

func Success(c \*gin.Context, data interface{}) {
c.JSON(http.StatusOK, gin.H{
"status": "success",
"data": data,
})
}

func Error(c \*gin.Context, statusCode int, message string) {
c.JSON(statusCode, gin.H{
"status": "error",
"message": message,
})
}

jadi nanti di handler tinggal manggil response.Success(c, data) atau response.Error(c, 404, "Not Found")

Yang perlu di folder response :

- Template response yang konsisten untuk aplikasi, contoh response sukses, response error, dll
- Memudahkan handler dalam menyusun response yang akan dikirim ke client
