import axios from "axios";

export const getProjectsApi = async () => {
    const response = await axios.get('/api/projects', {
        credentials: true,
    });
    return response.data;
}

export const createProjectApi = async (project) => {
    const response = await axios.post('/api/projects', project, {
        credentials: true,
    });
    return response.data;
}