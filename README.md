# dot-test
## Muhammad Rizal R - Backend Engineer Freelance

Ini adalah sebuah API sederhana yang digunakan untuk mengelola postingan dengan menerapkan syarat-syarat berikut:
1. Menggunakan HTTP VERB GET, POST, PUT, PATCH, dan DELETE.
2. Menyimpan data menggunakan ORM. Minimal harus terdapat satu relasi antar tabel dan DB Transaction.
3. Menggunakan pola yang umum digunakan. Saya memilih menggunakan Repository pattern.
4. Menambahkan penanganan error yang terpusat.
5. Terdapat fitur caching, dimana cache akan dihapus setiap kali ada permintaan dengan VERB selain GET.

## Alasan menggunakan Repository pattern pada project ini
Repository pattern, juga dikenal sebagai Clean Architecture, adalah pattern yang umum digunakan dalam pengembangan API. 
Beberapa alasan mengapa saya memilih pattern ini adalah sebagai berikut:

1. Pemisahan logika berdasarkan layer:
  - Layer Handler digunakan untuk proses komunikasi antara antarmuka dan service, dan semua proses I/O terjadi di layer ini.
  - Layer Usecase digunakan untuk semua logika bisnis yang terjadi.
  - Layer Repository bertanggung jawab untuk mengelola semua yang berhubungan dengan data.

2. Memudahkan dalam melakukan pengujian: Dengan menggunakan pattern ini, proses pengujian akan menjadi lebih mudah karena pengembang dapat membuat unit testing untuk setiap layer.
   Selain itu, dengan pemisahan logika, pengembang dapat dengan mudah membuat mock/mocking service yang diperlukan.
   
3. Telah menjadi pattern yang umum digunakan, sehingga memudahkan kerjasama antar pengembang. Ini juga didukung dengan arsitektur yang terstruktur.

5. Kode menjadi lebih mudah untuk digunakan kembali.

Selain menggunakan Repository Pattern, saya juga menerapkan Pattern DDD (Domain Driven Design), dimana setiap domain dipisahkan ke dalam folder-folder tersendiri. Ini dapat mempermudah ketika aplikasi menjadi besar dan kompleks, karena proses pemecahan menjadi microservice menjadi lebih mudah karena sudah dibedakan berdasarkan domain itu sendiri.


# Cara menjalankan program
1. git clone https://github.com/mrzalr/dot-test.git
2. cd dot-test
3. edit env.example file
4. import init.sql file
5. go run .
