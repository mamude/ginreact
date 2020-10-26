import React, { useContext, useState } from 'react'
import {
  Button,
  Container,
  CircularProgress,
  Grid,
  TextField,
  Typography,
  Snackbar,
} from '@material-ui/core'
import { useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import { Alert } from '@material-ui/lab'
import { LoginFormInput } from '../../common/interfaces/inputs'
import { loginCustomerSchema } from '../../common/validations/customer'
import { CustomerContext } from '../../context/Customer/index'
import { Wrapper } from './style'
import Api from '../../utils/api'

const LoginPage: React.FC = () => {
  const [error, setError] = useState('')
  const [open, setOpen] = useState(false)
  const [loading, setLoading] = useState(false)
  const { dispatch } = useContext(CustomerContext)
  const { register, handleSubmit, errors } = useForm<LoginFormInput>({
    resolver: yupResolver(loginCustomerSchema),
  })

  const onSubmit = async (data: LoginFormInput) => {
    setLoading(true)
    await Api.post('/customer/login', data)
      .then(response => {
        const payload = response.data
        dispatch({
          type: 'LOGIN_SUCCESS',
          payload,
        })
      })
      .catch(err => {
        setOpen(true)
        setLoading(false)
        setError(err.response.data.error)
      })
  }

  return (
    <Container component="main" maxWidth="xs">
      <Wrapper>
        <Snackbar open={open} onClose={() => setOpen(false)}>
          <Alert severity="error">{error}</Alert>
        </Snackbar>
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
          <Button
            startIcon={loading ? <CircularProgress size={30} /> : null}
            type="submit"
            fullWidth
            color="primary"
            disabled={loading}
          >
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
