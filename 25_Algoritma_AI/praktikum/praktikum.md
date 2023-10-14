- SOAL
Terdapat sekumpulan data mengenai tulisan dalam bentuk tweet mengenai sebuah kebijakan. Sekumpulan data tersebut ingin dikelompokkan berdasarkan sentimen dari tweet tersebut yaitu sentimen positif dan negatif. Jelaskan algoritma A.I. yang dapat digunakan untuk mengelompokkan tweet tersebut beserta alasannya.

- Jawaban 
untuk pengelompokan data yang berbentuk tweet dan dikelompokkan berdasarkan sentimen, saya menggunakan algoritma random forest. algortima ini sendiri merupakan ensemble yang digunakan untuk tugas klasifikasi dan regresi. Ensemble berarti algoritma ini menggabungkan beberapa model pembelajaran mesin (biasanya decision tree) untuk meningkatkan kinerja dan mengurangi overfitting. 

- Alasan 
 Saya memilih aloritma ini karena algoritma Random Forest adalah algoritma ensemble yang bekerja dengan baik untuk masalah klasifikasi teks. Ini memungkinkan Anda untuk mengatasi kompleksitas tanda (features) dalam tweet Anda. Ensemble model menggunakan banyak decision tree untuk memperbaiki performa dan mengatasi overfitting.

 - Tahapan penggunaan algoritma 
 
 1. Pengumpulan data : Tahap pertama adalah mengumpulkan data yang sesuai untuk tugas yang ingin Anda selesaikan. Ini bisa berupa data klasifikasi atau regresi, tergantung pada tujuan Anda.

 2. Preprocessing Data : Langkah ini melibatkan pembersihan dan persiapan data. Anda perlu menghapus data yang hilang atau kosong, menangani data yang tidak valid, dan melakukan pengkodean ulang atau transformasi jika diperlukan. Selain itu juga perlu membagi data menjadi data latih dan data uji.

 3. Membangun decision tree : Random Forest adalah kumpulan dari decision tree. Setiap decision tree dibangun dengan menggunakan sub-kumpulan (subset) dari data latih. Ini disebut sebagai "bootstrap sample" yang berarti Anda mengambil sampel data dengan penggantian dari data latih. Setiap tree dibangun berdasarkan subset data dan hanya menggunakan subset acak dari fitur (fitur yang dipilih secara acak). Ini membantu dalam mengurangi overfitting dan meningkatkan keragaman tree.

 4. Memilih Pembagian Terbaik : Ketika membangun setiap tree, algoritma akan mencari pemisahan terbaik (split) pada setiap node berdasarkan kriteria tertentu, seperti Gini impurity atau Entropy untuk klasifikasi, atau Mean Squared Error (MSE) untuk regresi.

 5. Memperoleh Decision tree : tree terus dibagi menjadi cabang atau simpul (nodes) hingga beberapa kondisi berhenti terpenuhi, seperti mencapai kedalaman maksimum atau jumlah sampel minimum pada simpul. Setelah tree selesai dibangun, Anda memiliki decision tree lengkap.

 6. Menjalankan semua Decision tree : Setelah Anda membangun sejumlah tree (biasanya ratusan atau ribuan), langkah selanjutnya adalah menjalankan semua tree tersebut untuk memprediksi data uji atau data baru. Dalam tugas klasifikasi, masing-masing tree memberikan prediksi kelas, dan prediksi akhir diperoleh melalui pemilihan mayoritas dari semua tree. Dalam tugas regresi, prediksi akhir adalah rata-rata dari semua prediksi tree.

 7. Evaluasi Model 
 Setelah Anda mendapatkan prediksi dari model Random Forest, Anda perlu mengevaluasi kinerjanya. Anda dapat menggunakan metrik seperti akurasi, presisi, recall, F1-score, atau MSE (untuk regresi) untuk mengukur seberapa baik model berkinerja pada data uji.

 8. Tuning Hyperparameter : Anda dapat melakukan tuning hyperparameter untuk mengoptimalkan model, seperti mengatur jumlah tree, kedalaman maksimum tree, jumlah fitur yang digunakan dalam setiap tree, dan lain-lain.
