// src/components/Background3D.js
import React from 'react';
import { Canvas } from '@react-three/fiber';
import { Stars, MeshWobbleMaterial, Sphere } from '@react-three/drei';

const Background3D = () => {
    return (
        <div className="absolute inset-0 -z-10">
            <Canvas camera={{ position: [0, 0, 5], fov: 60 }}>
                {/* Bintang sebagai latar belakang */}
                <Stars
                    radius={100}
                    depth={50}
                    count={5000}
                    factor={4}
                    saturation={0}
                    fade
                />
                {/* Pencahayaan */}
                <ambientLight intensity={0.5} />
                <pointLight position={[10, 10, 10]} />
                {/* Bola berputar */}
                <mesh rotation={[10, 15, 0]}>
                    <Sphere args={[1, 32, 32]}>
                        <MeshWobbleMaterial
                            attach="material"
                            color="#1e90ff"
                            speed={1}
                            factor={0.6}
                        />
                    </Sphere>
                </mesh>
            </Canvas>
        </div>
    );
};

export default Background3D;
