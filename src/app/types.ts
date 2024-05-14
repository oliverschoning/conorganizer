export type Event = {
    id: string;
    title: string;
    gameMaster: string;
    system: string;
    shortDescription: string;
    icons?: string[];
};

export type NewEvent = {
    id: string;
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
    unpublished: boolean;
    createdAt: string;
    createdBy: string;
    updateAt: string;
    updatedBy: string;
    subTitle: string;
    isSubmited: boolean;
};
