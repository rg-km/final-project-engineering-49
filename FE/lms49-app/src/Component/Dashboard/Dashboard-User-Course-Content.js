import React from "react";
import Navbar from "../Navbar/Navbar";
import "./Style/Dashboard-User-Course-Content.css";
import VideoJS from "./VideoJS";


const CourseContent = () => {
    return (
        <div>
             <Navbar/>
            <h1>HTML</h1>
            <VideoJS/>
            <p className="text">
            HTML adalah bahasa standar pemrogaman yang digunakan untuk membuat halaman website, yang diakses melalui internet. 
Singkatan dari "Hypertext Markup Language" atau "bahasa markup". HTML disusun berdasar kode dan simbol tertentu, yang 
dimasukkan dalam sebuah file atau dokumen. Sehingga bisa ditampilkan pada layar komputer. Dan bisa dipahami oleh para 
pengguna internet.            </p>
            <p className="text">
            Memahami setiap kata yang terkandung, hypertext sendiri dimaksudkan sebagai metode yang digunakan untuk berpindah 
laman web ke laman lain. Usai mengklik tulisan atau simbol yang muncul di halaman website. HTML digunakan untuk 
membuat dokumen elektronik (disebut halaman) yang ditampilkan di World Wide Web (www). Setiap halaman berisi 
serangkaian koneksi ke halaman lain yang disebut hyperlink. Setiap halaman web yang Anda lihat di Internet ditulis 
menggunakan satu versi kode HTML atau yang lain.            </p>
            <p className="text">   
Kode HTML memastikan format teks dan gambar yang tepat untuk browser Internet. Tanpa HTML, browser tidak akan tahu 
bagaimana menampilkan teks sebagai elemen atau memuat gambar atau elemen lainnya. HTML juga menyediakan struktur 
dasar halaman, di mana Cascading Style Sheets dihamparkan untuk mengubah tampilannya.            </p> 
        </div>
    )
}
export default CourseContent;