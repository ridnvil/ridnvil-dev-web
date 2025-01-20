import React, {useEffect} from 'react';
import {useDispatch} from "react-redux";
import Cookies from "js-cookie";
import {logout} from "../features/authSlice";
import ThemeSwitcher from "../components/ThemeSwitcher";

const Profile = () => {
    const dispatch = useDispatch();

    useEffect (() => {
        getProfile().then((res) => {
            console.log(res)
        }).catch(err => alert(err))
    }, [dispatch]);

    const handleLogout = () => {
        Cookies.remove('token');
        localStorage.removeItem("user")
        dispatch(logout());
    }

    const getProfile = () => {
        return new Promise((resolve, reject) => {
            const token = Cookies.get("token")
            fetch("/api/profile", {
                method: 'GET',
                headers: {
                    Authorization: `Bearer ${token}`,
                }
            })
                .then(async (res) => {
                    if (!res.ok) {
                        reject("error get profile")
                    }

                    resolve(await res.json())
                }).catch(err => reject(err))
        })
    }

    return (
        <div className="relative min-h-screen bg-transparent bg-white dark:bg-blue-950 p-8 flex flex-col items-center">
            <ThemeSwitcher />
            <h1 className='text-blue-950 dark:text-amber-50'>Profile</h1>
            <button className='text-blue-950 dark:text-amber-50 border' onClick={handleLogout}>Logout</button>
        </div>
    );
};

export default Profile;