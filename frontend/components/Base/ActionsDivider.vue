<template>
  <div class="divider">
    <div class="btn-group min-w-[180px] flex-nowrap">
      <button name="options" class="btn btn-sm btn-primary" @click="$emit('edit')">
        <Icon name="heroicons-pencil" class="h-5 w-5 mr-1" aria-hidden="true" />
        <span> Edit </span>
      </button>
      <button name="options" class="btn btn-sm btn-primary" @click="$emit('delete')">
        <Icon name="heroicons-trash" class="h-5 w-5 mr-1" aria-hidden="true" />
        <span> Delete </span>
      </button>
    </div>
  </div>
</template>
