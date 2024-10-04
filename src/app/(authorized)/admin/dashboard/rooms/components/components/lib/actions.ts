'use server';

import { getEventById, getPoolEventById } from '$app/(public)/components/lib/serverAction';
import { PoolName, type RoomName } from '$lib/enums';
import { getAuthorizedAuth } from '$lib/firebase/firebaseAdmin';
import { ConEvent, PoolEvent, EventRoom } from '$lib/types';
import { addDoc, collection, deleteDoc, doc, updateDoc, type DocumentReference } from 'firebase/firestore';
import { revalidatePath } from 'next/cache';

//// ******************* Do not delete ****************
//export async function removeFromPool(eventId: string, poolName: PoolName) {
//     const { db, user } = await getAuthorizedAuth();
//     if (db === null || user === null) {
//         return;
//     }
//     const conEvent: ConEvent = await getEventById(eventId);
//     if (conEvent.poolIds?.some((pool) => pool.poolName === poolName) === false) {
//         console.error('Event not found');
//         return;
//     }
//
//     const poolEventId = conEvent.poolIds?.find((pool) => pool.poolName === poolName)?.id;
//     if (!poolEventId) {
//         console.error('PoolEvent not found in parent event');
//         return;
//     }
//
//     const poolEvent: PoolEvent = await getPoolEventById(poolEventId);
//
//     if (!poolEvent) {
//         console.error('PoolEvent not found');
//         return;
//     }
//
//     try {
//         await deleteDoc(doc(db, 'pool-events', poolEventId));
//         console.log('Document deleted', poolEventId);
//     } catch (e) {
//         console.error('Error updating document: ', e);
//     }
//
//     const removePoolId = conEvent.poolIds.filter((pool) => pool.id !== poolEventId);
//     conEvent.poolIds = removePoolId;
//
//     if (poolName === PoolName.fridayEvening) {
//         conEvent.puljeFridayEvening = false;
//     }
//     if (poolName === PoolName.saturdayMorning) {
//         conEvent.puljeSaturdayMorning = false;
//     }
//     if (poolName === PoolName.saturdayEvening) {
//         conEvent.puljeSaturdayEvening = false;
//     }
//     if (poolName === PoolName.sundayMorning) {
//         conEvent.puljeSundayMorning = false;
//     }
//
//     conEvent.updateAt = Date.now().toString();
//
//     try {
//         await updateDoc(doc(db, 'events', eventId), conEvent);
//         console.log('Document updated', eventId);
//     } catch (e) {
//         console.error('Error updating document: ', e);
//     }
//
//     revalidatePath('/admin/dashboard/rooms', 'page');
// }
//// ******************* Do not delete ****************

export async function removeFromRoom(eventId: string, roomName: RoomName, poolName: PoolName) {
    const { db, user } = await getAuthorizedAuth();
    if (db === null || user === null) {
        return;
    }
    const conEvent: ConEvent = await getEventById(eventId);
    const roomPoolId = conEvent.poolIds?.find((pool) => pool.poolName === poolName);

    if (!roomPoolId) {
        console.error('Pool not found');
        return;
    }

    // const poolEvent: PoolEvent = await getPoolEventById(roomPoolId.id);
    //
    // if (!poolEvent) {
    //     console.error('PoolEvent not found');
    //     return;
    // }

    const roomId: string = conEvent.roomIds.find((room) => room.roomName === roomName)?.id ?? '';
    const poolId = conEvent.roomIds.find((room) => room.roomName === roomName)?.poolId ?? '';

    try {
        await deleteDoc(doc(db, 'pool-events', poolId, 'room', roomId));
        console.log('Document deleted', roomId);
    } catch (e) {
        console.error('Error updating document: ', e);
    }

    const removeRoomId = conEvent.roomIds.filter((room) => room.id !== roomId);
    console.log('removeRoomId: ', removeRoomId);
    conEvent.roomIds = removeRoomId;
    conEvent.updateAt = Date.now().toString();

    try {
        await updateDoc(doc(db, 'events', eventId), conEvent);
        console.log('Document updated', eventId);
    } catch (e) {
        console.error('Error updating document: ', e);
    }

    revalidatePath('/admin/dashboard/rooms', 'page');
}

