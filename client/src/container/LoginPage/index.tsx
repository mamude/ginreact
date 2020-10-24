import React, { useContext } from 'react'
import {
  Button,
  Container,
  Grid,
  TextField,
  Typography,
} from '@material-ui/core'
import { useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import { LoginFormInput } from '../../common/interfaces/inputs'
import { loginCustomerSchema } from '../../common/validations/customer'
import { CustomerContext } from '../../context/Customer/index'
import { Wrapper } from './style'
import Api from '../../utils/api'

const LoginPage: React.FC = () => {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const { state, dispatch } = useContext(CustomerContext)
  const { register, handleSubmit, errors } = useForm<LoginFormInput>({
    resolver: yupResolver(loginCustomerSchema),
  })

  const onSubmit = (data: LoginFormInput) => {
    Api.post('/customer/login', data).then(response => {
      const payload = response.data
      dispatch({
        type: 'LOGIN_SUCCESS',
        payload,
      })
    })
  }

  return (
    <Container component="main" maxWidth="xs">
      <Wrapper>
        <Typography variant="h6">Autenticação</Typography>
        <form autoComplete="off" onSubmit={handleSubmit(onSubmit)}>
          <TextField
            required
            inputRef={register}
            name="email"
            label="E-mail"
            margin="normal"
            fullWidth
            helperText={errors.email?.message}
          />
          <TextField
            required
            inputRef={register}
            name="password"
            type="password"
            label="Senha"
            margin="normal"
            fullWidth
            helperText={errors.password?.message}
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
