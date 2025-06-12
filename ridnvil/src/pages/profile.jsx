import React, {useEffect} from 'react';
import {useDispatch} from "react-redux";
import Cookies from "js-cookie";
import {logout} from "../features/authSlice";
import ThemeSwitcher from "../components/ThemeSwitcher";
import MyNavbar from "../components/MyNavbar";

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
        <div>
            <ThemeSwitcher />
            <div className="dark:bg-blue-950 dark:text-white text-gray-500">
                <MyNavbar />
                <div className='h-screen dark:bg-blue-950 p-10 flex items-center justify-center'>
                    <h1>Profile</h1>
                </div>
            </div>
        </div>
    );
};

export default Profile;