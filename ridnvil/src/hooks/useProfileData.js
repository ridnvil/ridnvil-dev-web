import {useQuery} from "@tanstack/react-query";
import {profileApi} from "../api/profile";
import {useProfileStore} from "../store/useProfileStore";

export const useProfileData = () => {
    const {profile, setProfile} = useProfileStore()
    return useQuery({
        queryKey: ["profileData"],
        queryFn: async () => {
            const res = await profileApi();
            setProfile(res)
            return res;
        },
        initialData: profile
    })
}