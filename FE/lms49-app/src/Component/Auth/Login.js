import React from "react";
import axios from "axios";
import "./Style/Login.css";
import {useState, useEffect} from "react";
import { useNavigate } from "react-router-dom";

function Login() {

  const [ email, setEmail ] = useState("");
  const [ password, setPassword ] = useState("");
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
      password: password
    }

    axios
    .post("http://localhost:8000/login", user)
    .then(res => {
      const response = res.data
      localStorage.setItem("token", response.token);
      navigate("/");
    })
    .catch(err => {
      setMessage(err.message);
    })

  }


  return (
    <div className="login">
      <div className="loginForm">
        <div className="containerLogin">
          <h6>{message}</h6>
          <h1>Login</h1>
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
            <label>Password</label>

            <input
              type="password"
              name="password"
              placeholder="Masukkan Password..."
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
            <div className="buttonLogin">
              <button type="submit">Login</button>
            </div>
          </form>

          <a href="/register">
            <u>Register</u>
          </a>
        </div>
      </div>
    </div>
  );
}

export default Login;