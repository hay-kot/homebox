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

  <div class="bg-neutral absolute shadow-xl top-0 h-[50vh] max-h-96 sm:h-[28vh] -z-10 w-full"></div>

  <BaseContainer cmp="header" class="py-6 max-w-none">
    <BaseContainer>
      <h2 class="mt-1 text-4xl font-bold tracking-tight text-neutral-content sm:text-5xl lg:text-6xl flex">
        HomeB
        <AppLogo class="w-12 -mb-4" style="padding-left: 3px; padding-right: 2px" />
        x
      </h2>
      <div class="ml-1 mt-2 text-lg text-neutral-content/75 space-x-2">
        <template v-for="link in links">
          <NuxtLink
            v-if="!link.action"
            :key="link.name"
            class="hover:text-base-content transition-color duration-200 italic"
            :to="link.href"
          >
            {{ link.name }}
          </NuxtLink>
          <button
            v-else
            :key="link.name + 'link'"
            for="location-form-modal"
            class="hover:text-base-content transition-color duration-200 italic"
            @click="link.action"
          >
            {{ link.name }}
          </button>
          <span v-if="!link.last" :key="link.name"> / </span>
        </template>
      </div>
      <div class="flex mt-6">
        <div class="dropdown">
          <label tabindex="0" class="btn btn-primary btn-sm">
            <span>
              <Icon name="mdi-plus" class="mr-1 -ml-1" />
            </span>
            Create
          </label>
          <ul tabindex="0" class="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-52">
            <li v-for="btn in dropdown" :key="btn.name">
              <button @click="btn.action">
                {{ btn.name }}
              </button>
            </li>
          </ul>
        </div>
      </div>
    </BaseContainer>
  </BaseContainer>
</template>
