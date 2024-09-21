'use server';

import { getAuthorizedAuth } from '$lib/firebase/firebaseAdmin';
import { Participant } from '$lib/types';
import { addDoc, collection } from 'firebase/firestore';

export type CrmRecord = {
    id: number;
    firstName: string;
    lastName: string;
    hash: string;
    email: {
        email: string;
        id: number;
    };
};

export type CrmData = {
    records: number;
    data: CrmRecord[];
};

export type CrmJson = {
    data: {
        eventTickets: EventTicket[];
    };
    errors: Error;
};

export type EventTicket = {
    id: number;
    category: string;
    category_id: number;
    order_id: number;
    crm: {
        first_name: string;
        last_name: string;
        id: number;
        email: string;
    };
};
export const GetTicketsByEmail = async (email: string | null | undefined) => {
    if ((typeof email === 'string') == false) {
        return;
    }

    //const allTickets = await GetTicketsFromCheckIn();

    const allTickets = [
        {
            id: 4098407,
            category: 'Bilett 1',
            category_id: 157049,
            crm: {
                first_name: 'Gerhard',
                last_name: 'Just-Olsen',
                id: 1555823,
                email: 'cinmay05@gmail.com',
            },
            order_id: 12743193,
        },
        {
            id: 4098408,
            category: 'Bilett 1 the sequel',
            category_id: 157057,
            crm: {
                first_name: 'Nils-Erik',
                last_name: 'Just-Olsen',
                id: 2585634,
                email: 'nils.erik.tarpan@gmail.com',
            },
            order_id: 12743193,
        },
        {
            id: 4098407,
            category: 'Bilett 2',
            category_id: 157049,
            crm: {
                first_name: 'Gerhard',
                last_name: 'Just-Olsen',
                id: 1555823,
                email: 'cinmay05@gmail.com',
            },
            order_id: 127431934,
        },
        {
            id: 4098408,
            category: 'Bilett nr 2 the sequel',
            category_id: 157057,
            crm: {
                first_name: 'Nils-Erik',
                last_name: 'Just-Olsen',
                id: 2585634,
                email: 'en annen email',
            },
            order_id: 127431934,
        },
        {
            id: 4098408,
            category: 'Bilett nr 3',
            category_id: 157057,
            crm: {
                first_name: 'Nils-Erik',
                last_name: 'Just-Olsen',
                id: 2585634,
                email: 'en annen email',
            },
            order_id: 127431937,
        },
    ];

    const emailTickets = allTickets?.filter((ticket) => ticket.crm.email === email);
    const ticketsWithOrderNumberFromEmail = allTickets?.filter((ticket) =>
        emailTickets?.some((emailTicket) => emailTicket.order_id === ticket.order_id)
    );

    return ticketsWithOrderNumberFromEmail;
};
export const ConvertTicketIdToParticipant = async (ticketId: number) => {
    console.log('ConvertTicketToParticipant ', ticketId);
    //Query ticket from CheckIn by id
    const tickets = await GetTicketsFromCheckIn();

    const ticket = tickets?.find((ticket) => ticket.id === ticketId);
    console.log(ticket, 'ticket');

    if (!ticket) {
        return 'No ticket found';
    }

    const result = await ConvertTicketToParticipant(ticket);
    return result;

    // const query = `{
    // eventTickets(customer_id: 13446, id: ${ticketId}, onlyCompleted: true) {
    //   id
    //   category
    //   category_id
    //   crm {
    //     first_name
    //     last_name
    //     id
    //     email
    //   }
    //   order_id
    // }
    // }`;
    //
    // const res = await fetch(
    //     `https://app.checkin.no/graphql?client_id=${process.env.CHECKIN_KEY}&client_secret=${process.env.CHECKIN_SECRET}`,
    //     {
    //         method: 'POST',
    //         body: JSON.stringify({ query }),
    //         headers: {
    //             'Content-Type': 'application/json',
    //         },
    //     }
    // );
    // if (res.status !== 200) {
    //     throw new Error('Failed to fetch tickets');
    // }
    // const queryResult: CrmJson | undefined = await res.json();
    // console.log(queryResult, 'queryResult');
    //
    // // const result = await ConvertTicketToParticipant(queryResult?.data.eventTickets[0] as EventTicket);
    // return result;
};

const ConvertTicketToParticipant = async (ticket: EventTicket) => {
    //Ckeck if ticket is already converted

    //Create participant from ticket
    const { db, user } = await getAuthorizedAuth();
    if (db === null || user === null) {
        return;
    }
    let participant: Partial<Participant> = {
        name: `${ticket.crm.first_name} ${ticket.crm.last_name}`,
        ticketEmail: ticket.crm.email,
        ticketId: ticket.id,
        orderId: ticket.order_id,
        ticketCategory: ticket.category,
        ticketCategoryId: ticket.category_id,
        createdAt: new Date().toISOString(),
        createdBy: user.email as string,
        updateAt: new Date().toISOString(),
        updatedBy: user.email as string,
    };
    console.log(participant, 'participant');

    //Save participant to firestore
    // let poolEventId = '';
    // try {
    //     const docRef = await addDoc(collection(db, 'pool-events'), poolEvent);
    //     console.log('Document written with ID: ', docRef.id);
    //     poolEventId = docRef.id;
    // } catch (e) {
    //     console.error('Error adding document: ', e);
    //     return;
    // }

    let participantId = '';
    try {
        const docRef = await addDoc(collection(db, 'participants'), participant);
        console.log('Document written with ID: ', docRef.id);
        participantId = docRef.id;
    } catch (e) {
        console.error('Error adding prticipant document: ', e);
        return;
    }
    return 'Convertion done';
};

export const GetTicketsFromCheckIn = async () => {
    const query = `{
    eventTickets(customer_id: 13446, id: 73685, onlyCompleted: true) {
      id
      category
      category_id
      crm {
        first_name
        last_name
        id
        email
      }
      order_id
    }
  }`;

    const res = await fetch(
        `https://app.checkin.no/graphql?client_id=${process.env.CHECKIN_KEY}&client_secret=${process.env.CHECKIN_SECRET}`,
        {
            method: 'POST',
            body: JSON.stringify({ query }),
            headers: {
                'Content-Type': 'application/json',
            },
        }
    );

    const queryResult: CrmJson | undefined = await res.json();
    //console.log(queryResult, 'queryResult');

    return queryResult?.data.eventTickets;
};
