import { createMuiTheme } from '@material-ui/core'

const theme = createMuiTheme({
  palette: {
    primary: {
      main: '#b71c1c',
    },
    secondary: {
      main: '#fb8c00',
    },
  },
  props: {
    MuiTextField: {
      variant: 'outlined',
      size: 'small',
    },
    MuiButton: {
      variant: 'contained',
    },
  },
  overrides: {
    MuiButton: {
      root: {
        marginTop: 16,
        marginBottom: 16,
      },
    },
  },
})

export default theme
