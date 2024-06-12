export default interface Pagination<T, D=any> {
  data: Array<T>;
  filter: D;
  paging: {
    page: number;
    limit: number;
    total: number;
    next_cursor: string;
  };
}