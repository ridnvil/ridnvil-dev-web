// src/components/ProfilePhoto.js
import React from 'react';
import profilePic from '../assets/profile.png'; // Pastikan Anda memiliki foto di src/assets/

const ProfilePhoto = () => {
    return (
        <div className="flex justify-center">
            <img
                src={profilePic}
                alt="Profile"
                className="w-64 h-auto rounded-full object-cover border-4 border-gray-300"
            />
        </div>
    );
};

export default ProfilePhoto;
