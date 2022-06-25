import axios from "axios";
import { useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import React from "react";
import Navbar from "../Navbar/Navbar"
import "./Style/ListMateriStudent.css";

function ListMateriStudent() {

  const [ message, setMessage ] = useState("");
  const [ token, setToken ] = useState("");
  const [ materi, setMateri ] = useState([]);
  const [ keyword, setKeyword ] = useState("");
  const [ page, setPage ] = useState(1);
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

  const search = (e) => {
    const page = 1
    const limit = 10
    e.preventDefault();
    setPage(1);
    axios
      .get(
        "http://localhost:8080/course/filter/" +
          page +
          "/" +
          limit +
          "?keyword=" +
          keyword,
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      )
      .then((res) => {
        const response = res.data;
        setMateri(response.data.Materi);
      })
      .catch((error) => {
        setMessage(error.message);
      });
  };

  
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

        <div className="SearchBarComponent">
          <form onSubmit={search}>
            <input
              type="text"
              placeholder="Your Keyword Please..."
              value={keyword}
              onChange={(e) => setKeyword(e.target.value)}
            />
            <div className="iconSearch">
              <button type="submit">Search</button>
            </div>
          </form>
        </div>

        <div className="mainList">
          {materi &&
            [...materi].reverse().map((m, index) => {
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