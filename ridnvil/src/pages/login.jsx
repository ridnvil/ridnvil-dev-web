import React, {useState} from 'react';
import ThemeSwitcher from "../components/ThemeSwitcher";
import logo from '../assets/profile.png'
import {useDispatch} from "react-redux";

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const dispatch = useDispatch();

    const handleLogin = () => {
        console.log(email, password)
    }

    return (
        <div className='h-screen bg-white dark:bg-blue-950 p-8 flex flex-col items-center'>
            <ThemeSwitcher />
            <div className='h-full'>
                <div className='gap-20 flex h-full flex-row items-center justify-center lg:justify-start'>
                    <div className="shrink-1 mb-12 grow-0 basis-auto md:mb-0 md:w-9/12 md:shrink-0 lg:w-6/12 xl:w-6/12">
                        <img
                            src={logo}
                            className="w-full rounded-full"
                            alt="Sample image"
                        />
                    </div>

                    <div className='flex flex-col'>
                        <div className='flex flex-row justify-between'>
                            <p className='mr-5 text-blue-950 dark:text-amber-50'>Email</p>
                            <input
                                value={email}
                                type="email"
                                onChange={(e) => setEmail(e.target.value)}
                                className="mb-6 border border-blue-950 dark:bg-amber-50 dark:border-amber-50 p-1 rounded"
                            ></input>
                        </div>

                        {/* <!--Password input--> */}
                        <div className='flex flex-row justify-between'>
                            <p className='mr-5 text-blue-950 dark:text-amber-50'>Password</p>
                            <input
                                type="password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                className="mb-6 border border-blue-950 dark:bg-amber-50 p-1 rounded"
                            ></input>
                        </div>

                        <button
                            type="button"
                            onClick={handleLogin}
                            className="rounded border bg-white dark:bg-blue-950 px-7 pb-2.5 pt-3 text-sm font-medium uppercase text-blue-950 dark:text-amber-50 hover:bg-amber-50 dark:hover:bg-blue-800"
                        >
                            Login
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Login;