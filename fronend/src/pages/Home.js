import React from "react";
import Ticket from "../component/forms/Ticket";
import '../styles/ticket.css'
import {Container} from "@mui/material";
import ButtonBases from "../component/buttonbases/ButtonBases";
import Test from "../component/Test";


function Home() {

    return (
        <>
            <Container maxWidth="sm" className=" rmdp-rtl">
                <row>
                    <Ticket/>
                </row>
            </Container>
            <ButtonBases/>
            {/*<Test/>*/}
        </>
    );
};

export default Home;
