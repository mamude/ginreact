import React, { useContext, useState } from 'react'
import { Snackbar } from '@material-ui/core'
import { Field, Form, Formik } from 'formik'
import { TextField } from 'formik-material-ui'
import { Alert } from '@material-ui/lab'
import { useHistory } from 'react-router-dom'
import { CreateCustomerInput } from '../../common/interfaces/inputs'
import { createCustomerSchema } from '../../common/validations/customer'
import { CustomerContext } from '../../context/Customer'
import PageWrapper from '../../components/PageWrapper'
import PageTitle from '../../components/PageTitle'
import SubmitButton from '../../components/SubmitButton'
import TextFieldPhoneMask from '../../components/TextFieldPhoneMask/index'
import Api from '../../utils/api'

const CreateAccountPage: React.FC = () => {
  const [error, setError] = useState('')
  const [open, setOpen] = useState(false)
  const [loading, setLoading] = useState(false)
  const { dispatch } = useContext(CustomerContext)
  const history = useHistory()

  const initialValues: CreateCustomerInput = {
    firstName: '',
    lastName: '',
    phone: '',
    email: '',
    password: '',
    passwordConfirm: '',
  }

  const onSubmit = async (data: CreateCustomerInput) => {
    setLoading(true)
    await Api.post('/customer/create', data)
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
      <PageTitle>Criar conta</PageTitle>
      <Formik
        initialValues={initialValues}
        validationSchema={createCustomerSchema}
        onSubmit={data => {
          onSubmit(data)
        }}
      >
        <Form>
          <Field component={TextField} name="firstName" label="Nome" />
          <Field component={TextField} name="lastName" label="Sobrenome" />
          <Field component={TextField} name="email" label="E-mail" />
          <Field component={TextFieldPhoneMask} name="phone" label="Celular" />
          <Field
            component={TextField}
            type="password"
            name="password"
            label="Senha"
          />
          <Field
            component={TextField}
            type="password"
            name="passwordConfirm"
            label="Confirmar Senha"
          />
          <SubmitButton name="Enviar" loading={loading} />
        </Form>
      </Formik>
    </PageWrapper>
  )
}

export default CreateAccountPage
