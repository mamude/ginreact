export type CustomerType = {
  id: number
  token: string
  email: string
  user: string
  sessionId: string
}

export type InitialStateType = {
  customer: CustomerType
}

export type CustomerAction = {
  type: 'LOGIN_SUCCESS'
  payload: {
    id: number
    token: string
    email: string
    user: string
    sessionId: string
  }
}
