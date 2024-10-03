import React from 'react';
import { motion } from 'framer-motion';

const FadeInSection = ({ children }) => {
    const variants = {
        hidden: { opacity: 0, y: 50 },
        visible: {
            opacity: 1,
            y: 0,
            transition: {
                duration: 2,
                ease: 'easeIn'
            }
        }
    };

    return (
        <motion.div
            initial="hidden"
            whileInView="visible"
            viewport={{ once: false, amount: 0.3 }}
            variants={variants}
            className="my-8"
        >
            {children}
        </motion.div>
    );
};

export default FadeInSection;
