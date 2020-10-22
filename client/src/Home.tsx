import React from 'react'
import {
  Container,
  Grid,
  Header,
  Icon,
  Image,
  Rating,
  Segment,
} from 'semantic-ui-react'
import styled from 'styled-components'
import { gql, useQuery } from '@apollo/client'
import logo from './images/logo.png'
import shop from './images/image.png'

const GridColumn = styled(Grid.Column)`
  height: 70px;
`

const Buttons = styled.div`
  margin: 20px;
`

const Logo = styled.img`
  object-fit: cover;
  width: 120px;
  height: 70px;
  margin-top: 5px;
`

const ShopLogo = styled(Image)`
  width: 150px !important;
`

const ShopName = styled.div`
  font-size: 1.1rem;
  font-weight: bolder;
`

const ShopDetail = styled.div`
  margin-top: 5px;
  font-size: 0.8rem;
  color: #a5a1a1;
`
const RatingPointuation = styled.strong`
  color: orange;
`

const MARKETS_QUERY = gql`
  query getMarkets {
    markets {
      id
      name
      rating
      deliveryTime
      deliveryTax
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
  categoryBusiness: Category
}

function LoadMarkets(): JSX.Element {
  const { loading, error, data } = useQuery(MARKETS_QUERY)
  if (loading) return <div>Loading...</div>
  if (error) return <div>{error.message}</div>

  return (
    <>
      <Header as="h2">Lojas</Header>
      <Grid columns={2} container stackable>
        {data.markets.map((market: Market) => (
          <Grid.Column key={market.id}>
            <Segment>
              <Grid>
                <Grid.Column width={4}>
                  <ShopLogo src={shop} />
                </Grid.Column>
                <Grid.Column width={12}>
                  <ShopName>{market.name}</ShopName>
                  <ShopDetail>
                    <Rating icon="star" defaultRating={5} />
                    <RatingPointuation>{market.rating}</RatingPointuation>
                    <span>{market.categoryBusiness.name}</span>
                  </ShopDetail>
                  <ShopDetail>{market.deliveryTax}</ShopDetail>
                </Grid.Column>
              </Grid>
            </Segment>
          </Grid.Column>
        ))}
      </Grid>
    </>
  )
}

const HomePage: React.FC = () => {
  return (
    <Container>
      <Grid columns={2}>
        <Grid.Row color="red">
          <GridColumn textAlign="left">
            <Logo src={logo} alt="Samir Mamude" />
          </GridColumn>
          <GridColumn textAlign="right">
            <Buttons>
              <Icon name="shop" size="large" />
              <Icon name="user" size="large" />
            </Buttons>
          </GridColumn>
        </Grid.Row>
      </Grid>
      <LoadMarkets />
    </Container>
  )
}

export default HomePage
