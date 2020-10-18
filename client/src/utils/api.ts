import axios from 'axios'

const Api = axios.create({
  baseURL: '/api/v1',
  withCredentials: true,
})

export default Api
