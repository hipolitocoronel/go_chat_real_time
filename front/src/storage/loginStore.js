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
