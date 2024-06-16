import React, { useEffect, useRef, useState } from 'react';
import { Link } from 'react-router-dom';
import Logo from '../assets/CIneU.png';
import Button from '../components/button';
import useIsAuthenticated from 'react-auth-kit/hooks/useIsAuthenticated';
import useSignOut from 'react-auth-kit/hooks/useSignOut';
import UserService from '../services/UserService';
import MovieIcon from '@mui/icons-material/Movie';
import HouseIcon from '@mui/icons-material/House';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';

const NavBar: React.FC<{ username: string }> = ({ username }) => {
  const [userName, setUserName] = useState<string>();
  const [showModal, setShowModal] = useState<boolean>(false);
  const isAuthenticated = useIsAuthenticated();
  useEffect(() => {
    if (username !== userName) setUserName(username);
  }, [username]);
  useEffect(() => {
    UserService.getProfile().then((res) => {
      setUserName(res.data.name);
    });
  }, []);
  return (
    <div className="h-[10%] w-full px-10 bg-[#151720] flex items-center justify-between text-white font-Montserrat">
      <div className="flex gap-x-5 items-center">
        <Link className="logo" to="/">
          <img
            src={Logo}
            alt="CineU logo"
            className="w-[100px] h-[100px] rounded-[100px]"
          />
        </Link>
        <Link to="/movies"> 
          <button className="text-lg font-semibold transition-all duration-[0.3s] ease-[ease-in-out] hover:text-[#03C04A]">Movies</button>
        </Link>
        <Link to="/cinemas"> 
          <button className="text-lg font-semibold transition-all duration-[0.3s] ease-[ease-in-out] hover:text-[#03C04A]">Cinemas</button>
        </Link>
        <button className="text-lg font-semibold transition-all duration-[0.3s] ease-[ease-in-out] hover:text-[#03C04A]">
          About us
        </button>
      </div>
      <div className="flex gap-x-5">
        {isAuthenticated && userName ? (
          <div>
            <Button
              text={`${userName}`}
              hollow={true}
              onClick={() => setShowModal(!showModal)}
            />
            {showModal && (
              <div className="z-50 absolute">
                <UserMenu />
              </div>
            )}
          </div>
        ) : (
          <Button
            text="Sign in/Sign up"
            hollow={true}
            onClick={() => (window.location.href = '/login')}
          />
        )}
        <Button
          text="Buy ticket"
          hollow={false}
          onClick={() => (window.location.href = '/movies')}
        />
      </div>
    </div>
  );
};

const NavBarMobile: React.FC<{ username: string }> = ({ username }) => {
  const [userName, setUserName] = useState<string>();
  const [showModal, setShowModal] = useState<boolean>(false);
  const [isMoviePressed, setIsMoviePressed] = useState(false);
  const [isCinemaPressed, setIsCinemaPressed] = useState(false);
  const [isProfilePressed, setIsProfilePressed] = useState(false);
  const buttons = [setIsMoviePressed, setIsCinemaPressed, setIsProfilePressed];
  const timeoutRef = useRef<number | null>(null);

  const handleTouchStart = (button: number) => {
    timeoutRef.current = window.setTimeout(() => {
      buttons[button](true);
    }, 50);
  };

  const handleTouchEnd = (button: number) => {
    if (timeoutRef.current) {
      clearTimeout(timeoutRef.current);
    }
    buttons[button](false);
  };

  useEffect(() => {
    if (username !== userName) setUserName(username);
  }, [username]);
  useEffect(() => {
    UserService.getProfile().then((res) => {
      setUserName(res.data.name);
    });
  }, []);
  return (
    <div className="bg-[#151720] w-full">
      <Link className="logo" to="/">
        <div className="w-full h-[5vh] flex justify-center items-center">
          <img src={Logo} alt="CineU logo" className="h-[200%]" />
        </div>
      </Link>
      <div className="w-full flex p-2">
        <Link
          to="/movies"
          className={`w-1/3 text-center text-white ${isMoviePressed ? 'bg-[#03C04A]' : 'bg-transparent'}`}
          onTouchStart={() => handleTouchStart(0)}
          onTouchEnd={() => handleTouchEnd(0)}
        >
          <MovieIcon />
        </Link>
        <Link
          to="/cinemas"
          className={`w-1/3 text-center text-white ${isCinemaPressed ? 'bg-[#03C04A]' : 'bg-transparent'}`}
          onTouchStart={() => handleTouchStart(1)}
          onTouchEnd={() => handleTouchEnd(1)}
        >
          <HouseIcon />
        </Link>
        <div
          className={`w-1/3 text-center text-white ${isProfilePressed ? 'bg-[#03C04A]' : 'bg-transparent'}`}
          onClick={() => setShowModal(!showModal)}
          onTouchStart={() => handleTouchStart(2)}
          onTouchEnd={() => handleTouchEnd(2)}
        >
          <AccountCircleIcon />
        </div>
      </div>
      {showModal && (
        <div className="w-full">
          <ProfileMenu
            setShowModal={setShowModal}
            showModal={showModal}
            username={userName!}
          />
        </div>
      )}
    </div>
  );
};

