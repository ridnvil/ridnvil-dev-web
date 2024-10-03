import React, {useEffect, useState} from 'react';
import FadeInSection from '../src/framemotion/FadeInSection';
import ProfilePhoto from "./contents/ProfilePhoto";
import Biodata from "./contents/Biodata";
import Pendidikan from "./contents/Pendidikan";
import AnimatedBackground from "./components/AnimatedBackground";
import FadeInLeft from "./framemotion/FadeInLeft";
import FadeInRight from "./framemotion/FadeInRight";
import SocialMediaIcons from "./components/SocialMediaIcons";
import cv from '../src/assets/CV.pdf';
import Portofolio from "./components/Portofolio";
import {FaThumbsUp} from "react-icons/fa";

function App() {
    const [clientInfo, setClientInfo] = useState(null)
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

    return (
        <div className="relative min-h-screen bg-transparent bg-opacity-85 p-8 flex flex-col items-center">
            <AnimatedBackground />
            <FadeInLeft>
                <h1 className="text-2xl font-bold text-amber-50 text-center">Hello, Welcome</h1>
                <p className="text-xl italic mb-5 text-amber-50 text-center">Thank for visit my personal website</p>
                <h1 className="text-4xl font-bold mb-2 text-amber-50 text-center">Backend Engineer</h1>
            </FadeInLeft>

            <div className="w-full max-w-5xl mt-2 flex flex-col items-center justify-center md:flex-row gap-10">

                <div className="md:w-1/2">
                    <FadeInLeft>
                        <Biodata/>
                    </FadeInLeft>
                </div>

                <div className="md:w-1/2">
                    <FadeInSection>
                        <ProfilePhoto/>
                    </FadeInSection>
                </div>

                <div className="md:w-1/2">
                    <FadeInRight>
                        <Pendidikan/>
                    </FadeInRight>
                </div>
            </div>

            <FadeInSection>
                <a className='bg-blue-950 text-amber-50 p-2 rounded shadow' href={cv} target='_blank'>Preview CV</a>
            </FadeInSection>

            <div className='space-x-0 md:space-x-4 w-full flex flex-col md:flex-row justify-center items-center'>
                <div className='flex items-center justify-center'>
                    <FadeInSection>
                        <SocialMediaIcons />
                    </FadeInSection>
                </div>
                <div>
                    <FadeInSection>
                        <Portofolio />
                    </FadeInSection>
                </div>
            </div>


        </div>
    );
}

export default App;
