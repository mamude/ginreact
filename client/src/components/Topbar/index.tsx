import React from 'react'
import { AppBar, Box, Button, Hidden, Toolbar } from '@material-ui/core'
import ExitToAppIcon from '@material-ui/icons/ExitToApp'
import ShoppingBasketIcon from '@material-ui/icons/ShoppingBasket'
import { ImgLogo } from './style'
import logo from '../../images/logo.png'

const Topbar: React.FC = () => {
  return (
    <AppBar position="fixed">
      <Toolbar>
        <ImgLogo src={logo} alt="Samir Mamude" />
        <Hidden mdDown>
          <Box flexGrow={1} />
          <Button variant="text" color="inherit" startIcon={<ExitToAppIcon />}>
            Entrar
          </Button>
          <Button
            variant="text"
            color="inherit"
            startIcon={<ShoppingBasketIcon />}
          >
            Sacola
          </Button>
        </Hidden>
      </Toolbar>
    </AppBar>
  )
}

export default Topbar
