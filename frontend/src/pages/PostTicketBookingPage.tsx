import { useEffect, useState } from 'react';
import useIsAuthenticated from 'react-auth-kit/hooks/useIsAuthenticated';
import { Link, useLocation } from 'react-router-dom';
import CustomDate from '../utils/date';
import Button from '../components/button';
import HomeIcon from '@mui/icons-material/Home';

const PostTicketBookingPage = () => {
  const isAuthenticated = useIsAuthenticated();
  const [data, setData] = useState<any>();
  const location = useLocation();
  useEffect(() => {
    if (location.state) {
      setData(location.state);
    } else {
      window.location.href = '/';
    }
  }, []);
  if (!isAuthenticated) {
    window.location.href = '/';
  } else
    return (
      <div className="w-full md:px-10">
        <div className='w-full bg-[#FDF7DC]'>
        <p className="w-full pt-5 font-Montserrat text-lg text-center font-bold">
          Thank you for choosing our service!
        </p>
        {data && (
          <div className="w-full flex px-2 lg:px-20 py-3">
          <div className="w-1/4 justify-center items-center flex">
            <img src={data.movie.poster} alt={data.movie.poster} className="w-full lg:px-10 aspect-auto" />
          </div>
          <div className="w-3/4">
            <div className="w-full flex flex-wrap justify-between">
              <div className="max-w-[33.33%] px-2 py-5">
                <p className="text-base lg:text-xl font-bold">{data.movie.title}</p>
                <p>{data.movie.rated}</p>
                <p>{data.movie.runtime} minutes</p>
              </div>
              <div className="px-2 py-5">
                <p className="text-base lg:text-xl font-bold">{data.show.auditorium.cinema.name}</p>
                <p>{data.show.auditorium.name}</p>
                <p>{CustomDate.format(data.show.date)}</p>
                <p>
                  {data.show.startTime.substring(0, 5)} ~{' '}
                  {data.show.endTime.substring(0, 5)}
                </p>
              </div>
              <div className="px-2 py-5">
              <div className='max-h-20 overflow-y-auto'>
              <p><span className='text-base lg:text-xl font-bold mr-1 w-full'>Seats:</span>{data.booked.join(", ")}</p>
              </div>
                
                <p><span className='text-base lg:text-xl font-bold mr-1 overflow-x-auto w-full'>Price:</span>{data.price.toLocaleString('vi-VN', {
                  style: 'currency',
                  currency: 'VND'
                })}</p>
              </div>
            </div>
            <div className="w-full flex justify-center items-center mt-2">
              <Button hollow={false}><Link to="/"><HomeIcon/>Home</Link></Button>
            </div>
          </div>
        </div>
        )}
        </div>
       
      </div>
    );
};

export default PostTicketBookingPage;
