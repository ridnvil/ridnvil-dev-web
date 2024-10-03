// src/components/SocialMediaIcons.js
import React from 'react';
import qbook from '../assets/qbook.jpg';
import inspektorat from '../assets/logopaniai.png';

const Portofolio = () => {
    return (
        <div className="flex space-x-5">
            {/* LinkedIn */}
            <a
                href="https://erp.qbook.id"
                title='Portofolio'
                target="_blank"
                rel="noopener noreferrer"
                className="text-blue-700 hover:text-blue-500 transition-colors duration-300"
            >
                <img src={qbook} className='w-[50px] rounded-full'  alt='qbook' />
            </a>
            {/* Telegram */}
            <a
                href="https://inspektorat.paniaikab.go.id"
                title='Portofolio'
                target="_blank"
                rel="noopener noreferrer"
                className="text-blue-400 hover:text-blue-300 transition-colors duration-300"
            >
                <img src={inspektorat} className='w-[45px]' alt='inspektorat'/>
            </a>
        </div>
    );
};

export default Portofolio;
