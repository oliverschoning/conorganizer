import { Box, Typography } from '@mui/material';
import ParticipantCard from './components/ParticipantCard';
import { Participant } from '$lib/types';

const participants = () => {
    const tmpParticipants: Participant[] = [
        {
            id: '1',
            name: 'Ola Nordmann',
            over18: true,
            notes: [],
            ticketId: '12345678',
            ticketType: 'Festivalpass Voksen (26 år +) Early-bird',
            ticketEmail: '',
            ticketStatus: 'Betalt',
            createdAt: '',
            createdBy: '',
            updateAt: '',
            updatedBy: '',
        },
        {
            id: '2',
            name: 'Kari Nordmann',
            over18: true,
            notes: [],
            ticketId: '12345678',
            ticketType: 'Festivalpass Ungdom/student (13-25/30år) Early-bird',
            ticketEmail: '',
            ticketStatus: 'Betalt',
            createdAt: '',
            createdBy: '',
            updateAt: '',
            updatedBy: '',
        },
        {
            id: '3',
            name: 'Per Nordmann',
            over18: false,
            notes: [],
            ticketId: '12345678',
            ticketType: 'SØNDAG Dagspass Barn (3-12)',
            ticketEmail: '',
            ticketStatus: 'Betalt',
            createdAt: '',
            createdBy: '',
            updateAt: '',
            updatedBy: '',
        },
    ];
    return (
        <Box>
            <Typography variant="h1">Participants</Typography>
            <Typography variant="h2">Under utvikilig. Leker med ekte data, Ikke trykk på ting. </Typography>
            <Box
                sx={{
                    display: 'grid',
                    gridTemplateColumns: 'repeat(auto-fit,minmax(306px, 1fr))',
                    gap: '1rem',
                }}
            >
                {tmpParticipants.map((participant) => (
                    <ParticipantCard key={participant.id} participant={participant} />
                ))}
            </Box>
        </Box>
    );
};
export default participants;
