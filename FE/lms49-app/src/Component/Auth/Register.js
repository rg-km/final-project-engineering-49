import React from "react";
import axios from "axios";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import "./Style/Register.css";

function Register() {

  const [ email, setEmail ] = useState("");
  const [ name, setName ] = useState("");
  const [ password, setPassword ] = useState("");
  const [ confirmPassword, setConfirmPassword ] = useState("");
  const [ message, setMessage ] = useState("");
  const navigate = useNavigate();


  useEffect(() => {
    if (localStorage.getItem("token")) {
      navigate("/");
    }
  }, [])


  const handleSubmit = (e) => {
    e.preventDefault();
    const user = {
      email: email,
      name: name,
      password: password,
      confirmPassword: confirmPassword
    }

    axios.post("http://localhost:8000/register", user)
      .then((res) => {
        navigate("/login");
      })
      .catch(err => {
        setMessage(err.message);
      })
  }


  return (
    <div className="register">
      <div className="registerForm">
        <div className="containerRegister">
          <h6>{message}</h6>
          <h1>Register</h1>
          <form onSubmit={handleSubmit}>
            <label>Email</label>
            <input
              type="email"
              name="email"
              placeholder="Masukkan Email..."
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />

            <label>Username</label>
            <input
              type="text"
              name="username"
              placeholder="Masukkan Nama..."
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
            />

            <label>Password</label>
            <input
              type="password"
              name="password"
              placeholder="Masukkan Password..."
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />

            <label>Retype Password</label>
            <input
              type="password"
              name="repassword"
              placeholder="Ketik Ulang Password..."
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
              required
            />

            <div className="buttonRegister">
              <button type="submit">Register</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}

export default Register;