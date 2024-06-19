# CineU

This is a booking website that allows users to book movie tickets from all cinemas within just one website. This project is made for the "Web Application Development" course.

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

Our website shows movie data from all cinemas, including movie information, showtimes from every cinema, and allows users to book movie tickets without accessing specific cinema websites. We also display all the cinemas, including their information, for easy user reference.

### Built With

For the frontend, we use TypeScript and React library to render the UI, React Router for routing, React Icon for icons, React Auth Kit for user authentication, and TailwindCSS for styling. For handling API calls, we use RestfulAPI and Axios.

![image](https://images.viblo.asia/7be53157-271e-4a02-81ef-1bd05ce05832.png)
![image](https://caodang.fpt.edu.vn/wp-content/uploads/Tailwind-Css.jpg)

For the backend, we use Golang, and our architecture follows the Clean Architecture pattern, which includes a business layer, a store layer, and a transport layer.

![image](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png)

### Backend Architecture

Our backend is designed using the Clean Architecture pattern, which promotes separation of concerns and ensures that the application is easy to maintain and test.

- **Business Layer:** This layer contains the core business logic and domain models of the application. It is independent of frameworks and other layers, ensuring that changes to the business logic do not affect other parts of the application.

- **Store Layer:** This layer handles data persistence and retrieval. It includes implementations for interacting with databases or any other form of storage. In our system, this layer interacts with Redis for caching and a database for persistent storage.

- **Transport Layer:** This layer deals with the transport mechanisms for interacting with the application. It includes the HTTP server, routes, and request handlers. This layer is responsible for processing incoming requests, delegating work to the business layer, and sending responses back to the client.

### Infrastructure Components

- **Redis for Caching:** Redis is used as an in-memory cache to speed up data access and reduce load on the backend server. It helps in storing frequently accessed data and session management.

- **Netdata for Monitoring:** Netdata is used for real-time performance monitoring of the system. It provides detailed insights into system health, resource usage, and application performance metrics.

- **Jaeger and OpenTelemetry for Tracing:** We use OpenTelemetry for generating and collecting trace data and Jaeger for visualizing and analyzing the traces. This helps in understanding the application's performance and diagnosing issues.

### Deployment

All components of the system are hosted on a virtual machine (VM) in the cloud, which provides flexibility, scalability, and reliability. The cloud environment ensures that the system can scale with increasing demand and provides high availability.

## Getting Started

Follow these steps to set up and run the project.

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Darkiee12/Cinema-Ticket-Booking-System-WAD.git
    ```
2. Install Docker:
    ```sh
    sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
    ```
3. Deploy using docker compose:
    ```sh
    docker compose up -d --build
    ```

## Usage

Upon accessing the website, you will land on the Homepage where you can view trending movies.

![alt text](./readme%20images/image-1.png)

The navigation bar at the top of the page allows you to access different sections: Movies, Cinemas, and About Us. The Signin/Signup button allows users to log in or create an account, and the Buy Ticket button navigates users to the movie ticket purchase page.

![alt text](./readme%20images/image-2.png)

### Movies Page
This page displays all current and upcoming movies. Users can click on the "Buy ticket" button for each movie to view more details and proceed with the booking.

![alt text](./readme%20images/image-3.png)

### Movie Detail Page
This page shows detailed information about the movie and available showtimes. Users can select a showtime and proceed to the seat selection page.

### Sign-in Page
Users must sign in to book tickets. If not signed in, users will be prompted to log in or register.

![alt text](./readme%20images/image-8.png)

### Seat Selection Page
After selecting a showtime, users can pick their seats. Information such as the movie, cinema details, showtime, selected seats, and total price is displayed. Once confirmed, users can book the tickets and are directed to a thank you page.

![alt text](./readme%20images/image-9.png)
![alt text](./readme%20images/image-10.png)

### Cinemas Page
This page lists all partner cinemas and their branches. Users can view more details about each cinema by clicking the "Detail" button.

![alt text](./readme%20images/image-11.png)
![alt text](./readme%20images/image-12.png)

### Profile Page
Logged-in users can access their profile page to view and update their information and view their booked tickets.

![alt text](./readme%20images/image-4.png)
![alt text](./readme%20images/image-5.png)
![alt text](./readme%20images/image-7.png)

## Acknowledgments

- [React](https://reactjs.org/)
- [TailwindCSS](https://tailwindcss.com/)
- [Golang](https://golang.org/)
- [Redis](https://redis.io/)
- [Netdata](https://www.netdata.cloud/)
- [Jaeger](https://www.jaegertracing.io/)
- [OpenTelemetry](https://opentelemetry.io/)

This project is a collective effort of the Web Application Development course participants. Special thanks to all contributors and supporters.

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
