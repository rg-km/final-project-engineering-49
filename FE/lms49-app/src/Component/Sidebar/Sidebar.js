import React, { useState } from "react";
import "./Style/Sidebar.css";
import { Link } from "react-router-dom";
import * as FaIcons from "react-icons/fa";
import { SidebarData } from "./SidebarData";


function Sidebar() {
    const [sidebar, setSidebar] = useState(false);
    const showSidebar = () => {
        setSidebar(!sidebar);
    }

    return (
        <>
            <div className="sidebar">
                <Link to="/" className="side-menu-icon" onClick={showSidebar}>
                    <FaIcons.FaBars />
                 </Link>
            </div>
            <div 
                className={sidebar ? "sidebar-container active": "sidebar-container"}
            >
                <ul className="sidebar-items" onClick={showSidebar}>
                    <li className="sidebar-toggle">
                        <Link to="/" className="side-menu-icon">
                            <FaIcons.FaWindowClose />
                        </Link>
                    </li>
                    {SidebarData.map((sidebaritem) => {
                    return (
                        <li key={sidebaritem.id} className={sidebaritem.cName}>
                            <Link to={sidebaritem.path}>
                                {sidebaritem.icon}
                                <span>{sidebaritem.title}</span>
                            </Link>
                        </li>
                        );
                    })}
                </ul>
            </div>
        </>
    );
}

export default Sidebar;