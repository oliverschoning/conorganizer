'use client';
import { signOutAndDeleteCookie } from '$lib/firebase/firebase';
import Button from '@mui/material/Button';
import { useRouter } from 'next/navigation';

const LogOutButton = () => {
    const router = useRouter();
    return (
        <Button
            onClick={async () => {
                await signOutAndDeleteCookie();
                router.replace('/');
            }}
            variant="outlined"
        >
            logg ut
        </Button>
    );
};

export default LogOutButton;
