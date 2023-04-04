# Challenge-1 Chapter 3 FGA Golang

## Soal

Buatlah sebuah service untuk melakukan POST data dalam format JSON dan secara acak setiap 15 detik dengan angka random antara 1-100 untuk value water dan wind. Gunakan url post berikut untuk menjalankan simulasi post request :

    https://jsonplaceholder.typicode.com/posts

Kemudian tampilkan pada terminal hasil postnya. Selain itu kalian harus menentukan status water dan wind tersebut.

Dengan ketentuan:

-   jika water dibawah 5 maka status aman
-   jika water antara 6 - 8 maka status siaga
-   jika water diatas 8 maka status bahaya
-   jika wind dibawah 6 maka status aman
-   jika wind antara 7 - 15 maka status siaga
-   jika wind diatas 15 maka status bahaya
-   value water dalam satuan meter
-   value wind dalam satuan meter per detik
