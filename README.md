# Simple Client-Server Communication 

This repository contains a simple example of client-server communication using sockets in a Windows environment. The repository includes two programs: a server and a client. The client sends a message to the server, and the server receives and displays the message.

## Prerequisites

- Windows operating system (tested on Windows 10)
- C++ development environment (e.g., Visual Studio, Code::Blocks, or MinGW for command line)
- Basic knowledge of C++ programming

## How to Run

Follow these steps to compile and run the server and client programs:

### Server

1. Open a command prompt.

2. Navigate to the directory containing the `server.cpp` file.

3. Compile the server program using a C++ compiler. For example, using g++:

   ```bash
   g++ server.cpp -o server -lws2_32
4. Run the program using   
   ```bash
   ./server
### Client
1. Now compile the Client program using a C++ compiler. For example, using g++:

   ```bash
   g++ client.cpp -o client -lws2_32
2. Run the program using
   ```bash
   ./client
### Output: 
   ![image](https://github.com/OmkarWadekar/Socket-Program-CPP/assets/87438761/878742bc-5a6a-497f-b216-02b76298b5d9)

