// src/components/FadeInFromSides.js
import React from 'react';
import { motion } from 'framer-motion';

const FadeInLeft = ({ children }) => {
    return (
        <motion.div
            className="my-5"
            initial={{opacity: 0, x: -100}} // Mulai dari opacity 0 dan x -100 (kiri)
            animate={{opacity: 1, x: 0}}    // Berakhir di opacity 1 dan x 0
            transition={{duration: 2}}
        >
            {children}
        </motion.div>
    );
};

export default FadeInLeft;
