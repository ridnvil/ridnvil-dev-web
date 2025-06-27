import {useQuery} from "@tanstack/react-query";
import {getProfileGlobal} from "../api/profileglobal";

export const useWelcomeProfileData = () => {
    return useQuery({
        queryKey: ['welcome'],
        queryFn: getProfileGlobal,
    })
}