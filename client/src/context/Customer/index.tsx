import React, { createContext, useEffect, useReducer } from 'react'
import { LayoutProps } from '../../common/interfaces/props'
import { CustomerType } from '../../common/types/customer'
import { getCustomer, setCustomer } from '../../utils/security'
import { customerReducer } from './reducer'

const initialState = {
  id: 0,
  token: '',
  email: '',
  user: '',
}

const CustomerContext = createContext<{
  state: CustomerType
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  dispatch: React.Dispatch<any>
}>({
  state: initialState,
  dispatch: () => null,
})

const CustomerProvider = ({ children }: LayoutProps) => {
  const localState = getCustomer() || ''
  const [state, dispatch] = useReducer(
    customerReducer,
    localState || initialState,
  )

  useEffect(() => {
    setCustomer(state)
  }, [state])

  return (
    <CustomerContext.Provider value={{ state, dispatch }}>
      {children}
    </CustomerContext.Provider>
  )
}

const CustomerConsumer = ({ children }: LayoutProps) => {
  return (
    <CustomerContext.Consumer>
      {context => {
        return children
      }}
    </CustomerContext.Consumer>
  )
}

export { CustomerContext, CustomerProvider, CustomerConsumer }
