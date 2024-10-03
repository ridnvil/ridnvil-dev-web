import React from 'react';
import { motion } from 'framer-motion';

const FadeInRight = ({children}) => {
    return (
        <motion.div
            className="my-8"
            initial={{opacity: 0, x: 100}} // Mulai dari opacity 0 dan x 100 (kanan)
            animate={{opacity: 1, x: 0}}    // Berakhir di opacity 1 dan x 0
            transition={{duration: 2}}
        >
            {children}
        </motion.div>
    );
};

export default FadeInRight;