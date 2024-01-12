import { useLoginStore } from "@/storage/loginStore";
import { storeToRefs } from "pinia";
import { googleAuthCodeLogin } from "vue3-google-login";

export const useLogin = () => {
  const loginStore = useLoginStore();

  const { formLogin, formRegister, dialogForgotPassword, formForgotPassword } =
    storeToRefs(loginStore);

  const loginWithGoogle = (resp) => {
    googleAuthCodeLogin().then((response) => {
      console.log("Handle the response", response);
    });
    console.log(resp);
  };

  return {
    //properties
    formLogin,
    formRegister,
    dialogForgotPassword,
    formForgotPassword,

    //methods
    loginWithGoogle,
  };
};
