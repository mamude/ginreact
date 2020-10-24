import React from 'react'
import { ThemeProvider } from '@material-ui/core'
import { Content, ContentContainer, Root, Wrapper } from './style'
import { IProps } from '../../common/interfaces/props'
import Topbar from '../Topbar'
import theme from '../../utils/theme'
import MenuBottom from '../MenuBottom/index'

const Layout: React.FC<IProps> = props => {
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
