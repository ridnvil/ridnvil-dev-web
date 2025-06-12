import React, {useEffect, useState} from 'react';
import MyNavbar from "../components/MyNavbar";
import ThemeSwitcher from "../components/ThemeSwitcher";
import {useDispatch} from "react-redux";
import {getExperiences} from "../api/experiences";

const Experience = () => {
    const dispatch = useDispatch();
    const [experiences, setExperiences] = useState([]);

    const fetchExperiences = async () => {
        return await getExperiences()
    }

    useEffect(() => {
        fetchExperiences().then(r => {
            setExperiences(r);
            console.log(r);
        })
    }, [])

    return (
        <div>
            <ThemeSwitcher />
            <div className='dark:bg-blue-950 dark:text-white'>
                <MyNavbar />
                <div className='h-screen dark:bg-blue-950 p-10 flex items-center justify-center'>
                    {experiences.map((experience) => (
                        <div key={experience.id} className='grid grid-cols-3 gap-4'>
                            <div className='bg-white dark:bg-gray-800 p-4 rounded-lg shadow-md'>
                                <h2 className='text-xl font-semibold'>{experience.position}</h2>
                                <p className='text-sm text-gray-500 dark:text-gray-400'>Company: {experience.company_name}</p>
                                <p className='text-sm text-gray-500 dark:text-gray-400'>Duration: {experience.start_years + "-" + experience.until_years}</p>
                                <p className='text-gray-600 dark:text-gray-300'>{experience.description}</p>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default Experience;