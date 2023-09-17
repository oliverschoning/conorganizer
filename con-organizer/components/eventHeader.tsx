import { faChessKing, faDiceD20, faHatWizard, faPalette } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Alert, Box, CardHeader, CardMedia, Chip, Typography } from '@mui/material';
import { gameType } from '@/lib/enums';
import { ConEvent } from '@/lib/types';

type Props = {
    conEvent: ConEvent;
};

const EventHeader = ({ conEvent }: Props) => {
    return (
        <>
            {conEvent?.published === false ? (
                <Alert severity="warning" sx={{ marginBottom: '1rem' }}>
                    Dette arrangementet er ikke publisert enda.
                </Alert>
            ) : null}
            <Box sx={{ backgroundImage: 'url(/placeholder.jpg)', backgroundSize: 'cover' }}>
                <CardHeader
                    title={conEvent?.title}
                    subheader={conEvent?.subtitle}
                    sx={{
                        background: 'linear-gradient(transparent, transparent, black)',
                        // height: '50vh',
                        height: '10em',
                        // minHeight: '300px',
                        // maxHeight: '500px',
                        color: 'white',
                        alignItems: 'end',
                    }}
                />
            </Box>

            <Box
                sx={{
                    display: 'grid',
                    justifyItems: 'start',
                    alignItems: 'center',
                    gridTemplateColumns: 'auto 1fr',
                    gap: '1em',
                }}
            >
                <span>
                    {conEvent?.gameType === gameType.roleplaying ? (
                        <Box sx={{ p: '1em', placeItems: 'center', display: 'grid', gap: '.5em' }}>
                            <FontAwesomeIcon icon={faDiceD20} fontSize="2em" color="red" />
                            <Typography variant="caption">Rollespill</Typography>
                        </Box>
                    ) : // <Chip
                    //     icon={<FontAwesomeIcon icon={faDiceD20} />}
                    //     label="Rollespill"
                    //     size="small"
                    //     variant="outlined"
                    // />
                    null}
                    {conEvent?.gameType === gameType.boardgame ? (
                        <Chip
                            icon={<FontAwesomeIcon icon={faChessKing} />}
                            label="Brettspill"
                            size="small"
                            variant="outlined"
                        />
                    ) : null}
                    {conEvent?.gameType === gameType.other ? (
                        <Chip
                            icon={<FontAwesomeIcon icon={faPalette} />}
                            label="Annet"
                            size="small"
                            variant="outlined"
                        />
                    ) : null}
                </span>
                <Box sx={{ display: 'grid' }}>
                    <span>{conEvent.gameSystem} </span>
                    <span>{conEvent.room} </span>
                    <span>{conEvent.pool} </span>
                </Box>
            </Box>
        </>
    );
};

export default EventHeader;
