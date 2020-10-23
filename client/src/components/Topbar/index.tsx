import React from 'react'
import { AppBar, Button, IconButton, Toolbar } from '@material-ui/core'
import MenuIcon from '@material-ui/icons/Menu'
import { Root, ShoppingName } from './style'

const Topbar: React.FC = () => {
  return (
    <Root>
      <AppBar position="fixed">
        <Toolbar>
          <IconButton edge="start" color="inherit" aria-label="menu">
            <MenuIcon />
          </IconButton>
          <ShoppingName variant="h6">E-commerce</ShoppingName>
          <Button color="inherit">Login</Button>
        </Toolbar>
      </AppBar>
    </Root>
  )
}

export default Topbar
