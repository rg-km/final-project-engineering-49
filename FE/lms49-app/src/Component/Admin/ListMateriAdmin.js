import React from "react";
import "./Style/ListMateriAdmin.css";
import DeleteIcon from "../../Assets/del_butt.png";
import EditIcon from "../../Assets/edit_butt.png";

function ListMateriAdmin() {
  return (
    <div className="container">
      <div className="App">
        <div className="gap">
          <text className="text">Courses</text>
        </div>
        <div>
          <button className="button">
            <text className="textButton">Tambah Data</text>
          </button>
        </div>
      </div>
      <div className="main">
        <div className="component">
          <div className="image">Gambar</div>
          <div className="footerBox">
            <div className="title">
              <text className="subTitle1">HTML</text>
              <text className="subTitle2">Admin 2</text>
            </div>
            <div>
              <button className="button1"> 
              <img src={DeleteIcon} alt="delete" />
               </button> 
              <button className="button2">
              <img src={EditIcon} alt="edit" />
              </button> 
            </div>
          </div>
        </div>
        <div className="component">
          <div className="image">Gambar</div>
          <div className="footerBox">
            <div className="title">
              <text className="subTitle1">HTML</text>
              <text className="subTitle2">Admin 2</text>
            </div>
            <div>
              <button className="button1"> 
              <img src={DeleteIcon} alt="delete" />
               </button> 
              <button className="button2">
              <img src={EditIcon} alt="edit" />
              </button> 
            </div>
          </div>
        </div>
        <div className="component">
          <div className="image">Gambar</div>
          <div className="footerBox">
            <div className="title">
              <text className="subTitle1">HTML</text>
              <text className="subTitle2">Admin 2</text>
            </div>
            <div>
            <button className="button1"> 
              <img src={DeleteIcon} alt="delete" />
               </button> 
              <button className="button2">
              <img src={EditIcon} alt="edit" />
              </button> 
            </div>
          </div>
        </div>
        <div className="component">
          <div className="image">Gambar</div>
          <div className="footerBox">
            <div className="title">
              <text className="subTitle1">HTML</text>
              <text className="subTitle2">Admin 2</text>
            </div>
            <div>
            <button className="button1"> 
              <img src={DeleteIcon} alt="delete" />
               </button> 
              <button className="button2">
              <img src={EditIcon} alt="edit" />
              </button> 
            </div>
          </div>
        </div>
        <div className="component">
          <div className="image">Gambar</div>
          <div className="footerBox">
            <div className="title">
              <text className="subTitle1">HTML</text>
              <text className="subTitle2">Admin 2</text>
            </div>
            <div>
            <button className="button1"> 
              <img src={DeleteIcon} alt="delete" />
               </button> 
              <button className="button2">
              <img src={EditIcon} alt="edit" />
              </button> 
            </div>
          </div>
        </div>
        <div className="component">
          <div className="image">Gambar</div>
          <div className="footerBox">
            <div className="title">
              <text className="subTitle1">HTML</text>
              <text className="subTitle2">Admin 2</text>
            </div>
            <div>
            <button className="button1"> 
              <img src={DeleteIcon} alt="delete" />
               </button> 
              <button className="button2">
              <img src={EditIcon} alt="edit" />
              </button> 
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default ListMateriAdmin;