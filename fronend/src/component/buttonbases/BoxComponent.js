import * as React from 'react';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';


export default function BoxComponent({className, classP}) {

    return (
        <Box component="span" sx={{
            bgcolor: 'background.paper',
            borderColor: 'text.primary',
            border: 1,
            borderRadius: '16px',
            width: 180,
            height: 50,
        }}>
            <Button>
                {className} ${classP}
            </Button>
        </Box>
    );
}