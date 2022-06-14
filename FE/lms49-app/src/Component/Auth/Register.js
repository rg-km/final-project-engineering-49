import React from "react";
import "./Style/Register.css";

function Register() {
  return (
    <div className="register">
      <div className="registerForm">
        <div className="container">
          <h1>Register</h1>
          <form>
            <label>Email</label>
            <input type="email" name="email" placeholder="Masukkan Email..." />

            <label>Username</label>
            <input type="text" name="username" placeholder="Masukkan Nama..." />

            <label>Password</label>
            <input
              type="password"
              name="password"
              placeholder="Masukkan Password..."
            />

            <label>Retype Password</label>
            <input
              type="password"
              name="repassword"
              placeholder="Ketik Ulang Password..."
            />

            <button type="submit">Register</button>
          </form>
        </div>
      </div>
    </div>
  );
}

export default Register;