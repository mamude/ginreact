import React from 'react'
import { Container } from '@material-ui/core'
import { BrowserRouter } from 'react-router-dom'
import Layout from '../../components/Layout'
import Navigation from '../../components/Navigation'

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <Layout>
        <Container fixed>
          <Navigation />
        </Container>
      </Layout>
    </BrowserRouter>
  )
}

export default App
