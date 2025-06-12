import React from 'react';

const Biodata = ({name, email, phone, address}) => {
    return (
        <div className="bg-white bg-opacity-85 p-6 rounded shadow-md text-right dark:bg-blue-950">
            <ul className="space-y-2 text-blue-950 dark:text-amber-50">
                <li><strong>Name:</strong> {name}</li>
                <li><strong>Email:</strong> {email}</li>
                <li><strong>Whatsapp:</strong> {phone}</li>
                <li><strong>Address:</strong> {address}</li>
            </ul>
        </div>
    );
};

export default Biodata;
