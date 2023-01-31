# MyTokuList

## Andhika Ditya Bagaskara D
## Link : https://mytokulist-production.up.railway.app/

## Public
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

#### example:

#### https://mytokulist-production.up.railway.app/categories
<img width="313" alt="image" src="https://user-images.githubusercontent.com/87473932/213652847-1f997677-da1d-4e76-9f4c-75fe607611d5.png">

#### https://mytokulist-production.up.railway.app/register
```json
{
    "username": "Test01",
    "email": "test@mail.com",
    "password": "test123"
}
```
<img width="948" alt="image" src="https://user-images.githubusercontent.com/87473932/213653210-c3166eca-4bb0-47cf-8633-0d93b9f12068.png">

#### https://mytokulist-production.up.railway.app/login
<img width="955" alt="image" src="https://user-images.githubusercontent.com/87473932/213654156-53ee81ea-e946-4a71-af0e-ebae8df9a565.png">

#### https://mytokulist-production.up.railway.app/logout
<img width="953" alt="image" src="https://user-images.githubusercontent.com/87473932/213654391-b9f588df-ae1c-46f6-a114-4558ea10b1c5.png">

## User
Berikut adalah list endpoint yang dapat diakses oleh user:

### /
| Method | Endpoint      | Description                                     |
|--------|---------------|-------------------------------------------------|
| GET    | /home         | Menampilkan home user                           |
| GET    | /list-user    | Menampilkan list user yang ada                  |
| PUT    | /new-password | Memperbarui password yang lama dengan yang baru |

### /WATCHLIST
| Method | Endpoint       | Description                        |
|--------|----------------|------------------------------------|
| GET    | /watchlist     | Menampilkan watchlist user         |
| POST   | /watchlist     | Menambahkan watchlist user         |
| DELETE | /watchlist/:id | Menghapus watchlist berdasarkan id |

<img width="965" alt="image" src="https://user-images.githubusercontent.com/87473932/215767644-aa1a89b4-fe32-40bd-bc37-0d7899c66f2a.png">

<img width="967" alt="image" src="https://user-images.githubusercontent.com/87473932/215767503-3f3f2b23-84a9-49e9-80fa-853485a6e71b.png">

## Admin
Berikut adalah list endpoint yang dapat diakses oleh admin:

### /ADMIN
| Method | Endpoint            | Description                                     |
|--------|---------------------|-------------------------------------------------|
| GET    | /admin/home         | Menampilkan home admin                          |
| GET    | /admin/list-admin   | Menampilkan list admin yang ada                 |
| GET    | /admin/list-user    | Menampilkan list user yang ada                  |
| PUT    | /admin/new-password | Memperbarui password yang lama dengan yang baru |

### /TYPES
| Method | Endpoint   | Description                         |
|--------|------------|-------------------------------------|
| POST   | /types     | Menambahkan tipe film baru          |
| PUT    | /types/:id | Mengupdate tipe film berdasarkan id |
| DELETE | /types:/id | Menghapus tipe film berdasarkan id  |

<img width="955" alt="image" src="https://user-images.githubusercontent.com/87473932/215766229-54b33e05-96f0-4aa4-88ae-e7b85a3cf3b7.png">

### /CATEGORIES
| Method | Endpoint        | Description                        |
|--------|-----------------|------------------------------------|
| POST   | /categories     | Menambahkan kategori baru          |
| PUT    | /categories/:id | Mengupdate kategori berdasarkan id |
| DELETE | /categories:/id | Menghapus kategori berdasarkan id  |

<img width="959" alt="image" src="https://user-images.githubusercontent.com/87473932/215766479-0c7709f8-5da1-4536-ba41-fa93b4918b18.png">


### /MOVIES
| Method | Endpoint    | Description                                                |
|--------|-------------|------------------------------------------------------------|
| POST   | /movies     | Menambahkan film baru                                      |
| PUT    | /movies/:id | Mengupdate film berdasarkan id                             |
| DELETE | /movies:/id | Menghapus film berdasarkan id                              |
| PATCH  | /movies/:id | Menambahkan kategori dan tipe ke tabel film berdasarkan id |

<img width="960" alt="image" src="https://user-images.githubusercontent.com/87473932/215766663-105289a7-f26a-4ee8-8931-705a8f45ad3f.png">

#### Perlu diperhatikan, method `PUT` pada endpoint di atas membutuhkan semua atribut `movies` kecuali `category_id` dan `types_id`, sebaliknya, method `PATCH` hanya membutuhkan atribut `category_id` dan `types_id`. Hal ini dikarenakan method `PATCH` dimaksudkan untuk meng-attach / update `category` dan `types` saja pada `movies`
