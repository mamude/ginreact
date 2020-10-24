import 'fontsource-roboto'
import React from 'react'
import ReactDOM from 'react-dom'
import { ApolloProvider } from '@apollo/client'
import { CssBaseline } from '@material-ui/core'
import { CustomerProvider } from './context/Customer'
import App from './container/App'
import Graphql from './utils/graphql'

ReactDOM.render(
  <ApolloProvider client={Graphql}>
    <CssBaseline />
    <CustomerProvider>
      <App />
    </CustomerProvider>
  </ApolloProvider>,
  document.getElementById('root'),
)
