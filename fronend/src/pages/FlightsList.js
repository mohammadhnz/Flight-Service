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
    const isLimited = false;
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
                                         bussinessP={flight.j_price} economyP={flight.y_price}
                                         firstClassP={flight.f_price}
                                         isLimited={isLimited} flightId = {flight.flight_id}

                        />
                    </>
                ))}
            </Grid>
        </div>
    );
};

export default Flight;
