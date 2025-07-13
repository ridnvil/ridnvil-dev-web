import React from 'react';
import MyNavbar from "../components/MyNavbar";
import ThemeSwitcher from "../components/ThemeSwitcher";
import {useProjectsData} from "../hooks/useProjectsData";
import Loading from "../components/Loading";
import {errorHandler} from "../components/ErrorHandler";
import {PencilSquareIcon, PlusCircleIcon, PlusIcon} from "@heroicons/react/24/solid";
import {TrashIcon} from "@heroicons/react/16/solid";

const Projects = () => {
    const { data, isLoading, error } = useProjectsData()

    if (isLoading) return <Loading />
    if (error) errorHandler(error);

    return (
        <div>
            <ThemeSwitcher />
            <div className='dark:bg-blue-950 dark:text-white'>
                <MyNavbar />
                {/* Main content area for projects make */}
                <div className='flex items-center justify-between mt-20 pl-10 pr-10'>
                    <div>
                        <button className='flex flex-row'>
                            <PlusIcon className='w-5' />Add Project
                        </button>
                    </div>
                    <div>
                        <h1 className='text-2xl'>Experience Projects</h1>
                    </div>
                </div>
                <div className='h-screen dark:bg-blue-950 p-10 flex items-start justify-center'>
                    <div className='grid grid-cols-3 gap-2'>
                        { data.map((project) => (
                            <div key={project.id} className='bg-gray-200 dark:bg-gray-800 p-4 rounded-lg shadow-md'>
                                <div className='flex items-center justify-between p-2'>
                                    <button className='bg-gray-400/30 p-2 rounded-full'>
                                        <PencilSquareIcon  className='w-4'/>
                                    </button>
                                    <button className='bg-gray-400/30 p-2 rounded-full'>
                                        <PlusCircleIcon  className='w-4'/>
                                    </button>
                                    <button className='text-red-500'>
                                        <TrashIcon  className='w-4'/>
                                    </button>
                                </div>
                                <div className='flex items-center justify-between bg-gray-100 dark:bg-gray-900 rounded-t-md p-5'>
                                    <h1 className='text-2xl'>{project.name}</h1>
                                    <div>
                                        <p>Project Done: {project.length}</p>
                                        <p className='text-right mb-5'>Company: {project.company}</p>
                                    </div>
                                </div>
                                <div className='dark:text-white p-5'>
                                    <p className='text-gray-600 dark:text-gray-500'>{project.description}</p>
                                </div>
                            </div>
                        )) }
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Projects;