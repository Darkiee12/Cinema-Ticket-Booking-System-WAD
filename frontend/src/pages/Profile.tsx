import useAuthUser from "react-auth-kit/hooks/useAuthUser";
import { useState } from "react";
import Button from "../components/button";
import UserService from "../services/UserService";

const UserInfo = () => {

    const user = useAuthUser();
    const email = (user as any).email;

    const [Update, setUpdate] = useState(false);
    const [gender, setGender] = useState((user as any).gender);
    const [phone, setPhone] = useState((user as any).phone);
    const [dob, setDob] = useState((user as any).date_of_birth);
    const [name, setName] = useState((user as any).name);

    const updateInfo ={
        date_of_birth: dob,
        gender: gender,
        name: name,
        phone: phone
    }

    const handleUpdate = () => {
        setUpdate(false);
        UserService
        .update(updateInfo)
        .then(() => {
            (user as any).date_of_birth = dob;
            (user as any).gender = gender;
            (user as any).name = name;
            (user as any).phone = phone;
        })
        .catch((error) => {
            console.log(error);
        })
        console.log(updateInfo);
    }
    return(
        <div>
            <p className="w-full text-center text-black text-[25px] font-semibold font-Montserrat pt-2">Profile</p>
            {!Update && (
                <div className="max-w-[1000px] h-max border-2 border-black mx-auto rounded-[10px]">
                    <p className="text-black text-xl font-medium font-Montserrat text-left p-2">Username: {name}</p>
                    <p className="text-black text-xl font-medium font-Montserrat text-left p-2">Email: {email}</p>
                    <p className="text-black text-xl font-medium font-Montserrat text-left p-2">Phone number: {phone}</p>
                    <p className="text-black text-xl font-medium font-Montserrat text-left p-2">Date of birth: {dob}</p>
                    <p className="text-black text-xl font-medium font-Montserrat text-left p-2">Gender: {gender}</p>
                    <div className="w-full flex justify-center pb-2">
                        <Button text="Update" hollow={false} onClick={() => setUpdate(true)} />
                    </div>
                </div>
            )}
            {Update && (
                <form className="max-w-[1000px] h-max border-2 border-black mx-auto rounded-[10px] flex flex-col justify-center">
                    <div className="p-2">
                        <p className="text-black text-lg font-medium font-Montserrat text-left">Username</p>
                        <input className="w-2/3 border-2 px-5 py-2 rounded-[10px] border-black bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="text" placeholder="New name" onChange={(e)=>setName(e.target.value)} />
                    </div>
                    <div className="p-2">
                        <p className="text-black text-lg font-medium font-Montserrat text-left">Phone</p>
                        <input className="w-2/3 border-2 px-5 py-2 rounded-[10px] border-black bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="text" placeholder="New phone" onChange={(e)=>setPhone(e.target.value)} />
                    </div>
                    <div className="p-2">
                        <p className="text-black text-lg font-medium font-Montserrat text-left">Date of birth (yyyy-mm-dd)</p>
                        <input className="w-2/3 border-2 px-5 py-2 rounded-[10px] border-black bg-[#FDF7DC] focus:border-green-400 focus:outline-none" type="text" placeholder="New date of birth" onChange={(e)=>setDob(e.target.value)} />
                    </div>
                    <div className="p-2">
                        <p className="text-black text-lg font-medium font-Montserrat text-left">Gender</p>
                        <select className="w-2/3 border-2 px-5 py-2 rounded-[10px] border-black bg-[#FDF7DC] focus:border-green-400 focus:outline-none" onChange={(e) => setGender(e.target.value)}>
                            <option value="None">None</option>
                            <option value="Male">Male</option>
                            <option value="Female">Female</option>
                            <option value="Other">Other</option>
                        </select>
                    </div>
                    <div className="w-full flex justify-center pb-2">
                        <div className="px-2">
                            <Button text="Save" hollow={false}  onClick={() => handleUpdate()}/>
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

const UserTicket = () => {

}

const Profile = () => {
    return (
        <div className='max-w-[1040px] h-full mx-auto bg-[#FDF7DC]'>
            <UserInfo />
        </div>
    );
}

export default Profile;