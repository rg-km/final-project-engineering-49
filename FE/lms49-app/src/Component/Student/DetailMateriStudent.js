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
    }, []);
    
    return (
      <div className="courseContent">
        <Navbar />
        <h1>KA</h1>
        <div className='videoo'>
                <ReactPlayer url='https://www.youtube.com/watch?v=kUMe1FH4CHE' controls={true}/>
        </div>
        <p className="textDes">test</p>
      </div>
    );
}
export default DetailMateriStudent;