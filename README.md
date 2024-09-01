# Con Organizer

## Description

This is the source code for the Regncon 2024 app. It's a work in progress, but the idea is to have a place to keep track
of your con schedule, your panels, and your expenses.

## referat fra post-mortem: sorteres inn i must-haves og nice-to-haves av neste styre

-   bedre loginsystem? Passwordless login?
-   tydeligere hva pameldingen gjor. tooltip? Bedre ord enn "pamelding"
-   klokkeslett klar for connet
-   bedre grensesnitt og automatisering av opprop? Forslag: liste med deltakere, liste med queue, mest populare forst
-   statistikk pa slutten av connet: hvor mange fikk spilt ting de hadde hvor lyst til? hvor mange dukket ikke opp? osv
-   britisk flagg som funker i alle nettlesere
-   pameldingsstatus tilbake til listviewen - hvordan med flere brukere? Bare prikker eller ikoner?
-   bash-script eller noe annen automatisering for billedbehandling. F.eks: https://www.npmjs.com/package/sharp
-   mulighet for at folk kan legge inn arrangement selv (som skjult, sa kan styret godkjenne og evt endre)
    -   obs: folk ma vare innlogget, sa de ma ha kjopt billett?
    -   obs: hvor mye tekstformatering skal vi ha med? Fet og kursiv, f.eks? Lenker?
    -   obs: fint hvis folk kan laste opp bilder selv og f.eks. bruke https://www.npmjs.com/package/cropperjs
    -   fint hvis folk kan laste opp modulen ogsa, og sjekke av for at den skal vare med i modulkonkurranse
    -   mulighet for a trekke arrangementet?
-   refakturere og/eller lage fra scratch med andre teknologier
-   brukertesting av grensesnittet, gi folk oppdrag som "meld deg pa the One Ring" og se hva som skjer

## Todo

### Must haves for publishing events

-   Fix cookie lifetime
-   Bug: Problems with hitting enter while editing event
-   Easy next and previous navigation
-   Fix bug in debounce
-   Convert my events to use new event system
-   Desktop view
-   Icons
-   Help text explaining how the algorithm works
-   Store interests
-   Connect to checkin system
-   Participants can choose favorite events
-   Display label on new or unread events
-   Debounce/fix text box, my events
-   Send emails from server
-   Create participant in database
-   Get tickets from checkin

#### Nice to haves for publishing events

-   Picture upload
-   Check preloading of pages
-   Fix small card layout

### Must haves for publishing full program

-   Make it very clear what day is displayed
-   Filter out my signups
-   Add is gm to event

## Nice to haves

-   Add keyboard navigation
-   Add command palatte
-   Add search
-   Tags
-   Add event to calendar
-   storskjermversjon som viser resten av dagen automatisk

## Done

-   Add none events to list
-   Match header and footer to design
-   Check if firebase token should be private
-   Add media queries mobile first
-   Add event images
-   Add error boundary
-   Save user choices for events
-   Add timeslots
-   Add rooms
-   Add favicon
-   Set page title
-   Authorization
-   Auth
-   query string for event id
-   Add event types
-   Add read more
-   Unlist event
-   Get event from database
-   Add new event
-   Edit event
