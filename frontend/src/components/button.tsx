import React from 'react';

interface ButtonProps {
  hollow?: boolean;
  onClick?: () => void;
  children?: React.ReactNode;
}

const Button: React.FC<ButtonProps> = ({
  hollow = false,
  onClick,
  children
}) => {
  return (
    <button
      className={`rounded-lg p-1 text-white font-bold border-2 border-[#03C04A] focus:outline-none focus:ring-2 focus:ring-${hollow ? 'transparent' : '#03C04A'} mt-1 text-sm md:text-lg py-1 px-1 md:py-1 md:px-5 focus:ring-opacity-50
                  ${hollow ? 'bg-transparent' : 'bg-[#03C04A]'} hover:bg-[#028A3D] active:bg-[#026A2F] transform transition duration-150 ease-in-out hover:scale-105 active:scale-95`}
      onClick={onClick}
      type="button"
    >
      {children}
    </button>
  );
};

export default Button;
