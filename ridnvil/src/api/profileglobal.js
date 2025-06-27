import axios from "axios";

export const getProfileGlobal = async () => {
    const response = await axios.get('/api/welcome', {
        credentials: true,
    });
    return response.data;
}