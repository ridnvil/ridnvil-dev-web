import React from 'react';
import {RouterProvider, createBrowserRouter} from "react-router-dom"
import Home from "./pages/home";
import Login from "./pages/login";
import ProtectedRoute from "./components/ProtectedRoute";
import Profile from "./pages/profile";
import Experience from "./pages/Experience";
import Projects from "./pages/Projects";
import Skills from "./pages/Skills";

function App() {
    const router = createBrowserRouter([
        { path: '/', element: <Home /> },
        { path: '/profile', element: <ProtectedRoute><Profile /></ProtectedRoute> },
        {path: '/experiences', element: <ProtectedRoute><Experience /></ProtectedRoute>},
        {path: '/projects', element: <ProtectedRoute><Projects /></ProtectedRoute>},
        {path: 'skills', element: <ProtectedRoute><Skills /></ProtectedRoute>},
        { path: '/login', element: <Login /> }
    ])
    return (
        <RouterProvider router={router}/>
    );
}

export default App;
