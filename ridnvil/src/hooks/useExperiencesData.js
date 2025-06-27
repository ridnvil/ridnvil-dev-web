import {useQuery} from "@tanstack/react-query";
import {getExperiences} from "../api/experiences";
import {useExperiencesStore} from "../store/useExperiencesStore";

export const useExperiencesData = () => {
    const { experiences, setExperiences } = useExperiencesStore();
    return useQuery({
        queryKey: ["experiences"],
        queryFn: async () => {
            const res = await getExperiences();
            setExperiences(res);
            return res;
        },
        initialData: experiences,
    })
}