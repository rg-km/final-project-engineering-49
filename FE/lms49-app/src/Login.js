import React from "react";

import "./Style/Login.css";

function Login() {
  return (
    <div className="login">
      <div className="loginForm">
        <div className="container">
          <h1>Login</h1>
          <form>
            <label>Email</label>

            <input type="email" name="email" placeholder="Masukkan Email..." />
            <label>Password</label>

            <input
              type="password"
              name="password"
              placeholder="Masukkan Password..."
            />

            <button type="submit">Login</button>
          </form>

          <a href="/">
            <u>Register</u>
          </a>
        </div>
      </div>
    </div>
  );
}

export default Login;