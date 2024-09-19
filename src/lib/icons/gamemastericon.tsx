'use client';
import React from 'react';
import { SvgIcon } from '@mui/material';

type Props = {
    color?: "inherit" | "primary" | "secondary" | "action" | "disabled" | "error" | "info" | "success" | "warning";
    size?: 'small' | 'medium' | 'large';
  }

const GamemasterIcon = ({ color = 'primary', size = 'medium' }:Props) => {
  return (
    <SvgIcon
      style={{ fontSize: size === 'small' ? "2rem" : size === 'large' ? "4rem" : 24 }}
      color={'primary'}
    >
    <svg
   id="Layer_1"
   viewBox="0 0 34.359078 39.6"
   version="1.1"
   width="34.359077"
   height="39.599998">
  <path
     d="m 29.889882,39.32 c -1.38,-0.13 -3.09,-0.39 -4.8,-0.41 -2.31,-0.02 -4.63,0.08 -6.93,0.25 -1.56,0.11 -3.09,0 -4.64,-0.02 -3.11,-0.06 -6.2300004,-0.09 -9.3400004,0.1 -3.18999998,0.2 -4.97999998,-2.3 -3.82999998,-5.35 1.05999998,-2.82 2.49999998,-5.43 4.82999998,-7.43 0.53,-0.45 0.52,-0.84 0.32,-1.46 -0.6,-1.82 -1.45,-3.54 -2.23,-5.29 -0.64,-1.44 -0.31,-2.13 1.2,-2.61 0.73,-0.23 1.49,-0.3 2.25,-0.22 0.73,0.07 0.9,-0.19 0.79,-0.89 -0.14,-0.82 -0.14,-1.67 0.11,-2.49 0.19,-0.6 -0.03,-0.87 -0.63,-1.01 -1.02,-0.23 -1.97,-0.65 -2.86,-1.17 -0.57,-0.34 -1.06,-0.76 -0.97,-1.51 0.08,-0.69 0.6,-0.95 1.18,-1.21 1.1,-0.48 2.16,-1.02 3.34,-1.3 0.56,-0.14 0.84,-0.53 1.02,-1.07 0.59,-1.73 1.6700004,-3.17 2.8000004,-4.57 1.23,-1.51999996 2.6,-1.78999996 4.48,-0.81999996 1.1,0.56999996 1.98,0.83999996 3.16,0 1.33,-0.94 3.39,-0.34 4.39,0.98999996 0.65,0.87 1.2,1.86 1.44,2.87 0.48,2.05 1.92,2.65 3.69,2.92 0.73,0.11 1.38,0.41 2,0.77 1.15,0.65 1.26,1.7 0.32,2.65 -0.7,0.71 -1.6,1.05 -2.53,1.27 -0.49,0.12 -0.82,0.31 -0.81,0.86 0.02,0.98 -0.27,1.89 -0.61,2.79 -0.32,0.85 0.03,1.04 0.82,1.04 0.75,0 1.52,0.08 2.26,0.23 1.26,0.26 1.79,1.24 1.3,2.42 -0.62,1.51 -1.26,3.01 -1.95,4.49 -0.44,0.94 -0.39,1.62 0.42,2.43 2.23,2.2 3.79,4.81 4.35,7.95 0.58,3.21 -0.78,4.84 -4.34,4.79 z m -8.91,-22.11 c 1.6,0.04 2.96,-1.22 3.42,-3.11 0.19,-0.77 0.08,-1.04 -0.84,-0.91 -4.07,0.59 -8.14,0.87 -12.23,0.11 -0.35,-0.07 -0.82,-0.13 -0.89,0.33 -0.12,0.79 -0.12,1.62 0.48,2.25 1.43,1.47 4.13,1.32 5.37,-0.27 0.74,-0.95 1.11,-0.97 1.78,0.02 0.71,1.05 1.68,1.54 2.91,1.58 z m 0.29,17.45 c 1.39,-3.76 1.94,-7.59 2.87,-11.33 0.11,-0.45 -0.06,-0.8 -0.55,-0.53 -1.15,0.62 -2.5,0.92 -3.44,1.89 -0.59,0.61 -1.34,1.29 -1.18,2.25 0.44,2.62 1.16,5.17 2.3,7.73 z m -7.65,0.1 c 0.59,-1.76 1.1,-3.34 1.66,-4.91 1.01,-2.83 0.19,-5.08 -2.41,-6.45 -0.48,-0.25 -1.09,-0.63 -1.57,-0.33 -0.47,0.3 -0.23,0.98 -0.14,1.46 0.47,2.26 1.01,4.51 1.5,6.77 0.24,1.11 0.46,2.23 0.95,3.46 z"
     id="path47" />
</svg>

    </SvgIcon>
  );
};

export default GamemasterIcon;