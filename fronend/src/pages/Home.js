import React from "react";
import Ticket from "../component/forms/Ticket";
import '../styles/ticket.css'
import ButtonBases from "../component/buttonbases/ButtonBases";
import DrawerAppBar from "../component/pageElements/DrawerAppBar";


function Home() {

    return (
        <>
            <DrawerAppBar/>
            <div maxWidth="sm" className=" rmdp-rtl">
                    <Ticket/>
            </div>
            <ButtonBases/>
        </>
    );
};

export default Home;
