import { defineStore } from "pinia";

export const useLoginStore = defineStore("form", {
  state: () => ({
    formLogin: {
      username: "",
      password: "",
    },
    formRegister: {
      name: "",
      username: "",
      email: "",
      password: "",
    },

    formForgotPassword: {
      email: "",
    },

    dialogForgotPassword: false,
  }),
  actions: {
    setFormLogin(formLogin) {
      this.formLogin = formLogin;
    },
    setFormRegister(formRegister) {
      this.formRegister = formRegister;
    },
  },
});
