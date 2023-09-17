'use client';

import { createTheme, ThemeOptions } from '@mui/material';

const muiDarkTheme: ThemeOptions = {
    palette: {
        mode: "dark",
        primary: {
            light: '#e0cfc9',
            main: '#a1887f',
            dark: '#000',
            contrastText: '#fff',
        },
        secondary: {
            light: '#ffd54f',
            main: '#ff8f00',
            dark: '#000',
            contrastText: '#000',
        },
    },
    typography: {
        h6: {
            fontWeight: 'bold',
        },
        h3: {
            fontWeight: 'bold',
            lineHeight: "1", // FUNKER FAEN IKKJE!
            fontFamily: "Garamond, serif", // IKKJE DETTE HELLER!
        },
    },
    components: {
        MuiButton: {
            styleOverrides: {
                root: ({ ownerState }) => ({
                    ...(ownerState.variant === 'contained' &&
                        ownerState.color === 'primary' && {
                           backgroundColor: 'primary.dark',
                            color: '#fff',
                        }),
                }),
            },
        },
    },
};

export const muiDark = createTheme(muiDarkTheme);
