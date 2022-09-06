<template>
  <div class="divider">
    <div class="btn-group min-w-[180px] flex-nowrap">
      <button @click="$emit('edit')" name="options" class="btn btn-sm btn-primary">
        <Icon name="heroicons-pencil" class="h-5 w-5 mr-1" aria-hidden="true" />
        <span> Edit </span>
      </button>
      <button @click="$emit('delete')" name="options" class="btn btn-sm btn-primary">
        <Icon name="heroicons-trash" class="h-5 w-5 mr-1" aria-hidden="true" />
        <span> Delete </span>
      </button>
    </div>
  </div>
</template>
