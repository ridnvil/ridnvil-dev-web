import {create} from "zustand";

export const useExperiencesStore = create((set) => ({
    experiences: [],
    setExperiences: (exp) => set({ experinces: exp }),
}));