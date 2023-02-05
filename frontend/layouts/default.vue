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
      <input id="my-drawer-2" v-model="drawerToggle" type="checkbox" class="drawer-toggle" />
      <div class="drawer-content justify-center bg-base-300 pt-20 lg:pt-0">
        <AppHeaderDecor class="-mt-10 hidden lg:block" />
        <!-- Button -->
        <div class="navbar z-[99] lg:hidden top-0 fixed bg-primary shadow-md drawer-button">
          <NuxtLink to="/home">
            <h2 class="mt-1 ml-1 text-3xl font-bold tracking-tight text-base-100 flex">
              HomeB
              <AppLogo class="w-8 -mb-3" />
              x
            </h2>
          </NuxtLink>
          <label for="my-drawer-2" class="btn btn-square btn-ghost ml-auto text-base-100 drawer-button lg:hidden">
            <Icon name="mdi-menu" class="h-6 w-6" />
          </label>
        </div>

        <slot></slot>
      </div>

      <!-- Sidebar -->
      <div class="drawer-side shadow-lg">
        <label for="my-drawer-2" class="drawer-overlay"></label>

        <!-- Top Section -->
        <div class="w-60 py-5 md:py-10 bg-base-200 flex flex-grow-1 flex-col">
          <div class="space-y-8">
            <div class="flex flex-col items-center gap-4">
              <p>Welcome, {{ username }}</p>
              <NuxtLink class="avatar placeholder" to="/home">
                <div class="bg-base-300 text-neutral-content rounded-full w-24 p-4">
                  <AppLogo />
                </div>
              </NuxtLink>
            </div>
            <div class="flex flex-col bg-base-200">
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
                <li v-for="n in nav" :key="n.id" class="text-xl" @click="unfocus">
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
          <button class="mt-auto mx-2 hover:bg-base-300 p-3 rounded-btn" @click="logout">Sign Out</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { useAuthStore } from "~~/stores/auth";
  import { useLabelStore } from "~~/stores/labels";
  import { useLocationStore } from "~~/stores/locations";

  const username = computed(() => authStore.self?.name || "User");

  // Preload currency format
  useFormatCurrency();

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

  const drawerToggle = ref();

  function unfocus() {
    // unfocus current element
    drawerToggle.value = false;
  }

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
      icon: "mdi-magnify",
      id: 3,
      active: computed(() => route.path === "/items"),
      name: "Search",
      to: "/items",
    },
    {
      icon: "mdi-map-marker",
      id: 4,
      active: computed(() => route.path === "/locations"),
      name: "Locations",
      to: "/locations",
    },
    {
      icon: "mdi-database",
      id: 2,
      name: "Import",
      action: () => {
        modals.import = true;
      },
    },
  ];

  function isMutation(method: string | undefined) {
    return method === "POST" || method === "PUT" || method === "DELETE";
  }
  function isSuccess(status: number) {
    return status >= 200 && status < 300;
  }

  const labelStore = useLabelStore();
  const reLabel = /\/api\/v1\/labels\/.*/gm;
  const rmLabelStoreObserver = defineObserver("labelStore", {
    handler: (resp, req) => {
      if (isMutation(req?.method) && isSuccess(resp.status) && resp.url.match(reLabel)) {
        labelStore.refresh();
      }
      console.debug("labelStore handler called by observer");
    },
  });

  const locationStore = useLocationStore();
  const reLocation = /\/api\/v1\/locations\/.*/gm;
  const rmLocationStoreObserver = defineObserver("locationStore", {
    handler: (resp, req) => {
      if (isMutation(req?.method) && isSuccess(resp.status) && resp.url.match(reLocation)) {
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

    eventBus.emit(EventTypes.ClearStores);
    navigateTo("/");
  }
</script>
