# fp-sanbercode-go-batch-41
# MyTokuList

## Andhika Ditya Bagaskara D
## Link : https://mytokulist-production.up.railway.app/

## Global
Berikut adalah list endpoint yang dapat diakses tanpa harus melakukan login terlebih dahulu:

### GET METHOD
| Method | Endpoint    | Description                              |
|--------|-------------|------------------------------------------|
| GET    | /types      | Menampilkan semua tipe film yang ada.    |
| GET    | /categories | Menampilkan semua kategori film yang ada |
| GET    | /movies     | Menampilkan semua film yang ada          |

### POST METHOD
| Method | Endpoint | Description                                                                                                                                 |
|--------|----------|---------------------------------------------------------------------------------------------------------------------------------------------|
| POST   | /signup  | Mendaftarkan user baru. Dengan melakukan signup, user baru akan dibuat ke dalam table users                                                 |
| POST   | /login   | Melakukan login user dengan menggunakan JWT. Kemudian, token JWT yang digenerate akan disimpan ke dalam cookie.                             |
| POST   | /logout  | Melakukan logout user dengan cara men-drop cookie yang disimpan. Untuk logout akan dilakukan pengecekan apakah user sudah masuk atau belum. |

