import Cookies from "js-cookie";

export const getExperiences = async () => {
    try {
        const token = Cookies.get('token')
        const response = await fetch('/api/experiences', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
            },
        });
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching experiences:', error);
        throw error;
    }
}