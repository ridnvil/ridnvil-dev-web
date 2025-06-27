import axios from "axios";

export const getExperiences = async () => {
    const response = await axios.get('/api/experiences', {
        credentials: true,
    });
    return response.data;
}