import { useEffect, useRef, useState } from "react";
import { Credential, Register as RegisterAccount } from "../models/user";
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
    <form className="w-full">
      <p className="text-[25px] font-semibold font-Montserrat md:text-2xl text-center">Sign in</p>
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
  )
}

const SignIn = () => {
  return(
    <div>

    </div>
  )

}

const Register = () => {
  const [email, setEmail] = useState("");
  const [validEmail, setValidEmail] = useState(false);
  const [emailFocus , setEmailFocus] = useState(false);

  const [password, setPassword] = useState("");
  const [validPassword, setValidPassword] = useState(false);
  const [passwordFocus, setPasswordFocus] = useState(false);

  const [matchPassword, setMatchPassword] = useState("");
  const [validMatchPassword, setValidMatchPassword] = useState(false);
  const [matchPasswordFocus, setMatchPasswordFocus] = useState(false);

  const [name, setName] = useState("");
  const [validName, setValidName] = useState(false);
  const [nameFocus, setNameFocus] = useState(false);

  const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
  const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$/;
  const nameRegex = /^[a-zA-Z0-9]{3,}$/;


  useEffect(() => {
    const result = emailRegex.test(email);
    console.log(result);
    setValidEmail(result);
  }, [email]);

  useEffect(() => {
    const result = passwordRegex.test(password);
    console.log(result);
    setValidPassword(result);
    const matchPasswordResult = password === matchPassword;
    setValidMatchPassword(matchPasswordResult);
  }, [password, matchPassword]);

  useEffect(() => {
    const result = nameRegex.test(name);
    console.log(result);
    setValidName(result);
  }, [name]);

  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
  }

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  }

  const handleUserNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value);
  }

  const handleMatchPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setMatchPassword(e.target.value);
  }

  const handleRegister = () => {
    if (!validEmail) {
      alert("Invalid email");
      return;
    }
    if (!validPassword) {
      alert("Invalid password");
      return;
    }
    if (!validMatchPassword) {
      alert("Password does not match");
      return;
    }
    if (!validName) {
      alert("Invalid name");
      return;
    }
    const register: RegisterAccount = {
      email,
      name,
      password
    }
    
    console.log (register);

    UserService
    .register(register)
    .then ((response) => {
      console.log(response.data);
      alert("Register success");
    }).catch((error) => {
      console.log(error);
    });
  }

  return (
    <form className="w-full pt-2">
      <p className="text-[25px] font-semibold font-Montserrat md:text-2xl text-center pb-2">Register</p>
      <div className="md:px-48 sm:px-12 px-10 py-3">
        <input className="w-full border-4 px-5 py-2 rounded-[10px] border-[#03C04A] bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="email" placeholder="Username" value={name} onChange={handleUserNameChange} onFocus={()=> setNameFocus(true)} onBlur={()=> setNameFocus(false)}/>
      </div>
      <div className="md:px-48 sm:px-12 px-10 pb-3 ">
        <input className="w-full border-4 px-5 py-2 rounded-[10px] border-[#03C04A] bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="email" placeholder="Email" value={email} onChange={handleEmailChange} onFocus={()=> setEmailFocus(true)} onBlur={()=> setEmailFocus(false)} />
      </div>
      <div className="md:px-48 sm:px-12 px-10 pb-3">
        <input className="w-full border-4 px-5 py-2 rounded-[10px] border-[#03C04A] bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="password" placeholder="Password" value={password} onChange={handlePasswordChange} onFocus={()=> setPasswordFocus(true)} onBlur={()=> setPasswordFocus(false)} />
      </div>
      <div className="md:px-48 sm:px-12 px-10">
        <input className="w-full border-4 px-5 py-2 rounded-[10px] border-[#03C04A] bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="password" placeholder="Re-enter password" value={matchPassword} onChange={handleMatchPasswordChange} onFocus={()=> setMatchPasswordFocus(true)} onBlur={()=> setMatchPasswordFocus(false)}/>
      </div>
      <div className="w-full flex justify-center items-center py-2">
        <Button text="Register" hollow={false} onClick={handleRegister} />
      </div>
    </form>
  );
}

const CredentialPage = () => {
  return(
    <div className="flex flex-col justify-center max-w-[1040px] h-screen mx-auto bg-[#FDF7DC]">
        <Login />
        <p className="text-black text-xl font-medium font-Montserrat text-center pt-2">Or if you do not have an account</p>
        <SignIn />
        <Register />
    </div>
  )
}
export default CredentialPage;