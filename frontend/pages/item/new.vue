<script setup>
  definePageMeta({
    middleware: ["auth"],
  });

  const show = reactive({
    identification: true,
    purchase: false,
    sold: false,
    extras: false,
  });

  const form = reactive({
    name: "",
    description: "",
    notes: "",

    // Item Identification
    serialNumber: "",
    modelNumber: "",
    manufacturer: "",

    // Purchase Information
    purchaseTime: "",
    purchasePrice: "",
    purchaseFrom: "",

    // Sold Information
    soldTime: "",
    soldPrice: "",
    soldTo: "",
    soldNotes: "",
  });

  function submit() {
    console.log("Submitted!");
  }
</script>

<template>
  <BaseContainer cmp="section">
    <BaseSectionHeader> {{ $t("item.new.title") }} </BaseSectionHeader>
    <form class="max-w-3xl mx-auto my-5 space-y-6" @submit.prevent="submit">
      <div class="divider collapse-title px-0 cursor-pointer">{{ $t("item.new.info") }}</div>
      <div class="bg-base-200 card">
        <div class="card-body">
          <FormTextField v-model="form.name" :label="$t('item.new.name')" />
          <FormTextArea v-model="form.description" :label="$t('item.new.desp')" limit="1000" />
        </div>
      </div>

      <div class="divider">
        <button class="btn btn-sm" @click="show.identification = !show.identification">
          {{ $t("item.new.product.button") }}
        </button>
      </div>
      <div v-if="show.identification" class="card bg-base-200">
        <div class="card-body grid md:grid-cols-2">
          <FormTextField v-model="form.serialNumber" :label="$t('item.new.product.serialNumber')" />
          <FormTextField v-model="form.modelNumber" :label="$t('item.new.product.modelNumber')" />
          <FormTextField v-model="form.manufacturer" :label="$t('item.new.product.manufacturer')" />
        </div>
      </div>
      <div class="">
        <button class="btn btn-sm" @click="show.purchase = !show.purchase">
          {{ $t("item.new.purchase.button") }}
        </button>
        <div class="divider"></div>
      </div>
      <div v-if="show.purchase" class="card bg-base-200">
        <div class="card-body grid md:grid-cols-2">
          <FormTextField v-model="form.purchaseTime" :label="$t('item.new.purchase.purchaseTime')" />
          <FormTextField v-model="form.purchasePrice" :label="$t('item.new.purchase.purchasePrice')" />
          <FormTextField v-model="form.purchaseFrom" :label="$t('item.new.purchase.purchaseFrom')" />
        </div>
      </div>

      <div class="divider">
        <button class="btn btn-sm" @click="show.sold = !show.sold">
          {{ $t("item.new.sold.button") }}
        </button>
      </div>
      <div v-if="show.sold" class="card bg-base-200">
        <div class="card-body">
          <div class="grid md:grid-cols-2 gap-2">
            <FormTextField v-model="form.soldTime" :label="$t('item.new.sold.soldTime')" />
            <FormTextField v-model="form.soldPrice" :label="$t('item.new.sold.soldPrice')" />
            <FormTextField v-model="form.soldTo" :label="$t('item.new.sold.soldTo')" />
          </div>
          <FormTextArea v-model="form.soldNotes" :label="$t('item.new.sold.soldNotes')" limit="1000" />
        </div>
      </div>
      <div class="divider">
        <button class="btn btn-sm" @click="show.extras = !show.extras">
          {{ $t("item.new.extras.button") }}
        </button>
      </div>
      <div v-if="show.extras" class="card bg-base-200">
        <div class="card-body">
          <FormTextArea v-model="form.notes" :label="$t('item.new.extras.notes')" limit="1000" />
        </div>
      </div>
    </form>
  </BaseContainer>
</template>
