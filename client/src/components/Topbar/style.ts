import styled from 'styled-components'
import { TextField, Typography } from '@material-ui/core'

export const Root = styled.div`
  flex-grow: 1;
`

export const TopButtons = styled.div`
  margin-left: 20px;
  width: 300px;
`

export const ShoppingName = styled(Typography)`
  flex-grow: 1;
`

export const ImgLogo = styled.img`
  object-fit: cover;
  margin-top: -60px;
  margin-bottom: -60px;
  margin-right: 45px;
`
