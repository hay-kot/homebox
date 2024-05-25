<script setup lang="ts">
  import { useRouteHash } from "@vueuse/router";
  import MdiAccount from "~icons/mdi/account";
  import MdiAccountPlus from "~icons/mdi/account-plus";
  import MdiLogin from "~icons/mdi/login";
  import MdiArrowRight from "~icons/mdi/arrow-right";
  import MdiLock from "~icons/mdi/lock";

  enum PageForms {
    Register = "register",
    Login = "login",
    ForgotPassword = "forgot-password",
  }

  useHead({
    title: "Homebox | Organize and Tag Your Stuff",
  });

  definePageMeta({
    layout: "empty",
    middleware: [
      () => {
        const ctx = useAuthContext();
        if (ctx.isAuthorized()) {
          return "/home";
        }
      },
    ],
  });

  const ctx = useAuthContext();

  const api = usePublicApi();
  const toast = useNotifier();

  const pageForm = useRouteHash(PageForms.Login);
  const pageFormStr = computed(() => {
    if (!pageForm.value) {
      return PageForms.Login;
    }

    return pageForm.value[0] === "#" ? pageForm.value.slice(1) : pageForm.value;
  });

  const route = useRoute();
  const router = useRouter();

  const username = ref("");
  const email = ref("");
  const password = ref("");
  const canRegister = ref(false);
  const remember = ref(false);

  const groupToken = computed<string>({
    get() {
      const params = route.query.token;

      if (typeof params === "string") {
        return params;
      }

      return "";
    },
    set(v) {
      router.push({
        query: {
          token: v,
        },
      });
    },
  });

  async function registerUser() {
    loading.value = true;
    const { error } = await api.register({
      name: username.value,
      email: email.value,
      password: password.value,
      token: groupToken.value,
    });

    if (error) {
      toast.error("Problem registering user");
      return;
    }

    toast.success("User registered");

    loading.value = false;
    pageForm.value = PageForms.Login;
  }

  onMounted(() => {
    if (groupToken.value !== "") {
      pageForm.value = PageForms.Register;
    }
  });

  const loading = ref(false);
  const loginPassword = ref("");

  async function login() {
    loading.value = true;
    const { error } = await ctx.login(api, email.value, loginPassword.value, remember.value);

    if (error) {
      toast.error("Invalid email or password");
      loading.value = false;
      return;
    }

    toast.success("Logged in successfully");

    navigateTo("/home");
    loading.value = false;
  }

  async function resetPassword() {
    if (email.value === "") {
      toast.error("Email is required");
      return;
    }

    const resp = await api.resetPasseord(email.value);
    if (resp.error) {
      toast.error("Problem resetting password");
      return;
    }

    toast.success("Password reset link sent to your email");
    return await Promise.resolve();
  }
</script>

<template>
  <NuxtLayout v-slot="{ status }" name="center-card">
    <div>
      <Transition name="slide-fade">
        <form v-if="pageFormStr === PageForms.Register" @submit.prevent="registerUser">
          <div class="card w-max-[500px] md:w-[500px] bg-base-100 shadow-xl">
            <div class="card-body">
              <h2 class="card-title text-2xl align-center">
                <MdiAccount class="mr-1 w-7 h-7" />
                Register
              </h2>
              <FormTextField v-model="email" label="Set your email?" />
              <FormTextField v-model="username" label="What's your name?" />
              <div v-if="!(groupToken == '')" class="pt-4 pb-1 text-center">
                <p>You're Joining an Existing Group!</p>
                <button type="button" class="text-xs underline" @click="groupToken = ''">
                  Don't Want To Join a Group?
                </button>
              </div>
              <FormPassword v-model="password" label="Set your password" />
              <PasswordScore v-model:valid="canRegister" :password="password" />
              <div class="card-actions justify-end">
                <button
                  type="submit"
                  class="btn btn-primary mt-2"
                  :class="loading ? 'loading' : ''"
                  :disabled="loading || !canRegister"
                >
                  Register
                </button>
              </div>
            </div>
          </div>
        </form>
        <form v-else-if="pageFormStr === PageForms.ForgotPassword" @submit.prevent="resetPassword">
          <div class="card w-max-[500px] md:w-[500px] bg-base-100 shadow-xl">
            <div class="card-body">
              <h2 class="card-title text-2xl align-center">
                <MdiAccount class="mr-1 w-7 h-7" />
                Reset Password
              </h2>
              <FormTextField v-model="email" label="Email" />
              <p class="text-sm text-base-content/50">
                If you have an account with us, we will send you a password reset link.
              </p>
              <div class="card-actions justify-end mt-4">
                <button
                  type="submit"
                  class="btn btn-primary btn-block"
                  :class="loading ? 'loading' : ''"
                  :disabled="loading"
                >
                  Reset Password
                </button>
              </div>
            </div>
          </div>
        </form>
        <form v-else @submit.prevent="login">
          <div class="card w-max-[500px] md:w-[500px] bg-base-100 shadow-xl">
            <div class="card-body">
              <h2 class="card-title text-2xl align-center">
                <MdiAccount class="mr-1 w-7 h-7" />
                Login
              </h2>
              <template v-if="status && status.demo">
                <p class="text-xs italic text-center">This is a demo instance</p>
                <p class="text-xs text-center"><b>Email</b> demo@example.com</p>
                <p class="text-xs text-center"><b>Password</b> demo</p>
              </template>
              <FormTextField v-model="email" label="Email" />
              <FormPassword v-model="loginPassword" label="Password" />
              <div class="max-w-[140px]">
                <FormCheckbox v-model="remember" label="Remember Me" />
              </div>
              <div class="card-actions justify-end">
                <button
                  type="submit"
                  class="btn btn-primary btn-block"
                  :class="loading ? 'loading' : ''"
                  :disabled="loading"
                >
                  Login
                </button>
              </div>
            </div>
          </div>
        </form>
      </Transition>
      <div class="text-center mt-6">
        <BaseButton
          v-if="status && status.allowRegistration"
          class="btn-primary btn-wide"
          :to="pageFormStr === PageForms.Register ? `#${PageForms.Login}` : `#${PageForms.Register}`"
        >
          <template #icon>
            <MdiAccountPlus v-if="pageFormStr === PageForms.Register" class="w-5 h-5 swap-off" />
            <MdiLogin v-else class="w-5 h-5 swap-off" />
            <MdiArrowRight class="w-5 h-5 swap-on" />
          </template>
          {{ pageFormStr === PageForms.Register ? "Login" : "Register" }}
        </BaseButton>
        <p v-else class="text-base-content italic text-sm inline-flex items-center gap-2">
          <MdiLock class="w-4 h-4 inline-block" />
          Registration Disabled
        </p>
        <NuxtLink :to="`#${PageForms.ForgotPassword}`">
          <p class="text-xs text-base-content/50 mt-2">Forgot your password?</p>
        </NuxtLink>
      </div>
    </div>
  </NuxtLayout>
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

  progress[value]::-webkit-progress-value {
    transition: width 0.5s;
  }
</style>
