import styled from 'styled-components'
import { CardMedia, Divider } from '@material-ui/core'

export const Separator = styled.span`
  margin-left: 5px;
  margin-right: 5px;
`

export const DividerHr = styled(Divider)`
  margin-top: 20px;
  margin-bottom: 20px;
`

export const Image = styled(CardMedia)`
  height: 100px;
  width: 100px;
`

export const ImageProduct = styled(CardMedia)`
  width: 60%;
`
