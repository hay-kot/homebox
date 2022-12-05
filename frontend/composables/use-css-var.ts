import { ComputedRef } from "vue";

type ColorType = "hsla";

export type VarOptions = {
  type: ColorType;
  transparency?: number;
  apply?: (value: string) => string;
};

export function useCssVar(name: string, options?: VarOptions): ComputedRef<string> {
  if (!options) {
    options = {
      type: "hsla",
      transparency: 1,
      apply: null,
    };
  }

  switch (options.type) {
    case "hsla": {
      return computed(() => {
        if (!document) {
          return "";
        }

        let val = getComputedStyle(document.documentElement).getPropertyValue(name);
        val = val.trim().split(" ").join(", ");

        if (options.transparency) {
          val += `, ${options.transparency}`;
        }

        return `hsla(${val})`;
      });
    }
  }
}
