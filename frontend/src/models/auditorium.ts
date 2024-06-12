import Cinema from './cinema';
export default interface Auditorium {
  cinema: Cinema;
  created_at: string;
  id: string;
  name: string;
  seats: number;
  status: number;
}