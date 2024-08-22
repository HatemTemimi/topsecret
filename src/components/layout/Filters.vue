<script setup lang="ts">
</script>

<template>
  <div class="bg-gray-700">
    <div class="mx-auto px-2 sm:px-6 lg:px-8">
      <div class="relative flex flex-row justify-between h-[4vh]">
        <div class="flex flex-shrink-0 items-center">
          <Input type="text" placeholder="search by location.." class="rounded text-center"/>
        </div>
        <div class="flex flex-shrink-0 items-center">
          <select class=" w-60 h-3/5 text-center
cursor-default rounded bg-white text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:outline-none focus:ring-2 focus:ring-indigo-500 sm:text-sm sm:leading-6
">
            <option>option 1</option>
            <option>option</option>
            <option>option</option>
          </select>
        </div>
        <div class="flex flex-shrink-0 items-center">
          <div class="text-white">Filter 1</div>
        </div>
        <div class="flex flex-shrink-0 items-center">
          <div class="text-white">Filter 1</div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>