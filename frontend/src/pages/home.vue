<script setup lang="ts">
  import { useUserApi } from '@/composables/use-api';
  import { useAuthStore } from '@/store/auth';
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
</script>

<template>
  <section class="max-w-7xl mx-auto">
    <header class="sm:px-6 py-2 lg:p-14 sm:py-6">
      <h2 class="mt-1 text-4xl font-bold tracking-tight text-gray-200 sm:text-5xl lg:text-6xl">Homebox</h2>
      <div class="ml-1 text-lg text-gray-400 space-x-2">
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
    </header>
  </section>
  <section class="max-w-7xl mx-auto sm:px-6 lg:px-14">
    {{ user }}
  </section>
</template>

<route lang="yaml">
name: home
</route>
