<script>
export const FORM_ITEM_INJECTION_KEY = Symbol();
</script>

<script setup>
import { provide } from "vue";
import { useId } from "radix-vue";
import { cn } from "/src/lib/utils";

const id = useId();
provide(FORM_ITEM_INJECTION_KEY, id);
</script>

<template>
  <div :class="cn('space-y-2', $attrs.class ?? '')">
    <slot />
  </div>
</template>
