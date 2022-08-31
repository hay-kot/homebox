<script setup lang="ts">
  import { useAuthStore } from '@/store/auth';
  import { type Location } from '@/api/classes/locations';
  import { Icon } from '@iconify/vue';
  import { useUserApi } from '@/composables/use-api';
  useHead({
    title: 'Homebox | Home',
  });

  const api = useUserApi();

  const user = ref({});

  onMounted(async () => {
    const { data } = await api.self();

    if (data) {
      user.value = data.item;
    }
  });

  const locations = ref<Location[]>([]);

  onMounted(async () => {
    const { data } = await api.locations.getAll();

    if (data) {
      console.log(data);
      locations.value = data.items;
    }
  });

  const authStore = useAuthStore();
  const router = useRouter();

  async function logout() {
    const { error } = await authStore.logout(api);

    if (error) {
      console.error(error);
      return;
    }

    router.push('/');
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
      action: () => {},
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
</script>

<template>
  <section class="max-w-6xl mx-auto">
    <header class="sm:px-6 py-2 lg:px-14 sm:py-6">
      <h2 class="mt-1 text-4xl font-bold tracking-tight text-base-content sm:text-5xl lg:text-6xl">Homebox</h2>
      <div class="ml-1 mt-2 text-lg text-base-content/50 space-x-2">
        <template v-for="link in links">
          <router-link
            v-if="!link.action"
            class="hover:text-base-content transition-color duration-200 italic"
            :to="link.href"
          >
            {{ link.name }}
          </router-link>
          <button v-else @click="link.action" class="hover:text-base-content transition-color duration-200 italic">
            {{ link.name }}
          </button>
          <span v-if="!link.last"> / </span>
        </template>
      </div>
      <div class="flex mt-6">
        <div class="dropdown">
          <label tabindex="0" class="btn btn-sm">
            <span>
              <Icon icon="mdi-plus" class="w-5 h-5 mr-2" />
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
    </header>
  </section>
  <section class="max-w-6xl mx-auto sm:px-6 lg:px-14">
    <div class="border-b border-gray-600 pb-3 mb-3">
      <h3 class="text-lg text-base-content font-medium leading-6">Storage Locations</h3>
    </div>
    <div class="grid grid-cols-3 gap-4">
      <a
        :href="`#${l.id}`"
        class="card bg-primary text-primary-content hover:-translate-y-1 focus:-translate-y-1 transition duration-300"
        v-for="l in locations"
      >
        <div class="card-body p-4">
          <h2 class="flex items-center gap-2">
            <Icon icon="mdi-light:home" class="h-5 w-5" height="25" />
            {{ l.name }}
            <span class="badge badge-accent badge-lg ml-auto text-accent-content text-lg">0</span>
          </h2>
        </div>
      </a>
    </div>
  </section>
</template>

<route lang="yaml">
name: home
</route>
