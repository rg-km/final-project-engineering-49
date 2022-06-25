import axios from "axios";
import { useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import React from "react";
import Navbar from "../Navbar/Navbar"
import "./Style/ListMateriStudent.css";

function ListMateriStudent() {

  const [message, setMessage] = useState("");
  const [token, setToken] = useState("");
  const [materi, setMateri] = useState([]);
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
    .catch((err) => {
      navigate("/login");
    })
    
    axios
    .get("http://localhost:8080/courses", {
      headers: { Authorization: `Bearer ${token}` },
    })
    .then((res) => {
      const response = res.data;
      setMateri(response.data.Materi);
    })
    .catch((err) => {
      setMessage(err.message);
    })

  }, [])

  const detailMateri = (id) => {
    navigate("/detailmateri/" + id);
  }

    return (
      <div className="containerListMateri">
        <div className="NavbarComponentListMateri">
          <Navbar />
        </div>
        <div className="AppListMateri">
          <div className="gapPage">
            <text className="textPage">Courses</text>
          </div>
        </div>
        <div className="mainList">  

          {materi && [...materi].reverse().map((m, index) => {
            return (
              <div className="componentList" key={index}>
                <div className="imageListMat">Gambar</div>
                <div className="footerBox">
                  <div className="title">
                    <text className="subTitle1">{m.Title}</text>
                    <text className="subTitle2">{m.Creator}</text>
                  </div>
                  <div>
                    <button
                      className="button1"
                      onClick={() => detailMateri(m.ID)}
                    >
                      Detail Materi
                    </button>
                    <button className="button2">Read Now</button> 
                  </div>
                </div>
              </div>
            );
          })}
        </div>
      </div>
    );
  }
  
export default ListMateriStudent;