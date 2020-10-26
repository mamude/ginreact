import React, { useState } from 'react'
import { Container, TextField, Typography } from '@material-ui/core'
import { useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import { CreateCustomerInput } from '../../common/interfaces/inputs'
import { createCustomerSchema } from '../../common/validations/customer'
import PageWrapper from '../../components/PageWrapper'
import SubmitButton from '../../components/SubmitButton'

const CreateAccountPage: React.FC = () => {
  const [error, setError] = useState('')
  const [open, setOpen] = useState(false)
  const [loading, setLoading] = useState(false)
  const { register, handleSubmit, errors } = useForm<CreateCustomerInput>({
    resolver: yupResolver(createCustomerSchema),
  })

  const onSubmit = async (data: CreateCustomerInput) => {
    console.log(data)
  }

  return (
    <Container component="main" maxWidth="xs">
      <PageWrapper>
        <Typography variant="h6">Criar conta</Typography>
        <form onSubmit={handleSubmit(onSubmit)}>
          <TextField
            required
            inputRef={register}
            name="firstName"
            label="Nome"
            margin="normal"
            fullWidth
            helperText={errors.firstName?.message}
          />
          <TextField
            required
            inputRef={register}
            name="lastName"
            label="Sobrenome"
            margin="normal"
            fullWidth
            helperText={errors.lastName?.message}
          />
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
            name="phone"
            label="Telefone"
            margin="normal"
            fullWidth
            helperText={errors.phone?.message}
          />
          <TextField
            required
            inputRef={register}
            type="password"
            name="password"
            label="Senha"
            margin="normal"
            fullWidth
            helperText={errors.password?.message}
          />
          <TextField
            required
            inputRef={register}
            type="password"
            name="password_confirm"
            label="Confirmar senha"
            margin="normal"
            fullWidth
            helperText={errors.passwordConfirm?.message}
          />
          <SubmitButton name="Confirmar" loading={loading} />
        </form>
      </PageWrapper>
    </Container>
  )
}

export default CreateAccountPage
