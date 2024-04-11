import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import Paper from '@mui/material/Paper';
import Typography from '@mui/material/Typography';
import styles from './page.module.scss';
import EventCardBig from './EventCardBig';
import EventCardSmall from './EventCardSmall';
import { getAll } from './serverAction';
import RealtimeEvents from './RealtimeEvents';

export default async function Home() {
	const events = await getAll();
	console.log(events);

	return (
		<>
			<img src="/placeholderlogo.png" alt="logo" />
			<Box
				sx={{
					display: 'flex',
					flexWrap: 'wrap',
					maxWidth: '320px',
					gap: '0.5rem',
					placeContent: 'center',
				}}
			>
				{events.map((event, i) => {
					if (i === 0) {
						return (
							<EventCardBig
								key={i}
								title={event.title}
								gameMaster={event.gameMaster}
								shortDescription={event.shortDescription}
								system={event.system}
							/>
						);
					}
					return (
						<EventCardSmall
							key={i}
							title={event.title}
							gameMaster={event.gameMaster}
							system={event.system}
						/>
					);
				})}
			</Box>
			{/* <EventCardBig
    					title="Hello world"
    					gameMaster="Gerhard Fajita"
    					shortDescription="Mord overalt! Kos! Gøy!"
    					system="Call of Chthuhlth"
    				/>
    				<Box sx={{ display: 'flex' }}>
    					<EventCardSmall title="Hi" gameMaster="Gardh Fajita2" system="Dungeons 2" />
    					<EventCardSmall title="Any% speedrun" gameMaster="Gorde Fajita3" system="Terraria" />
    				</Box> */}

			<RealtimeEvents />
		</>
	);
}
