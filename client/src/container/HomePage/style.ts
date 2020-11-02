import styled from 'styled-components'
import { CardContent, CardMedia } from '@material-ui/core'
import StarIcon from '@material-ui/icons/Star'

export const Image = styled(CardMedia)`
  height: 100px;
  width: 100px;
`

export const MarketContent = styled(CardContent)`
  padding: 12px;
  padding-bottom: 12px !important;
  font-size: 0.85rem;
`

export const MarketInfo = styled.div`
  display: flex;
  width: 100%;
  margin-top: 8px;
`

export const MarketRating = styled.span`
  font-size: 0.75rem;
  font-weight: bold;
  display: flex;
  align-items: center;
  line-height: 9px;
  color: #e7a74e;
`

export const StarRate = styled(StarIcon)`
  color: #e7a74e;
  font-size: 17px;
`

export const Separator = styled.span`
  margin-left: 5px;
  margin-right: 5px;
`
