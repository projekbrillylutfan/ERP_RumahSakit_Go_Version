<h1 align="center">
  <a href="https://gofiber.io">
    <picture>
      <source height="125" media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo-dark.svg">
      <img height="125" alt="Fiber" src="https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo.svg">
    </picture>
  </a>
  <br>
  <a href="https://pkg.go.dev/github.com/gofiber/fiber/v3#pkg-overview">
    <img src="https://img.shields.io/badge/%F0%9F%93%9A%20godoc-pkg-00ACD7.svg?color=00ACD7&style=flat-square">
  </a>
  <a href="https://goreportcard.com/report/github.com/gofiber/fiber/v3">
    <img src="https://img.shields.io/badge/%F0%9F%93%9D%20goreport-A%2B-75C46B?style=flat-square">
  </a>
  <a href="https://codecov.io/gh/gofiber/fiber" >
   <img alt="Codecov" src="https://img.shields.io/codecov/c/github/gofiber/fiber?token=3Cr92CwaPQ&style=flat-square&logo=codecov&label=codecov">
 </a>
  <a href="https://github.com/gofiber/fiber/actions?query=workflow%3ATest">
    <img src="https://img.shields.io/github/actions/workflow/status/gofiber/fiber/test.yml?branch=master&label=%F0%9F%A7%AA%20tests&style=flat-square&color=75C46B">
  </a>
    <a href="https://docs.gofiber.io">
    <img src="https://img.shields.io/badge/%F0%9F%92%A1%20fiber-docs-00ACD7.svg?style=flat-square">
  </a>
  <a href="https://gofiber.io/discord">
    <img src="https://img.shields.io/discord/704680098577514527?style=flat-square&label=%F0%9F%92%AC%20discord&color=00ACD7">
  </a>
</h1>
<p align="center">
  <em><b>Fiber</b> is an <a href="https://github.com/expressjs/express">Express</a> inspired <b>web framework</b> built on top of <a href="https://github.com/valyala/fasthttp">Fasthttp</a>, the <b>fastest</b> HTTP engine for <a href="https://go.dev/doc/">Go</a>. Designed to <b>ease</b> things up for <b>fast</b> development with <a href="https://docs.gofiber.io/#zero-allocation"><b>zero memory allocation</b></a> and <b>performance</b> in mind.</em>
</p>

---

# ERD untuk ERP Manajemen Rumah Sakit

Diagram ini menggambarkan Entity-Relationship Diagram (ERD) untuk sistem manajemen rumah sakit. Sistem ini mencakup berbagai entitas yang penting untuk pengelolaan rumah sakit, termasuk pasien, dokter, perawat, janji temu, rawat inap, kamar, obat, resep, resep detail, dan tagihan.

## Entitas dan Atribut

### 1. User
- **ID User**: int auto increment
- **Nama**: string not null
- **Alamat**: text 
- **Username**: string unik
- **Email**: string unik
- **Password**: string
- **Role**: enum
- **is email verifed**: boolean
- **Tanggal Lahir**: date
- **Jenis Kelamin**: string
- **Nomor Telepon**: string unik
- **Created At**: date
- **Update At**: date

### 2. Dokter
- **ID Dokter**: int
- **Nama**: string
- **Email**: string unik
- **Password**: string
- **Role**: enum
- **is email verifed**: boolean
- **Spesialisasi**: string
- **Nomor Telepon**: string
- **Created At**: date
- **Update At**: date

### 3. Perawat
- **ID Perawat**: int
- **Nama**: string
- **Email**: string unik
- **Username**: string unik
- **Password**: string
- **Role**: enum
- **Shift**: string
- **Nomor Telepon**: string
- **Created At**: date
- **Update At**: date

### 4. Janji Temu
- **ID Janji Temu**: int
- **ID Pasien**: int
- **ID Dokter**: int
- **Tanggal**: date
- **Waktu**: time

### 5. Rawat Inap
- **ID Rawat Inap**: int
- **ID Pasien**: int
- **ID Kamar**: int
- **Tanggal Masuk**: date
- **Tanggal Keluar**: date

### 6. Kamar
- **ID Kamar**: int
- **Jenis Kamar**: string
- **Tarif Per Hari**: decimal

### 7. Obat
- **ID Obat**: int
- **Nama Obat**: string
- **Deskripsi**: string
- **Harga**: decimal

### 8. Resep
- **ID Resep**: int
- **ID Pasien**: int
- **ID Dokter**: int
- **Tanggal**: date
- **Total Harga**: decimal

### 9. Resep Detail
- **ID Resep**: int
- **ID Obat**: int
- **Jumlah**: int
- **Harga**: decimal

### 10. Tagihan
- **ID Tagihan**: int
- **ID Pasien**: int
- **Tanggal**: date
- **Total Jumlah**: decimal

## Hubungan Antar Entitas

- **Pasien** `1:N` **Janji Temu**
- **Dokter** `1:N` **Janji Temu**
- **Pasien** `1:N` **Rawat Inap**
- **Kamar** `1:N` **Rawat Inap**
- **Pasien** `1:N` **Resep**
- **Dokter** `1:N` **Resep**
- **Resep** `1:N` **Resep Detail**
- **Obat** `1:N` **Resep Detail**
- **Pasien** `1:N` **Tagihan**

## Diagram ERD

![ERD Hospital Management System](ERD_Hospital_Management_System.png)

### Penjelasan Hubungan
1. **Pasien dan Janji Temu**: Seorang pasien dapat memiliki banyak janji temu, dan setiap janji temu dihadiri oleh satu pasien dan satu dokter.
2. **Pasien dan Rawat Inap**: Seorang pasien dapat dirawat di banyak kamar (berbeda waktu), dan setiap kamar dapat menampung banyak pasien (berbeda waktu).
3. **Pasien dan Resep**: Seorang pasien dapat memiliki banyak resep yang diberikan oleh dokter.
4. **Resep dan Resep Detail**: Setiap resep dapat terdiri dari banyak obat.
5. **Pasien dan Tagihan**: Seorang pasien dapat memiliki banyak tagihan yang terkait dengan perawatan dan pengobatan yang mereka terima.

## Cara Membaca Diagram

- Kotak mewakili entitas.
- Setiap entitas memiliki atribut yang tercantum di dalamnya.
- Garis penghubung antara entitas menunjukkan hubungan antara mereka, dengan simbol `1` dan `N` yang menunjukkan kardinalitas (misalnya, `1:N` untuk satu ke banyak).

## Contoh Visual

Untuk visualisasi diagram, gunakan alat seperti draw.io, Lucidchart, atau alat diagram lainnya untuk menggambar entitas sebagai kotak dan hubungkan dengan garis yang menunjukkan hubungan antar entitas.
