/* eslint no-redeclare: 0 */
import { WritableComputedRef } from "vue";

export function useRouteQuery(q: string, def: string[]): WritableComputedRef<string[]>;
export function useRouteQuery(q: string, def: string): WritableComputedRef<string>;
export function useRouteQuery(q: string, def: boolean): WritableComputedRef<boolean>;
export function useRouteQuery(q: string, def: number): WritableComputedRef<number>;
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function useRouteQuery(q: string, def: any): WritableComputedRef<any> {
  const route = useRoute();
  const router = useRouter();

  switch (typeof def) {
    case "string":
      if (route.query[q] === undefined) {
        router.push({ query: { ...route.query, [q]: def } });
      }

      return computed({
        get: () => {
          const qv = route.query[q];
          if (Array.isArray(qv)) {
            return qv[0];
          }
          return qv;
        },
        set: v => {
          const query = { ...route.query, [q]: v };
          router.push({ query });
        },
      });
    case "object": // array
      return computed({
        get: () => {
          const qv = route.query[q];
          if (Array.isArray(qv)) {
            return qv;
          }
          return [qv];
        },
        set: v => {
          const query = { ...route.query, [q]: v };
          router.push({ query });
        },
      });
    case "boolean":
      return computed({
        get: () => {
          const qv = route.query[q];
          if (Array.isArray(qv)) {
            return qv[0] === "true";
          }
          return qv === "true";
        },
        set: v => {
          const query = { ...route.query, [q]: `${v}` };
          router.push({ query });
        },
      });
    case "number":
      return computed({
        get: () => {
          const qv = route.query[q];
          if (Array.isArray(qv) && qv.length > 0) {
            return parseInt(qv[0] as string, 10);
          }
          return parseInt(qv as string, 10) || def;
        },
        set: v => {
          const query = { ...route.query, [q]: `${v}` };
          router.push({ query });
        },
      });
  }

  throw new Error("Invalid type");
}
