// src/components/Pendidikan.js
import React from 'react';

const Pendidikan = () => {
    return (
        <div className="bg-white bg-opacity-85 p-6 rounded shadow-md text-blue-950 dark:bg-blue-950 dark:text-white">
            <h2 className="text-2xl font-semibold mb-4">Education</h2>
            <ul className="space-y-4">
                <li>
                    <h3 className="text-xl font-medium">STMIK Dipanegara Makassar</h3>
                    <p className='text-sm'>Sarjana Sistem Informasi (2011 - 2015)</p>
                </li>
            </ul>
        </div>
    );
};

export default Pendidikan;
