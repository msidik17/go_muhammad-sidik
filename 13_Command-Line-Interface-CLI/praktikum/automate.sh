#!/bin/bash

#Menampung input yang dimasukkan
name_folder="$1"
facebook_url="$2"
linkedin_url="$3"

current_date=$(date "+%Y-%m-%d %H:%M:%S")

#Membuat folder utama dengan nama ditambah waktu dijalankan program
main_folder="$name_folder at $current_date"
mkdir "$main_folder"

#Masuk ke folder utama yang telah dibuat
cd "$main_folder"

#Membuat folder about_me dan folder personal dan profesional di dalamnya
mkdir -p "about_me/personal"
mkdir -p "about_me/professional"

#Membuat folder my_friends
mkdir -p "my_friends"

#Membuat folder my_system_info
mkdir -p "my_system_info"

#Membuat file facebook.txt dan linkedin.txt di dalam folder about_me dengan argumen yang ditentukan
echo "https://www.facebook.com/$facebook_url" > "about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$linkedin_url" > "about_me/professional/linkedin.txt"

#Mendownload list friend dan dimasukkan ke file list_of_my_friend.txt
curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt > "./my_friends/list_of_my_friends.txt"

#Memasukkan isi nama dan host ke file about_this_laptop.txt ke folder my_system_info
echo "My Username: $USER" > "./my_system_info/about_this_laptop.txt"
echo "With Host : $(uname -a)" >> "./my_system_info/about_this_laptop.txt"

#Mengecek ping googke sebanyak 3x dan memasukkan ke file internet_connection.txt
ping -c 3 google.com > "my_system_info/internet_connection.txt"

#Menampilkan struktur data dalam bentuk tree(pohon)
tree

#Menampilkna bahwa program selesai dijalankan
echo "Program berhasil dijalankan dan telah selesai"