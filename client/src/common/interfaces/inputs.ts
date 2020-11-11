export interface LoginFormInput {
  email: string
  password: string
  sessionId: string
}

export interface CreateCustomerInput {
  firstName: string
  lastName: string
  email: string
  phone: string
  password: string
  passwordConfirm: string
}
