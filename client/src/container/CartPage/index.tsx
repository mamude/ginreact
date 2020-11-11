import React from 'react'
import { gql, useQuery } from '@apollo/client'
import { Grid } from '@material-ui/core'
import { Link } from 'react-router-dom'
import { BackHome } from '../../components/Common/style'
import PageTitle from '../../components/PageTitle'
import PageWrapper from '../../components/PageWrapper'
import { CustomerConsumer } from '../../context/Customer'
import { getCustomer } from '../../utils/security'
import { ShoppingCart } from '../../common/interfaces/models'

const SHOPPING_CART_LIST = gql`
  query ListShoppingCart($sessionId: String!, $customerId: Int!) {
    shoppingCart(sessionId: $sessionId, customerId: $customerId) {
      id
      amount
      price
      market {
        name
      }
      product {
        name
        price
      }
    }
    shoppingCartSum(sessionId: $sessionId, customerId: $customerId) {
      amount
      total
    }
  }
`

const CartPage: React.FC = () => {
  const customer = getCustomer()
  const { loading, error, data } = useQuery(SHOPPING_CART_LIST, {
    variables: { sessionId: customer.sessionId, customerId: customer.id },
  })

  if (loading) return <div>Carregando...</div>
  if (error) return <div>{error.message}</div>

  return (
    <CustomerConsumer>
      <PageWrapper>
        <Link to="/">
          <BackHome />
        </Link>
        <PageTitle>Sacola</PageTitle>
        <Grid container spacing={4}>
          {data.shoppingCart.map((row: ShoppingCart) => (
            <Grid item key={row.id} xs={12}>
              {row.amount}
              {row.price}
            </Grid>
          ))}
        </Grid>
        <Grid container spacing={4}>
          <Grid item xs={12}>
            {data.shoppingCartSum.amount}
            <br />
            {data.shoppingCartSum.total.toFixed(2)}
          </Grid>
        </Grid>
      </PageWrapper>
    </CustomerConsumer>
  )
}
export default CartPage
