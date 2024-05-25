<template>
  <div>
    <Title>Password Reset</Title>
    <form @submit.prevent="resetPassword">
      <div class="card w-max-[500px] md:w-[500px] bg-base-100 shadow-xl">
        <div class="card-body">
          <h2 class="card-title text-2xl align-center">
            <MdiAccount class="mr-1 w-7 h-7" />
            Password Reset
          </h2>
          <FormPassword v-model="form.password" label="New Password" />
          <FormPassword v-model="form.passwordConfirm" label="Confirm Password" />
          <PasswordScore v-model:valid="form.requirementsMet" :password="form.password" />
          <div class="card-actions justify-end">
            <button type="submit" class="btn btn-primary mt-2" :class="loading ? 'loading' : ''">Reset Password</button>
          </div>
        </div>
      </div>
    </form>
    <div class="grid place-content-center pt-4">
      <NuxtLink to="/#login">
        <p class="text-xs text-base-content/50 mt-2">Account Login</p>
      </NuxtLink>
    </div>
  </div>
</template>

<script setup lang="ts">
  import MdiAccount from "~icons/mdi/account";

  const route = useRoute();

  definePageMeta({
    title: "Password Reset",
    layout: "center-card",
    middleware: [
      () => {
        const ctx = useAuthContext();
        if (ctx.isAuthorized()) {
          return "/home";
        }
      },
    ],
  });

  const toast = useNotifier();

  const loading = ref(false);

  const token = route.query.token;

  const form = reactive({
    requirementsMet: false,
    password: "",
    passwordConfirm: "",
  });
  function resetPassword() {
    if (token === undefined) {
      return toast.error("Invalid reset token");
    }

    if (form.password !== form.passwordConfirm) {
      return toast.error("Passwords do not match");
    }

    if (!form.requirementsMet) {
      return toast.error("Password does not meet requirements");
    }

    loading.value = true;
  }
</script>

<style scoped></style>
