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
        <div className='min-h-screen bg-white dark:bg-blue-950 p-8 flex flex-col items-center justify-center'>
            <ThemeSwitcher />
            <div className='felx items-center justify-center'>
                <div className='flex h-full flex-col items-center justify-center lg:justify-start'>
                    <div className='m-4'>
                        <img
                            src={logo}
                            className="w-[50px] rounded-full"
                            alt="Sample image"
                        />
                    </div>
                    <form onSubmit={handleSubmit}>
                        <div className='w-[350px] flex flex-col dark:bg-gray-900 bg-white rounded p-10 shadow-lg'>
                            <label htmlFor='email' className='dark:text-white'>Email</label>
                            <input
                                value={email}
                                type="email"
                                placeholder={email}
                                onChange={(e) => setEmail(e.target.value)}
                                className="mb-2 border border-white dark:bg-blue-950 dark:text-white dark:border-amber-50 p-1 rounded"
                            />
                            <label htmlFor='email' className='dark:text-white'>Password</label>
                            <input
                                type="password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                className="mb-10 border border-white dark:bg-blue-950 dark:text-white p-1 rounded"
                            />

                            <button
                                type="submit"
                                disabled={isLoading}
                                className="rounded border bg-white dark:bg-blue-950 px-7 pb-2.5 pt-3 text-sm font-medium uppercase text-blue-950 dark:text-amber-50 hover:bg-amber-50 dark:hover:bg-blue-950"
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