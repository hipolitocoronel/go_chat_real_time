<template>
  <div class="form-container">
    <Card class="w-[350px]">
      <CardHeader
        class="flex flex-col items-center justify-center text-center bg-gradient-to-r from-blue-400 to-teal-500 text-white py-8 rounded-t-lg"
      >
        <CardTitle
          class="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
          >Welcome!
        </CardTitle>
        <CardDescription>Go Real Time Chat</CardDescription>
      </CardHeader>
      <CardContent>
        <!-- ---------------- FORMULARIO ------------------ -->
        <form @submit="onSubmit">
          <!-- Username -->
          <FormField v-slot="{ componentField }" name="username">
            <FormItem class="my-8">
              <FormLabel>Username</FormLabel>
              <FormControl>
                <Input
                  type="text"
                  placeholder="Enter your username"
                  v-bind="componentField"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <!-- Password -->
          <FormField v-slot="{ componentField }" name="password">
            <FormItem class="my-8">
              <FormLabel>Password</FormLabel>
              <FormControl>
                <Input
                  type="password"
                  placeholder="Enter your password"
                  v-bind="componentField"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <Button
            type="submit"
            onsubmit="onSubmit"
            class="w-full bg-gradient-to-r from-blue-400 to-teal-500 text-white py-2 rounded-lg mt-3"
          >
            Login
          </Button>
        </form>
      </CardContent>
    </Card>
  </div>
</template>

<script setup>
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  CardFooter,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { toast } from "@/components/ui/toast";

import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

const formSchema = toTypedSchema(
  z.object({
    username: z.string().min(2).max(50),
    password: z.string().min(8).max(50),
  })
);

const { handleSubmit } = useForm({
  validationSchema: formSchema,
});

const onSubmit = handleSubmit((values) => {
  toast({
    title: "Login Success!",
    description: JSON.stringify(values, null, 2),
    variant: "destructive",
  });
});
</script>

<style scoped>
.form-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
</style>
