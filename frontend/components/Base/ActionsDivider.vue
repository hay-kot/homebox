<template>
  <div class="relative">
    <div class="absolute inset-0 flex items-center" aria-hidden="true">
      <div class="w-full border-t border-primary" />
    </div>
    <div class="relative flex justify-center">
      <span class="isolate inline-flex -space-x-px rounded-md shadow-sm">
        <div class="btn-group">
          <button @click="$emit('edit')" name="options" class="btn btn-sm btn-primary">
            <Icon name="heroicons-pencil" class="h-5 w-5 mr-1" aria-hidden="true" />
            <span> Edit </span>
          </button>
          <button @click="$emit('delete')" name="options" class="btn btn-sm btn-primary">
            <Icon name="heroicons-trash" class="h-5 w-5 mr-1" aria-hidden="true" />
            <span> Delete </span>
          </button>
        </div>
      </span>
    </div>
  </div>
</template>
