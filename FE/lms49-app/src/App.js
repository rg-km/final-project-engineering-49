import React from "react";
import Login from "./Component/Auth/Login";
import Register from "./Component/Auth/Register";
import CourseContent from "./Component/Dashboard/Dashboard-User-Course-Content";
import Navbar from "./Component/Navbar/Navbar";
import Sidebar from "./Component/Sidebar/Sidebar";
import Footer from "./Component/Footer/Footer";
import ListMateriAdmin from "./Component/Admin/ListMateriAdmin";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

export default function App() {
    return (
      <Router>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/" element={<CourseContent />} />
          <Route path="/register" element={<Register />} />
          <Route path="/navbar" element={<Navbar />} />
        </Routes>
      </Router>
    );
}


