import React from "react";
import "./Style/Register.css";

function Register() {
  return (
    <div className="register">
      <div className="registerForm">
        <div className="containerRegister">
          <h1>Register</h1>
          <form>
            <label>Email</label>
            <input type="email" name="email" placeholder="Masukkan Email..." required/>

            <label>Username</label>
            <input type="text" name="username" placeholder="Masukkan Nama..." required/>

            <label>Password</label>
            <input
              type="password"
              name="password"
              placeholder="Masukkan Password..."
              required
            />

            <label>Retype Password</label>
            <input
              type="password"
              name="repassword"
              placeholder="Ketik Ulang Password..."
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