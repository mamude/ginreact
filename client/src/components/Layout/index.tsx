import React from 'react'
import { ThemeProvider } from '@material-ui/core'
import { LayoutProps } from '../../common/interfaces/props'
import Topbar from '../Topbar'
import theme from '../../utils/theme'
import MenuBottom from '../MenuBottom/index'

const Layout: React.FC<LayoutProps> = props => {
  return (
    <ThemeProvider theme={theme}>
      <Topbar />
      {props.children}
      <MenuBottom />
    </ThemeProvider>
  )
}

export default Layout
