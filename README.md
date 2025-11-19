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

## **2. Instalasi**

1. Clone repository

```bash
git clone <repo_url>
cd gudang-app
