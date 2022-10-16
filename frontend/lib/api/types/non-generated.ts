export enum AttachmentTypes {
  Photo = "photo",
  Manual = "manual",
  Warranty = "warranty",
  Attachment = "attachment",
  Receipt = "receipt",
}

export type Result<T> = {
  item: T;
};

export type Results<T> = {
  items: T[];
};

export interface PaginationResult<T> {
  items: T[];
  page: number;
  pageSize: number;
  total: number;
}
