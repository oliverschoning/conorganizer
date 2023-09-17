import { faChess, faChessKing, faDiceD20, faHatWizard, faPalette } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Alert, Box, CardHeader, CardMedia, Chip, Typography } from '@mui/material';
import { gameType } from '@/lib/enums';
import { ConEvent } from '@/lib/types';

type Props = {
    conEvent: ConEvent;
    listView?: boolean;
};

const EventHeader = ({ conEvent, listView = false }: Props) => {
    return (
        <>
            {conEvent?.published === false ? (
                <Alert severity="warning" sx={{ marginBottom: '1rem' }}>
                    Dette arrangementet er ikke publisert enda.
                </Alert>
            ) : null}
            <Box
                sx={{
                    backgroundImage: 'url(/placeholder.jpg)',
                    backgroundSize: 'cover',
                }}
            >
                <Box
                    sx={{
                        background: 'linear-gradient(transparent, transparent, black)',
                        // height: '50vh',
                        height: listView ? '10em' : '25em',
                        // minHeight: '300px',
                        // maxHeight: '500px',
                        display: 'flex',
                        alignItems: 'end',
                    }}
                >
                    <Box
                        sx={{
                            color: 'white',
                            maxWidth: { xs: '95vw', md: '1080px' },
                            padding: '.7em',
                        }}
                    >
                        <Typography variant="h3">{conEvent?.title}</Typography>
                        <Typography variant="h4">{conEvent?.subtitle}</Typography>
                    </Box>
                </Box>
            </Box>

            <Box
                sx={{
                    display: 'grid',
                    justifyItems: 'start',
                    alignItems: 'center',
                    gridTemplateColumns: 'auto 1fr',
                    gap: '1em',
                    backgroundColor: 'white',
                    color: 'black',
                    width: { xs: '95vw', md: '1080px' },
                }}
            >
                <span>
                    {conEvent?.gameType === gameType.roleplaying ? (
                        <Box sx={{ p: '1em', placeItems: 'center', display: 'grid', gap: '.5em' }}>
                            <FontAwesomeIcon icon={faDiceD20} fontSize="2em" color="orangered" />
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
                        <Box sx={{ p: '1em', placeItems: 'center', display: 'grid', gap: '.5em' }}>
                            <FontAwesomeIcon icon={faChessKing} fontSize="2em" color="orangered" />
                            <Typography variant="caption">Brettspill</Typography>
                        </Box>
                    ) : null}
                    {conEvent?.gameType === gameType.other ? (
                        <Box sx={{ p: '1em', placeItems: 'center', display: 'grid', gap: '.5em' }}>
                            <FontAwesomeIcon icon={faPalette} fontSize="2em" color="orangered" />
                            <Typography variant="caption">Annet</Typography>
                        </Box>
                    ) : null}
                </span>
                <Box sx={{ display: 'grid', width: '100vw', maxWidth: '1080px' }}>
                    <span>{conEvent.gameSystem} </span>
                    <span>{conEvent.room} </span>
                    <span>{conEvent.pool} </span>
                    <span>{conEvent.host} </span>
                </Box>
            </Box>
        </>
    );
};

export default EventHeader;
