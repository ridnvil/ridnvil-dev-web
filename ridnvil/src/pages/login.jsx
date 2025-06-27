import React, {useState} from 'react';
import ThemeSwitcher from "../components/ThemeSwitcher";
import logo from '../assets/profile.png'
import {useAuthStore} from "../store/useAuthStore";
import {useAuthMutation} from "../hooks/useAuthData";
import {Link, useNavigate} from "react-router-dom";
import {errorHandler} from "../components/ErrorHandler";

const Login = () => {
    const { login } = useAuthStore()
    const navigate = useNavigate();
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const { mutate, isLoading, error } = useAuthMutation();

    const handleSubmit = (e) => {

        e.preventDefault();
        if (!email || !password) {
            alert('Email dan password tidak boleh kosong');
            return;
        }

        let data = {
            email,
            password,
        }

        mutate(data, {
            onSuccess: (data) => {
                login(data);
                navigate('/profile');
            },
            onError: (error) => {
                errorHandler(error);
            }
        })
    }

    if (error) {
        errorHandler(error);
    }

    return (
        <div className='h-screen bg-white dark:bg-blue-950 p-8 flex flex-col items-center'>
            <ThemeSwitcher />
            <div className='felx items-center justify-center'>
                <div className='flex flex-col items-center justify-center'>

                </div>
                <div className='flex h-full flex-col items-center justify-center lg:justify-start'>
                    <div className='m-4'>
                        <img
                            src={logo}
                            className="w-[200px] rounded-full"
                            alt="Sample image"
                        />
                    </div>
                    <form onSubmit={handleSubmit}>
                        <div className='w-[300px] flex flex-col dark:bg-black bg-white rounded p-10 shadow-lg'>
                            <input
                                value={email}
                                type="email"
                                onChange={(e) => setEmail(e.target.value)}
                                className="mb-6 border border-white dark:bg-black dark:text-white dark:border-amber-50 p-1 rounded"
                            />
                            <input
                                type="password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                className="mb-6 border border-white dark:bg-black dark:text-white p-1 rounded"
                            />

                            <button
                                type="submit"
                                disabled={isLoading}
                                className="rounded border bg-white dark:bg-black px-7 pb-2.5 pt-3 text-sm font-medium uppercase text-blue-950 dark:text-amber-50 hover:bg-amber-50 dark:hover:bg-blue-950"
                            >
                                {isLoading ? 'Logging in...' : 'Login'}
                            </button>
                        </div>
                    </form>
                    <div className='m-5'>
                        <Link to={"/"} className='text-white'>Back</Link>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Login;