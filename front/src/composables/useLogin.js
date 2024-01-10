import { useLoginStore } from "@/storage/loginStore";
import { storeToRefs } from "pinia";

export const useLogin = () => {
  const loginStore = useLoginStore();

  const { formLogin, formRegister, dialogForgotPassword, formForgotPassword } =
    storeToRefs(loginStore);

  return {
    //properties
    formLogin,
    formRegister,
    dialogForgotPassword,
    formForgotPassword,

    //methods
  };
};
