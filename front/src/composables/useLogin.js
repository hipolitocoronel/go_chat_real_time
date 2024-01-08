import { useLoginStore } from "@/storage/loginStore";
import { storeToRefs } from "pinia";

export const useLogin = () => {
  const loginStore = useLoginStore();

  const { formLogin } = storeToRefs(loginStore);

  return {
    //properties
    formLogin,
    //methods
  };
};
