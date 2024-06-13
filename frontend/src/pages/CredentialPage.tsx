import { useState } from "react";
import { Credential } from "../models/user";
import UserService from "../services/UserService";
import useSignIn from 'react-auth-kit/hooks/useSignIn';
import Button from "../components/button";
const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const signIn = useSignIn();
  const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
  }

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  }

  const handleLogin = () => {
    if (!emailRegex.test(email)) {
      alert("Invalid email")
      return;
    }
    const credentials: Credential = {
      email,
      password
    }
    UserService
      .login(credentials)
      .then((response) => {
        signIn({
          auth: {
            token: response.data.token
          },
          userState: {
            email: email,
            expiry: response.data.expiry
          }
        })
      }).catch((error) => {
        console.log(error);
      });
  }
  return (
    <div className="w-full px-10">
      <div className="w-full pt-5">
        <form className="w-full">
          <p className="w-full text-center font-bold text-sm md:text-2xl">Sign in</p>
          <div className="md:px-48 sm:px-12 px-10 py-3">
            <input className="w-full border-4 px-5 py-2 rounded-[10px] border-[#03C04A] bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="email" placeholder="Email" value={email} onChange={handleEmailChange} />
          </div>
          <div className="md:px-48 sm:px-12 px-10">
            <input className="w-full border-4 px-5 py-2 rounded-[10px] border-[#03C04A] bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="password" placeholder="Password" value={password} onChange={handlePasswordChange} />
          </div>
          <div className="w-full flex justify-center items-center py-2">
            <Button text="Sign in" hollow={false} onClick={handleLogin} />
          </div>
        </form>
      </div>
    </div>
  )
}

const SignIn = () => {
  return(
    <div>

    </div>
  )

}

const CredentialPage = () => {
  return(
    <div className="w-full h-full bg-[#FDF7DC]">
      <Login />
      <p className="py-5 w-full text-center">Or if you do not have an account</p>
      <SignIn />
    </div>
  )
}
export default CredentialPage;