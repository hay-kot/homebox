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

  const eventBus = useEventBus();
  eventBus.on(
    EventTypes.ClearStores,
    () => {
      labelStore.refresh();
      locationStore.refresh();
    },
    "stores"
  );

  onUnmounted(() => {
    rmLabelStoreObserver();
    rmLocationStoreObserver();
    eventBus.off(EventTypes.ClearStores, "stores");
  });
</script>
