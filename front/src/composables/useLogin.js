import { useLoginStore } from "@/storage/loginStore";
import { storeToRefs } from "pinia";

export const useLogin = () => {
  const loginStore = useLoginStore();

  const { username, password } = storeToRefs(loginStore);

  return {
    //properties
    username,
    password,

    //methods
  };
};
