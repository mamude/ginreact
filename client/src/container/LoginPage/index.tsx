import React from 'react'
import * as yup from 'yup'
import {
  Button,
  Container,
  Grid,
  TextField,
  Typography,
} from '@material-ui/core'
import { useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import { Wrapper } from './style'

const schema = yup.object().shape({
  username: yup.string().required(),
  password: yup.string().required().min(6),
})

interface LoginFormInput {
  username: string
  password: string
}

const onSubmit = (data: LoginFormInput) => console.log(data)

const LoginPage: React.FC = () => {
  const { register, handleSubmit, errors } = useForm<LoginFormInput>({
    resolver: yupResolver(schema),
  })

  return (
    <Container component="main" maxWidth="xs">
      <Wrapper>
        <Typography variant="h6">Autenticação</Typography>
        <form autoComplete="off" onSubmit={handleSubmit(onSubmit)}>
          <TextField
            inputRef={register}
            name="username"
            label="E-mail"
            margin="normal"
            fullWidth
          />
          <TextField
            inputRef={register}
            name="password"
            type="password"
            label="Senha"
            margin="normal"
            fullWidth
          />
          <Button type="submit" fullWidth color="primary">
            Entrar
          </Button>
          <Grid container>
            <Grid item xs />
            <Grid item>Criar conta</Grid>
          </Grid>
        </form>
      </Wrapper>
    </Container>
  )
}

export default LoginPage
