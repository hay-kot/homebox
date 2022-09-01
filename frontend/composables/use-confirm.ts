import { UseConfirmDialogReturn } from '@vueuse/core';
import { Ref } from 'vue';

type Store = UseConfirmDialogReturn<any, Boolean, Boolean> & {
  text: Ref<string>;
  setup: boolean;
};

const store: Partial<Store> = {
  text: ref('Are you sure you want to delete this item? '),
  setup: false,
};

/**
 * This function is used to wrap the ModalConfirmation which is a "Singleton" component
 * that is used to confirm actions. It's mounded once on the root of the page and reused
 * for every confirmation action that is required.
 *
 * This is in an experimental phase of development and may have unknown or unexpected side effects.
 */
export function useConfirm(): Store {
  if (!store.setup) {
    store.setup = true;
    const { isRevealed, reveal, confirm, cancel } = useConfirmDialog<any, Boolean, Boolean>();
    store.isRevealed = isRevealed;
    store.reveal = reveal;
    store.confirm = confirm;
    store.cancel = cancel;
  }

  async function openDialog(msg: string) {
    store.text.value = msg;
    return await store.reveal();
  }

  return {
    ...(store as Store),
    reveal: openDialog,
  };
}
