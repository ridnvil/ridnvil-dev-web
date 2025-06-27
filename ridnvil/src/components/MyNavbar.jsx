import React, {useState} from 'react';
import {Link, useNavigate} from "react-router-dom";
import {useAuthStore} from "../store/useAuthStore";
import {logoutApi} from "../api/auth";

const MyNavbar = () => {
    const navigate = useNavigate();
    const { logout } = useAuthStore();
    const handleLogout = async () => {
        await logoutApi().then((res) => {
            logout()
            navigate('/login');
        }).catch((err) => {
            console.log(err)
        })
    }

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