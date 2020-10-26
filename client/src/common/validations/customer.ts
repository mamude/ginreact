import * as yup from 'yup'

export const loginCustomerSchema = yup.object().shape({
  email: yup
    .string()
    .email('E-mail inválido!')
    .required('E-mail é um campo obrigatório!'),
  password: yup
    .string()
    .required('Senha é um campo obrigatório!')
    .min(6, 'Senha deve conter pelo menos 6 dígitos'),
})

export const createCustomerSchema = yup.object().shape({
  firstName: yup.string().required('Nome é um campo obrigatório!'),
  lastName: yup.string().required('Nome é um campo obrigatório!'),
  email: yup
    .string()
    .email('E-mail inválido!')
    .required('E-mail é um campo obrigatório!'),
  phone: yup.string().required('Telefone é um campo obrigatório!'),
  password: yup
    .string()
    .required('Senha é um campo obrigatório!')
    .min(6, 'Senha deve conter pelo menos 6 dígitos'),
  passwordConfirm: yup
    .string()
    .required('Confirmar senha é um campo obrigatório!')
    .when('password', {
      is: (password: string) => !!(password && password.length > 0),
      then: yup.string().oneOf([yup.ref('password')], 'Senhas não confirmam!'),
    }),
})
