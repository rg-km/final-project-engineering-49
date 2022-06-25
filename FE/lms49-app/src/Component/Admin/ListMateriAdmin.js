import axios from "axios";
import { useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import Navbar from "../Navbar/Navbar";
import React from "react";
import "./Style/ListMateriAdmin.css";
import DeleteIcon from "../../Assets/del_butt.png";
// import EditIcon from "../../Assets/edit_butt.png";

function ListMateriAdmin() {

  const [message, setMessage] = useState("");
  const [token, setToken] = useState("");
  const [materi, setMateri] = useState([]);
  const navigate = useNavigate();

  useEffect (() => {
    const token = localStorage.getItem("token");
    if (!token) {
      navigate("/login");
    }
    setToken(token);

    axios
      .get("http://localhost:8080/user", {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then((res) => {
        const response = res.data;
        if (response.data.Role != "admin") {
          navigate("/");
        }
      })
      .catch((error) => {
        navigate("/");
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

  const deleteMateri = (id) => {
    axios
      .delete("http://localhost:8080/course/delete/" + id, {
        headers: { Token: token },
      })
      .then((res) => {
        const response = res.data;
        setMessage(response.message);
      })
      .catch((error) => {
        setMessage(error.message);
      });
  };

  return (
    <div className="container">
      <div className="App">
        <div className="gap">
          <text className="text">Courses</text>
        </div>
        <div>
          <button className="button">
            <text className="textButton">Tambah Data</text>
          </button>
        </div>
      </div>
      <div className="main">

      {materi && [...materi].reverse().map((m, index) => {
        return (
          <div className="component" key={index}>
            <div className="image">Gambar</div>
            <div className="footerBox">
              <div className="title">
                <text className="subTitle1">{m.Title}</text>
                <text className="subTitle2">{m.Creator}</text>
              </div>
              <div>
                <button className="button1" onClick={() => deleteMateri(m.ID)}>
                  <img src={DeleteIcon} alt="delete" />
                </button>
                {/* <button className="button2">
                  <img src={EditIcon} alt="edit" />
                </button> */}
              </div>
            </div>
          </div>
        );
      })}
      </div>
    </div>
  );
}

export default ListMateriAdmin;