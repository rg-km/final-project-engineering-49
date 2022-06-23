import React from "react";
import "./Style/Navbar.css";
import { useState, useEffect } from "react";
import logofficon from "../../Assets/logout.png";
import SearchIcon from "../../Assets/search.png";

function Navbar() {

    const [dateState, setDateState] = useState(new Date());
    useEffect(() => {
      setInterval(() => setDateState(new Date()), 30000);
    }, []);

    return (
      <div className="containerNavbar">
        <div className="navbar">
          <div className="dateNavbar">
            <h3>
              {" "}
              {dateState.toLocaleDateString("en-GB", {
                day: "numeric",
                month: "short",
                year: "numeric",
              })}
            </h3>
          </div>

          <div className="profileNavbar">
            <a href="/logout">
              <img src={logofficon} alt="Avatar" className="avatarNavbar" />
            </a>
          </div>
        </div>
      </div>
    );
}

export default Navbar;