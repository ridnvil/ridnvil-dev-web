import React from 'react';
import MyNavbar from "../components/MyNavbar";
import ThemeSwitcher from "../components/ThemeSwitcher";

const Projects = () => {
    return (
        <div>
            <ThemeSwitcher />
            <div className='dark:bg-blue-950 dark:text-white'>
                <MyNavbar />
                {/* Main content area for projects make */}
                <div className='h-screen dark:bg-blue-950 p-10 flex items-center justify-center'>
                    <h1>Projects</h1>
                </div>
            </div>
        </div>
    );
};

export default Projects;