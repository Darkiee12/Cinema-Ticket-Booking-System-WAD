import useIsAuthenticated from 'react-auth-kit/hooks/useIsAuthenticated';
import { useEffect, useState } from 'react';
import Button from '../components/button';
import UserService from '../services/UserService';
import TicketService from '../services/TicketService';
import Ticket from '../models/ticket';
import Show from '../models/show';
import Movie from '../models/movie';
import User from '../models/user';
import CustomDate from '../utils/date';
import VisibilityIcon from '@mui/icons-material/Visibility';
import VisibilityOffIcon from '@mui/icons-material/VisibilityOff';
const UserInfo: React.FC<{ onUsernameUpdate: (username: string) => void }> = ({
  onUsernameUpdate,
}) => {
  const [user, setUser] = useState<User>();
  const [Update, setUpdate] = useState(false);

  useEffect(() => {
    UserService.getProfile().then((res) => {
      setUser(res.data);
    });
  }, [Update]);
  return (
    <div>
      <p className="w-full text-center text-black text-[25px] font-semibold font-Montserrat pt-2">
        Profile
      </p>

      {user &&
        (!Update ? (
          <div className="max-w-[1000px] h-max border-2lack mx-auto rounded-[10px]">
            <UserDetail user={user} />
            <div className="w-full flex justify-center pb-2">
              <Button
                hollow={false}
                onClick={() => setUpdate(true)}
              >Update</Button>
            </div>
          </div>
        ) : (
          <div>
            <UpdateInfo
              user={user}
              setUpdate={setUpdate}
              onUsernameUpdate={onUsernameUpdate}
            />
          </div>
        ))}
    </div>
  );
};

