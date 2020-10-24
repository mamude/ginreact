import React from 'react'
import { AppBar, Box, Button, Hidden, Toolbar } from '@material-ui/core'
import ExitToAppIcon from '@material-ui/icons/ExitToApp'
import ShoppingBasketIcon from '@material-ui/icons/ShoppingBasket'
import { ImgLogo, Root, TopButtons } from './style'
import logo from '../../images/logo.png'

const Topbar: React.FC = () => {
  return (
    <Root>
      <AppBar position="fixed">
        <Toolbar>
          <ImgLogo src={logo} alt="Samir Mamude" />
          <Hidden mdDown>
            <Box flexGrow={1} />
            <Button color="inherit" startIcon={<ExitToAppIcon />}>
              Entrar
            </Button>
            <Button color="inherit" startIcon={<ShoppingBasketIcon />}>
              Sacola
            </Button>
          </Hidden>
        </Toolbar>
      </AppBar>
    </Root>
  )
}

export default Topbar
