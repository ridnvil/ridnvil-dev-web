// src/components/GradientAnimation.js
import React, { useRef } from 'react';
import { Canvas, useFrame } from '@react-three/fiber';
import { shaderMaterial } from '@react-three/drei';
import * as THREE from 'three';
import { extend } from '@react-three/fiber';

const vertexShader = `
  varying vec2 vUv;
  void main() {
    vUv = uv;
    gl_Position = projectionMatrix * modelViewMatrix * vec4(position, 1.0);
  }
`;

const fragmentShader = `
  precision mediump float;
  uniform float time;
  uniform vec2 resolution;
  varying vec2 vUv;

  void main() {
    vec2 st = gl_FragCoord.xy / resolution.xy;
    float color = 0.0;

    color = 0.5 + 0.5 * sin(time + st.x * 10.0);
    gl_FragColor = vec4(vec3(color), 1.0);
  }
`;

// Membuat custom shader material
const GradientShaderMaterial = shaderMaterial(
    {
        time: 0,
        resolution: new THREE.Vector2(),
    },
    // Vertex Shader
    vertexShader,
    fragmentShader
);

// Menambahkan material ke Three.js
extend({ GradientShaderMaterial });

const GradientAnimation = () => {
    const materialRef = useRef();

    useFrame(({ clock, size }) => {
        if (materialRef.current) {
            materialRef.current.time = clock.getElapsedTime();
            materialRef.current.resolution = new THREE.Vector2(size.width, size.height);
        }
    });

    return (
        <Canvas
            className="w-full h-screen"
            style={{ position: 'absolute', top: 0, left: 0, zIndex: -1 }}
        >
            <mesh>
                <planeBufferGeometry args={[2, 2, 1, 1]} />
                <gradientShaderMaterial ref={materialRef} />
            </mesh>
        </Canvas>
    );
};

export default GradientAnimation;
