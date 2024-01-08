import { defineStore } from "pinia";

export const useLoginStore = defineStore("form", {
  state: () => ({
    formLogin: {
      username: "",
      password: "",
    },
  }),
  actions: {
    setFormLogin(formLogin) {
      this.formLogin = formLogin;
    },
  },
});
