import Cookies from "js-cookie";

export const getProfileGlobal = async () => {
    try {
        const response = await fetch('/api/welcome');
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data = await response.json();
        return data;
    } catch (error) {
        console.error("Error fetching global profile data:", error);
        throw error;
    }
}