const UserDetail: React.FC<{ user: User }> = ({ user }) => {
  const [hiddenPhone, setHiddenPhone] = useState<boolean>(true);
  const [hiddenEmail, setHiddenEmail] = useState<boolean>(true);
  const dob = CustomDate.format(user.date_of_birth);
  return (
    <div className="w-full overflow-x-auto">
      <table className="table-auto w-full border-collapse">
        <tbody>
          <tr>
            <td className="p-2">Username:</td>
            <td className="p-2">{user.name}</td>
            <td className="p-2"></td>
          </tr>
          <tr>
            <td className="p-2">Email:</td>
            <td className="p-2">
              {hiddenEmail
                ? "*".repeat((user.email.split('@')[0]).length)+"@"+user.email.split('@')[1]
                : user.email}
            </td>
            <td className="p-2">
              <button
                onClick={() => {
                  setHiddenEmail(!hiddenEmail);
                }}
              >
                {hiddenEmail ? <VisibilityOffIcon /> : <VisibilityIcon />}
              </button>
            </td>
            <td className="p-2"></td>
          </tr>
          <tr>
            <td className="p-2">Phone number:</td>
            <td className="p-2">
              {hiddenPhone
                ? '*'.repeat(user.phone.length - 4) + user.phone.slice(-4)
                : user.phone}
            </td>
            <td className="p-2">
              <button
                onClick={() => {
                  setHiddenPhone(!hiddenPhone);
                }}
              >
                {hiddenPhone ? <VisibilityOffIcon /> : <VisibilityIcon />}
              </button>
            </td>
          </tr>
          <tr>
            <td className="p-2">Date of birth:</td>
            <td className="p-2">{dob}</td>
            <td className="p-2"></td>
          </tr>
          <tr>
            <td className="p-2">Gender:</td>
            <td className="p-2">{user.gender}</td>
            <td className="p-2"></td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

const UpdateInfo: React.FC<{
  user: User;
  setUpdate: (value: boolean) => void;
  onUsernameUpdate: (username: string) => void;
}> = ({ user, setUpdate, onUsernameUpdate }) => {
  const [newGender, setNewGender] = useState('');
  const [newPhone, setNewPhone] = useState('');
  const [newDob, setNewDob] = useState('');
  const [newName, setNewName] = useState('');
  const [alert, setAlert] = useState<{ text: string; color: string }>();
  const updateInfo = {
    date_of_birth: newDob || user.date_of_birth,
    gender: newGender || user.gender,
    name: newName || user.name,
    phone: newPhone || user.phone,
  };

  const handleUpdate = () => {
    setUpdate(false);
    UserService.update(updateInfo)
      .then(() => {
        onUsernameUpdate(updateInfo.name);
        setAlert({ text: 'Update successfully', color: 'text-[#03C04A]' });
      })
      .catch((error) => {
        console.log(error);
      });
    console.log(updateInfo);
  };
  return (
    <form className="max-w-[1000px] h-max border-2lack mx-auto rounded-[10px] flex flex-col justify-center">
      {alert && (
        <p className={`w-full text-center ${alert.color}`}>{alert.text}</p>
      )}
      <div className="p-2">
        <p className="text-black text-lg font-medium font-Montserrat text-left">
          Username
        </p>
        <input
          className="w-2/3 border-2 px-5 py-2 rounded-[10px]lack bg-[#FDF7DC] focus:border-green-400 focus:outline-none"
          type="text"
          placeholder={`Your current username is ${user.name}`}
          onChange={(e) => setNewName(e.target.value)}
        />
      </div>
      <div className="p-2">
        <p className="text-black text-lg font-medium font-Montserrat text-left">
          Phone
        </p>
        <input
          className="w-2/3 border-2 px-5 py-2 rounded-[10px]lack bg-[#FDF7DC] focus:border-green-400 focus:outline-none"
          type="text"
          placeholder="New phone"
          onChange={(e) => setNewPhone(e.target.value)}
        />
      </div>
      <div className="p-2">
        <p className="text-black text-lg font-medium font-Montserrat text-left">
          Date of birth (yyyy-mm-dd)
        </p>
        <input
          className="w-2/3 border-2 px-5 py-2 rounded-[10px]lack bg-[#FDF7DC] focus:border-green-400 focus:outline-none"
          type="date"
          placeholder="New date of birth"
          onChange={(e) => setNewDob(e.target.value)}
        />
      </div>
      <div className="p-2">
        <p className="text-black text-lg font-medium font-Montserrat text-left">
          Gender
        </p>
        <select
          className="w-2/3 border-2 px-5 py-2 rounded-[10px]lack bg-[#FDF7DC] focus:border-green-400 focus:outline-none"
          onChange={(e) => setNewGender(e.target.value)}
        >
          <option value="None">None</option>
          <option value="Male">Male</option>
          <option value="Female">Female</option>
          <option value="Other">Other</option>
        </select>
      </div>
      <div className="w-full flex justify-center pb-2">
        <div className="px-2">
          <Button hollow={false} onClick={() => handleUpdate()}>Save</Button>
        </div>
        <div>
          <Button
            hollow={false}
            onClick={() => setUpdate(false)}
          >Cancel</Button>
        </div>
      </div>
    </form>
  );
};

const UserTicket = () => {
  const [tickets, setTickets] = useState<Ticket[]>([]);
  const [shows, setShows] = useState<Show[]>([]);
  const [movies, setMovies] = useState<Movie[]>([]);
  const [startTime, setStartTime] = useState<Date[]>([]);
  const [endTime, setEndTime] = useState<Date[]>([]);

  const fetchTickets = async () => {
    const resTicket = await TicketService.getByUser();
    const setData = async () => {
      setTickets(
        resTicket.data.sort((a, b) => {
          const dateA = new Date(a.show.date).getTime();
          const dateB = new Date(b.show.date).getTime();
          if (dateA !== dateB) {
            return dateA - dateB;
          }
          if (a.show.id !== b.show.id) {
            return parseInt(a.show.id) - parseInt(b.show.id);
          }

          return a.seat_number - b.seat_number;
        }),
      );
      setShows(resTicket.data.map((ticket: any) => ticket.show));
      resTicket.data.map((ticket: any) => {
        setMovies((movies) => [...movies, ticket.show.movie]);
        setStartTime((startTime) => [
          ...startTime,
          new Date(ticket.show.startTime),
        ]);
        setEndTime((endTime) => [...endTime, new Date(ticket.show.startTime)]);
      });
    };
    await setData();
  };
  useEffect(() => {
    fetchTickets();
  }, []);

  return (
    <div>
      <p className="w-full text-center text-black text-[25px] font-semibold font-Montserrat pt-2">
        Tickets
      </p>
      <div className="max-w-[1000px] h-max border-2lack mx-auto rounded-[10px]">
        {tickets.map((ticket, index) => (
          <div className="flex p-2" key={index}>
            <div>
              <img
                src={movies[index].poster}
                alt={movies[index].title}
                className="w-32 h-48"
              />
            </div>
            <div>
              <p className="text-black text-xl font-medium font-Montserrat text-left p-2">
                {movies[index].title}
              </p>
              <p className="text-black text-xl font-medium font-Montserrat text-left p-2">
                {movies[index].rated}
              </p>
            </div>
            <div>
              <p className="text-black text-xl font-medium font-Montserrat text-left p-2">
                {shows[index].auditorium.cinema.name}
              </p>
              <p className="text-black text-xl font-medium font-Montserrat text-left p-2">
                {shows[index].auditorium.name}
              </p>
              <p className="text-black text-xl font-medium font-Montserrat text-left p-2">
                {shows[index].date +
                  ' ' +
                  shows[index].startTime +
                  '-' +
                  shows[index].endTime}
              </p>
            </div>
            <div>
              <p className="text-black text-xl font-medium font-Montserrat text-left p-2">
                Seat: {ticket.seat_number}
              </p>
              <p className="text-black text-xl font-medium font-Montserrat text-left p-2">
                Total: {(50_000).toLocaleString('vi-VN', {
                style: 'currency',
                currency: 'VND'
              })}
              </p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

const Profile: React.FC<{ onUsernameUpdate: (username: string) => void }> = ({
  onUsernameUpdate,
}) => {
  const isAuthenticated = useIsAuthenticated();

  if (!isAuthenticated) {
    setTimeout(() => {
      window.location.href = '/login';
    }, 2000);
    return (
      <div className="w-full text-center">
        You need to login to view this page
      </div>
    );
  } else
    return (
      <div className="max-w-[1040px] h-full mx-auto bg-[#FDF7DC]">
        <UserInfo onUsernameUpdate={onUsernameUpdate} />
        <UserTicket />
      </div>
    );
};
export default Profile;
