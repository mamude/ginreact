import React from 'react'
import 'fontsource-roboto'
import ReactDOM from 'react-dom'
import { ApolloProvider } from '@apollo/client'
import { CssBaseline } from '@material-ui/core'
import App from './container/App'
import Graphql from './utils/graphql'

ReactDOM.render(
  <ApolloProvider client={Graphql}>
    <CssBaseline />
    <App />
  </ApolloProvider>,
  document.getElementById('root'),
)
