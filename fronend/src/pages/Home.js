import React from "react";
import Ticket from "../component/forms/Ticket";
import '../styles/ticket.css'
import ButtonBases from "../component/buttonbases/ButtonBases";
import DrawerAppBar from "../component/pageElements/DrawerAppBar";
import Grid from "@mui/material/Grid";


function Home() {

    return (
        <div className=" rmdp-rtl">
            <DrawerAppBar/>
            <Grid container justifyContent="center" sx={{height: '55vh'}}>
                    <Ticket/>
            </Grid>
            <Grid justifyContent="center">
                <ButtonBases/>
            </Grid>
        </div>
    );
};

export default Home;
