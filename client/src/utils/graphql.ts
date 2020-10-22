import { ApolloClient, InMemoryCache } from '@apollo/client'

const Graphql = new ApolloClient({
  uri: '/api/v1/query',
  cache: new InMemoryCache(),
})

export default Graphql
