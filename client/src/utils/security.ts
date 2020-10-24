import SecureLS from 'secure-ls'
import { CustomerType } from '../common/types/customer'

const ls = new SecureLS({ encodingType: 'aes' })

export function setToken(token: string) {
  ls.set('token', token)
}

export function getToken() {
  return ls.get('token')
}

export function setCustomer(state: CustomerType) {
  ls.set('customer', state)
}

export function getCustomer(): CustomerType {
  return ls.get('customer')
}

export function clearSession() {
  ls.removeAll()
}
