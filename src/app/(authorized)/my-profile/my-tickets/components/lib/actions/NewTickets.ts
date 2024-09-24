import { Participant } from '$lib/types';
import { User } from 'firebase/auth';
import { EventTicket } from './actions';
import { generateParticipant } from './Helpers';

export const NewTickets = async (tickets: EventTicket[], participants: Participant[], user: User) => {
    console.log('Assigning new tickets to participants');
    // console.log('tickes', tickets, 'participants', participants, 'user', user);

    const newParticipants: Participant[] = [];
    tickets.forEach((ticket) => {
        const newParticipant = generateParticipant(ticket.id, tickets, user.email as string);
        newParticipant.users = [user.uid];
        newParticipants.push(newParticipant);
    });
    return newParticipants;
};
