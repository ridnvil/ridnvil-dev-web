// src/components/SocialMediaIcons.js
import React from 'react';
import inspektorat from '../assets/logopaniai.png';

const Portofolio = () => {
    return (
        <div className="flex space-x-5">
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
