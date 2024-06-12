import { useState, useEffect } from "react";

const NavBar = () => {
  return(
    <div className="h-[10%] w-full px-20 bg-[#151720] flex items-center justify-between text-white">
      <div className="flex gap-x-5">
        <div className="logo"></div>
        <div className="">Movies</div>
        <div className="">Cinemas</div>
        <div className="">About us</div>
      </div>
      <div className="flex gap-x-5">
        <div className="border-2 rounded-lg border-[#03C04A] p-1">Sign in/Sign up</div>
        <div className="rounded-lg bg-[#03C04A] p-1">Buy ticket</div>
      </div>
    </div>
  )
};

export default NavBar;