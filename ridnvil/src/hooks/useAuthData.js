import {useMutation} from "@tanstack/react-query";
import axios from "axios";

const loginAuth = async (credential) => {
    const response = await axios.post('/api/login', {email: credential.email, password: credential.password});
    if (response.status !== 200) {
        throw new Error('Network response was not ok');
    }
    return response.data;
}

export const useAuthMutation = () => {
    return useMutation({
        mutationFn: loginAuth,
    })
}