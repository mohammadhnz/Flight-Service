import {
    BrowserRouter,
    Routes,
    Route,
} from "react-router-dom";
import './App.css';
import Home from "./pages/Home";
import SignIn from "./pages/SignIn";
import About from "./pages/About";
import DrawerAppBar from "./component/pageElements/DrawerAppBar";
import SignUp from "./pages/SignUp";
import FlightsList from "./pages/FlightsList";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home/>}/>
                <Route path="/sign-in" element={<SignIn/>}/>
                <Route path="/about" element={<About/>}/>
                <Route path="/sign-up" element={<SignUp/>}/>
                <Route path="/flights" element={<FlightsList/>}/>
            </Routes>
        </BrowserRouter>
    );
}

export default App;