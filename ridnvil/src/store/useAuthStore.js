import {create} from "zustand";

export const useAuthStore = create((set) => ({
    isLoggedIn: false,
    user: null,
    login: (dataUser) => set({ isLoggedIn: true, user: dataUser }),
    logout: () => set({ isLoggedIn: false, user: null }),
}));