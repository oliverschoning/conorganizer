'use client';

import React from 'react';
import AddCircleIcon from '@mui/icons-material/AddCircle';
import { IconButton } from '@mui/material';
import { CollectionReference, DocumentData } from 'firebase/firestore';
import EditDialog from './editDialog';

type Props = {};

const AddEvent = ({}: Props) => {
    const [open, setOpen] = React.useState(false);
    const handleClose = () => {
        setOpen(false);
    };

    return (
        <>
            <IconButton
                aria-label="add"
                color="error"
                size="large"
                onClick={() => {
                    setOpen(true);
                }}
            >
                <AddCircleIcon />
            </IconButton>
            <EditDialog open={open} handleClose={handleClose} />
        </>
    );
};

export default AddEvent;
