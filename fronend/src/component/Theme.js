import {createMuiTheme} from "@mui/material";

const theme = createMuiTheme({
  overrides: {
    MuiCssBaseline: {
      "@global": {
        "*, *::before, *::after": {
          boxSizing: "content-box",
        },

        body: {
          backgroundColor: "#37373b",
        },
        AppBar : {
          backgroundColor: "#37373b",
          color: "#37373b"
        }
      },
    },
  },
});
export default theme;