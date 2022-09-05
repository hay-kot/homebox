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
    <div class="fill-primary min-w-full absolute top-0 z-[-1]">
      <div class="bg-primary flex-col flex min-h-[20vh]" />
      <svg
        class="fill-primary drop-shadow-xl"
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 1440 320"
        preserveAspectRatio="none"
      >
        <path
          fill-opacity="1"
          d="M0,32L80,69.3C160,107,320,181,480,181.3C640,181,800,107,960,117.3C1120,128,1280,224,1360,272L1440,320L1440,0L1360,0C1280,0,1120,0,960,0C800,0,640,0,480,0C320,0,160,0,80,0L0,0Z"
        ></path>
      </svg>
    </div>
    <div>
      <header class="p-4 sm:px-6 lg:p-14 sm:py-6 sm:flex sm:items-end mx-auto">
        <div>
          <h2 class="mt-1 text-4xl font-bold tracking-tight text-neutral-content sm:text-5xl lg:text-6xl flex">
            HomeB
            <AppLogo class="w-12 -mb-4" style="padding-left: 3px; padding-right: 2px" />
            x
          </h2>
          <p class="ml-1 text-lg text-base-content/50">Track, Organize, and Manage your Shit.</p>
        </div>
        <div class="flex mt-6 sm:mt-0 gap-4 ml-auto text-neutral-content">
          <a class="tooltip" data-tip="Project Github" href="https://github.com/hay-kot/homebox" target="_blank">
            <Icon name="mdi-github" class="h-8 w-8" />
          </a>
          <a href="https://twitter.com/haybytes" class="tooltip" data-tip="Follow The Developer" target="_blank">
            <Icon name="mdi-twitter" class="h-8 w-8" />
          </a>
          <a href="/" class="tooltip" data-tip="Join The Discord">
            <Icon name="mdi-discord" class="h-8 w-8" />
          </a>
          <a href="/" class="tooltip" data-tip="Read The Docs">
            <Icon name="mdi-folder" class="h-8 w-8" />
          </a>
        </div>
      </header>
      <div class="grid p-6 sm:place-items-center min-h-[50vh]">
        <div>
          <Transition name="slide-fade">
            <form v-if="registerForm" @submit.prevent="registerUser">
              <div class="card w-max-[500px] md:w-[500px] bg-base-100 shadow-xl">
                <div class="card-body">
                  <h2 class="card-title text-2xl align-center">
                    <Icon name="heroicons-user" class="mr-1 w-7 h-7" />
                    Register
                  </h2>
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
                  <h2 class="card-title text-2xl align-center">
                    <Icon name="heroicons-user" class="mr-1 w-7 h-7" />
                    Login
                  </h2>
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
          <div class="text-center mt-6">
            <button
              @click="toggleLogin"
              class="text-base-content text-lg hover:bg-primary hover:text-primary-content px-3 py-1 rounded-xl transition-colors duration-200"
            >
              {{ registerForm ? 'Already a User? Login' : 'Not a User? Register' }}
            </button>
          </div>
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
