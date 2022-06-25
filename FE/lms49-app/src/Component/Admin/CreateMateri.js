import React from 'react'
import Sidebar from '../Sidebar/Sidebar'
import SidebarAdmin from './SidebarAdmin'
import ListMateriAdmin from './ListMateriAdmin';
import TambahData from './TambahData';
import "./Style/CreateMateri.css";

function CreateMateri() {
  return (
    <div className="CreateMateri">
      <div className="CreateMateriForm">
        <div className="containerCreateMateri">
          <h2>ADD COURSE</h2>
          <form>
            <label>Nama Course</label>
            <input
              type="text"
              name="nama_course"
              placeholder="Input Judul..."
              required
            />

            <label>Deskripsi</label>
            <input
              type="text"
              name="deskripsi"
              placeholder="Input Deskripsi..."
              required
            />

            <label>Link Video</label>
            <input
              type="text"
              name="link_vid"
              placeholder="Input Link..."
              required
            />

            <label>Creator</label>
            <input
              type="text"
              name="creator_name"
              placeholder="Input Crator..."
              required
            />

            <div className="buttonCreate">
              <button type="submit">Submit</button>
            </div>
            <div className="buttonCreate12">
              <button type="submit">Cancel</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}

export default CreateMateri;