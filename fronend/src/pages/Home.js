import React from "react";
import Ticket from "../component/forms/Ticket";
import '../styles/ticket.css'
import {Container} from "@mui/material";
import ButtonBases from "../component/buttonbases/ButtonBases";
import Test from "../component/Test";
import DrawerAppBar from "../component/pageElements/DrawerAppBar";


function Home() {

    return (
        <>
            <DrawerAppBar/>
            <div maxWidth="sm" className=" rmdp-rtl">
                    <Ticket/>
            </div>
            <ButtonBases/>
            {/*<Test/>*/}
        </>
    );
};

export default Home;
