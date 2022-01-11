export interface BasicPageParams {
  page: number;
  page_ize: number;
  order_field: string;
  order_desc: boolean;
}

export interface BasicFetchResult<T> {
  list: T[];
  total: number;
}
