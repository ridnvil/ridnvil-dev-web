import {create} from "zustand";

export const useProfileStore = create((set) => ({
    profile: null,
    setProfile: (userData) => set({isLoggedIn: userData !== null, profile: userData}),
}));