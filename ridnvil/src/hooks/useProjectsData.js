import {useQuery} from "@tanstack/react-query";
import {getProjectsApi} from "../api/projects";

export const useProjectsData = () => {
    return useQuery({
        queryKey: ['projects'],
        queryFn: getProjectsApi
    })
}