import axios from "axios";

export const profileApi = async () => {
    const response = await axios.get("/api/profile", {
        credentials: true,
    });
    return response.data;
}