import React from 'react'
import PageTitle from '../../components/PageTitle'
import PageWrapper from '../../components/PageWrapper'
import { CustomerConsumer } from '../../context/Customer'

const CartPage: React.FC = () => {
  return (
    <CustomerConsumer>
      <PageWrapper>
        <PageTitle>Sacola</PageTitle>
      </PageWrapper>
    </CustomerConsumer>
  )
}
export default CartPage
