<template>
  <div class="dark">
    <router-view />
    <Toaster />
  </div>
</template>

<script setup>
import Toaster from "@/components/ui/toast/Toaster.vue";
</script>

<style scoped></style>
