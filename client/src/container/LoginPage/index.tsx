import React, { useContext, useState } from 'react'
import { Container, Grid, Typography, Snackbar } from '@material-ui/core'
import { Field, Form, Formik } from 'formik'
import { TextField } from 'formik-material-ui'
import { Alert } from '@material-ui/lab'
import { LoginFormInput } from '../../common/interfaces/inputs'
import { loginCustomerSchema } from '../../common/validations/customer'
import { CustomerContext } from '../../context/Customer/index'
import { Wrapper } from './style'
import Api from '../../utils/api'
import SubmitButton from '../../components/SubmitButton'

const LoginPage: React.FC = () => {
  const [error, setError] = useState('')
  const [open, setOpen] = useState(false)
  const [loading, setLoading] = useState(false)
  const { dispatch } = useContext(CustomerContext)

  const initialValues: LoginFormInput = {
    email: '',
    password: '',
  }

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
        <Formik
          initialValues={initialValues}
          validationSchema={loginCustomerSchema}
          onSubmit={data => {
            onSubmit(data)
          }}
        >
          <Form>
            <Field component={TextField} name="email" label="E-mail" />
            <Field
              component={TextField}
              type="password"
              name="password"
              label="Senha"
            />
            <SubmitButton name="Entrar" loading={loading} />
            <Grid container>
              <Grid item xs />
              <Grid item>Criar conta</Grid>
            </Grid>
          </Form>
        </Formik>
      </Wrapper>
    </Container>
  )
}

export default LoginPage
