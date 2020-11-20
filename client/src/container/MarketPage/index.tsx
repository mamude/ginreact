import React from 'react'
import { Link, useHistory, useParams } from 'react-router-dom'
import { gql, useMutation, useQuery } from '@apollo/client'
import { Box, Card, CardActionArea, Grid, Typography } from '@material-ui/core'
import { Product } from '../../common/interfaces/models'
import { getCustomer } from '../../utils/security'
import {
  DividerHr,
  Separator,
  ImageProduct,
  BackHome,
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

const ADD_TO_CART = gql`
  mutation AddToCart(
    $customerId: Int!
    $marketId: Int!
    $productId: Int!
    $amount: Int!
    $sessionId: String!
  ) {
    addToCart(
      customerId: $customerId
      marketId: $marketId
      productId: $productId
      amount: $amount
      sessionId: $sessionId
    ) {
      amount
      price
    }
  }
`

const MarketPage: React.FC = () => {
  const customer = getCustomer()
  const { id } = useParams<{ id: string }>()
  const history = useHistory()
  const { loading, error, data } = useQuery(MARKETS_QUERY, {
    variables: { id },
  })
  const [shoppingCart] = useMutation(ADD_TO_CART)

  if (loading) return <div>Carregando...</div>
  if (error) return <div>{error.message}</div>

  const addProductToCart = (product: Product) => {
    shoppingCart({
      variables: {
        customerId: customer.id,
        marketId: id,
        productId: product.id,
        amount: 1,
        sessionId: customer.sessionId,
      },
    }).finally(() => {
      history.push('/cart')
    })
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
          {`${data.market.distance.toFixed(1)} km`}
          <Separator>•</Separator>
        </MarketInfo>
        <MarketInfo>
          <Deliver>
            {`ENTREGA R$ ${data.market.deliveryTax.toFixed(2)}`}
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
                    <ProductPrice>{`R$ ${product.price}`}</ProductPrice>
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
