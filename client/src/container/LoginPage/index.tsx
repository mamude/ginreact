import React, { useContext, useState } from 'react'
import { useHistory } from 'react-router-dom'
import { Grid, Snackbar } from '@material-ui/core'
import { Field, Form, Formik } from 'formik'
import { TextField } from 'formik-material-ui'
import { Alert } from '@material-ui/lab'
import { LoginFormInput } from '../../common/interfaces/inputs'
import { loginCustomerSchema } from '../../common/validations/customer'
import { CustomerContext } from '../../context/Customer/index'
import PageWrapper from '../../components/PageWrapper'
import PageTitle from '../../components/PageTitle'
import SubmitButton from '../../components/SubmitButton'
import Api from '../../utils/api'

const LoginPage: React.FC = () => {
  const history = useHistory()
  const [error, setError] = useState('')
  const [open, setOpen] = useState(false)
  const [loading, setLoading] = useState(false)
  const { dispatch } = useContext(CustomerContext)

  const initialValues: LoginFormInput = {
    email: '',
    password: '',
    sessionId: '512bf895-48a9-45fc-b03d-054f6b652ad9',
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
        history.push('/')
      })
      .catch(err => {
        setOpen(true)
        setLoading(false)
        setError(err.response.data.error)
      })
  }

  return (
    <PageWrapper>
      <Snackbar open={open} onClose={() => setOpen(false)}>
        <Alert severity="error">{error}</Alert>
      </Snackbar>
      <PageTitle>Autenticação</PageTitle>
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
    </PageWrapper>
  )
}

export default LoginPage
