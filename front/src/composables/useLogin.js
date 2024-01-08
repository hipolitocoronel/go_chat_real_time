import { useLoginStore } from "@/storage/loginStore";
import { storeToRefs } from "pinia";

export const useLogin = () => {
  const loginStore = useLoginStore();

  const { formLogin, formRegister } = storeToRefs(loginStore);

  return {
    //properties
    formLogin,
    formRegister,

    //methods
  };
};
