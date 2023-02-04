import * as React from 'react';
import AspectRatio from '@mui/joy/AspectRatio';
import Link from '@mui/joy/Link';
import Card from '@mui/joy/Card';
import Chip from '@mui/joy/Chip';
import Typography from '@mui/joy/Typography';
import airplane from './img/airplane.png'
import Button from "@material-ui/core/Button";
import {Grid3x3Outlined} from "@mui/icons-material";
import Grid from "@mui/material/Grid";
import {Paper} from "@mui/material";
import * as PropTypes from "prop-types";
import Box from "@mui/material/Box";
import BoxComponent from "../buttonbases/BoxComponent";
import {useNavigate} from "react-router-dom";

function Item(props) {
    return null;
}

Item.propTypes = {children: PropTypes.node};
export default function InteractiveCard({
                                            origin,
                                            destination,
                                            departure_local_time,
                                            arrival_local_time,
                                            hour,
                                            minute,
                                            isLimited,
                                            className,
                                            economyP,
                                            bussinessP,
                                            firstClassP
                                        }) {
    const dep_time = new Date(departure_local_time).toJSON().slice(0, 10);
    const arr_time = new Date(arrival_local_time).toJSON().slice(0, 10);
    const navigate = useNavigate();
    const handleClick = (event) => {
        navigate('/about')
    }
    return (
        <Card
            variant="outlined"
            orientation="horizontal"
            sx={{
                width: 1000,
                gap: 1,
                '&:hover': {boxShadow: 'md', borderColor: 'neutral.outlinedHoverBorder'},
            }}
        >
            <AspectRatio ratio="1" sx={{width: 80}}>
                <img src={airplane} loading="lazy" alt=""/>
            </AspectRatio>
            <div>
                <Typography level="h2" fontSize="lg" id="card-description" sx={{fontWeight: 'md'}}>
                    از {origin} به  {destination}
                </Typography>

                <Typography level="h2" fontSize="lg" id="card-description" sx={{fontWeight: 'md'}}>
                    {dep_time} ------ {arr_time}
                </Typography>
                <Typography fontSize="sm" aria-describedby="card-description" mb={1}>
                    {minute} : {hour}
                </Typography>
                <Grid container direction="row" justifyContent="flex-end" alignItems="flex-end" spacing={1}
                      columns={12}>
                    <Grid container
                          justifyContent="center"
                          alignItems="center"
                          item xs={4}
                    >
                        <BoxComponent textAlign="center" className="Economy" classP={economyP}/>
                    </Grid>
                    <Grid container
                          justifyContent="center"
                          alignItems="center"
                          item xs={4}>
                        <BoxComponent  textAlign="center" className="Business" classP={bussinessP}/>
                    </Grid>
                    <Grid container justifyContent="center"
                          alignItems="center"
                          item xs={4}>
                        <BoxComponent onClick={handleClick} textAlign="center" className="First Class" classP={firstClassP}/>
                    </Grid>
                </Grid>
                {isLimited && (
                    <div className="text-danger" style={{color:"red"}}>ظرفیت محدود است</div>
                )}
                {/*<Button variant="outlined" className="btn btn-primary">*/}
                {/*    <Link*/}
                {/*        overlay*/}
                {/*        underline="none"*/}
                {/*        href={className}*/}
                {/*        sx={{color: 'text.tertiary'}}*/}
                {/*    >خرید*/}
                {/*    </Link>*/}
                {/*</Button>*/}
            </div>
        </Card>
    );
}
