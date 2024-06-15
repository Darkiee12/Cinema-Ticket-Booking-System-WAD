import { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import Logo from "../assets/CIneU.png";
import Button from '../components/button';
import useIsAuthenticated from 'react-auth-kit/hooks/useIsAuthenticated'
import useSignOut from 'react-auth-kit/hooks/useSignOut';
import useAuthUser from 'react-auth-kit/hooks/useAuthUser';

const NavBar = () => {
  const [userName, setUserName] = useState<string>();
  const [showModal, setShowModal] = useState<boolean>(false);
  const isAuthenticated = useIsAuthenticated()
  const auth = useAuthUser();
  const navigate = useNavigate();
  
 
  useEffect(() => {
    if(isAuthenticated){
      setUserName((auth as any).name);
    } else{
      navigate("/login")
    }
  }, []);
  

  

  return (
    <div className="h-[10%] w-full px-10 bg-[#151720] flex items-center justify-between text-white font-Montserrat">
      <div className="flex gap-x-5 items-center">
        <Link className="logo" to="/">
          <img src={Logo} alt="CineU logo" className="w-[100px] h-[100px] rounded-[100px]" />
        </Link>
        <button className="text-lg font-semibold transition-all duration-[0.3s] ease-[ease-in-out] hover:text-[#03C04A]"><Link to="/movies">Movies</Link></button>
        <button className="text-lg font-semibold transition-all duration-[0.3s] ease-[ease-in-out] hover:text-[#03C04A]">Cinemas</button>
        <button className="text-lg font-semibold transition-all duration-[0.3s] ease-[ease-in-out] hover:text-[#03C04A]">About us</button>
      </div>
      <div className="flex gap-x-5">
        {isAuthenticated && userName ? (
          <div>
            <Button text={`${userName}`} hollow={true} onClick={() => setShowModal(!showModal)} />
            {showModal && (<div className='z-50 absolute'><UserMenu /></div>)}
          </div>
        ) : (
          <Button text="Sign in/Sign up" hollow={true} onClick={() => window.location.href = "/login"} />
        )}
        <Button text="Buy ticket" hollow={false} onClick={() => window.location.href = "/movies"} />
      </div>
    </div>
  );
};

const UserMenu = () => {
  const signOut = useSignOut();
  const handleSignOut = () => {
    signOut();
    window.location.href = "/login";
  };
  const Menu = [
    <Button text="Profile" hollow={true} onClick={() => window.location.href = "/profile"} />,
    <Button text="Sign out" hollow={true} onClick={handleSignOut} />
  ]
  
  return(
    <div className='w-full bg-[#151720] p-4 shadow-lg mt-1 flex justify-center items-center'>
      <ul>
        {Menu.map((item, index) => (
          <li key={index} className='my-1'>{item}</li>
        ))}
      </ul>
    </div>
  )
}
export default NavBar;
