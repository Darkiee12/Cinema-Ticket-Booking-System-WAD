import {Link} from "react-router-dom";
import Logo from "../assets/CIneU.png";
import Button from "./button";
import { useEffect, useState } from "react";
import UserService from "../services/UserService";
const NavBar = () => {
  
  const [userName, setUserName] = useState("");
  const [token, setToken] = useState("");

  useEffect(() => {
    UserService
    .getProfile()
    .then((response) => {
      console.log(response.data);
      setUserName(response.data.name);
    }).catch((error) => {
      console.log(error);
    });
    setToken(localStorage.getItem("token") || "");
    console.log("Token: "+token);
  }, []);  

  return(
    <div className="h-[10%] w-full px-10 bg-[#151720] flex items-center justify-between text-white font-Montserrat">
      <div className="flex gap-x-5 items-center">
        <Link className="logo" to="/">
          <img src={Logo} alt="CineU logo" className="w-[100px] h-[100px] rounded-[100px]" />
        </Link>
        <button className="text-lg font-semibold transition-all duration-[0.3s] ease-[ease-in-out] hover:text-[#03C04A]"><Link to="/movies" >Movies</Link></button>
        <button className="text-lg font-semibold transition-all duration-[0.3s] ease-[ease-in-out] hover:text-[#03C04A]">Cinemas</button>
        <button className="text-lg font-semibold transition-all duration-[0.3s] ease-[ease-in-out] hover:text-[#03C04A]">About us</button>
      </div>
      <div className="flex gap-x-5">
        {userName && <Button text={`Hello, ${userName}`} hollow={true} onClick={() => window.location.href = "/profile"} /> }
        {!userName && <Button text="Sign in/Sign up" hollow={true} onClick={() => window.location.href = "/login"} />}
        <Button text="Buy ticket" hollow={false} onClick={() => window.location.href = "/movies"} />
      </div>
    </div>
  )
};

export default NavBar;