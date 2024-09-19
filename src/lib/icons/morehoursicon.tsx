'use client';
import React from 'react';
import { SvgIcon } from '@mui/material';

type Props = {
    color?: "inherit" | "primary" | "secondary" | "action" | "disabled" | "error" | "info" | "success" | "warning";
    size?: 'small' | 'medium' | 'large';
  }

const MoreHoursIcon = ({ color = 'primary', size = 'medium' }:Props) => {
  return (
    <SvgIcon
      style={{ fontSize: size === 'small' ? "2rem" : size === 'large' ? "4rem" : 24 }}
      color={'primary'}
    >
    <svg
   id="Layer_1"
   viewBox="0 0 58.865624 39.6"
   version="1.1"
   width="58.865623"
   height="39.599998">
  <g
     id="Layer_5"
     transform="translate(-41.95048,-94.989998)">
    <path
       d="m 61.17,134.22 c -2.96,-0.54 -5.95,-0.55 -8.93,-0.47 -2.57,0.07 -5.08,-0.32 -7.58,-0.76 -1.3,-0.23 -2.23,-1.02 -2.47,-2.43 -0.36,-2.1 -0.21,-4.21 -0.07,-6.32 0.31,-4.71 -0.33,-9.41 -0.13,-14.11 0.13,-3.15 -0.04,-6.29 -0.03,-9.44 0,-1.89 1.08,-3.28 2.9,-3.92 2.75,-0.96 5.57,-1.03 8.39,-0.6 2.43,0.38 4.83,0.44 7.26,0.11 3.81,-0.5 7.63,-0.64 11.47,-0.28 2.12,0.2 4.19,-0.56 6.31,-0.5 5.32,0.14 10.68,-0.79 15.95,0.69 0.87,0.25 1.81,0.2 2.7,0.49 2.08,0.68 2.85,1.44 2.9,3.68 0.06,2.43 0.34,4.85 0.41,7.27 0.09,3.11 0.42,6.21 0.46,9.3 0.04,3.29 0.28,6.61 -0.13,9.9 -0.05,0.39 -0.03,0.8 -0.03,1.2 -0.05,2.91 -0.71,3.83 -3.59,4.25 -3.98,0.58 -7.94,1.22 -12,1.33 -3.4,0.09 -6.84,-0.26 -10.23,0.35 -0.99,0.18 -2.02,0.03 -2.97,-0.19 -1.84,-0.43 -3.63,-0.26 -5.46,-0.03 -1.54,0.2 -3.08,0.31 -4.63,0.46 h -0.48 z m -3.19,-7.9 c 2.86,0.38 5.23,-0.75 7.32,-2.43 1.75,-1.41 3.13,-3.11 3.33,-5.58 0.29,-3.48 0.42,-6.87 -1.54,-10.01 -1.64,-2.62 -3.7,-4.69 -6.83,-5.35 -2.6,-0.54 -5.23,-0.57 -7.81,0.19 -1.94,0.57 -3.69,1.43 -4.89,3.2 -2.25,3.33 -3.06,6.97 -2.43,10.91 0.56,3.57 2.63,6.12 5.83,7.65 2.2,1.04 4.54,1.78 7.02,1.42 z m 11.82,-8.26 c -0.11,0.56 0.23,1.48 0.36,2.42 0.37,2.76 2.3,4.7 5.02,4.96 3.11,0.3 6.67,-2.31 7.06,-5.7 0.49,-4.31 -3.79,-7.3 -7.05,-6.09 -0.8,0.3 -0.7,-0.31 -0.6,-0.69 0.43,-1.65 0.87,-3.31 1.73,-4.81 0.81,-1.42 1.83,-2.7 2.82,-4 0.63,-0.82 0.4,-1.9 -0.41,-2.08 -0.82,-0.19 -1.63,-0.24 -2.46,0.23 -1.78,1.01 -2.94,2.53 -3.9,4.27 -1.91,3.46 -2.2,7.28 -2.58,11.5 z m 16.87,-0.67 c 0.09,0.85 0.12,1.61 0.26,2.36 0.24,1.29 1.53,1.02 2.32,1.5 0.61,0.37 1.02,0.13 1.34,-0.27 0.52,-0.64 1.03,-1.3 0.97,-2.23 -0.07,-1.03 -0.02,-2.07 -0.27,-3.09 -0.18,-0.72 0.1,-1.27 0.95,-1.3 0.56,-0.02 1.14,-0.09 1.66,0.03 0.93,0.22 1.54,-0.05 2.36,-0.58 1.27,-0.82 0.87,-1.69 0.58,-2.62 -0.24,-0.77 -0.78,-1.28 -1.67,-1.42 -0.83,-0.13 -1.67,-0.27 -2.5,-0.21 -0.97,0.07 -1.3,-0.28 -1.4,-1.27 -0.13,-1.21 0.12,-2.49 -0.55,-3.63 -0.86,-1.47 -2.77,-1.53 -3.66,-0.1 -0.47,0.76 -0.45,1.59 -0.47,2.43 -0.05,2.72 -0.07,2.72 -2.71,2.95 -0.04,0 -0.08,0 -0.12,0 -0.82,0.06 -1.79,0.18 -1.97,1.06 -0.19,0.95 -0.4,2.06 0.47,2.91 0.7,0.69 1.6,0.61 2.47,0.63 1.71,0.05 2.18,0.55 2.06,2.24 -0.02,0.23 -0.1,0.47 -0.13,0.61 z"
       id="path414" />
    <path
       d="m 47.75,114.31 c 0.15,-2.53 0.98,-5.17 2.95,-7.39 0.79,-0.89 1.74,-1.49 2.91,-1.75 1.57,-0.35 3.19,-0.37 4.75,-0.32 2.81,0.1 4.87,1.7 6.13,4.14 1.61,3.1 2.35,6.42 0.93,9.78 -1.35,3.2 -6.01,7.11 -10.88,5.09 -4.08,-1.69 -6.94,-5.26 -6.79,-9.56 z m 7.35,-2.07 c -0.36,1.98 1.09,3.27 2.27,4.7 0.58,0.71 1.17,1.42 2.16,1.59 0.52,0.09 0.86,0.11 1.24,-0.5 0.68,-1.1 0.6,-1.9 -0.78,-3.17 -0.96,-0.88 -1.32,-1.76 -1.29,-3.01 0.03,-1.59 -0.13,-3.19 -0.94,-4.62 -0.23,-0.41 -0.53,-0.98 -1.04,-0.83 -0.51,0.15 -0.92,0.64 -1.05,1.22 -0.33,1.47 -0.63,2.95 -0.56,4.62 z"
       id="path416" />
    <path
       d="m 73.32,118.43 c -0.08,-0.84 0.18,-1.44 1.03,-1.8 0.59,-0.25 1.17,-0.47 1.81,-0.57 1.99,-0.32 2.62,0.2 2.4,2.2 -0.16,1.54 -0.85,2.82 -2.33,3.5 -1.03,0.47 -1.74,0.19 -2.28,-0.8 -0.32,-0.6 -0.6,-1.72 -0.63,-2.53 z"
       id="path418" />
    </g>
</svg>

    </SvgIcon>
  );
};

export default MoreHoursIcon;