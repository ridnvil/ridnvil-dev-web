import React from 'react';
import {useDispatch} from "react-redux";
import Cookies from "js-cookie";
import {logout} from "../features/authSlice";
import {Link} from "react-router-dom";

const MyNavbar = () => {
    const dispatch = useDispatch();
    const handleLogout = () => {
        Cookies.remove('token');
        localStorage.removeItem("user")
        dispatch(logout());
    };

    return (
        <nav className="">
            <div className="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
                <div className='relative flex h-16 items-center justify-between'>
                    <div className='flex flex-1 items-center justify-between'>
                        <div className="flex flex-shrink-0 items-center">
                            <Link to="/profile" className="text-2xl font-bold text-gray-900 dark:text-gray-100">
                                Ridnvil
                            </Link>
                        </div>
                        <div>
                            <ul className="flex space-x-5 dark:text-gray-100 text-gray-900">
                                <li>
                                    <Link to="/profile" className="dark:hover:text-white hover:text-blue-500">Home</Link>
                                </li>
                                <li>
                                    <Link to="/experiences" className="dark:hover:text-white hover:text-blue-500">Experiences</Link>
                                </li>
                                <li>
                                    <Link to="/projects" className="dark:hover:text-white hover:text-blue-500">Projects</Link>
                                </li>
                                <li>
                                    <Link to="/skills" className="dark:hover:text-white hover:text-blue-500">Skills</Link>
                                </li>
                                <li>
                                    <button onClick={handleLogout} className="dark:hover:text-white hover:text-blue-500">Logout</button>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
        </nav>
    );
};

export default MyNavbar;