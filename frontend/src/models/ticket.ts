import User from './user'

export default interface Ticket {
  created_at: string
  id: string
  seat_number: number
  show_id: number
  status: number
  timestamp: string
  updated_at: string
  user: User
};
export interface bookingTicket {
  seat_number: number
  show_id: number
};
