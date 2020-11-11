import { CustomerAction, CustomerType } from '../../common/types/customer'
import { setCustomer } from '../../utils/security'

export function customerReducer(
  state: CustomerType,
  action: CustomerAction,
): CustomerType {
  switch (action.type) {
    case 'LOGIN_SUCCESS':
      setCustomer(action.payload)
      return {
        id: action.payload.id,
        token: action.payload.token,
        email: action.payload.email,
        user: action.payload.user,
        sessionId: '',
      }
    default:
      return state
  }
}
