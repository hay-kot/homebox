<script setup lang="ts">
  import TextField from '@/components/Form/TextField.vue';
  import { useNotifier } from '@/composables/use-notifier';
  import { usePublicApi } from '@/composables/use-api';
  import { useAuthStore } from '~~/stores/auth';
  useHead({
    title: 'Homebox | Organize and Tag Your Stuff',
  });

  definePageMeta({
    layout: 'empty',
  });

  const authStore = useAuthStore();
  if (!authStore.isTokenExpired) {
    navigateTo('/home');
  }

  const registerFields = [
    {
      label: "What's your name?",
      value: '',
    },
    {
      label: "What's your email?",
      value: '',
    },
    {
      label: 'Name your group',
      value: '',
    },
    {
      label: 'Set your password',
      value: '',
      type: 'password',
    },
    {
      label: 'Confirm your password',
      value: '',
      type: 'password',
    },
  ];

  const api = usePublicApi();

  async function registerUser() {
    loading.value = true;
    // Print Values of registerFields

    const { error } = await api.register({
      user: {
        name: registerFields[0].value,
        email: registerFields[1].value,
        password: registerFields[3].value,
      },
      groupName: registerFields[2].value,
    });

    if (error) {
      toast.error('Problem registering user');
      return;
    }

    toast.success('User registered');

    loading.value = false;
    loginFields[0].value = registerFields[1].value;
    registerForm.value = false;
  }

  const loginFields = [
    {
      label: 'Email',
      value: '',
    },
    {
      label: 'Password',
      value: '',
      type: 'password',
    },
  ];

  const toast = useNotifier();
  const loading = ref(false);

  async function login() {
    loading.value = true;
    const { data, error } = await api.login(loginFields[0].value, loginFields[1].value);

    if (error) {
      toast.error('Invalid email or password');
    } else {
      toast.success('Logged in successfully');

      authStore.$patch({
        token: data.token,
        expires: data.expiresAt,
      });

      navigateTo('/home');
    }
    loading.value = false;
  }

  const registerForm = ref(false);
  function toggleLogin() {
    registerForm.value = !registerForm.value;
  }
</script>

<template>
  <div>
    <header class="sm:px-6 py-2 lg:p-14 sm:py-6">
      <h2 class="mt-1 text-4xl font-bold tracking-tight text-base-content sm:text-5xl lg:text-6xl flex">
        HomeB
        <AppLogo class="w-12 -mb-4" style="padding-left: 3px; padding-right: 2px" />
        x
      </h2>
      <p class="ml-1 text-lg text-base-content/50">Track, Organize, and Manage your Shit.</p>
    </header>
    <div class="grid p-6 sm:place-items-center min-h-[50vh]">
      <div>
        <Transition name="slide-fade">
          <form v-if="registerForm" @submit.prevent="registerUser">
            <div class="card w-max-[500px] md:w-[500px] bg-base-100 shadow-xl">
              <div class="card-body">
                <h2 class="card-title">Register</h2>
                <TextField
                  v-for="field in registerFields"
                  v-model="field.value"
                  :label="field.label"
                  :key="field.label"
                  :type="field.type"
                />
                <div class="card-actions justify-end">
                  <button
                    type="submit"
                    class="btn btn-primary mt-2"
                    :class="loading ? 'loading' : ''"
                    :disabled="loading"
                  >
                    Register
                  </button>
                </div>
              </div>
            </div>
          </form>
          <form v-else @submit.prevent="login">
            <div class="card w-max-[500px] md:w-[500px] bg-base-100 shadow-xl">
              <div class="card-body">
                <h2 class="card-title">Login</h2>
                <TextField
                  v-for="field in loginFields"
                  v-model="field.value"
                  :label="field.label"
                  :key="field.label"
                  :type="field.type"
                />
                <div class="card-actions justify-end mt-2">
                  <button type="submit" class="btn btn-primary" :class="loading ? 'loading' : ''" :disabled="loading">
                    Login
                  </button>
                </div>
              </div>
            </div>
          </form>
        </Transition>
        <div class="text-center mt-4">
          <button @click="toggleLogin" class="text-primary-content text-lg">
            {{ registerForm ? 'Already a User? Login' : 'Not a User? Register' }}
          </button>
        </div>
      </div>
    </div>
    <div class="min-w-full absolute bottom-0 z-[-1]">
      <svg class="fill-primary" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="none" viewBox="0 0 1440 320">
        <path
          fill-opacity="1"
          d="M0,32L30,42.7C60,53,120,75,180,80C240,85,300,75,360,80C420,85,480,107,540,128C600,149,660,171,720,160C780,149,840,107,900,90.7C960,75,1020,85,1080,122.7C1140,160,1200,224,1260,234.7C1320,245,1380,203,1410,181.3L1440,160L1440,320L1410,320C1380,320,1320,320,1260,320C1200,320,1140,320,1080,320C1020,320,960,320,900,320C840,320,780,320,720,320C660,320,600,320,540,320C480,320,420,320,360,320C300,320,240,320,180,320C120,320,60,320,30,320L0,320Z"
        />
      </svg>
      <div class="bg-primary flex-col flex min-h-[32vh]">
        <div class="mt-auto mx-auto mb-8">
          <p class="text-center text-gray-200">&copy; 2022 Contents. All Rights Reserved. Haybytes LLC</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="css" scoped>
  .slide-fade-enter-active {
    transition: all 0.2s ease-out;
  }

  .slide-fade-enter-from,
  .slide-fade-leave-to {
    position: absolute;
    transform: translateX(20px);
    opacity: 0;
  }
</style>
