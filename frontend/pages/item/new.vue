<script setup>
  definePageMeta({
    layout: 'home',
  });

  const show = reactive({
    identification: true,
    purchase: false,
    sold: false,
    extras: false,
  });

  const form = reactive({
    name: '',
    description: '',
    notes: '',

    // Item Identification
    serialNumber: '',
    modelNumber: '',
    manufacturer: '',

    // Purchase Information
    purchaseTime: '',
    purchasePrice: '',
    purchaseFrom: '',

    // Sold Information
    soldTime: '',
    soldPrice: '',
    soldTo: '',
    soldNotes: '',
  });

  function submit() {}
</script>

<template>
  <BaseContainer is="section">
    <BaseSectionHeader> Add an Item To Your Inventory </BaseSectionHeader>
    <form @submit.prevent="submit" class="max-w-3xl mx-auto my-5 space-y-6">
      <div class="divider collapse-title px-0 cursor-pointer">Required Information</div>
      <div class="bg-base-200 card">
        <div class="card-body">
          <FormTextField label="Name" v-model="form.name" />
          <FormTextArea label="Description" v-model="form.description" limit="1000" />
        </div>
      </div>

      <div class="divider">
        <button class="btn btn-sm" @click="show.identification = !show.identification">Product Information</button>
      </div>
      <div class="card bg-base-200" v-if="show.identification">
        <div class="card-body grid md:grid-cols-2">
          <FormTextField label="Serial Number" v-model="form.serialNumber" />
          <FormTextField label="Model Number" v-model="form.modelNumber" />
          <FormTextField label="Manufacturer" v-model="form.manufacturer" />
        </div>
      </div>
      <div class="">
        <button class="btn btn-sm" @click="show.purchase = !show.purchase">Purchase Information</button>
        <div class="divider"></div>
      </div>
      <div class="card bg-base-200" v-if="show.purchase">
        <div class="card-body grid md:grid-cols-2">
          <FormTextField label="Purchase Time" v-model="form.purchaseTime" />
          <FormTextField label="Purchase Price" v-model="form.purchasePrice" />
          <FormTextField label="Purchase From" v-model="form.purchaseFrom" />
        </div>
      </div>

      <div class="divider">
        <button class="btn btn-sm" @click="show.sold = !show.sold">Sold Information</button>
      </div>
      <div class="card bg-base-200" v-if="show.sold">
        <div class="card-body">
          <div class="grid md:grid-cols-2 gap-2">
            <FormTextField label="Sold Time" v-model="form.soldTime" />
            <FormTextField label="Sold Price" v-model="form.soldPrice" />
            <FormTextField label="Sold To" v-model="form.soldTo" />
          </div>
          <FormTextArea label="Sold Notes" v-model="form.soldNotes" limit="1000" />
        </div>
      </div>
      <div class="divider">
        <button class="btn btn-sm" @click="show.extras = !show.extras">Extras</button>
      </div>
      <div class="card bg-base-200" v-if="show.extras">
        <div class="card-body">
          <FormTextArea label="Notes" v-model="form.notes" limit="1000" />
        </div>
      </div>
    </form>
  </BaseContainer>
</template>
