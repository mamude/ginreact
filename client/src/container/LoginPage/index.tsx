import React from 'react';
import { Container, Typography } from '@material-ui/core';
import { Wrapper } from '../../components/Layout/style';

const LoginPage: React.FC = () => {
  return (
    <Container component="main">
      <Wrapper>
        <Typography variant="h3">Autenticação</Typography>
      </Wrapper>
    </Container>
  )
}

export default LoginPage;
