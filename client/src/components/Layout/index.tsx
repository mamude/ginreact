import React, { ReactNode } from 'react'
import { ThemeProvider } from '@material-ui/core'
import { Content, ContentContainer, Root, Wrapper } from './style'
import Topbar from '../Topbar'
import theme from '../../utils/theme'
import MenuBottom from '../MenuBottom/index'

interface LayoutProps {
  children: ReactNode
}

const Layout: React.FC<LayoutProps> = props => {
  return (
    <ThemeProvider theme={theme}>
      <Root>
        <Topbar />
        <Wrapper>
          <ContentContainer>
            <Content>{props.children}</Content>
          </ContentContainer>
        </Wrapper>
      </Root>
      <MenuBottom />
    </ThemeProvider>
  )
}

export default Layout
