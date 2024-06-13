import { Link } from "react-router-dom";
const NavBar = () => {
  return(
    <div className="h-[10%] w-full px-10 bg-[#151720] flex items-center justify-between text-white font-Montserrat">
      <div className="flex gap-x-5 items-center">
        <div className="logo">
          <img src="./src/assets/CIneU.png" alt="CineU logo" className="w-[100px] h-[100px] rounded-[100px]" />
        </div>
        <button className="text-lg font-semibold"><Link to="/movies" >Movies</Link></button>
        <button className="text-lg font-semibold">Cinemas</button>
        <button className="text-lg font-semibold">About us</button>
      </div>
      <div className="flex gap-x-5">
        <div className="border-2 rounded-[10px] border-[#03C04A] p-2 text-lg font-semibold">Sign in/Sign up</div>
        <div className="rounded-[10px] bg-[#03C04A] p-2 text-lg font-semibold border-2 border-[#03C04A]">Buy ticket</div>
      </div>
    </div>
  )
};

export default NavBar;