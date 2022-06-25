import React from 'react'
import Sidebar from '../Sidebar/Sidebar'
import SidebarAdmin from './SidebarAdmin'
import ListMateriAdmin from './ListMateriAdmin';
import TambahData from './TambahData';

function CreateMateri() {
  return (
    <div className='CreateMateri'>
        {/* <Sidebar /> */}
        {/* <SidebarAdmin /> */}
        {/* <ListMateriAdmin /> */}
        <TambahData />
    </div>
  )
}

export default CreateMateri;