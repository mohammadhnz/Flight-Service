import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import CssBaseline from '@mui/material/CssBaseline';
import Divider from '@mui/material/Divider';
import Drawer from '@mui/material/Drawer';
import IconButton from '@mui/material/IconButton';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemText from '@mui/material/ListItemText';
import MenuIcon from '@mui/icons-material/Menu';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import {Link} from "@mui/material";
import {useNavigate} from "react-router-dom";
import About from "../../pages/About";
import { red } from '@mui/material/colors';
import {MuiThemeProvider} from "@material-ui/core";
import theme from "../Theme";

const drawerWidth = 240;
const navItems = ['ثبت نام', 'ورود', 'تماس با ما'];

function DrawerAppBar(props) {
    const {window} = props;
    const [mobileOpen, setMobileOpen] = React.useState(false);

    const handleDrawerToggle = () => {
        setMobileOpen((prevState) => !prevState);
    };

    const navigate = useNavigate();

    function handleClick() {
        navigate("/sign-up");
    }

    const drawer = (
        <Box onClick={handleDrawerToggle} sx={{textAlign: 'center'}}>
            <Typography variant="h6" sx={{my: 2}}>
                MUI
            </Typography>
            <Divider/>
            <List>
                {navItems.map((item) => (<ListItem key={item} disablePadding>
                    <ListItemButton sx={{textAlign: 'center'}}>
                        <ListItemText primary={item}/>
                    </ListItemButton>
                </ListItem>))}
            </List>
        </Box>
    );

    const container = window !== undefined ? () => window().document.body : undefined;

    return (
        <Box sx={{display: 'flex', flexGrow: 1}}>
            <MuiThemeProvider theme={theme}/>
            <AppBar component="nav" style={{ background: '#37373b' }}>
                <Toolbar>
                    <IconButton
                        color="#fff"
                        aria-label="open drawer"
                        edge="start"
                        onClick={handleDrawerToggle}
                        sx={{mr: 2, display: {sm: 'none'}}}
                    >
                        <MenuIcon/>
                    </IconButton>
                    <Typography
                        variant="h6"
                        component="div"
                        sx={{flexGrow: 1, display: {xs: 'none', sm: 'block'}}}
                    >
                        پرواز
                    </Typography>
                    <Box sx={{display: {xs: 'none', sm: 'block'}}}>
                        {navItems.map((item) => (
                            <>
                                <Button onClick={handleClick} key={item} sx={{color: '#fff'}}>
                                    {item}
                                </Button>
                                {/*<Button component={Link} to="/about" key={item} sx={{color: '#fff'}}>*/}
                                {/*    {item}*/}
                                {/*</Button>*/}
                                {/*<Routes>*/}
                                {/*    <Route path="/about" element={<SignIn/>}/>*/}
                                {/*</Routes>*/}
                            </>
                        ))}

                    </Box>
                </Toolbar>
            </AppBar>
            <Box component="nav">
                <Drawer
                    container={container}
                    variant="temporary"
                    open={mobileOpen}
                    onClose={handleDrawerToggle}
                    ModalProps={{
                        keepMounted: true, // Better open performance on mobile.
                    }}
                    sx={{
                        display: {xs: 'block', sm: 'none'},
                        '& .MuiDrawer-paper': {boxSizing: 'border-box', width: drawerWidth},
                    }}
                >
                    {drawer}
                </Drawer>
            </Box>
            <Box component="main" sx={{p: 3}}>
                <Toolbar/>
            </Box>
        </Box>);
}

// DrawerAppBar.propTypes = {
//   /**
//    * Injected by the documentation to work in an iframe.
//    * You won't need it on your project.
//    */
//   window: PropTypes.func,
// };

export default DrawerAppBar;
