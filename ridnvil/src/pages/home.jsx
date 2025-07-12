import React, {useEffect, useState} from 'react';
import FadeInLeft from "../framemotion/FadeInLeft";
import Biodata from "../contents/Biodata";
import FadeInSection from "../framemotion/FadeInSection";
import ProfilePhoto from "../contents/ProfilePhoto";
import FadeInRight from "../framemotion/FadeInRight";
import Pendidikan from "../contents/Pendidikan";
import cv from "../assets/CV.pdf";
import SocialMediaIcons from "../components/SocialMediaIcons";
import Portofolio from "../components/Portofolio";
import ThemeSwitcher from "../components/ThemeSwitcher";
import IndodexBannerSmall from "../components/IndodexBannerSmall";
import {getProfileGlobal} from "../api/profileglobal";
import {useWelcomeProfileData} from "../hooks/useWelcomeProfileData";
import Loading from "../components/Loading";
import {errorHandler} from "../components/ErrorHandler";

const Home = () => {
    const { data, isLoading, error } = useWelcomeProfileData();

    useEffect(() => {
        getDetailsClient().then(async (res) => {
            await sendToServer(res).catch(err => console.log(err))
        }).catch(err => console.log(err))
    }, []);

    const getDetailsClient = async () => {
        return new Promise((resolve, reject) => {
            const url = 'https://ipapi.co/json/'
            fetch(url).then(async (res)=> {
                if (!res.ok) {
                    reject("Error")
                }
                resolve(await res.json())
            }).catch(err => {
                reject(err)
            })
        })
    }

    const sendToServer = (data) => {
        return new Promise((resolve, reject) => {
            fetch("/api/client", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data)
            }).then(async (res) => {
                if (!res.ok) {
                    reject("error")
                }
                resolve(await res.json());
            }).catch(err => {
                reject(err)
            })
        })
    }

    if (isLoading) {
        return <Loading />
    }

    if (error) errorHandler(error);

    return (
        <div className="relative bg-transparent bg-white dark:bg-gray-900 p-8 flex flex-col items-center">
            <div>
                <ThemeSwitcher />
            </div>
            <div className='flex flex-col items-center justify-center'>
                <FadeInLeft>
                    <h1 className="text-2xl font-bold text-blue-950 dark:text-amber-50 text-center">Hello, Welcome</h1>
                    <p className="text-xl italic mb-5 text-blue-950 dark:text-amber-50 text-center">Thank for visit my personal website</p>
                    <h1 className="text-4xl font-bold text-blue-950 mb-2 dark:text-amber-50 text-center">{data.position}</h1>
                </FadeInLeft>

                <div className="w-full max-w-5xl mt-2 flex flex-col items-center justify-center md:flex-row gap-10">

                    <div className="md:w-1/2">
                        <FadeInLeft>
                            <Biodata name={data.name} email={data.email} phone={data.phone} address={data.address} />
                            <div className='mt-4'>
                                {/*<IndodexBannerSmall />*/}
                            </div>
                        </FadeInLeft>
                    </div>

                    <div className="md:w-1/2">
                        <FadeInSection>
                            <ProfilePhoto/>
                        </FadeInSection>
                    </div>

                    <div className="md:w-1/2">
                        <FadeInRight>
                            <Pendidikan institude={data.education_ref[0].Institution} major={data.education_ref[0].Major} duration={data.education_ref[0].EducationLength}/>
                            <div className='mt-4'>
                                {/*<IndodexBannerSmall />*/}
                            </div>
                        </FadeInRight>
                    </div>
                </div>

                <FadeInSection>
                    <a className='light:bg-blue-950 dark:text-amber-50 p-2 rounded shadow' href={cv} target='_blank'>Preview CV</a>
                </FadeInSection>

                <div className='flex flex-col items-center justify-center'>
                    <FadeInSection>
                        <Portofolio />
                    </FadeInSection>
                </div>

                <div className='flex flex-col items-center justify-center'>
                    <FadeInSection>
                        <SocialMediaIcons />
                    </FadeInSection>
                </div>
            </div>
        </div>
    );
};

export default Home;