<template>
  <div>
    <AppToast />
    <AppHeader />
    <main>
      <slot></slot>
    </main>
  </div>
</template>

<script lang="ts" setup>
  import { useItemStore } from "~~/stores/items";
  import { useLabelStore } from "~~/stores/labels";
  import { useLocationStore } from "~~/stores/locations";
  /**
   * Store Provider Initialization
   */

  const labelStore = useLabelStore();
  const reLabel = /\/api\/v1\/labels\/.*/gm;
  const rmLabelStoreObserver = defineObserver("labelStore", {
    handler: r => {
      if (r.status === 201 || r.url.match(reLabel)) {
        labelStore.refresh();
      }
      console.debug("labelStore handler called by observer");
    },
  });

  const locationStore = useLocationStore();
  const reLocation = /\/api\/v1\/locations\/.*/gm;
  const rmLocationStoreObserver = defineObserver("locationStore", {
    handler: r => {
      if (r.status === 201 || r.url.match(reLocation)) {
        locationStore.refresh();
      }
      console.debug("locationStore handler called by observer");
    },
  });

  const itemStore = useItemStore();
  const reItem = /\/api\/v1\/items\/.*/gm;
  const rmItemStoreObserver = defineObserver("itemStore", {
    handler: r => {
      if (r.status === 201 || r.url.match(reItem)) {
        itemStore.refresh();
      }
      console.debug("itemStore handler called by observer");
    },
  });

  const eventBus = useEventBus();
  eventBus.on(
    EventTypes.ClearStores,
    () => {
      labelStore.refresh();
      itemStore.refresh();
      locationStore.refresh();
    },
    "stores"
  );

  onUnmounted(() => {
    rmLabelStoreObserver();
    rmLocationStoreObserver();
    rmItemStoreObserver();
    eventBus.off(EventTypes.ClearStores, "stores");
  });
</script>
