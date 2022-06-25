import React from "react";
import Login from "./Component/Auth/Login";
import Register from "./Component/Auth/Register";
import Navbar from "./Component/Navbar/Navbar";
import StudentDashboard from "./Component/Student/StudentDashboard";
import Sidebar from "./Component/Sidebar/Sidebar";
import Footer from "./Component/Footer/Footer";
import Logout from "./Component/Auth/Logout";
import ListMateriStudent from "./Component/Student/ListMateriStudent";
import DetailMateriStudent from "./Component/Student/DetailMateriStudent";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import CreateMateri from "./Component/Admin/CreateMateri";
import ListMateriAdmin from "./Component/Admin/ListMateriAdmin";

export default function App() {
    return (
      <Router>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/logout" element={<Logout />} />
          <Route path="/" element={<StudentDashboard />} />
          <Route path="/listmaterisiswa" element={<ListMateriStudent />} />
          <Route path="/detailmateristudent/:id" element={<DetailMateriStudent />} />
          <Route path="/admin/materi/create" element={<CreateMateri />} />
          <Route path="/admin/materi" element={<ListMateriAdmin />} />
        </Routes>
      </Router>
    );
}


