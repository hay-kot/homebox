<script setup lang="ts">
  definePageMeta({
    middleware: ["auth"],
  });

  const route = useRoute();
  const assetId = computed<string>(() => route.params.id as string);
  await navigateTo("/assets/" + assetId.value, { replace: true, redirectCode: 301 });
</script>
