import axios from "axios";

export const logoutApi = async () => {
    const response = await axios.post("/api/logout", {}, {
        credentials: true,
    })

    return response.data;
}