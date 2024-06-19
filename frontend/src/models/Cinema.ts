import { Owner } from "./user";

export default interface Cinema {
  id: string;
  name: string;
  address: string;
  created_at: string;
  email: string;
  capacity: number;
  owner: Owner;
  owner_id: Owner['id'];
  status: string;
  phone_number: string;
  banner: string;
}