<script setup lang="ts">
  import { Detail } from "~~/components/global/DetailsSection/types";
  import { DaisyTheme } from "~~/composables/use-preferences";
  import { useAuthStore } from "~~/stores/auth";

  definePageMeta({
    layout: "home",
  });
  useHead({
    title: "Homebox | Profile",
  });

  const { setTheme } = useTheme();

  type ThemeOption = {
    label: string;
    value: DaisyTheme;
  };

  const themes: ThemeOption[] = [
    {
      label: "Garden",
      value: "garden",
    },
    {
      label: "Light",
      value: "light",
    },
    {
      label: "Cupcake",
      value: "cupcake",
    },
    {
      label: "Bumblebee",
      value: "bumblebee",
    },
    {
      label: "Emerald",
      value: "emerald",
    },
    {
      label: "Corporate",
      value: "corporate",
    },
    {
      label: "Synthwave",
      value: "synthwave",
    },
    {
      label: "Retro",
      value: "retro",
    },
    {
      label: "Cyberpunk",
      value: "cyberpunk",
    },
    {
      label: "Valentine",
      value: "valentine",
    },
    {
      label: "Halloween",
      value: "halloween",
    },
    {
      label: "Forest",
      value: "forest",
    },
    {
      label: "Aqua",
      value: "aqua",
    },
    {
      label: "Lofi",
      value: "lofi",
    },
    {
      label: "Pastel",
      value: "pastel",
    },
    {
      label: "Fantasy",
      value: "fantasy",
    },
    {
      label: "Wireframe",
      value: "wireframe",
    },
    {
      label: "Black",
      value: "black",
    },
    {
      label: "Luxury",
      value: "luxury",
    },
    {
      label: "Dracula",
      value: "dracula",
    },
    {
      label: "Cmyk",
      value: "cmyk",
    },
    {
      label: "Autumn",
      value: "autumn",
    },
    {
      label: "Business",
      value: "business",
    },
    {
      label: "Acid",
      value: "acid",
    },
    {
      label: "Lemonade",
      value: "lemonade",
    },
    {
      label: "Night",
      value: "night",
    },
    {
      label: "Coffee",
      value: "coffee",
    },
    {
      label: "Winter",
      value: "winter",
    },
  ];

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
      {
        name: "Invitation Code",
        text: "",
        slot: "invitation",
      },
      {
        name: "Change Password",
        text: "",
        slot: "change-password",
      },
      {
        name: "Delete Profile",
        text: "",
        slot: "delete-profile",
      },
    ] as Detail[];
  });

  const confirm = useConfirm();

  async function deleteProfile() {
    const result = await confirm.open(
      "Are you sure you want to delete your account? If you are the last member in your group all your data will be deleted. This action cannot be undone."
    );

    if (result.isCanceled) {
      return;
    }

    console.log("delete profile");
  }
</script>

<template>
  <BaseContainer class="flex flex-col gap-4 mb-6">
    <BaseCard>
      <template #title>
        <BaseSectionHeader>
          <Icon name="mdi-fill" class="mr-2 text-base-600" />
          <span class="text-base-600"> User Profile </span>
          <template #description> Invite users, and manage your account. </template>
        </BaseSectionHeader>
      </template>

      <DetailsSection :details="details">
        <template #invitation>
          <BaseButton class="ml-auto" size="sm"> Generate Invite Link </BaseButton>
        </template>
        <template #change-password>
          <BaseButton class="ml-auto" size="sm"> Change Password </BaseButton>
        </template>
        <template #delete-profile>
          <BaseButton class="ml-auto btn-error" size="sm" @click="deleteProfile"> Delete Profile </BaseButton>
        </template>
      </DetailsSection>
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
  </BaseContainer>
</template>

<style scoped></style>
