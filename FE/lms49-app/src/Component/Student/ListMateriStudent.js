import axios from "axios";
import React from "react";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router";
import "./Style/ListMateriStudent.css";
import Navbar from "../Navbar/Navbar";
import Sidebar from "../Sidebar/Sidebar";
import Footer from "../Footer/Footer";
import Gambar from "../../Assets/couses_1.png";

function ListMateriStudent() {
    return (
      <div className="listMateriStudent">
        <div className="ComponentNavbar">
          <Navbar />
        </div>
        <div className="courseContainer">
            <div className="CourseImage">
                <img src={Gambar} alt="course" />
            </div>
            <div className="CourseInfoContainer">
                <div className="CourseInfoText">
                    <h5>HTML</h5>
                    <h6>Admin 2</h6>
                </div>
                <div className="CourseInfoButton">
                    <button>Read Now</button>
                </div>
            </div>
        </div>
      </div>
    );
}

export default ListMateriStudent;