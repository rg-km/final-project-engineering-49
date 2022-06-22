import React from "react";
import Navbar from "../Navbar/Navbar";
import "./Style/Dashboard-User-Course-Content.css";
import VideoJS from "./VideoJS";


const CourseContent = () => {
    return (
      <div className="courseContent">
        <Navbar />
        <h1>HTML</h1>
        <VideoJS />
        <p className="textDes">
          Aliquip enim cillum qui dolore veniam. Quis aute reprehenderit est
          adipisicing aute tempor. Pariatur occaecat dolor labore sint magna.
          Magna enim qui ea proident.
        </p>
        <p className="textDes">
          Est esse labore cupidatat pariatur exercitation labore ipsum culpa sit
          cillum consectetur consectetur duis est. Sit aliqua dolor cillum
          exercitation sit minim. Laborum incididunt qui minim quis cupidatat
          exercitation cillum non reprehenderit dolore. Incididunt esse id
          aliquip ex ipsum eiusmod fugiat enim.
        </p>
        <p className="textDes">
          Duis laborum est consequat incididunt ut. Qui ullamco ullamco aliqua
          est ullamco nisi aliqua. Minim cillum amet sit minim mollit consequat
          dolor velit do labore. Ullamco minim tempor irure id id in.
        </p>
      </div>
    );
}
export default CourseContent;