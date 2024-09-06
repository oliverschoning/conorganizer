'use client';
import React from 'react';
import { SvgIcon } from '@mui/material';

type Props = {
    color?: "inherit" | "primary" | "secondary" | "action" | "disabled" | "error" | "info" | "success" | "warning";
    size?: 'small' | 'medium' | 'large';
  }

const BoardGameIcon = ({ color = 'primary', size = 'medium' }:Props) => {
  return (
    <SvgIcon
      style={{ fontSize: size === 'small' ? "2rem" : size === 'large' ? "4rem" : 24 }}
      color={'primary'}
    >
    <svg
   id="Layer_8"
   viewBox="0 0 34.353153 39.599998"
   version="1.1"
   width="34.353153"
   height="39.599998">
 <path
       d="m 75.63,39.03 c -3.33,0.27 -6.97,-0.12 -10.59,0.29 -1.88,0.21 -3.76,-0.1 -5.61,-0.27 -1.95,-0.18 -3.8,0.46 -5.72,0.28 -0.28,-0.03 -0.56,-0.02 -0.83,-0.05 -1.69,-0.23 -2.39,-1.48 -1.63,-3.06 0.91,-1.89 2.25,-3.15 4.52,-3.44 2.47,-0.32 4.93,-0.18 7.39,-0.29 3.15,-0.15 6.28,0.28 9.43,0 2.91,-0.25 5.6,0.65 7.66,2.93 0.44,0.49 0.76,1.04 0.93,1.65 0.22,0.8 -0.14,1.65 -0.87,1.69 -1.47,0.07 -2.92,0.46 -4.69,0.27 z"
       id="path6" />
    <path
       d="m 53.03,12.53 c 2.75,-0.37 5.52,0.19 8.26,-0.24 0.34,-0.05 0.74,0.01 1.06,0.14 0.91,0.37 1.21,0.09 1.2,-0.87 -0.02,-2.02 0.02,-2.02 -2,-2.13 C 60.33,9.36 59.82,8.89 59.72,7.68 59.67,7.04 59.66,6.4 59.68,5.77 59.74,3.64 59.94,3.48 62.11,3.7 c 1.65,0.17 1.73,0.09 1.6,-1.53 -0.12,-1.49 0.2,-1.86 1.69,-1.92 0.44,-0.02 0.88,-0.04 1.32,-0.03 1.76,0.04 2.03,0.27 2.2,2.02 0.1,0.98 0.55,1.59 1.56,1.61 1.43,0.02 1.82,1.01 1.93,2.13 0.08,0.82 -0.09,1.66 -0.17,2.49 -0.03,0.32 -0.22,0.48 -0.54,0.55 -0.62,0.12 -1.24,0.27 -1.87,0.41 -0.43,0.09 -1.09,-0.16 -1.13,0.56 -0.04,0.69 -0.21,1.5 0.57,1.95 0.7,0.4 1.54,0.78 2.28,0.51 2.05,-0.75 4.11,-0.13 6.15,-0.23 0.99,-0.05 1.99,-0.1 2.98,0.15 1.92,0.48 2.76,1.92 2.08,3.8 -0.52,1.42 -1.16,2.8 -1.73,4.2 -0.95,2.33 -1.49,4.8 -2.57,7.08 -0.43,0.9 -0.93,1.76 -1.4,2.64 -0.21,0.39 -0.59,0.53 -0.97,0.64 -1.16,0.33 -2.32,0.58 -3.54,0.51 -3.42,-0.2 -6.85,-0.43 -10.27,-0.59 -1.07,-0.05 -2.11,0.33 -3.19,0.35 -1.14,0.02 -2.24,-0.57 -3.41,-0.11 -0.31,0.12 -0.62,-0.31 -0.75,-0.66 -1.3,-3.57 -3.5,-6.71 -4.88,-10.23 -0.52,-1.33 -0.92,-2.72 -1.28,-4.11 -0.51,-1.98 0.51,-3.27 2.58,-3.33 0.56,-0.02 1.12,0 1.68,0 z"
       id="path8" />
</svg>

    </SvgIcon>
  );
};

export default BoardGameIcon;