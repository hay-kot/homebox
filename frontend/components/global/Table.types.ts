export type TableHeader = {
  text: string;
  value: string;
  sortable?: boolean;
  align?: "left" | "center" | "right";
};

export type TableData = Record<string, any>;
