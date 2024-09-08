'use server';

import { getEventById } from '$app/(public)/components/lib/serverAction';
import { PoolName } from '$lib/enums';
import { getAuthorizedAuth } from '$lib/firebase/firebaseAdmin';
import { ConEvent, PoolEvent } from '$lib/types';
import { addDoc, collection, doc, updateDoc } from 'firebase/firestore';

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
        published: false,
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
        createdAt: Date.now().toString(),
        createdBy: user.uid,
        updateAt: Date.now().toString(),
        updatedBy: user.uid,
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
    if (poolName === PoolName.saturdayEvening) {
        conEvent.puljeSaturdayEvening = true;
    }

    conEvent.updateAt = Date.now().toString();
    conEvent.updatedBy = user.uid;

    try {
        await updateDoc(doc(db, 'events', eventId), conEvent);
        console.log('Document updated');
    } catch (e) {
        console.error('Error updating document: ', e);
    }
}
