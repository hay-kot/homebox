<template>
  <div>
    <!--
    Confirmation Modal is a singleton used by all components so we render
    it here to ensure it's always available. Possibly could move this further
    up the tree
   -->
    <ModalConfirm />
    <AppImportDialog v-model="modals.import" />
    <ItemCreateModal v-model="modals.item" />
    <LabelCreateModal v-model="modals.label" />
    <LocationCreateModal v-model="modals.location" />
    <AppToast />
    <div class="drawer drawer-mobile">
      <input id="my-drawer-2" type="checkbox" class="drawer-toggle" />
      <div class="drawer-content justify-center bg-base-300">
        <AppHeaderDecor class="-mt-10" />
        <slot></slot>

        <!-- Button -->
        <label for="my-drawer-2" class="btn btn-primary drawer-button lg:hidden">Open drawer</label>
      </div>

      <!-- Sidebar -->
      <div class="drawer-side overflow-visible shadow-lg w-60 flex flex-col justify-center bg-base-200 py-10">
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
              <div class="dropdown overflow visible w-40">
                <label tabindex="0" class="btn btn-primary btn-block text-lg text-no-transform">
                  <span>
                    <Icon name="mdi-plus" class="mr-1 -ml-1" />
                  </span>
                  Create
                </label>
                <ul tabindex="0" class="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-40">
                  <li v-for="btn in dropdown" :key="btn.name">
                    <button @click="btn.action">
                      {{ btn.name }}
                    </button>
                  </li>
                </ul>
              </div>
            </div>
            <ul class="flex flex-col mx-auto gap-2 w-40 menu">
              <li v-for="n in nav" :key="n.id" class="text-xl">
                <NuxtLink
                  v-if="n.to"
                  class="rounded-btn"
                  :to="n.to"
                  :class="{
                    'bg-secondary text-secondary-content': n.active?.value,
                  }"
                >
                  <Icon :name="n.icon" class="h-6 w-6 mr-4" />
                  {{ n.name }}
                </NuxtLink>
                <button v-else class="rounded-btn" @click="n.action">
                  <Icon :name="n.icon" class="h-6 w-6 mr-4" />
                  {{ n.name }}
                </button>
              </li>
            </ul>
          </div>
        </div>

        <!-- Bottom -->
        <button class="mt-auto mb-6" @click="logout">Sign Out</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { useAuthStore } from "~~/stores/auth";
  import { useLabelStore } from "~~/stores/labels";
  import { useLocationStore } from "~~/stores/locations";

  const modals = reactive({
    item: false,
    location: false,
    label: false,
    import: false,
  });

  const dropdown = [
    {
      name: "Item / Asset",
      action: () => {
        modals.item = true;
      },
    },
    {
      name: "Location",
      action: () => {
        modals.location = true;
      },
    },
    {
      name: "Label",
      action: () => {
        modals.label = true;
      },
    },
  ];

  const route = useRoute();

  const nav = [
    {
      icon: "mdi-home",
      active: computed(() => route.path === "/home"),
      id: 0,
      name: "Home",
      to: "/home",
    },
    {
      icon: "mdi-account",
      id: 1,
      active: computed(() => route.path === "/profile"),
      name: "Profile",
      to: "/profile",
    },
    {
      icon: "mdi-document",
      id: 3,
      active: computed(() => route.path === "/items"),
      name: "Items",
      to: "/items",
    },
    {
      icon: "mdi-database",
      id: 2,
      name: "Import",
      action: () => {
        modals.import = true;
      },
    },
    // {
    //   icon: "mdi-database-export",
    //   id: 5,
    //   name: "Export",
    //   action: () => {
    //     console.log("Export");
    //   },
    // },
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

  const authStore = useAuthStore();
  const api = useUserApi();

  async function logout() {
    const { error } = await authStore.logout(api);
    if (error) {
      return;
    }

    navigateTo("/");
  }
</script>
