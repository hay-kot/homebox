<script setup lang="ts">
  import { Detail } from "~~/components/global/DetailsSection/types";
  import { useAuthStore } from "~~/stores/auth";
  import { themes } from "~~/lib/data/themes";
  import { currencies, Currency } from "~~/lib/data/currency";

  definePageMeta({
    middleware: ["auth"],
  });
  useHead({
    title: "Homebox | Profile",
  });

  const api = useUserApi();
  const confirm = useConfirm();
  const notify = useNotifier();

  // Currency Selection
  const currency = ref<Currency>(currencies[0]);

  watch(currency, () => {
    if (group.value) {
      group.value.currency = currency.value.code;
    }

    console.log(group.value);
  });

  const currencyExample = computed(() => {
    const formatter = new Intl.NumberFormat("en-US", {
      style: "currency",
      currency: currency.value ? currency.value.code : "USD",
    });

    return formatter.format(1000);
  });

  const { data: group } = useAsyncData(async () => {
    const { data } = await api.group.get();
    return data;
  });

  // Sync Initial Currency
  watch(group, () => {
    if (group.value) {
      const found = currencies.find(c => c.code === group.value.currency);
      if (found) {
        currency.value = found;
      }
    }
  });

  async function updateGroup() {
    const { data, error } = await api.group.update({
      name: group.value.name,
      currency: group.value.currency,
    });

    if (error) {
      notify.error("Failed to update group");
      return;
    }

    group.value = data;
    notify.success("Group updated");
  }

  const pubApi = usePublicApi();
  const { data: status } = useAsyncData(async () => {
    const { data } = await pubApi.status();

    return data;
  });

  const { setTheme } = useTheme();

  const auth = useAuthStore();

  const details = computed(() => {
    return [
      {
        name: "Name",
        text: auth.self?.name || "Unknown",
      },
      {
        name: "Email",
        text: auth.self?.email || "Unknown",
      },
    ] as Detail[];
  });

  async function deleteProfile() {
    const result = await confirm.open(
      "Are you sure you want to delete your account? If you are the last member in your group all your data will be deleted. This action cannot be undone."
    );

    if (result.isCanceled) {
      return;
    }

    const { response } = await api.user.delete();

    if (response?.status === 204) {
      notify.success("Your account has been deleted.");
      auth.logout(api);
      navigateTo("/");
    }

    notify.error("Failed to delete your account.");
  }

  const token = ref("");
  const tokenUrl = computed(() => {
    if (!window) {
      return "";
    }

    return `${window.location.origin}?token=${token.value}`;
  });

  async function generateToken() {
    const date = new Date();

    const { response, data } = await api.group.createInvitation({
      expiresAt: new Date(date.setDate(date.getDate() + 7)),
      uses: 1,
    });

    if (response?.status === 201) {
      token.value = data.token;
    }
  }

  const passwordChange = reactive({
    loading: false,
    dialog: false,
    current: "",
    new: "",
    isValid: false,
  });

  function openPassChange() {
    passwordChange.dialog = true;
  }

  async function changePassword() {
    passwordChange.loading = true;
    if (!passwordChange.isValid) {
      return;
    }

    const { error } = await api.user.changePassword(passwordChange.current, passwordChange.new);

    if (error) {
      notify.error("Failed to change password.");
      passwordChange.loading = false;
      return;
    }

    notify.success("Password changed successfully.");
    passwordChange.dialog = false;
    passwordChange.new = "";
    passwordChange.current = "";
    passwordChange.loading = false;
  }

  async function ensureAssetIDs() {
    const { isCanceled } = await confirm.open(
      "Are you sure you want to ensure all assets have an ID? This will take a while and cannot be undone."
    );

    if (isCanceled) {
      return;
    }

    const result = await api.actions.ensureAssetIDs();

    if (result.error) {
      notify.error("Failed to ensure asset IDs.");
      return;
    }

    notify.success(`${result.data.completed} assets have been updated.`);
  }
</script>

