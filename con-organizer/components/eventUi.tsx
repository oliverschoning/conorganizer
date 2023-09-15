"use client";

import React from "react";
import Typography from "@mui/material/Typography";
import { FormControl, FormLabel, RadioGroup, FormControlLabel, Radio, FormControlLabelProps, styled, useRadioGroup, Box, Card, CardHeader, CardContent, CardMedia, Button, CardActions, Chip, Collapse, IconButton, Divider } from "@mui/material";
import { ConEvent } from "@/lib/types";
import parse from 'html-react-parser';
import { faDiceD20 } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import CloseIcon from '@mui/icons-material/Close';
import EventHeader from "./eventHeader";


interface Props {
    conEvent: ConEvent;
}

const EventUi = ({ conEvent }: Props) => {

    interface StyledFormControlLabelProps extends FormControlLabelProps {
        checked: boolean;
    }

    const StyledFormControlLabel = styled((props: StyledFormControlLabelProps) => <FormControlLabel {...props} />)(
        ({ theme, checked }) => ({
            '.MuiFormControlLabel-label': checked && {
                color: theme.palette.primary.main,
            },
        })
    );

    function MyFormControlLabel(props: FormControlLabelProps) {
        const radioGroup = useRadioGroup();

        let checked = false;

        if (radioGroup) {
            checked = radioGroup.value === props.value;
        }

        return <StyledFormControlLabel checked={checked} {...props} />;
    }

    return (
        <Card
        sx={{ minHeight: '900px', maxWidth: '440px' }}>
        <EventHeader conEvent={conEvent} />

        <Divider />
        <Typography className='pl-4 pr-4'>{parse(conEvent?.description || '')}</Typography>

        <Divider />
        <CardContent>
            <FormControl className="p-4">
                <FormLabel id="demo-row-radio-buttons-group-label">Puljepåmelding</FormLabel>
                <RadioGroup
                    row
                    aria-labelledby="demo-row-radio-buttons-group-label"
                    name="row-radio-buttons-group"
                    defaultValue="NotInterested"
                >
                    <MyFormControlLabel
                        value="NotInterested"
                        control={<Radio size="small" />}
                        label="Ikke intresert"
                    />
                    <MyFormControlLabel value="IfIHaveTo" control={<Radio size="small" />} label="Hvis jeg må" />
                    <MyFormControlLabel value="IWantTo" control={<Radio size="small" />} label="Har lyst" />
                    <MyFormControlLabel
                        value="RealyWantTo"
                        control={<Radio size="small" />}
                        label="Har veldig lyst"
                    />
                </RadioGroup>
            </FormControl>
        </CardContent>
        <Divider />
    </Card>
    );
};

export default EventUi;
