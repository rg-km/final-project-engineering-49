import React from "react";
import Login from "./Component/Auth/Login";
import Register from "./Component/Auth/Register";
import CourseContent from "./Component/Dashboard/Dashboard-User-Course-Content";
import Navbar from "./Component/Navbar/Navbar";
import Sidebar from "./Component/Sidebar/Sidebar";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import Footer from "./Component/Footer/Footer";
import ListMateriAdmin from "./Component/Admin/ListMateriAdmin";

export default function App() {
    return (
    <div className="Sidebar">
        <ListMateriAdmin />
        {/* <Footer /> */}
        <BrowserRouter>
            <Sidebar />
            <Switch>
                <Route></Route>
            </Switch> 
        </BrowserRouter>
    </div>
    )
}


