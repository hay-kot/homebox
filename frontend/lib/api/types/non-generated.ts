export enum AttachmentTypes {
  Photo = "photo",
  Manual = "manual",
  Warranty = "warranty",
  Attachment = "attachment",
}

export type Result<T> = {
  item: T;
};

export type Results<T> = {
  items: T[];
};
