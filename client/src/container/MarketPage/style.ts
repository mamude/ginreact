import styled from 'styled-components'
import StarIcon from '@material-ui/icons/Star'
import { CardContent } from '@material-ui/core'

export const MarketInfo = styled.div`
  display: flex;
  width: 100%;
  color: #c5c5c5;
`

export const MarketContent = styled(CardContent)`
  padding: 12px;
  padding-bottom: 12px !important;
  font-size: 0.85rem;
  width: 100%;
`

export const MarketRating = styled.span`
  font-size: 1rem;
  font-weight: bold;
  display: flex;
  align-items: center;
  line-height: 9px;
  color: #e7a74e;
  justify-content: flex-end;
`

export const Deliver = styled.span`
  border: 1px solid #c5c5c5;
  color: #717171;
  font-weight: bold;
  padding: 4px;
  font-size: 0.75rem;
`

export const StarRate = styled(StarIcon)`
  color: #e7a74e;
`

export const ProductName = styled.span`
  font-size: 14px;
`

export const ProductDescription = styled.span`
  display: flex;
  font-size: 11px;
  width: 100%;
  margin-top: 5px;
  color: #7d7c7c;
`

export const ProductPrice = styled.span`
  font-size: 16px;
  display: flex;
  width: 100%;
  margin-top: 5px;
`
