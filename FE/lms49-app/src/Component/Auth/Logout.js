import React from "react";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const Logout = () => {

    const navigate = useNavigate();

    useEffect(() => {
        localStorage.removeItem("token");
        navigate("/login");
    }, [])
}

export default Logout;