<template>
  <div>
    <BaseModal v-model="passwordChange.dialog">
      <template #title> Change Password </template>

      <FormTextField v-model="passwordChange.current" label="Current Password" type="password" />
      <FormTextField v-model="passwordChange.new" label="New Password" type="password" />
      <PasswordScore v-model:valid="passwordChange.isValid" :password="passwordChange.new" />

      <div class="flex">
        <BaseButton
          class="ml-auto"
          :loading="passwordChange.loading"
          :disabled="!passwordChange.isValid"
          @click="changePassword"
        >
          Submit
        </BaseButton>
      </div>
    </BaseModal>

    <BaseContainer class="flex flex-col gap-4 mb-6">
      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-account" class="mr-2 -mt-1 text-base-600" />
            <span class="text-base-600"> User Profile </span>
            <template #description> Invite users, and manage your account. </template>
          </BaseSectionHeader>
        </template>

        <DetailsSection :details="details" />

        <div class="p-4">
          <div class="flex gap-2">
            <BaseButton size="sm" @click="openPassChange"> Change Password </BaseButton>
            <BaseButton size="sm" @click="generateToken"> Generate Invite Link </BaseButton>
          </div>
          <div v-if="token" class="pt-4 flex items-center pl-1">
            <CopyText class="mr-2 btn-primary" :text="tokenUrl" />
            {{ tokenUrl }}
          </div>
          <div v-if="token" class="pt-4 flex items-center pl-1">
            <CopyText class="mr-2 btn-primary" :text="token" />
            {{ token }}
          </div>
        </div>
      </BaseCard>

      <BaseCard>
        <template #title>
          <BaseSectionHeader class="pb-0">
            <Icon name="mdi-accounts" class="mr-2 -mt-1 text-base-600" />
            <span class="text-base-600"> Group Settings </span>
            <template #description>
              Shared Group Settings. You may need to refresh your browser for some settings to apply.
            </template>
          </BaseSectionHeader>
        </template>

        <div v-if="group" class="p-5 pt-0">
          <FormSelect v-model="currency" label="Currency Format" :items="currencies" />
          <p class="m-2 text-sm">Example: {{ currencyExample }}</p>

          <div class="mt-4 flex justify-end">
            <BaseButton @click="updateGroup"> Update Group </BaseButton>
          </div>
        </div>
      </BaseCard>

      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-fill" class="mr-2 text-base-600" />
            <span class="text-base-600"> Theme Settings </span>
            <template #description>
              Theme settings are stored in your browser's local storage. You can change the theme at any time. If you're
              having trouble setting your theme try refreshing your browser.
            </template>
          </BaseSectionHeader>
        </template>

        <div class="px-4 pb-4">
          <div class="rounded-box grid grid-cols-1 gap-4 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
            <div
              v-for="theme in themes"
              :key="theme.value"
              class="border-base-content/20 hover:border-base-content/40 outline-base-content overflow-hidden rounded-lg border outline-2 outline-offset-2"
              :data-theme="theme.value"
              :data-set-theme="theme.value"
              data-act-class="outline"
              @click="setTheme(theme.value)"
            >
              <div :data-theme="theme.value" class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
                <div class="grid grid-cols-5 grid-rows-3">
                  <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                  <div class="bg-base-300 col-start-1 row-start-3"></div>
                  <div class="bg-base-100 col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                    <div class="font-bold">{{ theme.label }}</div>
                    <div class="flex flex-wrap gap-1">
                      <div class="bg-primary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                        <div class="text-primary-content text-sm font-bold">A</div>
                      </div>
                      <div class="bg-secondary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                        <div class="text-secondary-content text-sm font-bold">A</div>
                      </div>
                      <div class="bg-accent flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                        <div class="text-accent-content text-sm font-bold">A</div>
                      </div>
                      <div class="bg-neutral flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                        <div class="text-neutral-content text-sm font-bold">A</div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </BaseCard>

      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-warning" class="mr-2 -mt-1 text-base-600" />
            <span class="text-base-600"> Actions </span>
            <template #description>
              Apply Actions to your inventory in bulk. These are irreversible actions. Be careful.
            </template>
          </BaseSectionHeader>

          <div class="py-4 border-t-2 border-gray-300">
            <div class="grid grid-cols-1 md:grid-cols-4 gap-10">
              <div class="col-span-3">
                <h4>Manage Asset IDs</h4>
                <p class="text-sm">
                  Ensures that all items in your inventory have a valid asset_id field. This is done by finding the
                  highest current asset_id field in the database and applying the next value to each item that has an
                  unset asset_id field. This is done in order of the created_at field.
                </p>
              </div>
              <BaseButton class="btn-primary mt-auto" @click="ensureAssetIDs"> Ensure Asset IDs </BaseButton>
            </div>
          </div>
        </template>
      </BaseCard>

      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-delete" class="mr-2 -mt-1 text-base-600" />
            <span class="text-base-600"> Delete Account</span>
            <template #description> Delete your account and all it's associated data </template>
          </BaseSectionHeader>

          <div class="py-4 border-t-2 border-gray-300">
            <BaseButton class="btn-error" @click="deleteProfile"> Delete Account </BaseButton>
          </div>
        </template>
      </BaseCard>
    </BaseContainer>
    <footer v-if="status" class="text-center w-full bottom-0 pb-4">
      <p class="text-center text-sm">Version: {{ status.build.version }} ~ Build: {{ status.build.commit }}</p>
    </footer>
  </div>
</template>

<style scoped></style>
