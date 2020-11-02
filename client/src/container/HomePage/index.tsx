import React from 'react'
import { Card, Grid } from '@material-ui/core'
import { gql, useQuery } from '@apollo/client'
import { CustomerConsumer } from '../../context/Customer'
import PageWrapper from '../../components/PageWrapper/index'
import PageTitle from '../../components/PageTitle'
import image from '../../images/image.png'
import {
  Image,
  MarketContent,
  MarketInfo,
  MarketRating,
  Separator,
  StarRate,
} from './style'

const MARKETS_QUERY = gql`
  query getMarkets {
    markets {
      id
      name
      rating
      deliveryTime
      deliveryTax
      distance
      categoryBusiness {
        name
      }
    }
  }
`

interface Category {
  id: number
  name: string
}

interface Market {
  id: number
  name: string
  rating: number
  deliveryTax: number
  deliveryTime: string
  distance: number
  categoryBusiness: Category
}

const HomePage: React.FC = () => {
  const { loading, error, data } = useQuery(MARKETS_QUERY)
  if (loading) return <div>Loading...</div>
  if (error) return <div>{error.message}</div>
  return (
    <CustomerConsumer>
      <PageWrapper>
        <PageTitle>Lojas</PageTitle>
        <Grid container spacing={4}>
          {data.markets.map((market: Market) => (
            <Grid key={market.id} item xs={12} md={6}>
              <Card style={{ display: 'flex' }}>
                <Image image={image} title="Loja" />
                <MarketContent>
                  <strong>{market.name}</strong>
                  <MarketInfo>
                    <MarketRating>
                      <StarRate />
                      {market.rating.toFixed(1)}
                    </MarketRating>
                    <Separator>•</Separator>
                    {market.categoryBusiness.name}
                    <Separator>•</Separator>
                    {market.distance.toFixed(1)}
                    &nbsp;km
                  </MarketInfo>
                  <MarketInfo>
                    {market.deliveryTime}
                    <Separator>•</Separator>
                    R$&nbsp;
                    {market.deliveryTax.toFixed(2)}
                  </MarketInfo>
                </MarketContent>
              </Card>
            </Grid>
          ))}
        </Grid>
      </PageWrapper>
    </CustomerConsumer>
  )
}

export default HomePage
