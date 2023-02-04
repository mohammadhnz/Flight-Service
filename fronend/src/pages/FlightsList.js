import React from "react";
import '../styles/ticket.css'
import DrawerAppBar from "../component/pageElements/DrawerAppBar";
import InteractiveCard from "../component/pageElements/InteractiveCard";
import flights from "../static/flights.json";
import Grid from "@mui/material/Grid";
import '../styles/ticket.css';
import Ticket from "../component/forms/Ticket";


function Flight() {
    const flightsData = flights;
    return (
        <div className=" rmdp-rtl">
            <DrawerAppBar/>
            <Grid center  sx={{ height: '55vh' }}>
                <Ticket/>
            </Grid>
            <Grid container justifyContent="center">
                {flightsData.map((flight) => (
                    <>
                        <InteractiveCard origin={flight.origin} destination={flight.destination}
                                         departure_local_time={flight.departure_local_time}
                                         arrival_local_time={flight.arrival_local_time}
                                         hour={flight.duration.hours} minute={flight.duration.minutes}
                                         isLimited={flight.is_limited_capacity} className={flight.class_name}
                                         bussinessP={flight.price.Business} economyP={flight.price.Economy}
                                         firstClassP={flight.price.First_Class}
                        />
                    </>
                ))}
            </Grid>
        </div>
    );
};

export default Flight;
