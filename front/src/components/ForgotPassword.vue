<template>
  <AlertDialog :open="dialogForgotPassword">
    <form @submit="onSubmit">
      <AlertDialogContent>
        <AlertDialogTitle class="text-center mb-3"
          >Forgot Password 🔓</AlertDialogTitle
        >
        <AlertDialogDescription>
          Please enter your email address to receive a verification code.
        </AlertDialogDescription>
        <!-- FORMULARIO -->

        <FormField v-slot="{ componentField }" name="email">
          <FormItem class="my-8">
            <FormLabel>Email</FormLabel>
            <FormControl>
              <Input
                class="text-xs muted-foreground font-thin"
                type="email"
                placeholder="Enter your email"
                v-bind="componentField"
                v-model="formForgotPassword.email"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <AlertDialogFooter>
          <AlertDialogCancel @click="dialogForgotPassword = false">
            Cancel
          </AlertDialogCancel>

          <Button
            type="submit"
            @click="onSubmit"
            class="w-[30%] bg-gradient-to-r from-blue-400 to-teal-500 text-white py-2 rounded-lg my-3"
          >
            Send
          </Button>
        </AlertDialogFooter>
      </AlertDialogContent>
    </form>
  </AlertDialog>
</template>
<script setup>
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";

import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { toast } from "@/components/ui/toast";
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

import { useLogin } from "@/composables/useLogin";

//FORMULARIO
import { useForm, validate } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

const formSchema = toTypedSchema(
  z.object({
    email: z.string().email(),
  })
);

const { handleSubmit } = useForm({
  validationSchema: formSchema,
});

const onSubmit = handleSubmit((values) => {
  console.log("entre submit forgot");
  //guardo los valores en el storage
  formForgotPassword.email = values.email;

  //alerta
  toast({
    title: "Email sent!",
    description: JSON.stringify(values, null, 2),
    variant: "destructive",
  });
});

const { dialogForgotPassword, formForgotPassword } = useLogin();
</script>

<style scoped></style>