export async function convertToPoolEvent(eventId: string, poolName: PoolName) {
    const { db, user } = await getAuthorizedAuth();
    if (db === null || user === null) {
        return;
    }
    const conEvent: ConEvent = await getEventById(eventId);
    if (conEvent.poolIds?.some((pool) => pool.poolName === poolName)) {
        console.log('Event already in pool');
        return;
    }

    const poolEvent: PoolEvent = {
        poolName: poolName,
        published: conEvent.published,
        parentEventId: eventId,
        title: conEvent.title,
        gameMaster: conEvent.gameMaster,
        system: conEvent.system,
        shortDescription: conEvent.shortDescription,
        description: conEvent.description,
        gameType: conEvent.gameType,
        isSmallCard: false,
        participants: conEvent.participants,
        childFriendly: conEvent.childFriendly,
        possiblyEnglish: conEvent.possiblyEnglish,
        adultsOnly: conEvent.adultsOnly,
        lessThanThreeHours: conEvent.lessThanThreeHours,
        moreThanSixHours: conEvent.moreThanSixHours,
        beginnerFriendly: conEvent.beginnerFriendly,
        additionalComments: conEvent.additionalComments,
        volunteersPossible: conEvent.volunteersPossible,
        bigImageURL: conEvent.bigImageURL ?? '',
        smallImageURL: conEvent.smallImageURL ?? '',
        createdAt: Date.now().toString(),
        createdBy: user.uid,
        updateAt: Date.now().toString(),
        updatedBy: user.uid,
        isInterestChoiceClosed: false
    };
    let poolEventId = '';
    try {
        const docRef = await addDoc(collection(db, 'pool-events'), poolEvent);
        console.log('Document written with ID: ', docRef.id);
        poolEventId = docRef.id;
    } catch (e) {
        console.error('Error adding document: ', e);
        return;
    }

    const poolChildRef = {
        id: poolEventId,
        poolName: poolName,
        createdAt: Date.now().toString(),
        createdBy: user.uid,
        updateAt: Date.now().toString(),
        updatedBy: user.uid,
    };
    console.log('poolChildRef: ', poolChildRef);
    if (!conEvent.poolIds) {
        conEvent.poolIds = [];
    }
    conEvent.poolIds.push(poolChildRef);

    if (poolName === PoolName.fridayEvening) {
        conEvent.puljeFridayEvening = true;
    }
    if (poolName === PoolName.saturdayMorning) {
        conEvent.puljeSaturdayMorning = true;
    }
    if (poolName === PoolName.saturdayEvening) {
        conEvent.puljeSaturdayEvening = true;
    }
    if (poolName === PoolName.sundayMorning) {
        conEvent.puljeSundayMorning = true;
    }

    conEvent.roomIds = conEvent.roomIds ?? [];
    conEvent.updateAt = Date.now().toString();
    conEvent.updatedBy = user.uid;

    try {
        await updateDoc(doc(db, 'events', eventId), conEvent);
        console.log('Document updated');
    } catch (e) {
        console.error('Error updating document: ', e);
    }
    revalidatePath('/admin/dashboard/rooms', 'page');
}

export async function addToRoom(eventId: string, roomName: RoomName, poolName: PoolName) {
    const { db, user } = await getAuthorizedAuth();
    if (db === null || user === null) {
        return;
    }
    await convertToPoolEvent(eventId, poolName);

    const conEvent: ConEvent = await getEventById(eventId);
    const roomPoolId = conEvent.poolIds?.find((pool) => pool.poolName === poolName);

    const poolEvent: PoolEvent = await getPoolEventById(roomPoolId?.id ?? '');
    if (!poolEvent) {
        console.error('PoolEvent not found');
        return;
    }
    const room: Omit<EventRoom, 'id'> = {
        name: roomName,
        eventId: eventId,
        poolId: poolEvent?.id ?? '',
        players: [],
        createdAt: Date.now().toString(),
        createdBy: user.uid,
        updateAt: Date.now().toString(),
        updatedBy: user.uid,
    };

    if (poolEvent.id) {
        let roomDocument: Awaited<ReturnType<typeof addDoc>> | undefined;
        try {
            roomDocument = await addDoc(collection(db, 'pool-events', poolEvent.id, 'room'), room);
            console.log('Room collection created');
        } catch (e) {
            console.error('Error room collection creation failed: ', e);
        }

        try {
            const roomChildRef: Partial<ConEvent> = {
                roomIds: [
                    ...conEvent.roomIds,
                    {
                        id: roomDocument?.id ?? '',
                        roomName: roomName,
                        poolId: poolEvent.id,
                        poolName: poolName,
                        createdAt: Date.now().toString(),
                        createdBy: user.uid,
                        updateAt: Date.now().toString(),
                        updatedBy: user.uid,
                    },
                ],
            };
            await updateDoc(doc(db, 'events', eventId) as DocumentReference<ConEvent, ConEvent>, roomChildRef);
            console.log('Event updated');
        } catch (e) {
            console.error('Error updating document: ', e);
        }
    }
    revalidatePath('/admin/dashboard/rooms', 'page');
}
