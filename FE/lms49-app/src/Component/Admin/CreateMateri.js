import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import React from 'react'
import "./Style/CreateMateri.css";

function CreateMateri() {

  const [ message, setMessage ] = useState("");
  const [ token, setToken ] = useState("");
  const [ title, setTitle ] = useState("");
  const [ contain, setContain ] = useState("");
  const [ link, setLink ] = useState("");
  const [ creator, setCreator ] = useState("");
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
      .then((res) => {
        const response = res.data;
        if (response.data.Role != "admin") {
          navigate("/");
        }
      })
      .catch((error) => {
        navigate("/");
      });
  }, [])

  const handleSubmit = (event) => {
    event.preventDefault();
    const data = {
      title: title,
      contain: contain,
      file_name: link,
      creator: creator,
    };
    axios
      .post("http://localhost:8080/course", data, {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then((res) => {
        setMessage("Success create course");
      })
      .catch((error) => {
        console.log(error);
        setMessage(error.message);
      });
  };

  return (
    <div className="CreateMateri">
      <div className="CreateMateriForm">
        <div className="containerCreateMateri">
          <h6>{message}</h6>
          <h2>ADD COURSE</h2>
          <form onSubmit={handleSubmit}>
            <label>Nama Course</label>
            <input
              type="text"
              name="nama_course"
              placeholder="Input Judul..."
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              required
            />

            <label>Deskripsi</label>
            <input
              type="text"
              name="deskripsi"
              placeholder="Input Deskripsi..."
              value={contain}
              onChange={(e) => setContain(e.target.value)}
              required
            />

            <label>Link Video</label>
            <input
              type="text"
              name="link_vid"
              placeholder="Input Link..."
              value={link}
              onChange={(e) => setLink(e.target.value)}
              required
            />

            <label>Creator</label>
            <input
              type="text"
              name="creator_name"
              placeholder="Input Crator..."
              value={creator}
              onChange={(e) => setCreator(e.target.value)}
              required
            />

            <div className="buttonCreate">
              <button type="submit">Submit</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}

export default CreateMateri;