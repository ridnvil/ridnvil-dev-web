import Swal from "sweetalert2";
import {useState} from "react";

export const errorHandler = (message) => {
    const showAlert = () => {
        const dark = localStorage.getItem("theme");
        Swal.fire({
            title: 'Error!',
            text: message,
            icon: 'error',
            background: dark === "dark" ?'#333': '#fff',
            color: dark === "dark" ? '#fff': '#000',
            confirmButtonText: 'OK',
        });
    };

    return showAlert();
};