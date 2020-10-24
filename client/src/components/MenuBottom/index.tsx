import React from 'react'
import {
  BottomNavigation,
  BottomNavigationAction,
  Hidden,
} from '@material-ui/core'
import HomeIcon from '@material-ui/icons/Home'
import SearchIcon from '@material-ui/icons/Search'
import AssignmentIcon from '@material-ui/icons/Assignment'
import PersonIcon from '@material-ui/icons/Person'
import { Footer } from './style'

const MenuBottom: React.FC = () => {
  return (
    <Hidden lgUp>
      <Footer>
        <BottomNavigation showLabels>
          <BottomNavigationAction
            label="Início"
            value="inicio"
            icon={<HomeIcon />}
          />
          <BottomNavigationAction
            label="Busca"
            value="busca"
            icon={<SearchIcon />}
          />
          <BottomNavigationAction
            label="Pedidos"
            value="pedido"
            icon={<AssignmentIcon />}
          />
          <BottomNavigationAction
            label="Perfil"
            value="perfil"
            icon={<PersonIcon />}
          />
        </BottomNavigation>
      </Footer>
    </Hidden>
  )
}

export default MenuBottom
