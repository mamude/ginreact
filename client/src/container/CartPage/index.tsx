import React from 'react'
import { gql, useMutation, useQuery } from '@apollo/client'
import {
  Button,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@material-ui/core'
import { Link, useHistory } from 'react-router-dom'
import { BackHome } from '../../components/Common/style'
import PageTitle from '../../components/PageTitle'
import PageWrapper from '../../components/PageWrapper'
import { ButtonPlus, ButtonRemove } from '../../components/Button'
import { CustomerConsumer } from '../../context/Customer'
import { getCustomer } from '../../utils/security'
import { ShoppingCart } from '../../common/interfaces/models'
import { BoldText, Cell } from './style'

const SHOPPING_CART_LIST = gql`
  query ListShoppingCart($sessionId: String!, $customerId: Int!) {
    shoppingCart(sessionId: $sessionId, customerId: $customerId) {
      id
      amount
      price
      product {
        id
        name
        price
      }
    }
    shoppingCartSum(sessionId: $sessionId, customerId: $customerId) {
      amount
      tax
      subtotal
      total
      marketId
      market
    }
  }
`
const SHOPPING_CARD_UPDATE = gql`
  mutation UpdateShoppingCart($cartId: Int!, $productId: Int!) {
    updateItemCart(cartId: $cartId, productId: $productId) {
      id
      price
      amount
    }
  }
`

const SHOPPING_CARD_REMOVE_ITEM = gql`
  mutation RemoveItemShoppingCart($cartId: Int!, $productId: Int!) {
    removeItemCart(cartId: $cartId, productId: $productId) {
      id
      price
      amount
    }
  }
`

const CartPage: React.FC = () => {
  const customer = getCustomer()
  const history = useHistory()
  const { loading, error, data } = useQuery(SHOPPING_CART_LIST, {
    variables: { sessionId: customer.sessionId, customerId: customer.id },
  })
  const [updateShoppingCart] = useMutation(SHOPPING_CARD_UPDATE)
  const [removeItemShoppingCart] = useMutation(SHOPPING_CARD_REMOVE_ITEM)

  if (loading) return <div>Carregando...</div>
  if (error) return <div>{error.message}</div>

  const addIem = (cart: ShoppingCart) => {
    updateShoppingCart({
      variables: {
        cartId: cart.id,
        productId: cart.product.id,
      },
    })
  }
  const removeItem = (cart: ShoppingCart) => {
    removeItemShoppingCart({
      variables: {
        cartId: cart.id,
        productId: cart.product.id,
      },
    })
  }

  const backToMarket = () => {
    history.push(`/market/${data.shoppingCartSum.marketId}`)
  }

  return (
    <CustomerConsumer>
      <PageWrapper>
        <Link to="/">
          <BackHome />
        </Link>
        <PageTitle>Sacola</PageTitle>
        <TableContainer component={Paper}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>
                  <BoldText variant="h6">
                    {data.shoppingCartSum.market}
                  </BoldText>
                </TableCell>
                <TableCell />
                <TableCell />
              </TableRow>
            </TableHead>
            <TableBody>
              {data.shoppingCart.map((cart: ShoppingCart) => (
                <TableRow key={cart.id}>
                  <Cell>{`${cart.amount}x ${cart.product.name}`}</Cell>
                  <Cell align="right">
                    <ButtonPlus onClick={() => addIem(cart)} />
                    <ButtonRemove onClick={() => removeItem(cart)} />
                  </Cell>
                  <Cell align="right">{`R$ ${cart.price.toFixed(2)}`}</Cell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
        <Button variant="text" color="secondary" onClick={backToMarket}>
          Adicionar mais items
        </Button>
        <TableContainer>
          <Table>
            <TableBody>
              <TableRow>
                <Cell>SubTotal:</Cell>
                <Cell align="right">{`R$ ${data.shoppingCartSum.subtotal}`}</Cell>
              </TableRow>
              <TableRow>
                <Cell>Taxa de Entrega:</Cell>
                <Cell align="right">{`R$ ${data.shoppingCartSum.tax}`}</Cell>
              </TableRow>
              <TableRow>
                <Cell>
                  <BoldText variant="h6">Total:</BoldText>
                </Cell>
                <Cell align="right">
                  <BoldText variant="h6">{`R$ ${data.shoppingCartSum.total}`}</BoldText>
                </Cell>
              </TableRow>
            </TableBody>
          </Table>
        </TableContainer>
      </PageWrapper>
    </CustomerConsumer>
  )
}
export default CartPage
