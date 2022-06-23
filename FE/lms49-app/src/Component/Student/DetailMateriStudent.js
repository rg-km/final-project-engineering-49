import axios from "axios";
import React from "react";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useParams } from "react-router-dom";
import Navbar from "../Navbar/Navbar";
import "./Style/DetailMateriStudent.css";
import ReactPlayer from 'react-player';


const DetailMateriStudent = () => {

    const [ message, setMessage ] = useState("");
    const [ token, setToken ] = useState("");
    const [ title, setTitle ] = useState("");
    const [ contain, setContain ] = useState("");
    const [ link, setLink ] = useState("");
    const params = useParams();
    const navigate = useNavigate();


    useEffect(() => {
        const token = localStorage.getItem("token");
        if (!token) {
            navigate("/login");
        }
        setToken(token);

        axios
            .get("http://localhost:8080/course/id/" + params.id, {
                headers: { Authorization: `Bearer ${token}` },
            })
            .then((res) => {
                const response = res.data;
                setTitle(response.data.Title);
                setContain(response.data.Contain);
                setLink(response.data.Link);
            })
            .catch((err) => {
                setMessage(err.message);
            })
    }, []);
    
    return (
      <div className="courseContent">
        <Navbar />
        <h5>{message}</h5>
        <h1>{title}</h1>
        <div className="videoo">
          <ReactPlayer url={link} controls={true} />
        </div>
        <p className="textDes">{contain}</p>
      </div>
    );
}
export default DetailMateriStudent;