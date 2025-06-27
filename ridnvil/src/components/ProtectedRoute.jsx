import React from 'react';
import { Navigate } from 'react-router-dom';
import {useAuthStore} from "../store/useAuthStore";

function ProtectedRoute({ children }) {
    const { isLoggedIn, user } = useAuthStore();

    if (user) {
        if (!isLoggedIn) {
            return <Navigate to="/login" replace />;
        }
    }
    return children;
}

export default ProtectedRoute;