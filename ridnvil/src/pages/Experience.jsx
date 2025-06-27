import React, {useEffect, useState} from 'react';
import MyNavbar from "../components/MyNavbar";
import ThemeSwitcher from "../components/ThemeSwitcher";
import {useExperiencesData} from "../hooks/useExperiencesData";
import Loading from "../components/Loading";
import {errorHandler} from "../components/ErrorHandler";

const Experience = () => {
    const { data, isLoading, error } = useExperiencesData();

    if (isLoading) return <Loading />;
    if (error) errorHandler(error);

    return (
        <div>
            <ThemeSwitcher />
            <div className='dark:bg-blue-950 dark:text-white'>
                <MyNavbar />
                <div className='flex items-center justify-center mt-20'>
                    <h1 className='text-2xl'>Experinces</h1>
                </div>
                <div className='h-screen dark:bg-blue-950 p-10 flex items-start justify-center gap-5'>
                    {data.map((experience) => (
                        <div key={experience.id} className='w-1/2'>
                            <div className='bg-white dark:bg-gray-800 p-4 rounded-lg shadow-md'>
                                <h2 className='text-xl font-semibold'>{experience.position}</h2>
                                <p className='text-sm text-gray-500 dark:text-gray-400'>Company: {experience.company_name}</p>
                                <p className='text-sm text-gray-500 dark:text-gray-400'>Duration: {experience.start_years + "-" + experience.until_years}</p>
                                <p className='text-gray-600 dark:text-gray-500'>{experience.description}</p>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default Experience;