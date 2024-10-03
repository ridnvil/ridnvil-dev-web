// src/components/SocialMediaIcons.js
import React from 'react';
import { FaLinkedin, FaTelegram, FaGithub } from 'react-icons/fa';

const SocialMediaIcons = () => {
    return (
        <div className="flex space-x-4">
            {/* LinkedIn */}
            <a
                href="https://www.linkedin.com/in/rid-wan-57047b137"
                title='Linkedin'
                target="_blank"
                rel="noopener noreferrer"
                className="text-blue-700 hover:text-blue-500 transition-colors duration-300"
            >
                <FaLinkedin size={50} />
            </a>
            {/* Telegram */}
            <a
                href="https://t.me/rid_17"
                title='Telegram'
                target="_blank"
                rel="noopener noreferrer"
                className="text-blue-400 hover:text-blue-300 transition-colors duration-300"
            >
                <FaTelegram size={50} />
            </a>
            {/* GitHub */}
            <a
                href="https://github.com/ridnvil"
                title='Github'
                target="_blank"
                rel="noopener noreferrer"
                className="text-gray-100 hover:text-gray-500 transition-colors duration-300"
            >
                <FaGithub size={50} />
            </a>
        </div>
    );
};

export default SocialMediaIcons;
