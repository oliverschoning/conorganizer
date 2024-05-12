'use client';
import Typography from '@mui/material/Typography';
import Paper from '@mui/material/Paper';
import TextField from '@mui/material/TextField';
import FormControl from '@mui/material/FormControl';
import FormControlLabel from '@mui/material/FormControlLabel';
import FormLabel from '@mui/material/FormLabel';
import Radio from '@mui/material/Radio';
import RadioGroup from '@mui/material/RadioGroup';
import TextareaAutosize from '@mui/material/TextareaAutosize';
import FormGroup from '@mui/material/FormGroup';
import Checkbox from '@mui/material/Checkbox';
import Button from '@mui/material/Button';
import Confetti from 'react-confetti';
import { useEffect, useState } from 'react';
import { doc, onSnapshot, setDoc } from 'firebase/firestore';
import { db, firebaseAuth } from '$lib/firebase/firebase';
import type { NewEvent } from '$app/types';
import Grid2 from '@mui/material/Unstable_Grid2/Grid2';
import { onAuthStateChanged, type Unsubscribe, type User } from 'firebase/auth';
type Props = {
    params: {
        id: string;
    };
    searchParams: {
        userId: string;
    };
};
const Create = ({ params: { id } }: Props) => {
    const [isExploding, setIsExploding] = useState(false);
    const [newEvent, setNewEvent] = useState<NewEvent>();
    const [user, setUser] = useState<User | null>();
    const newEventDocRef = doc(db, 'users', user?.uid ?? '_', 'my-events', id);

    useEffect(() => {
        let unsubscribe: Unsubscribe | undefined;
        unsubscribe = onSnapshot(newEventDocRef, (snapshot) => {
            setNewEvent(snapshot.data() as NewEvent);
        });
        return () => {
            unsubscribe?.();
        };
    }, [user]);

    useEffect(() => {
        const unsubscribe = onAuthStateChanged(firebaseAuth, (user) => {
            console.log(user, 'user');
            setUser(user);
        });
        return () => {
            unsubscribe;
        };
    }, []);

    return newEvent ?
            <Grid2
                sx={{ marginBlock: '1rem' }}
                container
                component="form"
                spacing="2rem"
                onChange={(e) => {
                    const { value: inputValue, name: inputName, checked, type } = e.target as HTMLInputElement;

                    let value: string | boolean = inputValue;
                    let name: string = inputName;

                    if (type === 'checkbox') {
                        value = checked;
                    }

                    if (type === 'radio') {
                        name = 'gameType';
                        value = inputName;
                    }
                    if (user?.email !== null) {
                        const payload: Partial<NewEvent> = {
                            [name]: value,
                            updateAt: new Date(Date.now()).toString(),
                            updatedBy: user?.email,
                        };

                        setDoc(newEventDocRef, { ...newEvent, ...payload });
                    }
                }}
            >
                {isExploding && (
                    <Confetti
                        onConfettiComplete={() => {
                            setIsExploding(!isExploding);
                        }}
                        numberOfPieces={2000}
                        recycle={false}
                        height={document.body.scrollHeight}
                        width={document.body.scrollWidth}
                    />
                )}
                <Paper sx={{ marginInline: 2, padding: '1rem', marginBottom: '2rem' }}>
                    <Grid2 container gap="3rem">
                        <Typography variant="h1">Meld på arrangement til Regncon XXXII 2024</Typography>
                        <Typography>
                            Takk for at du vil arrangere eit spel på Regncon, anten det er brettspel, kortspel,
                            rollespel eller anna, så sett vi enormt pris på ditt bidrag. Fyll inn skjemaet så godt du
                            kan, og ikkje vere redd for å ta kontakt med Regnconstyret på{' '}
                            <a href="mailto:regncon@gmail.com">regncon@gmail.com</a>
                            om du skulle lure på noko!
                        </Typography>
                    </Grid2>
                </Paper>

                <Grid2 xs={12}>
                    <Paper>
                        <TextField
                            name="title"
                            label="Tittel på spelmodul / arrangement"
                            value={newEvent.title}
                            variant="outlined"
                            required
                            fullWidth
                        />
                    </Paper>
                </Grid2>
                <Grid2 xs={12}>
                    <Paper sx={{ padding: '1rem' }}>
                        <TextField
                            name="subTitle"
                            value={newEvent.subTitle}
                            label="Her kan du fylle ut ei kort skildring av modulen."
                            variant="outlined"
                            required
                            fullWidth
                        />
                    </Paper>
                </Grid2>
                <Grid2 xs={12} sm={6} md={3}>
                    <Paper>
                        <TextField
                            type="email"
                            name="email"
                            value={newEvent.email}
                            label="E-postadresse"
                            variant="outlined"
                            required
                            fullWidth
                        />
                    </Paper>
                </Grid2>
                <Grid2 xs={12} sm={6} md={3}>
                    <Paper>
                        <TextField
                            name="name"
                            value={newEvent.name}
                            label="Arrangørens namn (Ditt namn)"
                            variant="outlined"
                            required
                            fullWidth
                        />
                    </Paper>
                </Grid2>
                <Grid2 xs={12} sm={6} md={3}>
                    <Paper>
                        <TextField
                            type="phone"
                            name="phone"
                            value={newEvent.phone}
                            label="Kva telefonnummer kan vi nå deg på?"
                            variant="outlined"
                            required
                            fullWidth
                        />
                    </Paper>
                </Grid2>
                <Grid2 xs={12} sm={6} md={3}>
                    <Paper>
                        <TextField
                            name="system"
                            label="Spillsystem"
                            value={newEvent.system}
                            variant="outlined"
                            fullWidth
                        />
                    </Paper>
                </Grid2>
                <Grid2 xs={12}>
                    <Paper sx={{ padding: '1rem' }}>
                        <FormControl fullWidth>
                            <FormLabel>Skildring av modulen (tekst til programmet):</FormLabel>
                            <TextareaAutosize minRows={5} name="description" value={newEvent.description} />
                        </FormControl>
                    </Paper>
                </Grid2>
                <Grid2 xs={12} sm={4}>
                    <Paper sx={{ padding: '1rem' }}>
                        <FormControl fullWidth>
                            <FormLabel>Kva type spel er det?</FormLabel>
                            <RadioGroup
                                value={newEvent.gameType}
                                aria-labelledby="demo-controlled-radio-buttons-group"
                                name="controlled-radio-buttons-group"
                            >
                                <FormControlLabel
                                    value="rolePlaying"
                                    control={<Radio name="rolePlaying" />}
                                    label="rollespel"
                                />
                                <FormControlLabel
                                    value="boardGame"
                                    control={<Radio name="boardGame" />}
                                    label="Brettspel"
                                />
                                <FormControlLabel
                                    value="cardGame"
                                    control={<Radio name="cardGame" />}
                                    label="Kortspel"
                                />
                                <FormControlLabel value="other" control={<Radio name="other" />} label="Annet" />
                            </RadioGroup>
                        </FormControl>
                    </Paper>
                </Grid2>
                <Grid2 xs={12} sm={4}>
                    <Paper sx={{ padding: '1rem', height: '100%' }}>
                        <TextField
                            type="number"
                            name="participants"
                            value={newEvent.participants}
                            label="Maks antall deltakere"
                            variant="outlined"
                            required
                            fullWidth
                        />
                    </Paper>
                </Grid2>
                <Grid2 xs={12} sm={4}>
                    <Paper sx={{ padding: '1rem' }}>
                        <FormGroup>
                            <FormLabel>Kva for pulje kan du arrangere i?</FormLabel>
                            <FormControlLabel
                                control={<Checkbox checked={newEvent.fridayEvening} />}
                                name="fridayEvening"
                                label="Fredag Kveld"
                            />
                            <FormControlLabel
                                control={<Checkbox checked={newEvent.saturdayMorning} />}
                                name="saturdayMorning"
                                label="Lørdag Morgen"
                            />
                            <FormControlLabel
                                control={<Checkbox checked={newEvent.saturdayEvening} />}
                                name="saturdayEvening"
                                label="Lørdag Kveld"
                            />
                            <FormControlLabel
                                control={<Checkbox checked={newEvent.sundayMorning} />}
                                name="sundayMorning"
                                label="Søndag Morgen"
                            />
                        </FormGroup>
                    </Paper>
                </Grid2>
                <Grid2 xs={12}>
                    <Paper sx={{ padding: '1rem' }}>
                        <FormGroup>
                            <FormLabel>Kryss av for det som gjeld</FormLabel>
                            <FormControlLabel
                                control={<Checkbox name="moduleCompetition" checked={newEvent.moduleCompetition} />}
                                label="Eg vil vere med på modulkonkurransen husk å sende modulen til moduler@regncon.no innen første september!)"
                            />
                            <FormControlLabel
                                control={<Checkbox name="childFriendly" checked={newEvent.childFriendly} />}
                                label="Arrangementet passer for barn"
                            />
                            <FormControlLabel
                                control={<Checkbox name="adultsOnly" checked={newEvent.adultsOnly} />}
                                label="Arrangementet passer berre for vaksne (18+)"
                            />
                            <FormControlLabel
                                control={<Checkbox name="beginnerFriendly" checked={newEvent.beginnerFriendly} />}
                                label="Arrangementet er nybyrjarvenleg"
                            />
                            <FormControlLabel
                                control={<Checkbox name="possiblyEnglish" checked={newEvent.possiblyEnglish} />}
                                label="Arrangementet kan haldast på engelsk"
                            />
                            <FormControlLabel
                                control={<Checkbox name="volunteersPossible" checked={newEvent.volunteersPossible} />}
                                label="Andre kan halda arrangementet"
                            />
                            <FormControlLabel
                                control={<Checkbox name="lessThanThreeHours" checked={newEvent.lessThanThreeHours} />}
                                label="Eg trur arrangementet vil vare kortare enn 3 timer"
                            />
                            <FormControlLabel
                                control={<Checkbox name="moreThanSixHours" checked={newEvent.moreThanSixHours} />}
                                label="Eg trur arrangementet vil vare lenger enn 6 timer"
                            />
                        </FormGroup>
                    </Paper>
                </Grid2>
                <Grid2 xs={12}>
                    <Paper sx={{ padding: '1rem' }}>
                        <FormControl fullWidth>
                            <FormLabel>Merknader: Er det noko anna du vil vi skal vite?</FormLabel>
                            <TextareaAutosize
                                minRows={3}
                                name="additionalComments"
                                value={newEvent.additionalComments}
                            />
                        </FormControl>
                    </Paper>
                </Grid2>
                <Grid2 xs={12}>
                    <Paper sx={{ padding: '1rem' }}>
                        <Typography>
                            Skjemaet vert lagra automatisk, men om du likevel vil trykke på ein knapp, så er det ein
                            her. :)
                        </Typography>
                        <Button onClick={() => setIsExploding(!isExploding)}>Send inn</Button>
                    </Paper>
                </Grid2>
            </Grid2>
        :   null;
};
export default Create;
