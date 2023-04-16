import { createTheme } from '@mui/material/styles'
import { red, grey } from '@mui/material/colors'

// Create a theme instance.
export const lightTheme = createTheme({
  palette: {
    mode: 'light',
    primary: {
      main: '#556cd6',
    },
    secondary: {
      main: red.A400,
    },
    error: {
      main: red.A400,
    },
  },
})

export const darkTheme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: '#556cd6',
      light: '#f1f5f9',
    },
    secondary: {
      main: '#94a3b8',
    },
    error: {
      main: red.A400,
    },
    background: {
      default: grey[900],
    },
    text: {
      primary: grey[900],
      secondary: grey[800],
    },
  },
})
