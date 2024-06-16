import useAuthUser from "react-auth-kit/hooks/useAuthUser";
import { useEffect, useState } from "react";
import Button from "../components/button";
import UserService from "../services/UserService";
import TicketService from "../services/TicketService";
import Ticket from "../models/ticket";
import Show from "../models/show";
import Movie from "../models/movie";
import User from "../models/user";
import CustomDate from "../utils/date";
import useIsAuthenticated from 'react-auth-kit/hooks/useIsAuthenticated'

const UserInfo = () => {
    const user = useAuthUser();
    const isAuthenticated = useIsAuthenticated()
    const [userProfile, setUserProfile] = useState<User>();
    const [Update, setUpdate] = useState(false);
    const [newGender, setNewGender] = useState("");
    const [newPhone, setNewPhone] = useState("");
    const [newDob, setNewDob] = useState("");
    const [newName, setNewName] = useState("");
    const [info, setInfo] = useState<{ text: string, color: string }>();

    useEffect(() => {
        UserService
            .getProfile()
            .then((res) => {
                setUserProfile(res.data);
            })
    }, [user]);

    const updateInfo = {
        date_of_birth: newDob,
        gender: newGender,
        name: newName,
        phone: newPhone
    }

    const handleUpdate = () => {
        setUpdate(false);
        UserService
            .update(updateInfo)
            .then(() => {
                setInfo({
                    text: "Update successfully",
                    color: "bg-[#03C04A]"
                });
            })
            .catch((error) => {
                console.log(error);
            })
        console.log(updateInfo);
    }
    if(!isAuthenticated){
        return(
            <div>
                
            </div>
        )
    }
    else return (
        <div>
            <p className="w-full text-center text-black text-[25px] font-semibold font-Montserrat pt-2">Profile</p>
            {!Update && (
                <div>
                    <CurrentUserInfo user={userProfile!} />
                    <div className="w-full flex justify-center pb-2">
                        <Button text="Update" hollow={false} onClick={() => setUpdate(true)} />
                    </div>
                </div>
            )}
            {Update && (
                <form className="max-w-[1000px] h-max border-2 border-black mx-auto rounded-[10px] flex flex-col justify-center">
                    <div className="p-2">
                        <p className="text-black text-lg font-medium font-Montserrat text-left">Username</p>
                        <input className="w-2/3 border-2 px-5 py-2 rounded-[10px] border-black bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="text" placeholder="New name" onChange={(e) => setNewName(e.target.value)} />
                    </div>
                    <div className="p-2">
                        <p className="text-black text-lg font-medium font-Montserrat text-left">Phone</p>
                        <input className="w-2/3 border-2 px-5 py-2 rounded-[10px] border-black bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="text" placeholder="New phone" onChange={(e) => setNewPhone(e.target.value)} />
                    </div>
                    <div className="p-2">
                        <p className="text-black text-lg font-medium font-Montserrat text-left">Date of birth</p>
                        <input className="w-2/3 border-2 px-5 py-2 rounded-[10px] border-black bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="date" placeholder="New date of birth" onChange={(e) => setNewDob(e.target.value)} />
                    </div>
                    <div className="p-2">
                        <p className="text-black text-lg font-medium font-Montserrat text-left">Gender</p>
                        <select className="w-2/3 border-2 px-5 py-2 rounded-[10px] border-black bg-[#FDF7DC] focus:border-green-400 focus:outline-none" onChange={(e) => setNewGender(e.target.value)}>
                            <option value="None">None</option>
                            <option value="Male">Male</option>
                            <option value="Female">Female</option>
                            <option value="Other">Other</option>
                        </select>
                    </div>
                    <div className="w-full flex justify-center pb-2">
                        <div className="px-2">
                            <Button text="Save" hollow={false} onClick={() => handleUpdate()} />
                        </div>
                        <div>
                            <Button text="Cancel" hollow={false} onClick={() => setUpdate(false)} />
                        </div>
                    </div>
                </form>
            )}

        </div>
    )
}

const CurrentUserInfo: React.FC<{ user: User }> = ({ user }) => {
    const dob = CustomDate.format(user.date_of_birth);
    return (
        <div className="max-w-[1000px] h-max border-2 border-black mx-auto rounded-[10px]">
            <p className="text-black text-xl font-medium font-Montserrat text-left p-2">Username: {user.name}</p>
            <p className="text-black text-xl font-medium font-Montserrat text-left p-2">Email: {user.email}</p>
            <p className="text-black text-xl font-medium font-Montserrat text-left p-2">Phone number: {user.phone}</p>
            <p className="text-black text-xl font-medium font-Montserrat text-left p-2">Date of birth: {dob}</p>
            <p className="text-black text-xl font-medium font-Montserrat text-left p-2">Gender: {user.gender}</p>
        </div>
    )
}

const UserTicket = () => {
    const [tickets, setTickets] = useState<Ticket[]>([]);
    const [shows, setShows] = useState<Show[]>([]);
    const [movies, setMovies] = useState<Movie[]>([]);
    const [startTime, setStartTime] = useState<Date[]>();
    const [endTime, setEndTime] = useState<Date[]>();

    const fetchTickets = async () => {
        const resTicket = await TicketService.getByUser();
        const setData = async () => {
            setTickets(resTicket.data);
            setShows(resTicket.data.map((ticket) => ticket.show));
            setMovies(shows.map((show) => show.movie));
            setStartTime(shows.map((show) => new Date(show.startTime)));
            setEndTime(shows.map((show) => new Date(show.endTime)));
        }
        setData();
    }
    useEffect(() => {
        fetchTickets();
    }, []);
    // movies.map((movie, index) => {console.log(movie)});
    return (
        <div>
            <p className="w-full text-center text-black text-[25px] font-semibold font-Montserrat pt-2">Ticket</p>
            <div className="max-w-[1000px] h-max border-2 border-black mx-auto rounded-[10px]">
                {
                    tickets.map((ticket, index) => (
                        <div className="flex p-2" key={index}>
                            {/* <div>
                                <img src={movies[index].poster} alt={movies[index].title} className="w-32 h-48" />
                            </div>    
                            <div>
                                <p className="text-black text-xl font-medium font-Montserrat text-left p-2">{movies[index].title}</p>
                                <p className="text-black text-xl font-medium font-Montserrat text-left p-2">{movies[index].rated}</p>
                            </div> */}
                            <div>
                                <p className="text-black text-xl font-medium font-Montserrat text-left p-2">{shows[index].auditorium.cinema.name}</p>
                                <p className="text-black text-xl font-medium font-Montserrat text-left p-2">{shows[index].auditorium.name}</p>
                                <p className="text-black text-xl font-medium font-Montserrat text-left p-2">{shows[index].date + " " + shows[index].startTime + "-" + shows[index].endTime}</p>
                            </div>
                            <div>
                                <p className="text-black text-xl font-medium font-Montserrat text-left p-2">Seat: {ticket.seat_number}</p>
                                <p className="text-black text-xl font-medium font-Montserrat text-left p-2">Total: 50000VND</p>
                            </div>
                        </div>
                    ))
                }
            </div>
        </div>
    )
}

const Profile = () => {
    return (
        <div className='max-w-[1040px] h-full mx-auto bg-[#FDF7DC]'>
            <UserInfo />
            <UserTicket />
        </div>
    );
}

export default Profile;