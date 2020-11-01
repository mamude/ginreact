import React, { useState } from 'react'
import { Container, Typography } from '@material-ui/core'
import { Field, Form, Formik } from 'formik'
import { TextField } from 'formik-material-ui'
import PageWrapper from '../../components/PageWrapper'
import { CreateCustomerInput } from '../../common/interfaces/inputs'
import { createCustomerSchema } from '../../common/validations/customer'
import SubmitButton from '../../components/SubmitButton'
import TextFieldPhoneMask from '../../components/TextFieldPhoneMask/index'

const CreateAccountPage: React.FC = () => {
  const [loading, setLoading] = useState(false)

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
  }

  return (
    <Container component="main" maxWidth="xs">
      <PageWrapper>
        <Typography variant="h6">Criar conta</Typography>
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
            <Field
              component={TextFieldPhoneMask}
              name="phone"
              label="Celular"
            />
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
    </Container>
  )
}

export default CreateAccountPage
