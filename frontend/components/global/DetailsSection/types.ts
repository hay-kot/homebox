export type StringLike = string | number | boolean;

export type DateDetail = {
  name: string;
  text: string | Date;
  slot?: string;
  type: "date";
};

export type Detail = {
  name: string;
  text: StringLike;
  slot?: string;
  type?: "text";
};
