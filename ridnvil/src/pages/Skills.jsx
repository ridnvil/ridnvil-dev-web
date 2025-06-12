import React from 'react';
import MyNavbar from "../components/MyNavbar";

const Skills = () => {
    return (
        <div className='dark:bg-blue-950 dark:text-white'>
            <MyNavbar />
            <div className='h-screen dark:bg-blue-950 p-10 flex items-center justify-center'>
                <h1>Skills</h1>
            </div>
        </div>
    );
};

export default Skills;