<script setup lang="ts">
import CardContainer from "@/components/card/CardGrid.vue";
import Filters from "@/components/layout/Filters.vue";
import Map from "@/components/map/Map.vue";
import NavBar from "@/components/layout/Navbar.vue";
import {useRoute, useRouter} from "vue-router";
import {computed, onMounted, ref} from "vue";


const router = useRouter();
const route = useRoute();
const currentCenter = ref({})


const query = computed(() => route.query)

onMounted(async () => {
  await router.isReady();
  currentCenter.value = route.query
});


</script>

<template>
  <NavBar/>
  <Filters/>
  <div class="flex flex-row gap-4 w-full h-[91vh] overflow-y-hidden">
    <Map :center="route.query"/>
    <card-container/>
  </div>
</template>
