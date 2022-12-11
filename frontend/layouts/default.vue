<template>
  <div>
    <AppToast />
    <div class="drawer drawer-mobile">
      <input id="my-drawer-2" type="checkbox" class="drawer-toggle" />
      <div class="drawer-content justify-center">
        <AppHeaderDecor class="-mt-10" />
        <slot></slot>

        <!-- Button -->
        <label for="my-drawer-2" class="btn btn-primary drawer-button lg:hidden">Open drawer</label>
      </div>

      <!-- Sidebar -->
      <div class="drawer-side shadow-lg w-60 flex flex-col justify-center py-10" style="background: white">
        <label for="my-drawer-2" class="drawer-overlay"></label>
        <!-- Top Section -->
        <div class="space-y-8">
          <div class="flex flex-col items-center gap-4">
            <p>Kotelman House</p>
            <NuxtLink class="avatar placeholder" to="/home">
              <div class="bg-base-100 text-neutral-content rounded-full w-36">
                <span class="text-6xl text-base-content">HK</span>
              </div>
            </NuxtLink>
          </div>
          <div class="flex flex-col">
            <div class="mx-auto w-40 mb-6">
              <BaseButton class="btn-block btn-primary text-xl">
                <template #icon>
                  <Icon name="mdi-plus-circle" class="h-6 w-6" />
                </template>
                Create
              </BaseButton>
            </div>
            <ul class="flex flex-col mx-auto gap-y-8">
              <li v-for="n in nav" :key="n.id" class="text-xl">
                <NuxtLink v-if="n.to" :to="n.to">
                  <span class="mr-4">
                    <Icon :name="n.icon" class="h-5 w-5" />
                  </span>
                  {{ n.name }}
                </NuxtLink>
                <button v-else @click="n.action">
                  <span class="mr-4">
                    <Icon :name="n.icon" class="h-5 w-5" />
                  </span>
                  {{ n.name }}
                </button>
              </li>
            </ul>
          </div>
        </div>


        <!-- Bottom -->
        <button class="mt-auto mb-6">Sign Out</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { useLabelStore } from "~~/stores/labels";
  import { useLocationStore } from "~~/stores/locations";

  /**
   * Store Provider Initialization
   */

  const nav = [
    {
      icon: "mdi-account",
      id: 1,
      name: "Profile",
      to: "/profile",
    },
    {
      icon: "mdi-document",
      id: 3,
      name: "Items",
      to: "/items",
    },
    {
      icon: "mdi-database",
      id: 2,
      name: "Import",
      action: () => {},
    },
    {
      icon: "mdi-database-export",
      id: 5,
      name: "Export",
      action: () => {},
    },
  ];

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
        locationStore.refreshChildren();
        locationStore.refreshParents();
      }
      console.debug("locationStore handler called by observer");
    },
  });

  const eventBus = useEventBus();
  eventBus.on(
    EventTypes.ClearStores,
    () => {
      labelStore.refresh();
      locationStore.refreshChildren();
      locationStore.refreshParents();
    },
    "stores"
  );

  onUnmounted(() => {
    rmLabelStoreObserver();
    rmLocationStoreObserver();
    eventBus.off(EventTypes.ClearStores, "stores");
  });
</script>
