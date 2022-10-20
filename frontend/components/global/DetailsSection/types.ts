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

type LinkDetail = BaseDetail & {
  type: "link";
  text: string;
  href: string;
};

export type CustomDetail = DateDetail | CurrencyDetail | LinkDetail;

export type Detail = BaseDetail & {
  text: StringLike;
  type?: "text";
};

export type Details = Array<Detail | CustomDetail>;
