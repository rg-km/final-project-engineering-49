import axios from "axios";
import React from "react";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useParams } from "react-router-dom";
import Navbar from "../Navbar/Navbar";
import "./Style/Dashboard-User-Course-Content.css";
import VideoJS from "./VideoJS";


const CourseContent = () => {

    const [message, setMessage] = useState("");
    const [token, setToken] = useState("");
    const [materiId, setMateriId] = useState(0);
    const [title, setTitle] = useState("");
    const [contain, setContain] = useState("");
    const params = useParams();
    const navigate = useNavigate();

    useEffect(() => {
      const token = localStorage.getItem("token");
      if (!token) {
        navigate("/login");
      }
      setToken(token);
      setMateriId(params.id);

      axios
        .get("http://localhost:8080/course" + params.id, {
          headers: { Token: token },
        })
        .then((res) => {
          const response = res.data;
          setTitle(response.Title);
          setContain(response.Contain);
        })
        .catch((err) => {
          setMessage(err.message);
        });
    }, []);

    
    return (
      <div className="courseContent">
        <Navbar />
        <h1>{title}</h1>
        <VideoJS />
        <p className="textDes">{contain}</p>
      </div>
    );
}
export default CourseContent;