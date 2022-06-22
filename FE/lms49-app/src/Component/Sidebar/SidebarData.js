import * as FaIcons from "react-icons/fa";
import * as RiIcons from "react-icons/ri";

export const SidebarData = [
    {
        id: 1,
        title: "Home",
        cName: "sidebar-item",
        icon: <FaIcons.FaHome />,
        path: "/",
    },
    {
        id: 2,
        title: "Dashboard",
        cName: "sidebar-item",
        icon: <RiIcons.RiDashboardLine />,
        path: "/dashboard",
    },
    {
        id: 3,
        title: "Courses",
        cName: "sidebar-item",
        icon: <FaIcons.FaFileCode />,
        path: "/courses",
    },
    {
        id: 4,
        title: "Logut",
        cName: "sidebar-item",
        icon: <RiIcons.RiLogoutBoxLine />,
        path: "/logut",
    },
]