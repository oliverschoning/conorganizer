export type EventCardProps = {
    title: string;
    gameMaster: string;
    system: string;
    shortDescription: string;
    icons?: string[];
    backgroundImage?: string;
    myEventBar?: boolean;
    myEventBarSubmitted?: boolean;
    myEventDocId?: string;
};

export type ConEvent = {
    published: boolean;
    id?: string;
    title: string;
    gameMaster: string;
    system: string;
    shortDescription: string;
    description: string;
    icons?: string[];
    email: string;
    name: string;
    phone: string;
    gameType: string;
    isSmallCard: boolean;
    participants: number;
    puljeFridayEvening: boolean;
    puljeSaturdayMorning: boolean;
    puljeSaturdayEvening: boolean;
    puljeSundayMorning: boolean;
    unwantedFridayEvening: boolean;
    unwantedSaturdayMorning: boolean;
    unwantedSaturdayEvening: boolean;
    unwantedSundayMorning: boolean;
    moduleCompetition: boolean;
    childFriendly: boolean;
    possiblyEnglish: boolean;
    adultsOnly: boolean;
    volunteersPossible: boolean;
    lessThanThreeHours: boolean;
    moreThanSixHours: boolean;
    beginnerFriendly: boolean;
    additionalComments: string;
    createdAt: string;
    createdBy: string;
    updateAt: string;
    updatedBy: string;
    subTitle: string;
};

export type ConEventParent = {
    id?: string;
    childEvents: string[];
    title: string;
    gameMaster: string;
    system: string;
    shortDescription: string;
    description: string;
    email: string;
    name: string;
    phone: string;
    gameType: string;
    isSmallCard: boolean;
    participants: number;
    unwantedFridayEvening: boolean;
    unwantedSaturdayMorning: boolean;
    unwantedSaturdayEvening: boolean;
    unwantedSundayMorning: boolean;
    moduleCompetition: boolean;
    childFriendly: boolean;
    possiblyEnglish: boolean;
    adultsOnly: boolean;
    volunteersPossible: boolean;
    lessThanThreeHours: boolean;
    moreThanSixHours: boolean;
    beginnerFriendly: boolean;
    additionalComments: string;
    createdAt: string;
    createdBy: string;
    updateAt: string;
    updatedBy: string;
    subTitle: string;
};

export type ConEventPulje = {
    published: boolean;
    id?: string;
    parentEventId: string;
    title: string;
    gameMaster: string;
    system: string;
    shortDescription: string;
    description: string;
    gameType: string;
    isSmallCard: boolean;
    participants: number;
    childFriendly: boolean;
    possiblyEnglish: boolean;
    adultsOnly: boolean;
    lessThanThreeHours: boolean;
    moreThanSixHours: boolean;
    beginnerFriendly: boolean;
    additionalComments: string;
    createdAt: string;
    createdBy: string;
    updateAt: string;
    updatedBy: string;
};

export type MyNewEvent = {
    id: string;
    eventDocId?: string;
    email: string;
    name: string;
    phone: string;
    title: string;
    system: string;
    gameType: string;
    participants: number;
    fridayEvening: boolean;
    saturdayMorning: boolean;
    saturdayEvening: boolean;
    sundayMorning: boolean;
    moduleCompetition: boolean;
    childFriendly: boolean;
    possiblyEnglish: boolean;
    adultsOnly: boolean;
    volunteersPossible: boolean;
    lessThanThreeHours: boolean;
    moreThanSixHours: boolean;
    beginnerFriendly: boolean;
    description: string;
    additionalComments: string;
    createdAt: string;
    createdBy: string;
    updateAt: string;
    updatedBy: string;
    subTitle: string;
    isSubmitted: boolean;
    isRead?: boolean;
    isAccepted?: boolean;
};
