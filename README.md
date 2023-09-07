# Demo Ticket Booking

This repository showcases a demonstration of a ticket booking system. The system allows users to book tickets for
various events and demonstrates the implementation of database locking using asynchronous processing with Hibiken Asynq.
Additionally, it utilizes OpenTelemetry for observability and monitoring.

## Technologies Used

The ticket booking system demo is built using the following technologies:

- Programming Language: Go
- Web Framework: Go Fiber
- Database: Postgres
- PubSub: Redis
- Asynchronous Processing: Hibiken Asynq
- Observability: OpenTelemetry
- Containerization: Docker
- Orchestration: Docker Compose

## Prerequisites

Before running the demo, ensure that you have the following prerequisites installed:

- Docker

## Installation

Follow the steps below to install and run the demo:

1. Clone this repository:

   `````bash
   git clone https://github.com/iniudin/demo-ticket-booking.git
   ```

2. Navigate to the application directory:

   ````bash
   cd demo-ticket-booking
   ```

4. Run the application using Docker Compose:

   ````bash
   docker-compose up --build
   ```

   The application will be accessible at `http://localhost:8080`.

## Usage

To use the ticket booking system demo, follow these steps:

1. Open a web browser and access `http://localhost:8080`.

2. Browse the available events and select the desired event.

3. Choose the number of tickets and select seats.

4. Proceed to the booking process.

5. The system will process the booking asynchronously using Hibiken Asynq, ensuring database locking to prevent
   conflicts.

6. Once the booking is completed, the system will display the booking confirmation.

7. You can view the booked tickets and their details in the order details.

## Observability (WIP)

The ticket booking system demo incorporates OpenTelemetry for observability and monitoring. It provides insights into
the system's performance, traces, and logs. You can access the observability data through the OpenTelemetry compatible
tools and dashboards.

## Contribution

If you would like to contribute to this project, please follow these steps:

1. Fork this repository.

2. Create a new branch for your feature (`git checkout -b new-feature`).

3. Make the necessary changes.

4. Commit your changes (`git commit -am 'Add new feature'`).

5. Push to the branch (`git push origin new-feature`).

6. Create a new pull request.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

Feel free to explore and experiment with the ticket booking system demo. If you encounter any issues or have suggestions
for improvements, please open an issue or submit a pull request.

```
MIT License

Copyright (c) 2023 Ahmad Syaifudin <sysfdn@pm.me>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```