import type { WritableComputedRef } from "vue";

export function useMinLoader(ms = 500): WritableComputedRef<boolean> {
  const loading = ref(false);

  const locked = ref(false);

  const minLoading = computed({
    get: () => loading.value,
    set: value => {
      if (value) {
        loading.value = true;

        if (!locked.value) {
          locked.value = true;
          setTimeout(() => {
            locked.value = false;
          }, ms);
        }
      }

      if (!value && !locked.value) {
        loading.value = false;
      } else if (!value && locked.value) {
        setTimeout(() => {
          loading.value = false;
        }, ms);
      }
    },
  });
  return minLoading;
}
