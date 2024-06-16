export default interface User {
  id: string
  status: number
  created_at: Date
  date_of_birth: string
  email: string
  gender: string
  name: string
  role: string
  phone: string
}

export interface Credential {
  email: string
  password: string
}

export interface Register {
  email: string
  name: string
  password: string
}

export interface Account {
  created_at: string
  expiry: number
  token: string
}

export interface Owner extends User {
  tier: string
}
