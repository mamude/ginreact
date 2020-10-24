import React from 'react'
import { Container, Typography } from '@material-ui/core'
import { Wrapper } from '../../components/Layout/style'

const LoginPage: React.FC = () => {
  return (
    <Container component="main">
      <Wrapper>
        <Typography variant="h5">Autenticação</Typography>
      </Wrapper>
    </Container>
  )
}

export default LoginPage