const ProfileMenu: React.FC<{
  setShowModal: (showModal: boolean) => void;
  showModal: boolean;
  username: string;
}> = ({ setShowModal, showModal, username }) => {
  const [isProfilePressed, setIsProfilePressed] = useState(false);
  const [isSignOutPressed, setIsSignOutPressed] = useState(false);
  const [isSignInPressed, setIsSignInPressed] = useState(false);
  const buttons = [
    setIsProfilePressed,
    setIsSignOutPressed,
    setIsSignInPressed,
  ];
  const timeoutRef = useRef<number | null>(null);
  const isAuthenticated = useIsAuthenticated();
  const signOut = useSignOut();
  const handleSignOut = () => {
    signOut();
    window.location.href = '/login';
  };
  const handleTouchStart = (button: number) => {
    timeoutRef.current = window.setTimeout(() => {
      buttons[button](true);
    }, 50);
  };

  const handleTouchEnd = (button: number) => {
    if (timeoutRef.current) {
      clearTimeout(timeoutRef.current);
    }
    buttons[button](false);
  };
  return isAuthenticated ? (
    <ul>
      <li className="text-center text-white w-full py-2">Hello {username}</li>
      <li
        className={`text-center text-white w-full py-2 ${isProfilePressed ? 'bg-[#03C04A]' : 'bg-transparent'}`}
        onTouchStart={() => handleTouchStart(0)}
        onTouchEnd={() => handleTouchEnd(0)}
      >
        <Link to="/profile" onClick={() => setShowModal(!showModal)}>
          Profile
        </Link>
      </li>
      <li
        className={`text-center text-white w-full py-2 ${isSignOutPressed ? 'bg-[#03C04A]' : 'bg-transparent'}`}
        onTouchStart={() => handleTouchStart(1)}
        onTouchEnd={() => handleTouchEnd(1)}
      >
        <button
          onClick={() => {
            setShowModal(!showModal);
            handleSignOut();
          }}
        >
          Sign out
        </button>
      </li>
    </ul>
  ) : (
    <div
      className={`text-center text-white w-full py-2 ${isSignInPressed ? 'bg-[#03C04A]' : 'bg-transparent'}`}
      onTouchStart={() => handleTouchStart(2)}
      onTouchEnd={() => handleTouchEnd(2)}
    >
      <Link to="/login" onClick={() => setShowModal(!showModal)}>
        Log in
      </Link>
    </div>
  );
};

const UserMenu = () => {
  const signOut = useSignOut();
  const handleSignOut = () => {
    signOut();
    window.location.href = '/login';
  };
  const Menu = [
    <Button
      text="Profile"
      hollow={true}
      onClick={() => (window.location.href = '/profile')}
    />,
    <Button text="Sign out" hollow={true} onClick={handleSignOut} />,
  ];

  return (
    <div className="w-full bg-[#151720] p-4 shadow-lg mt-1 flex justify-center items-center">
      <ul>
        {Menu.map((item, index) => (
          <li key={index} className="my-1">
            {item}
          </li>
        ))}
      </ul>
    </div>
  );
};
export default NavBar;
export { NavBarMobile };
