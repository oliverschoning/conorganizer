import { Card } from '@mui/material';
import { Route } from 'next';
import Link from 'next/link';
import { ConEvent } from '@/models/types';
import EventHeader from './EventHeader';

type Props = {
    conEvent: ConEvent;
};

const EventCard = ({ conEvent }: Props) => {
    // if (conEvent.id === '16ezgmkATFy0afYRw4d7') {
    //     throw new Error(
    //         'lorem ipsum dolor sit amet, consectetur adip lorem ipsum dolor sit amet. lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet. lorem dolor sit am et. Lorem ipsum dolor sit am et. lorem dolor sit am et. Lorem ipsum dolor, lorem ipsum dolor sit am et. Lorem ipsum'
    //     );
    // }
    return (
        <Card
            key={conEvent.id}
            component={Link}
            href={`/event/${conEvent.id}` as Route}
            sx={{
                width: { xs: '100vw', md: '500px' },
                cursor: 'pointer',
                opacity: conEvent?.published === false ? '50%' : '',
                textDecoration: 'none',
            }}
        >
            <EventHeader conEvent={conEvent} listView={true} />
        </Card>
    );
};

export default EventCard;
