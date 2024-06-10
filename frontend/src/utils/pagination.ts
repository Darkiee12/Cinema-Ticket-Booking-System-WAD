export default interface Pagination<T, D=string> {
  data: Array<T>;
  filter: D;
  paging: string;
}