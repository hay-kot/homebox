import type { Ref } from "vue";

type TreeState = Record<string, boolean>;

const store: Record<string, Ref<TreeState>> = {};

export function newTreeKey(): string {
  return Math.random().toString(36).substring(2);
}

export function useTreeState(key: string): Ref<TreeState> {
  if (!store[key]) {
    store[key] = ref({});
  }

  return store[key];
}
