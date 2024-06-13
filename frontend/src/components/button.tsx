
import React from 'react';

interface ButtonProps {
  hollow?: boolean;
  text?: string;
  onClick?: () => void;
}

const Button: React.FC<ButtonProps> = ({ hollow = false, text = '', onClick }) => {
  return (
    <button
      className={`rounded-lg p-1 text-white font-bold border-2 border-[#03C04A] focus:outline-none focus:ring-2 focus:ring-${hollow ? 'transparent' : '#03C04A'} py-2 px-5 focus:ring-opacity-50
                  ${hollow ? 'bg-transparent' : 'bg-[#03C04A]'} hover:bg-[#028A3D] active:bg-[#026A2F] transform transition duration-150 ease-in-out hover:scale-105 active:scale-95`}
      onClick={onClick}
      type="button"
    >
      {text}
    </button>
  );
};

export default Button;
