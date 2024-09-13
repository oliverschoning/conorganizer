import type { ConEvent } from '$lib/types';
import { faUserSecret, faScroll } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { NavigateBefore } from '@mui/icons-material';
import { type SxProps, Paper, Box, Typography, Chip, type Theme } from '@mui/material';
import MuiMarkdown from 'mui-markdown';
import Image from 'next/image';
import blekksprut2 from '$public/blekksprut2.jpg';
import BackButton from '$app/(authorized)/components/BackButton';
import InterestSelector from './components/InterestSelector';

type Props = {
    event: ConEvent;
    prevNavigationId?: string;
    nextNavigationId?: string;
};

const MainEventBig = ({ event, prevNavigationId, nextNavigationId }: Props) => {
    console.log(typeof window === 'undefined' ? 'server' : 'client');
    const paragraphStyle: SxProps<Theme> = { margin: '1rem 0' };
    return (
        <Paper
            elevation={0}
            sx={{
                backgroundColor: 'black',
                '&': {
                    '--image-height': '15.1429rem',
                    '--slider-interest-width': '19.5714rem',
                    '--event-margin-left': '4rem',
                },
            }}
        >
            <Box
                sx={{
                    display: 'grid',
                    '& > *': { gridColumn: '1 / 2', gridRow: '1 / 2' },
                }}
            >
                <Box
                    component={Image}
                    src={blekksprut2}
                    alt="noe alt-tekst"
                    sx={{
                        width: '100%',
                        height: 'auto',
                        maxHeight: 'var(--image-height)',
                        maxWidth: '100%',
                        // aspectRatio: '3.3 / 2',
                    }}
                    placeholder="blur"
                    loading="lazy"
                />
                <Box
                    sx={{
                        background: `linear-gradient(0deg, black, transparent)`,
                        maxHeight: 'var(--image-height)',
                    }}
                >
                    <Box
                        sx={{
                            display: 'grid',
                            gridTemplateRows: '2rem 1fr',
                            height: '100%',
                            wordBreak: 'break-word',
                        }}
                    >
                        <Box sx={{ placeSelf: 'start' }}>
                            <BackButton />
                        </Box>
                        <Box sx={{ display: 'grid', gridTemplateColumns: '1fr 0.5fr' }}>
                            <Box
                                sx={{
                                    display: 'grid',
                                    gridTemplateRows: '1fr 1fr',
                                    gap: '1rem',
                                    marginBottom: '2rem',
                                }}
                            >
                                <Typography
                                    variant="h1"
                                    sx={{ marginInlineStart: 'calc(var(--event-margin-left) - 1.5rem)' }}
                                >
                                    {event.title || 'Tittel'}
                                </Typography>

                                <Box
                                    sx={{
                                        display: 'grid',
                                        gridTemplateColumns: 'repeat(2, max-content)',
                                        placeContent: 'start',
                                        marginInlineStart: 'var(--event-margin-left)',
                                    }}
                                >
                                    <Box sx={{ display: 'flex', gap: '0.8rem', placeItems: 'center' }}>
                                        <FontAwesomeIcon icon={faUserSecret} size="2x" style={{ color: '#ff7c7c' }} />
                                        <Box>
                                            <Typography component="span" sx={{ color: 'primary.main' }}>
                                                {event.icons?.includes('rollespill') ? 'Gamemaster' : 'Arrangør'}
                                            </Typography>
                                            <Typography variant="body1" margin={0}>
                                                {event.gameMaster || 'Navn'}
                                            </Typography>
                                        </Box>
                                    </Box>
                                    <Box sx={{ display: 'flex', gap: '0.8rem', placeItems: 'center' }}>
                                        <FontAwesomeIcon icon={faScroll} size="2x" style={{ color: '#ff7c7c' }} />
                                        <Box>
                                            <Typography component="span" sx={{ color: 'primary.main' }}>
                                                System
                                            </Typography>
                                            <Typography variant="body1" margin={0}>
                                                {event.system || 'System'}
                                            </Typography>
                                        </Box>
                                    </Box>
                                </Box>
                            </Box>
                            <Typography
                                sx={{
                                    corsor: 'pointer',
                                    ...paragraphStyle,
                                    marginBottom: '1rem',
                                    textAlign: 'center',
                                }}
                            >
                                {event.shortDescription || 'Kort beskrivelse'}
                            </Typography>
                        </Box>
                    </Box>
                </Box>
            </Box>
            <Box
                sx={{
                    padding: '0 1rem 1rem 0rem',
                    display: 'grid',
                    gridTemplateColumns: '1fr 1fr',
                    marginInlineStart: 'var(--event-margin-left)',
                }}
            >
                <Box>
                    <Box
                        sx={{
                            display: 'flex',
                            flexWrap: 'wrap',
                            gap: '.5em',
                            overflowX: 'auto',
                            paddingBottom: '0.35rem',
                        }}
                    >
                        {/* {event.icons?.map((tag) => (
                                <Chip label={tag} key={tag} color="primary" icon={<NavigateBefore />} />
                            ))} */}

                        <Chip label={'tag'} key={'tag'} color="primary" variant="outlined" icon={<NavigateBefore />} />
                        <Chip
                            label={'tatagtagtagg'}
                            key={'tag'}
                            color="primary"
                            variant="outlined"
                            icon={<NavigateBefore />}
                        />
                        <Chip
                            label={'tagtagtag'}
                            key={'tag'}
                            color="primary"
                            variant="outlined"
                            icon={<NavigateBefore />}
                        />
                        <Chip
                            label={'tagtagtag'}
                            key={'tag'}
                            color="primary"
                            variant="outlined"
                            icon={<NavigateBefore />}
                        />
                        <Chip
                            label={'tagtagtag'}
                            key={'tag'}
                            color="primary"
                            variant="outlined"
                            icon={<NavigateBefore />}
                        />
                        <Chip
                            label={'tagtagtag'}
                            key={'tag'}
                            color="primary"
                            variant="outlined"
                            icon={<NavigateBefore />}
                        />
                    </Box>
                    <InterestSelector />
                    <Box>
                        <Box
                            sx={{
                                display: 'grid',
                                gridTemplateColumns: !prevNavigationId || !nextNavigationId ? '1fr' : '1fr 1fr',
                                placeItems: 'space-between',
                                marginBlockStart: '1rem',
                                position: 'relative',
                                width: '100%',
                            }}
                        >
                            {/* <NavigatePreviousLink previousNavigationId={prevNavigationId} />
                                <NavigateNextLink nextNavigationId={nextNavigationId} /> */}
                        </Box>
                    </Box>
                </Box>

                <MuiMarkdown>{event.description || '# Lang beskrivelse'}</MuiMarkdown>
            </Box>
        </Paper>
    );
};

export default MainEventBig;
