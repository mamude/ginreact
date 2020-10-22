import React from 'react'
import ReactDOM from 'react-dom'
import { ApolloProvider } from '@apollo/client'
import App from './App'
import Graphql from './utils/graphql'
import 'semantic-ui-css/semantic.min.css'
import './index.css'

ReactDOM.render(
  <React.StrictMode>
    <ApolloProvider client={Graphql}>
      <App />
    </ApolloProvider>
  </React.StrictMode>,
  document.getElementById('root'),
)
