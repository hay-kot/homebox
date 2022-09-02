<script lang="ts" setup>
  import { useAuthStore } from '~~/stores/auth';

  const authStore = useAuthStore();
  const api = useUserApi();

  async function logout() {
    const { error } = await authStore.logout(api);
    if (error) {
      return;
    }

    navigateTo('/');
  }

  const links = [
    {
      name: 'Home',
      href: '/home',
    },
    {
      name: 'Logout',
      action: logout,
      last: true,
    },
  ];

  const modals = reactive({
    item: false,
    location: false,
    label: false,
  });

  const dropdown = [
    {
      name: 'Item / Asset',
      action: () => {
        modals.item = true;
      },
    },
    {
      name: 'Location',
      action: () => {
        modals.location = true;
      },
    },
    {
      name: 'Label',
      action: () => {
        modals.label = true;
      },
    },
  ];
</script>

<template>
  <!--
    Confirmation Modal is a singleton used by all components so we render
    it here to ensure it's always available. Possibly could move this further
    up the tree
   -->
  <ModalConfirm />
  <ItemCreateModal v-model="modals.item" />
  <LabelCreateModal v-model="modals.label" />
  <LocationCreateModal v-model="modals.location" />

  <BaseContainer is="header" class="py-6">
    <h2 class="mt-1 text-4xl font-bold tracking-tight text-base-content sm:text-5xl lg:text-6xl flex">
      HomeB
      <AppLogo class="w-12 -mb-4" style="padding-left: 3px; padding-right: 2px" />
      x
    </h2>
    <div class="ml-1 mt-2 text-lg text-base-content/50 space-x-2">
      <template v-for="link in links">
        <NuxtLink
          v-if="!link.action"
          class="hover:text-base-content transition-color duration-200 italic"
          :to="link.href"
        >
          {{ link.name }}
        </NuxtLink>
        <button
          for="location-form-modal"
          v-else
          @click="link.action"
          class="hover:text-base-content transition-color duration-200 italic"
        >
          {{ link.name }}
        </button>
        <span v-if="!link.last"> / </span>
      </template>
    </div>
    <div class="flex mt-6">
      <div class="dropdown">
        <label tabindex="0" class="btn btn-sm">
          <span>
            <Icon name="mdi-plus" class="mr-1 -ml-1" />
          </span>
          Create
        </label>
        <ul tabindex="0" class="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-52">
          <li v-for="btn in dropdown">
            <button @click="btn.action">
              {{ btn.name }}
            </button>
          </li>
        </ul>
      </div>
    </div>
  </BaseContainer>
</template>
