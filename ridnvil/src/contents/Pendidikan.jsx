// src/components/Pendidikan.js
import React from 'react';

const Pendidikan = ({institude, major, duration}) => {
    return (
        <div className="bg-white bg-opacity-85 p-6 rounded shadow-md text-blue-950 dark:bg-blue-950 dark:text-white">
            <h2 className="text-2xl font-semibold mb-4">Education</h2>
            <ul className="space-y-4">
                <li>
                    <h3 className="text-xl font-medium">{institude}</h3>
                    <p className='text-sm'>{major} ({duration})</p>
                </li>
            </ul>
        </div>
    );
};

export default Pendidikan;
