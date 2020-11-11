import React from 'react'
import { Link, useParams } from 'react-router-dom'
import { gql, useQuery } from '@apollo/client'
import { Box, Card, CardActionArea, Grid, Typography } from '@material-ui/core'
import { Product } from '../../common/interfaces/models'
import { getCustomer } from '../../utils/security'
import {
  DividerHr,
  Separator,
  ImageProduct,
} from '../../components/Common/style'
import { CustomerConsumer } from '../../context/Customer'
import {
  Deliver,
  MarketContent,
  MarketInfo,
  MarketRating,
  ProductName,
  ProductDescription,
  ProductPrice,
  StarRate,
  BackHome,
} from './style'
import PageWrapper from '../../components/PageWrapper'
import image from '../../images/image.png'

const MARKETS_QUERY = gql`
  query getMarket($id: Int!) {
    market(id: $id) {
      id
      name
      rating
      deliveryTime
      deliveryTax
      distance
      categoryBusiness {
        name
      }
      products {
        id
        name
        description
        price
      }
    }
  }
`

const MarketPage: React.FC = () => {
  const customer = getCustomer()
  const { id } = useParams<{ id: string }>()
  const { loading, error, data } = useQuery(MARKETS_QUERY, {
    variables: { id },
  })
  if (loading) return <div>Carregando...</div>
  if (error) return <div>{error.message}</div>

  const addProductToCart = (product: Product) => {
    console.log(product, customer)
  }

  return (
    <CustomerConsumer>
      <PageWrapper>
        <Link to="/">
          <BackHome />
        </Link>
        <MarketRating>
          <StarRate />
          {data.market.rating.toFixed(1)}
        </MarketRating>
        <Typography variant="h5">{data.market.name}</Typography>
        <MarketInfo>
          {data.market.categoryBusiness.name}
          <Separator>•</Separator>
          {data.market.deliveryTime}
          <Separator>•</Separator>
          {data.market.distance.toFixed(1)}
          &nbsp;km
          <Separator>•</Separator>
        </MarketInfo>
        <MarketInfo>
          <Deliver>
            ENTREGA R$ &nbsp;
            {data.market.deliveryTax.toFixed(2)}
          </Deliver>
        </MarketInfo>
        <DividerHr />
        <Grid container spacing={4}>
          {data.market.products.map((product: Product) => (
            <Grid key={product.id} item xs={12} md={6}>
              <CardActionArea onClick={() => addProductToCart(product)}>
                <Card style={{ display: 'flex' }}>
                  <MarketContent>
                    <ProductName>{product.name}</ProductName>
                    <ProductDescription>
                      {product.description.substring(0, 80)}
                    </ProductDescription>
                    <ProductPrice>
                      R$&nbsp;
                      {product.price}
                    </ProductPrice>
                  </MarketContent>
                  <Box flexGrow={1} />
                  <ImageProduct image={image} title={product.name} />
                </Card>
              </CardActionArea>
            </Grid>
          ))}
        </Grid>
      </PageWrapper>
    </CustomerConsumer>
  )
}

export default MarketPage
