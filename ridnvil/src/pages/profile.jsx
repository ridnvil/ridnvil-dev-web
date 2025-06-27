import React from 'react';
import ThemeSwitcher from "../components/ThemeSwitcher";
import MyNavbar from "../components/MyNavbar";
import Loading from "../components/Loading";
import {useProfileData} from "../hooks/useProfileData";
import {useProfileStore} from "../store/useProfileStore";
import {errorHandler} from "../components/ErrorHandler";

const Profile = () => {
    const { data, isLoading, error } = useProfileData();
    const {profile} = useProfileStore();

    if (isLoading) return <Loading />;
    if (error) errorHandler(error);

    return (
        <div>
            <ThemeSwitcher />
            <div className="dark:bg-blue-950 dark:text-white text-gray-500">
                <MyNavbar />
                <div className='h-screen dark:bg-blue-950 p-10 flex items-center justify-center'>
                    <p>Profile</p>
                </div>
            </div>
        </div>
    );
};

export default Profile;