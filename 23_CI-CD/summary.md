23 CI/CD

pada materi ini saya mempelajari tentang CI/CD

CI/CD adalah singkatan dari Continuous Integration (CI) dan Continuous Delivery (CD), dua praktik penting dalam pengembangan perangkat lunak modern. Ini adalah pendekatan yang digunakan untuk memastikan perangkat lunak dapat dikembangkan, diuji, dan diterapkan secara otomatis dan berkelanjutan.

Continuous Integration (CI)
Continuous Integration adalah praktik pengembangan perangkat lunak di mana pengembang secara teratur menggabungkan kode yang telah mereka ubah ke dalam repositori bersama atau pusat kode. Setiap kali perubahan kode disubmit, sistem otomatis akan melakukan sejumlah tindakan, seperti kompilasi kode, pengujian unit, dan analisis statis. Tujuannya adalah untuk memastikan bahwa perubahan kode baru tidak merusak fungsionalitas yang ada dan berinteraksi dengan kode yang sudah ada dengan benar.

- Automatisasi: Semua tindakan yang terkait dengan penggabungan kode harus otomatis, termasuk pengujian dan pembangunan.
- Integrasi Berkelanjutan: Kode baru harus diintegrasikan dengan kode yang sudah ada secara teratur, misalnya beberapa kali sehari.
- Pengujian Otomatis: Pengujian otomatis, seperti pengujian unit dan integrasi, harus dilakukan untuk memeriksa perubahan kode.
- Pemberian Laporan: Hasil pengujian dan laporan kesalahan harus segera dikomunikasikan kepada pengembang.

Continuous Delivery (CD)
Continuous Delivery adalah praktik yang lebih luas yang melibatkan otomatisasi proses pengujian, pembangunan, dan pengiriman perangkat lunak ke lingkungan produksi atau lingkungan yang setara secara cepat dan otomatis. CD bertujuan untuk menjaga perangkat lunak dalam kondisi selalu siap untuk dirilis, meskipun keputusan untuk merilis atau tidak tetap dalam kendali manusia.

- Otomatisasi Pengiriman: Proses pengiriman perangkat lunak harus sepenuhnya otomatis, termasuk pengujian akhir, penerapan, dan konfigurasi.
- Lingkungan yang Reproduksi: Lingkungan produksi harus disimulasikan secara akurat untuk menguji perangkat lunak.
- Pemeriksaan Manual: Meskipun pengiriman otomatis, manusia masih dapat memeriksa perangkat lunak sebelum merilisnya, jika diperlukan.
- Dukungan Rollback: Harus ada kemampuan untuk dengan cepat mengembalikan perangkat lunak ke versi sebelumnya jika ada masalah saat rilis.

Kombinasi CI dan CD menciptakan alur kerja pengembangan yang kuat, di mana perubahan kode diuji secara otomatis, dan perangkat lunak selalu siap untuk dirilis ke produksi dengan risiko minimum. Ini membantu mengurangi kesalahan, meningkatkan kualitas perangkat lunak, dan mempercepat waktu pengiriman produk. CI/CD sering digunakan dalam konteks pengembangan perangkat lunak berbasis cloud dan aplikasi web, tetapi juga dapat diterapkan dalam berbagai jenis proyek pengembangan perangkat lunak.
