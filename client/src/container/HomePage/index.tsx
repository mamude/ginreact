import React from 'react'
import { Card, CardActionArea, Grid } from '@material-ui/core'
import { gql, useQuery } from '@apollo/client'
import { CustomerConsumer } from '../../context/Customer'
import { Market } from '../../common/interfaces/models'
import { Image, Separator } from '../../components/Common/style'
import PageWrapper from '../../components/PageWrapper/index'
import PageTitle from '../../components/PageTitle'
import image from '../../images/image.png'
import { MarketContent, MarketInfo, MarketRating, StarRate } from './style'

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

const HomePage: React.FC = () => {
  const { loading, error, data } = useQuery(MARKETS_QUERY)
  if (loading) return <div>Carregando...</div>
  if (error) return <div>{error.message}</div>
  return (
    <CustomerConsumer>
      <PageWrapper>
        <PageTitle>Lojas</PageTitle>
        <Grid container spacing={4}>
          {data.markets.map((market: Market) => (
            <Grid key={market.id} item xs={12} md={6}>
              <CardActionArea component="a" href={`market/${market.id}`}>
                <Card style={{ display: 'flex' }}>
                  <Image image={image} title={market.name} />
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
                      {`${market.distance.toFixed(1)} km`}
                    </MarketInfo>
                    <MarketInfo>
                      {market.deliveryTime}
                      <Separator>•</Separator>
                      {`R$ ${market.deliveryTax.toFixed(2)}`}
                    </MarketInfo>
                  </MarketContent>
                </Card>
              </CardActionArea>
            </Grid>
          ))}
        </Grid>
      </PageWrapper>
    </CustomerConsumer>
  )
}

export default HomePage
