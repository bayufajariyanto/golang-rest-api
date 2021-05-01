# Simple Rest Api menggunakan golang
1. Import db_go.sql
2. Install golang di https://golang.org/
3. Cek versi golang di terminal dengan `go version`
4. Clone project dan masuk ke dalam folder project
5. Jalankan `go run .`

## REST

URL : `http://localhost:8080`

- `GET` Tampilkan semua data
  
  route `/users`
- `POST` Tambah data
  
  route `/user/{id}`
  raw 
  {
    "name" : "Data Baru"
  }
- `PUT` Ubah data

  route `/user/{id}`
  raw
  {
    "name" : "Data diperbarui"
  }
  
- `DELETE` Hapus data
  
  route `/user/{id}`
