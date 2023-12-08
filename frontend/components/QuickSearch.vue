<template>
  <BaseModal v-model="isActive">
    <template #title> Quick Search </template>
    <div class="flex flex-wrap md:flex-nowrap gap-4 items-end">
      <div class="w-full">
        <FormTextField v-model="query" placeholder="Search" trigger-focus @keyup.prevent.enter="quickSearch" />
      </div>
      <BaseButton class="btn-block md:w-auto" @click.prevent="quickSearch">
        <template #icon>
          <Icon name="mdi-search" />
        </template>
        Search
      </BaseButton>
    </div>
  </BaseModal>
</template>

<script setup lang="ts">
  const query = ref("");
  const { isActive } = useQuickSearch();
  const router = useRouter();

  function quickSearch() {
    router.push({
      path: "/items",
      query: {
        q: query.value,
      },
    });

    isActive.value = false;
    query.value = "";
  }
</script>
