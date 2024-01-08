import { defineStore } from "pinia";

export const useLoginStore = defineStore("form", {
  state: () => ({
    username: "",
    password: "",
  }),
  actions: {
    setUsername(username) {
      this.username = username;
    },
    setPassword(password) {
      this.password = password;
    },
  },
});
