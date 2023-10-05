20 Clean Architecture

pada materi ini saya memperlajari konsep dasar pada clean architecture.
Clean Architecture adalah suatu pendekatan dalam pengembangan perangkat lunak yang dirancang untuk menciptakan struktur proyek yang terorganisir dengan baik dan mudah dipahami. Tujuannya adalah untuk memisahkan berbagai komponen perangkat lunak agar mudah dikelola, diuji, dan dikembangkan ulang tanpa perlu mengubah bagian-bagian lainnya. Konsep ini pertama kali diperkenalkan oleh Robert C. Martin, seorang tokoh terkenal dalam dunia pengembangan perangkat lunak.

Pada Clean Architecture, kode proyek dibagi menjadi beberapa lapisan (layers) yang berbeda. Setiap lapisan memiliki tanggung jawabnya sendiri dan harus terisolasi dengan baik dari lapisan-lapisan lainnya. Lapisan-lapisan tersebut biasanya mencakup:

- Presentation Layer : Lapisan ini bertanggung jawab untuk berinteraksi dengan pengguna. Ini bisa berupa antarmuka pengguna (user interface) atau API yang digunakan untuk berkomunikasi dengan pengguna. Lapisan ini seharusnya hanya berisi kode terkait tampilan dan interaksi dengan pengguna.
- Business Layer : Lapisan ini berisi logika bisnis dari aplikasi. Ini adalah tempat di mana aturan bisnis utama didefinisikan dan diimplementasikan. Lapisan ini harus mandiri dan tidak tergantung pada detail teknis atau infrastruktur tertentu.
- Data Layer : Lapisan ini bertanggung jawab untuk mengakses dan menyimpan data. Ini mencakup akses ke database, penyimpanan berkas, pemanggilan API eksternal, dan sejenisnya. Lapisan ini seharusnya terisolasi sepenuhnya dari lapisan bisnis dan presentasi.
- nfrastructure Layer : Lapisan ini berisi semua detail teknis seperti framework, library, dan alat yang digunakan dalam proyek. Lapisan ini berfungsi sebagai penghubung antara lapisan lainnya dan dunia luar.

Prinsip utama dari Clean Architecture adalah bahwa arus kontrol (control flow) selalu dimulai dari lapisan presentasi dan kemudian bergerak ke dalam lapisan bisnis, dengan lapisan data hanya digunakan saat diperlukan. Ini memungkinkan pengujian unit yang lebih baik, memudahkan perubahan, dan memastikan bahwa bisnis logic dapat digunakan kembali tanpa harus terkait dengan teknologi tertentu.

Clean Architecture juga mendorong penggunaan prinsip-prinsip desain seperti Dependency Inversion Principle (DIP) dan Single Responsibility Principle (SRP) untuk menjaga proyek tetap bersih dan terstruktur dengan baik. Dengan demikian, Clean Architecture membantu menciptakan perangkat lunak yang lebih skalabel, mudah dipelihara, dan berorientasi pada bisnis.
