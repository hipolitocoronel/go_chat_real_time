<script setup>
import { useFormField } from "./useFormField";
import { cn } from "/src/lib/utils";

const { formDescriptionId } = useFormField();
</script>

<template>
  <p
    :id="formDescriptionId"
    :class="cn('text-[0.8rem] text-muted-foreground', $attrs.class ?? '')"
  >
    <slot />
  </p>
</template>
