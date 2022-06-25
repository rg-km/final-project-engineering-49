import React from "react";
import "./Style/Sidebar.css";
import LogoIcon from "../../Assets/Logo.png";
import MenuIcon from "../../Assets/menu-regular-24.png";
import SearchIcon from "../../Assets/search-regular-24.png";
// import HomeIcon from "../../Assets/home-alt.png";
// import GridIcon from "../../Assets/grid-alt.png";
import CodeIcon from "../../Assets/code-block.png";
// import LogoutIcon from "../../Assets/log-out.png";


function Sidebar() {
    return (
            <div className="sidebar">
               <div className="logo_content">
                <div className="logo">
                    {/* <a href="#">
                        <img src={MenuIcon} alt="menu" className="menu"/>    
                    </a> */}
                    <img src={LogoIcon} alt="logo" className="logo" />
                </div>
                    <ul className="nav_list">
                        <li>
                            <a href="#">
                                {/* <img src={HomeIcon} alt="home" className="home"/>
                                <span className="links_name">Home</span> */}
                            </a>
                            {/* <span className="tooltip">Home</span> */}
                        </li>
                        <li>
                            {/* <a href="#">
                                <img src={GridIcon} alt="grid" className="grid"/>
                                <span className="links_name">Admin</span>
                            </a> */}
                            {/* <span className="tooltip">Dashboard</span> */}
                        </li>
                        <li>
                            <a href="#">
                                <img src={CodeIcon} alt="code" className="code"/>
                                <span className="links_name">Courses</span>
                            </a>
                            {/* <span className="tooltip">Courses</span> */}
                        </li>
                        <li>
                            <a href="#">
                                {/* <img src={LogoutIcon} alt="logout" className="logout"/>
                                <span className="links_name">Log-Out</span> */}
                            </a>
                            {/* <span className="tooltip">Dashboard</span> */}
                        </li>
                    </ul>
               </div>
            </div>
    );
}

export default Sidebar;