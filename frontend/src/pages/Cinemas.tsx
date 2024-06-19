import CinemaService from "../services/CinemaService";

import { useState, useEffect } from "react";

import Cinema from "../models/cinema";

import { Link } from 'react-router-dom';
import Button from '../components/button';

const CinemaUnit: React.FC<{ cinema: Cinema }> = ({cinema}) => {
    return (
        <div className="">
          <div className="w-full">
            <img
              src={cinema.banner}
              alt={cinema.name}
              className="w-full min-h-[16rem]"
            />
          </div>
          <div>
            <p className="truncate text-xl font-bold text-center py-1">{cinema.name}</p>
          </div>
          <Link to={{ pathname: `/cinemas/${cinema.id}` }} state={cinema} className="flex justify-center items-center">
            <Button hollow={false}>Detail</Button>
          </Link>
        </div>
      );
}

const CinemaList: React.FC<{ cinemas: Cinema[] }> = ({cinemas}) => {
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        setTimeout(() => {
        setLoading(false);
        }, 2000);
    }, []);

    return (
        <div className="grid grid-cols-3 md:grid-cols-4 gap-3 md:gap-4 md:p-5 pt-2">
        {loading ? (
            <>
            {[...Array(10)].map((_, index) => (
                <div key={index} className="skeleton-row animate-pulse w-full">
                <div className="skeleton-img bg-[#f4e8b0] h-[16rem]"></div>
                <div className="skeleton-text w-3/4 h-[1rem] mt-2 bg-[#f4e8b0]"></div>
                <div className="skeleton-text w-1/2 h-[1rem] mt-2 bg-[#f4e8b0]"></div>
                <div className="skeleton-text rounded-lg p-1 w-1/2 h-[2rem] bg-[#f4e8b0] mt-2"></div>
                </div>
            ))}
            </>
        ) : (
            cinemas.map((cinema) => <CinemaUnit key={cinema.id} cinema={cinema} />)
        )}
        </div>
    );
}

const Cinemas = () => {
    const [cinemas, setCinemas] = useState<Cinema[]>([]);
    useEffect(() => {
      CinemaService
      .getAll()
      .then((res) => {
        setCinemas(res.data);
      });
    }, []);

    return (
        <div className="max-w-[1040px] h-screen mx-auto bg-[#FDF7DC]">
            <p className="w-full text-center text-black text-[25px] font-semibold font-Montserrat pt-2">Cinemas</p>
            <CinemaList cinemas={cinemas} />
        </div>
    );
}

export default Cinemas;