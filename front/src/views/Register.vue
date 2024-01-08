<template>
  <div class="form-container">
    <Card class="w-[355px]">
      <CardHeader
        class="flex flex-col items-center justify-center text-center bg-gradient-to-r from-blue-400 to-teal-500 text-white py-8 rounded-t-lg"
      >
        <CardTitle
          class="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
          >Register!
        </CardTitle>
        <CardDescription>
          Create your account and start chatting 🧑‍💻
        </CardDescription>
      </CardHeader>
      <CardContent class="mt-3">
        <form @submit="onSubmit">
          <!-- Name -->
          <FormField v-slot="{ componentField }" name="name">
            <FormItem class="my-8">
              <FormLabel>Name</FormLabel>
              <FormControl>
                <Input
                  class="text-xs muted-foreground font-thin"
                  type="text"
                  placeholder="Enter your name"
                  v-bind="componentField"
                  v-model="formRegister.name"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <!-- Username -->
          <FormField v-slot="{ componentField }" name="username">
            <FormItem class="my-8">
              <FormLabel>Username</FormLabel>
              <FormControl>
                <Input
                  class="text-xs muted-foreground font-thin"
                  type="text"
                  placeholder="Enter your username"
                  v-bind="componentField"
                  v-model="formRegister.username"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <!-- Email -->
          <FormField v-slot="{ componentField }" name="email">
            <FormItem class="my-8">
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input
                  class="text-xs muted-foreground font-thin"
                  type="email"
                  placeholder="Enter your email"
                  v-bind="componentField"
                  v-model="formRegister.email"
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
                  class="text-xs muted-foreground font-thin"
                  type="password"
                  placeholder="Enter your password"
                  v-bind="componentField"
                  v-model="formRegister.password"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <!-- SIGNUP BUTTON-->
          <Button
            type="submit"
            onsubmit="onSubmit"
            class="w-full bg-gradient-to-r from-blue-400 to-teal-500 text-white py-2 rounded-lg my-3"
          >
            <!-- loader -->
            <!-- TODO: configure loader -->
            <Loader2
              v-if="false"
              class="animate-spin mr-2"
              size="20"
              color="white"
            />
            Sign up
          </Button>
          <Separator class="my-4" />
          <!-- already have an account? Login -->
          <div class="flex justify-center">
            <Label class="text-sm text-muted-foreground font-thin">
              Already have an account?
            </Label>
            <!-- TODO: add router link inside the label to go to "/register" route -->
            <Label class="text-sm text-blue-400 ml-1" @click="$router.push('/')"
              >Login</Label
            >
          </div>
        </form>
      </CardContent>
    </Card>
  </div>
</template>

<script setup>
import { Button } from "@/components/ui/button";
import { Loader2 } from "lucide-vue-next";
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
import { Separator } from "@/components/ui/separator";

//FORMULARIO
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

//STORAGE
import { useLogin } from "@/composables/useLogin";

const { formRegister } = useLogin();

const formSchema = toTypedSchema(
  z.object({
    username: z.string().min(2).max(50),
    name: z.string().min(2).max(50),
    email: z.string().email(),
    password: z.string().min(8).max(50),
  })
);

const { handleSubmit } = useForm({
  validationSchema: formSchema,
});

const onSubmit = handleSubmit((values) => {
  //guardo los valores en el storage
  formRegister.username = values.username;
  formRegister.name = values.name;
  formRegister.email = values.email;
  formRegister.password = values.password;

  //alerta
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
