# Gudang App

Aplikasi **manajemen gudang** sederhana menggunakan **Golang Fiber**, **GORM**, dan **JWT Authentication**.  

Fitur:  
- Role-based access: `admin` & `staff`  
- Admin: CRUD barang + transaksi  
- Staff: hanya transaksi masuk/keluar, lihat barang & transaksi  
- Transaksi masuk/keluar otomatis update stok  

---

## **1. Teknologi**

- Golang  
- Fiber (web framework)  
- GORM (ORM untuk MySQL)  
- JWT (JSON Web Token)  
- MySQL / MariaDB  

---

## **2. Instalasi & API URL**

1. Clone repository

```bash
git clone <repo_url>
cd gudang-app
```

2. Setup Go modules

```bash
go mod tidy
```

3. Buat database `gudang_db` & jalankan query:

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE,
    password VARCHAR(255),
    role ENUM('admin','staff')
);

CREATE TABLE barang (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(100),
    stok INT,
    lokasi VARCHAR(50)
);

CREATE TABLE transaksi (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_barang INT,
    jenis ENUM('masuk','keluar'),
    jumlah INT,
    keterangan VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (id_barang) REFERENCES barang(id)
);
```

4. Jalankan server

```bash
go run main.go
```

Server berjalan di:  

```
http://localhost:3000
```

**Base API URL untuk semua endpoint:**

```
http://localhost:3000/api
```

---

## **3. Endpoint API**

### **Auth**

| Endpoint        | Method | Body JSON                                             | Keterangan          |
|-----------------|--------|------------------------------------------------------|-------------------|
| `/register`     | POST   | `{ "username": "admin1", "password": "123", "role": "admin" }` | Register user      |
| `/login`        | POST   | `{ "username": "admin1", "password": "123" }`       | Login, return JWT token |

**Header JWT**:  
```
Authorization: Bearer <token>
```

---

### **Barang**

| Endpoint           | Method | Role      | Body JSON / Params                                   | Keterangan                  |
|-------------------|--------|-----------|----------------------------------------------------|-----------------------------|
| `/barang`          | GET    | admin/staff | -                                                 | Lihat semua barang          |
| `/barang`          | POST   | admin      | `{ "nama":"Laptop", "stok":10, "lokasi":"Rak A1" }` | Tambah barang baru          |
| `/barang/:id`      | PUT    | admin      | `{ "nama":"Laptop Pro", "stok":15, "lokasi":"Rak A1" }` | Update barang              |
| `/barang/:id`      | DELETE | admin      | -                                                 | Hapus barang                |

---

### **Transaksi**

| Endpoint           | Method | Role           | Body JSON / Params                                         | Keterangan                     |
|-------------------|--------|----------------|------------------------------------------------------------|--------------------------------|
| `/transaksi`       | GET    | admin/staff    | -                                                          | Lihat semua transaksi          |
| `/transaksi`       | POST   | admin/staff    | `{ "id_barang":1, "jenis":"masuk/keluar", "jumlah":5, "keterangan":"Catatan" }` | Tambah transaksi, stok update otomatis |

**Logika stok:**

- `jenis = masuk` → stok bertambah  
- `jenis = keluar` → stok berkurang  
- Stok tidak cukup → error 400  

---

## **4. Role-based Access**

| Role   | GET Barang | POST Barang | PUT Barang | DELETE Barang | POST Transaksi | GET Transaksi |
|--------|------------|-------------|------------|---------------|----------------|---------------|
| Admin  | ✅         | ✅          | ✅         | ✅            | ✅             | ✅            |
| Staff  | ✅         | ❌          | ❌         | ❌            | ✅             | ✅            |

---

## **5. Contoh Request (cURL)**

### Login
```bash
curl -X POST http://localhost:3000/api/login -H "Content-Type: application/json" -d '{"username":"admin1","password":"123"}'
```

### Tambah Barang (Admin)
```bash
curl -X POST http://localhost:3000/api/barang -H "Content-Type: application/json" -H "Authorization: Bearer <token_admin>" -d '{"nama":"Laptop","stok":10,"lokasi":"Rak A1"}'
```

### Buat Transaksi (Staff/Admin)
```bash
curl -X POST http://localhost:3000/api/transaksi -H "Content-Type: application/json" -H "Authorization: Bearer <token_staff>" -d '{"id_barang":1,"jenis":"keluar","jumlah":3,"keterangan":"Dikirim ke client"}'
```

---

## **6. Catatan**

- Semua route setelah login **dilindungi JWT**  
- Role-based access untuk membatasi staff vs admin  
- Stok barang otomatis update saat transaksi  

---

## **7. License**
MIT License
