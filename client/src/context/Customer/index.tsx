import React, { createContext, useReducer } from 'react'
import { IProps } from '../../common/interfaces/props'
import { CustomerType } from '../../common/types/customer'
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

const CustomerProvider = ({ children }: IProps) => {
  const [state, dispatch] = useReducer(customerReducer, initialState)

  return (
    <CustomerContext.Provider value={{ state, dispatch }}>
      {children}
    </CustomerContext.Provider>
  )
}

export { CustomerContext, CustomerProvider }
