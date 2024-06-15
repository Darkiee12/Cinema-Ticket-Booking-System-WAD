import { useEffect, useState } from "react";
import { Credential, Register as RegisterAccount } from "../models/user";
import UserService from "../services/UserService";
import useSignIn from 'react-auth-kit/hooks/useSignIn';
import Button from "../components/button";
const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [info, setInfo] = useState<string>("");
  const signIn = useSignIn();
  const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
  }

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  }

  const handleLoginSuccess = () => {
    setInfo("Login success. Redirecting to homepage...");
    setTimeout(() => {
      window.location.href = "/";
    }, 2000);
  }
  const handleLogin = async () => {
    try {
      if (!emailRegex.test(email)) {
        alert("Invalid email");
        return;
      }

      const credentials: Credential = { email, password };
      const response = await UserService.login(credentials);
      const userProfile = (await UserService.getProfile(response.data.token));

      // Update authentication state
      signIn({
        auth: {
          token: response.data.token
        },
        userState: {
          name: userProfile.data.name,
          email: userProfile.data.email,
          expiry: response.data.expiry
        }
      });

      handleLoginSuccess();
    } catch (error: any) {
      if (error.response && error.response.status === 401) {
        setInfo("Either email or password is incorrect! Please try again!");
      } else {
        console.error("Login error:", error);
        setInfo("An error occurred during login. Please try again later.");
      }
    }
  };

  return (
    <div className="w-full">
      <form className="w-full">
        <p className="text-[25px] font-semibold font-Montserrat md:text-2xl text-center">Sign in</p>
        <p className="text-[#03C04A] text-sm italic font-medium font-Montserrat text-center pt-2">{info}</p>
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

  )
}

const Register = () => {
  const [email, setEmail] = useState("");
  const [validEmail, setValidEmail] = useState(false);
  const [emailFocus, setEmailFocus] = useState(false);
  const [password, setPassword] = useState("");
  // const [validPassword, setValidPassword] = useState(false);
  const [passwordFocus, setPasswordFocus] = useState(false);
  const [matchPassword, setMatchPassword] = useState("");
  const [validMatchPassword, setValidMatchPassword] = useState(false);
  const [matchPasswordFocus, setMatchPasswordFocus] = useState(false);
  const [name, setName] = useState("");
  const [validName, setValidName] = useState(false);
  const [nameFocus, setNameFocus] = useState(false);
  const [info, setInfo] = useState<string>("");
  const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
  // const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$/;
  const nameRegex = /^[a-zA-Z0-9]{3,}$/;

  useEffect(() => {
    const result = emailRegex.test(email);
    console.log(result);
    setValidEmail(result);
  }, [email]);

  useEffect(() => {
    // const result = passwordRegex.test(password);
    // setValidPassword(result);
    const matchPasswordResult = password === matchPassword;
    setValidMatchPassword(matchPasswordResult);
  }, [password, matchPassword]);

  useEffect(() => {
    const result = nameRegex.test(name);
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
      setInfo("Email is invalid. Please enter a valid email address.")
      return;
    }
    // if (!validPassword) {
    //   setInfo("Invalid password. Password has to be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, and one number.");
    //   return;
    // }
    if (!validMatchPassword) {
      setInfo("Password does not match");
      return;
    }
    if (!validName) {
      setInfo("Invalid name. Name has to be at least 3 characters long.");
      return;
    }
    const register: RegisterAccount = {
      email,
      name,
      password
    }
    UserService
      .register(register)
      .then(() => {
        setInfo("Register success. Please login to continue!");
        setTimeout(() => {
          window.scrollTo(0, 0);
        }, 2000);
      }).catch((error) => {
        console.log(error);
      });
  }

  return (
    <div className="w-full">
      <form className="w-full pt-2">
        <p className="text-[25px] font-semibold font-Montserrat md:text-2xl text-center pb-2">Register</p>
        {info.length > 0 && <p className="text-[#03C04A] text-sm italic font-medium font-Montserrat text-center pt-2">{info}</p>}
        <div className="md:px-48 sm:px-12 px-10 py-3">
          <input className="w-full border-4 px-5 py-2 rounded-[10px] border-[#03C04A] bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="email" placeholder="Username" value={name} onChange={handleUserNameChange} onFocus={() => setNameFocus(true)} onBlur={() => setNameFocus(false)} />
        </div>
        <div className="md:px-48 sm:px-12 px-10 pb-3 ">
          <input className="w-full border-4 px-5 py-2 rounded-[10px] border-[#03C04A] bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="email" placeholder="Email" value={email} onChange={handleEmailChange} onFocus={() => setEmailFocus(true)} onBlur={() => setEmailFocus(false)} />
        </div>
        <div className="md:px-48 sm:px-12 px-10 pb-3">
          <input className="w-full border-4 px-5 py-2 rounded-[10px] border-[#03C04A] bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="password" placeholder="Password" value={password} onChange={handlePasswordChange} onFocus={() => setPasswordFocus(true)} onBlur={() => setPasswordFocus(false)} />
        </div>
        <div className="md:px-48 sm:px-12 px-10">
          <input className="w-full border-4 px-5 py-2 rounded-[10px] border-[#03C04A] bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="password" placeholder="Re-enter password" value={matchPassword} onChange={handleMatchPasswordChange} onFocus={() => setMatchPasswordFocus(true)} onBlur={() => setMatchPasswordFocus(false)} />
        </div>
        <div className="w-full flex justify-center items-center py-2">
          <Button text="Register" hollow={false} onClick={handleRegister} />
        </div>
      </form>

    </div>
  );
}

const CredentialPage = () => {
  return (
    <div className="flex flex-col justify-center max-w-[1040px] h-screen mx-auto bg-[#FDF7DC]">
      <Login />
      <p className="text-black text-xl font-medium font-Montserrat text-center pt-2">Or if you do not have an account</p>
      <Register />
    </div>
  )
}
export default CredentialPage;