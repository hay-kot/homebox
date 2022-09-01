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

  const dropdown = [
    {
      name: 'Location',
      action: () => {
        modal.value = true;
      },
    },
    {
      name: 'Item / Asset',
      action: () => {},
    },
    {
      name: 'Label',
      action: () => {},
    },
  ];

  // ----------------------------
  // Location Stuff
  // Should move to own component
  const locationLoading = ref(false);
  const locationForm = reactive({
    name: '',
    description: '',
  });

  const locationNameRef = ref(null);
  const triggerFocus = ref(false);
  const modal = ref(false);

  whenever(
    () => modal.value,
    () => {
      triggerFocus.value = true;
    }
  );

  async function createLocation() {
    locationLoading.value = true;
    const { data } = await api.locations.create(locationForm);

    if (data) {
      navigateTo(`/location/${data.id}`);
    }

    locationLoading.value = false;
    modal.value = false;
    locationForm.name = '';
    locationForm.description = '';
    triggerFocus.value = false;
  }
</script>

<template>
  <ModalConfirm />
  <BaseModal v-model="modal">
    <template #title> Create Location </template>
    <form @submit.prevent="createLocation">
      <FormTextField
        :trigger-focus="triggerFocus"
        ref="locationNameRef"
        :autofocus="true"
        label="Location Name"
        v-model="locationForm.name"
      />
      <FormTextField label="Location Description" v-model="locationForm.description" />
      <div class="modal-action">
        <BaseButton type="submit" :loading="locationLoading"> Create </BaseButton>
      </div>
    </form>
  </BaseModal>
  <BaseContainer is="header" class="py-6">
    <h2 class="mt-1 text-4xl font-bold tracking-tight text-base-content sm:text-5xl lg:text-6xl">Homebox</h2>
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
            <Icon name="mdi-plus" class="w-5 h-5 mr-2" />
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
