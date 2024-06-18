# CIneU
This is a booking website that allow user to book movie ticket from all cinema within just one website. This project is made for "Web application development" course

![Logo](https://scontent.fsgn5-12.fna.fbcdn.net/v/t39.30808-6/448510917_4518499311707564_6026758865277843418_n.jpg?_nc_cat=103&ccb=1-7&_nc_sid=5f2048&_nc_ohc=iZl5UpbiNVYQ7kNvgGcXKsM&_nc_ht=scontent.fsgn5-12.fna&oh=00_AYDgB8k6AIibY5Y_uAs8cG6nCjSc1-HAXqt4tOa8dqmQJQ&oe=667623AF)


<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

## About The Project
![alt text](./readme%20images/image.png)

Our website will show the movie data from all cinema including movie's informations, movie's showtime from every cinema and allow user to book movie ticket without being have access to a specific cinema's website and we also display all the cinema including their information for user to check easily.

### Built With
For Frontend, we use Typescipt language and React library to render the UI and React Router for routing and React Icon for icon and React Auth Kit for authenticating user and TailwindCSS for decorating. For handling API, we use RestfulAPI and Axios

![image](https://images.viblo.asia/7be53157-271e-4a02-81ef-1bd05ce05832.png)
![image](https://caodang.fpt.edu.vn/wp-content/uploads/Tailwind-Css.jpg)

For Backend, we use Golang language to build 

![image](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png)

## Getting Started
Following these step to help you run the project

### Installation

1. Clone the repo
    ```sh
   git clone https://github.com/Darkiee12/Cinema-Ticket-Booking-System-WAD.git
   ```
2. Install NPM packages
   ```sh
   npm install
   ```
   or if you use bun
   ```sh
   bun install
   ```
3. Point to the "frontend" folder and type in terminal
   ```sh
   npm run dev
   ```
   or if you already have Bun then
   ```sh
   bun run dev
   ```
## Usage
At first you will landing on the Homepage, at here you can view some trending movies

![alt text](./readme%20images/image-1.png)

On the navigation bar at the top page, you can go to Movies page, where showing all the Current premiere movies and all Upcoming movies, Cinemas page, where you can view all Cinemas that we are working with, and About us page for you to acknowledge more information about our project. On the right side of navigation bar, you can click on Signin/Signup button so that you can go to Signin page and the Buy ticket button will navigate you to the buy movies ticket page.

![alt text](./readme%20images/image-2.png)

At Movies page, you can view all available movies and you can click on Buy ticket button at each movie's section to go to the Movie's detail page.

![alt text](./readme%20images/image-3.png)

At Movie's detail page, you can view more information about the movies and at the below will show all showtimes of the film regarding to the selected day. You can select 1 showtime and you will be navigated to the Seat select page.

If you are not sign in yet, then you will be asked to sign in at the sign in page.

![alt text](./readme%20images/image-8.png)

At here you can sign in or register new account. Once you have logged in, you will be able to buy ticket.

![alt text](./readme%20images/image-4.png)
![alt text](./readme%20images/image-5.png)
![alt text](./readme%20images/image-7.png)

After login and you select the movie's showtime you will go to seat select page. At here you can pick some seats by click on the seat and at the below will show some extra information including movie's information, the cinema's detail as well as the showtime and the number of selected seat and the total price. Once you confirm the booking, then you can click on Book button to finish the booking and you will be navigated to thank you page for reviewing the booking.

![alt text](./readme%20images/image-9.png)
![alt text](./readme%20images/image-10.png)

At Cinemas page, you can view all cinemas that are working with us as well as their branches and you can click on detail button at each cinema's section to view more information about that cinema.

![alt text](./readme%20images/image-11.png)
![alt text](./readme%20images/image-12.png)

Once you have logged in, you can go to Profile page. At this page, you can view more user's information as well as the booked tickets' information and you can click on update button to update user's information.

## Contributing

Nguyen Mach Khang Huy - ITCSU21072
* Figma, UI/UX desgin
* Frontend developer
* QA

Le Tuan Phuc - ITCSIU21096
* Frontend developer
* Database design
* Server deployment, maintenance

Hoang Nhan Kiet - ITITUN22037
* Backend developer
* Database design
