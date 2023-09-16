'use client';

import React from 'react';
import { Card, IconButton } from '@mui/material';
import AddCircleIcon from '@mui/icons-material/AddCircle';
import EditDialog from './editDialog';
import { CollectionReference, DocumentData } from 'firebase/firestore';

interface Props {
    collectionRef: CollectionReference<DocumentData, DocumentData>;
}

const AddEvent = ({ collectionRef }: Props) => {
    const [open, setOpen] = React.useState(false);
    const handleClose = () => {
        setOpen(false);
    };

    return (
        <>
            <Card>
                <h1 className="text-4xl">Con Organizer</h1>
            </Card>
            <IconButton
                className="fixed right-0 bottom-0 z-50"
                aria-label="add"
                color="error"
                size="large"
                onClick={() => {
                    setOpen(true);
                }}
            >
                <AddCircleIcon />
            </IconButton>
            <EditDialog open={open} collectionRef={collectionRef} handleClose={handleClose} />
        </>
    );
};

export default AddEvent;
