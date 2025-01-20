import React from 'react';
import {RouterProvider, createBrowserRouter} from "react-router-dom"
import Home from "./pages/home";
import Login from "./pages/login";
import ProtectedRoute from "./components/ProtectedRoute";
import Profile from "./pages/profile";

function App() {
    const router = createBrowserRouter([
        { path: '/', element: <Home /> },
        { path: '/profile', element: <ProtectedRoute><Profile /></ProtectedRoute> },
        { path: '/login', element: <Login /> }
    ])
    return (
        <RouterProvider router={router}/>
    );
}

export default App;
