import * as yup from 'yup'

export const loginCustomerSchema = yup.object().shape({
  email: yup.string().required('E-mail é um campo obrigatório!'),
  password: yup
    .string()
    .required('Senha é um campo obrigatório!')
    .min(6, 'Senha deve conter pelo menos 6 dígitos'),
})
