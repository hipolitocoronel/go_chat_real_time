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
      <CardContent class="mt-3">
        <!-- G00GLE BUTTON -->
        <Button
          class="w-full my-3 rounded-lg"
          variant="outline"
          @click="logInWithGoogle"
        >
          <img src="@/assets/google.png" alt="Google" class="mr-2 w-[8%]" />
          Continue with Google
        </Button>
        <div class="flex w-[42%] items-center">
          <Separator class="my-4" />
          <p class="mx-3 text-sm text-muted-foreground">OR</p>
          <Separator class="my-4" />
        </div>
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
                  v-model="formLogin.username"
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
                  v-model="formLogin.password"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          </FormField>
          <!-- LOGIN BUTTON-->
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
            Login
          </Button>
          <!-- Forgot your password? -->
          <div class="flex justify-center">
            <Label class="text-sm text-muted-foreground">
              Forgot your password?
            </Label>
          </div>
          <Separator class="my-4" />
          <!-- dont have an account? Register -->
          <div class="flex justify-center">
            <Label class="text-sm text-muted-foreground">
              Don't have an account?
            </Label>
            <!-- TODO: add router link inside the label to go to "/register" route -->
            <Label class="text-sm text-blue-400 ml-1">Register</Label>
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

const { formLogin } = useLogin();

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
  //guardo los valores en el storage
  formLogin.username = values.username;
  formLogin.password = values.password;

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
