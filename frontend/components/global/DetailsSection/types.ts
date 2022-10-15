export type StringLike = string | number | boolean;

type BaseDetail = {
  name: string;
  slot?: string;
};

type DateDetail = BaseDetail & {
  type: "date";
  text: Date | string;
};

type CurrencyDetail = BaseDetail & {
  type: "currency";
  text: string;
};

export type CustomDetail = DateDetail | CurrencyDetail;

export type Detail = BaseDetail & {
  text: StringLike;
  type?: "text";
};

export type Details = Array<Detail | CustomDetail>;
