import  { useState, useEffect } from 'react';

import CinemaService from '../services/CinemaService';

import Cinema from "../models/cinema";

const CinemaDetail = () => {
    const [cinema, setCinema] = useState<Cinema>();

    const id = location.pathname.split('/').pop()?.toString();
    console.log("id",id);

    useEffect(() => {
        CinemaService.getById(id!)
        .then((res) => {
            console.log("data",res.data);
            setCinema(res.data);
        })
        .catch(() => {
            console.log('error');   
            window.location.href = '/';
        });
    }, []);

    console.log("cinema",cinema);

    return cinema && (
        <div className="max-w-[1040px] h-screen mx-auto bg-[#FDF7DC]">
            <p className="w-full text-center text-black text-[25px] font-semibold font-Montserrat pt-2">Cinema detail</p>
            <div className="flex items-center">
                <div className="w-1/3 flex justify-center items-center">
                    <img src={cinema.banner} alt={cinema.name} className="w-full aspect-auto" />
                </div>
                <div className="w-2/3 p-2">
                    <p className="text-xl font-bold">{cinema.name}</p>
                    <p className="text-lg">{cinema.address}</p>
                    <p className="text-lg">{cinema.phone_number}</p>
                    <p className="text-lg">{cinema.email}</p>
                    {/* <p className="text-lg">{cinema!.website}</p> */}
                    <p className="text-lg">Auditoriums: {cinema.capacity}</p>
                </div>
            </div>
        </div>
    );
}
export default CinemaDetail;