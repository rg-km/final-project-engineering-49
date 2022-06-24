import axios from "axios";
import React from "react";
import "./Style/StudentDashboard.css";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import Navbar from "../Navbar/Navbar";
import Sidebar from "../Sidebar/Sidebar";
import Footer from "../Footer/Footer";
import Course from "../../Assets/course.png";
import StuImg from "../../Assets/siswa.png";

function StudentDashboard() {

    const [message, setMessage] = useState("");
    const [countOfMateri, setCountOfMateri] = useState(0);
    const [countOfStudent, setCountOfStudent] = useState(0);
    const [token, setToken] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (!token) {
            navigate("/login");
        }
        setToken(token);

        axios
        .get("http://localhost:8080/user", {
            headers: { Authorization: `Bearer ${token}` },
        })
        .then((res) => {})
        .catch((error) => {
        navigate("/login");
        });

        axios
          .get("http://localhost:8080/course/count", {
            headers: { Authorization: `Bearer ${token}` },
          })
          .then((res) => {
            const response = res.data;
            setCountOfMateri(response.data);
          })
          .catch((error) => {
            setMessage(error.message);
          });
        
        axios
          .get("http://localhost:8080/user/count", {
            headers: { Authorization: `Bearer ${token}` },
          })
          .then((res) => {
            const response = res.data;
            setCountOfStudent(response.data);
          })
          .catch((err) => {
            setMessage(err.message);
          });

    }, [])

    return (
      <div className="studentDashboard">
        <div className="NavbarComponent">
          <Navbar />
        </div>
        <div className="SidebarComponent">
          <Sidebar />
        </div>
        <div className="BoxContainer">
          <div className="HeaderDashPage">
            <h6>{message}</h6>
            <h2>Dashboard</h2>
          </div>
          <div className="BoxStatContainer">
            <div className="BoxStatCount">
              <div className="textStat">
                <h4>Jumlah Course</h4>
                <h5>{countOfMateri}</h5>
              </div>
              <img className="logoCourse" src={Course} />
            </div>
            <div className="BoxStatCount siswaStat">
              <div className="textStat">
                <h4>Jumlah Siswa</h4>
                <h5>{countOfStudent}</h5>
              </div>
              <img className="logoSiswa" src={StuImg} />
            </div>
          </div>
          <div className="BoxMessage">
            <h3>Selamat Datang Mahasiswa!</h3>
            <p>
              Di web ini kalian dapat menyaksikan materi dalam bentuk Video yang telah di susun oleh kampus loh!
            </p>
          </div>
          {/* <div className="ButtonToList">
            <a href="/listmaterisiswa">Lihat Materi</a>
          </div> */}
        </div>
        <div className="FooterComponent">
          <Footer />
        </div>
      </div>
    );
}

export default StudentDashboard;