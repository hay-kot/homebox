/**
 * OutType is the base type that is returned from the API.
 * In contains the common fields that are included with every
 * API response that isn't a bulk result
 */
export type OutType = {
  id: string;
  createdAt: string;
  updatedAt: string;
};

export type Details = {
  name: string;
  description: string;
};

export type Results<T> = {
  items: T[];
};
