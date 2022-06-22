import React from "react";
import "./Style/StudentDashboard.css";
import Navbar from "../Navbar/Navbar";
import Sidebar from "../Sidebar/Sidebar";
import Footer from "../Footer/Footer";
import Course from "../../Assets/course.png";

function StudentDashboard() {
    return (
      <div className="studentDashboard">
        <div className="NavbarComponent">
          <Navbar />
        </div>
        <div className="BoxContainer">
          <div className="HeaderDashPage">
            <h1>Dashboard</h1>
          </div>
          <div className="BoxStatCount">
            <div className="textStat">
                <h4>Jumlah Course</h4>
                <h5>0</h5>
            </div>
            <img className="logoCourse" src={Course}/>
          </div>
          <div className="BoxMessage">
            <h1>Box Pesan</h1>
          </div>
        </div>
      </div>
    );
}

export default StudentDashboard;