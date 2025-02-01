-- Rooms
INSERT INTO rooms (name) VALUES
('Room A'),
('Room B');

INSERT INTO rooms_map_data (name, x, y) VALUES
('Room A', 10, 10),
('Room B', 20, 20);

-- Ticket types
INSERT INTO ticket_types (name) VALUES
('Weekend Pass'),
('Day Pass');

-- Users
INSERT INTO users (email, is_admin) VALUES
('[email protected]', TRUE),
('[email protected]', FALSE),
('[email protected]', FALSE);

-- Ticketholders
INSERT INTO ticketholders (
    first_name,
    last_name,
    ticket_type,
    is_over_18,
    order_id,
    ticket_id,
    ticket_email,
    order_email,
    ticket_category_id
) VALUES
('John', 'Doe', 'Weekend Pass', TRUE, 1001, 10001, '[email protected]', '[email protected]', 'CAT1'),
('Jane', 'Smith', 'Day Pass', FALSE, 1002, 10002, '[email protected]', '[email protected]', 'CAT2'),
('Alice', 'Johnson', 'Weekend Pass', TRUE, 1003, 10003, '[email protected]', '[email protected]', 'CAT3');

-- Linking ticketholders to users
INSERT INTO ticketholders_users (ticketholder_id, user_id) VALUES
(1, 2),
(2, 3);

-- Suggested events (ensure no NULL for IDs; these will auto-increment to 1, 2, 3, 4)
INSERT INTO suggested_event (
    host,
    title,
    description,
    image_url,
    system,
    max_players,
    child_friendly,
    adults_only,
    beginner_friendly,
    experienced_only,
    can_be_run_in_english,
    long_running,
    short_running,
    submitted_time
) VALUES
(2, 'Dungeon Quest', 'A thrilling dungeon crawler!', 'https://example.com/image1.jpg', 'D&D 5e', 5, 1, 0, 1, 0, 1, 0, 1, '2025-01-01T10:00:00Z'),
(3, 'Space Battle', 'An epic sci-fi scenario!', 'https://example.com/image2.jpg', 'Starfinder', 6, 0, 1, 0, 0, 1, 1, 0, '2025-02-02T12:00:00Z'),
(2, 'Medieval Tournament', 'Jousting and archery events', 'https://example.com/mt1.jpg', 'Pathfinder', 8, 1, 0, 1, 0, 1, 0, 1, '2025-03-03T15:00:00Z'),
(3, 'Cyber Heist', 'High-tech infiltration mission', 'https://example.com/ch1.jpg', 'Shadowrun', 6, 0, 1, 0, 1, 1, 1, 0, '2025-04-04T09:00:00Z');

-- Events (referencing suggested_event_id = 1..4, so no NULL)
INSERT INTO events (
    suggested_event_id,
    title,
    description,
    image_url,
    system,
    host_name,
    host,
    room_name,
    pulje_name,
    max_players,
    child_friendly,
    adults_only,
    beginner_friendly,
    experienced_only,
    can_be_run_in_english,
    long_running,
    short_running
) VALUES
(1, 'Dungeon Quest - Part 1', 'Exploring the initial caves', 'https://example.com/dq1.jpg', 'D&D 5e', 'Adam the DM', 2, 'Room A', 'Fredag kveld', 5, 1, 0, 1, 0, 1, 0, 1),
(1, 'Dungeon Quest - Part 2', 'Deeper dungeon exploration', 'https://example.com/dq2.jpg', 'D&D 5e', 'Adam the DM', 2, 'Room B', 'Lørdag morgen', 5, 1, 0, 1, 0, 1, 0, 1),
(2, 'Space Battle - Outpost Assault', 'Attack on a remote station', 'https://example.com/sb1.jpg', 'Starfinder', 'Eve the Captain', 3, 'Room A', 'Lørdag kveld', 6, 0, 1, 0, 0, 1, 1, 0),
(3, 'Medieval Tournament', 'Jousting and archery events', 'https://example.com/mt1.jpg', 'Pathfinder', 'Sir Gregory', 2, 'Room B', 'Søndag morgen', 8, 1, 0, 1, 0, 1, 0, 1),
(4, 'Cyber Heist', 'High-tech infiltration mission', 'https://example.com/ch1.jpg', 'Shadowrun', 'Jane the Hacker', 3, 'Room A', 'Fredag kveld', 6, 0, 1, 0, 1, 1, 1, 0);

-- Interests
INSERT INTO interests (ticketholder_id, event_id, interest_level) VALUES
(1, 1, 'Middels interessert'),
(2, 1, 'Veldig interessert'),
(1, 2, 'Litt interessert');

-- Events players
INSERT INTO events_players (event_id, ticketholder_id, interest_level) VALUES
(1, 1, 'Veldig interessert'),
(1, 2, 'Middels interessert');