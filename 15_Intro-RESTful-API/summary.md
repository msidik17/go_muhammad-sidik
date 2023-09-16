15 Intro RESTful API

pada materi ini saya belajar banyak tentang API.
API adalah singkatan dari Application Programming Interface, yang merupakan sebuah set aturan dan protokol yang memungkinkan berbagai perangkat lunak atau komponen sistem untuk berinteraksi satu sama lain. API mengizinkan program komputer berkomunikasi dan berbagi data dengan cara yang terstruktur dan aman.

API digunakan untuk memungkinkan berbagai program berkomunikasi satu sama lain. Ini memungkinkan pengembang untuk memanfaatkan fungsionalitas yang ada tanpa perlu membangun semuanya dari awal.

ada beberapa element utama API seperti :

- EndPoint
  API sering kali memiliki beberapa titik akhir (endpoints) yang merepresentasikan operasi tertentu yang dapat dijalankan. Misalnya, sebuah API untuk aplikasi cuaca dapat memiliki endpoint untuk mendapatkan cuaca saat ini dan endpoint lainnya untuk mendapatkan ramalan cuaca.

- HTTP Method
  API menggunakan metode HTTP seperti GET, POST, PUT, dan DELETE untuk berinteraksi dengan data. Sebagai contoh, HTTP GET digunakan untuk mengambil data, sedangkan HTTP POST digunakan untuk mengirimkan data baru.

- Format Data
  Data yang dikirim dan diterima oleh API sering menggunakan format data yang telah ditentukan, seperti JSON (JavaScript Object Notation) atau XML (eXtensible Markup Language). JSON adalah format yang paling umum digunakan dalam API web modern.

- Authentication
  Authentication Untuk menjaga keamanan, sebagian besar API memerlukan otentikasi. Pengguna atau aplikasi perlu memberikan kredensial (seperti token API) untuk mengakses API tersebut.

dan ada beberapa jenis API seperti :

- API RESTful
  Ini adalah gaya arsitektur API yang umum digunakan untuk mengembangkan API web. API RESTful beroperasi dengan prinsip-prinsip seperti penggunaan metode HTTP (GET, POST, PUT, DELETE), penggunaan URL untuk mengakses sumber daya, dan penggunaan status kode HTTP untuk menunjukkan hasil operasi.

- API SOAP
  API Simple Object Access Protocol (SOAP) menggunakan protokol XML untuk pertukaran pesan. Ini lebih kompleks dan berat daripada REST, dan sering digunakan dalam aplikasi bisnis dan enterprise.

- API GraphQl
  Ini adalah jenis API yang memungkinkan klien untuk mengambil hanya data yang mereka butuhkan, sehingga mengurangi overfetching dan underfetching data.
