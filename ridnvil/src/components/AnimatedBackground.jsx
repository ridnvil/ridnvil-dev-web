// src/components/AnimatedBackground.js
import React from 'react';
import { motion } from 'framer-motion';

const AnimatedBackground = () => {
    return (
        <motion.div
            className="absolute inset-0 -z-10"
            initial={{ backgroundPosition: '0% 50%' }}
            animate={{ backgroundPosition: '100% 50%' }}
            transition={{
                repeat: Infinity,
                repeatType: 'loop',
                duration: 10,
                ease: 'linear',
            }}
            style={{
                background: 'radial-gradient(circle, rgba(2,0,36,1) 0%, rgba(9,9,121,1) 35%, rgba(3,115,177,1) 100%)',
                backgroundSize: '400% 400%',
            }}
        ></motion.div>
    );
};

export default AnimatedBackground;
