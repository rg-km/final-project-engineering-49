import axios from 'axios';
import React from 'react';
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useParams } from "react-router-dom";
import ReactPlayer from 'react-player';


function VideoJS() {

    const [ message, setMessage ] = useState("");
    const [ link, setLink ] = useState("");
    const [ token, setToken ] = useState("");
    const [ materiId, setMateriId ] = useState(0);
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
            headers: { Token: token}
        })
        .then((res) => {
            const response = res.data
            setLink(response.FileName);
        })
        .catch ((err) => {
            setMessage(err.message);
        })
    }, []);

    return (
        <div className='videoo'>
            
                <ReactPlayer url={link} controls={true}/>
            
        </div>
    );
}
export default VideoJS;