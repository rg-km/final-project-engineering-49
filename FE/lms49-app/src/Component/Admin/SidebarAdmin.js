import React from "react";
import "./Style/SidebarAdmin.css";
import LogoIcon from "../../Assets/Logo.png";
import CodeIcon from "../../Assets/code-block.png";


function SidebarAdmin() {
    return (
        <div className="sidebaradmin">
            <div className="nav_content">
                <div className="nav">
                    <a href="#">
                        <img src ={LogoIcon} alt="nav" className="nav" />
                    </a>
                </div>
                  <ul className="side-list">
                    <li>
                        <a href="#">
                        <img src={CodeIcon} alt="code" className="code"/>
                        <span className="links_name">Courses</span>
                        </a>
                         {/* <span className="tooltip">Courses</span> */}
                    </li>
                  </ul>
            </div>
        </div>
    )
}

export default SidebarAdmin;