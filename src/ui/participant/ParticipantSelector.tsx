'use client';
import { Participant } from '$lib/types';
import { useState, useEffect } from 'react';
import { Menu, MenuItem, Button } from '@mui/material';
import ParticipantAvatar from './ParticipantAvatar';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';

const ParticipantSelector = () => {
    const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
    const open = Boolean(anchorEl);

    const handleClick = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };

    // State to store participants
    const [participants, setParticipants] = useState<Participant[]>([]);

    // Fetch participants from localStorage when component mounts
    useEffect(() => {
        const storedParticipants = localStorage.getItem('newParticipants');
        if (storedParticipants) {
            setParticipants(JSON.parse(storedParticipants));
        }
    }, []);

    if (!participants || participants.length === 0) {
        return <Button href="/my-profile/my-tickets">Hent billett</Button>;
    }

    const selectedParticipant = participants[0];

    return (
        <>
            <Button
                aria-controls={open ? 'participant-menu' : undefined}
                aria-haspopup="true"
                aria-expanded={open ? 'true' : undefined}
                onClick={handleClick}
                variant="text"
                sx={{ textDecoration: 'none', textTransform: 'none' }}
            >
                <ParticipantAvatar
                    firstName={selectedParticipant.firstName}
                    lastName={selectedParticipant.lastName}
                    small
                />
                '
                <ExpandMoreIcon />
            </Button>
            <Menu
                id="participant-menu"
                anchorEl={anchorEl}
                open={open}
                onClose={handleClose}
                MenuListProps={{
                    'aria-labelledby': 'participant-button',
                }}
            >
                {participants.map((participant) => (
                    <MenuItem key={participant.id} onClick={handleClose}>
                        <ParticipantAvatar firstName={participant.firstName} lastName={participant.lastName} />
                    </MenuItem>
                ))}
            </Menu>
        </>
    );
};

export default ParticipantSelector;
