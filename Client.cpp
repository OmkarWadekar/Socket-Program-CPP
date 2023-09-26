#include <iostream>
#include <winsock2.h>

int main()
{
    // Initialize Winsock
    WSADATA wsaData;
    if (WSAStartup(MAKEWORD(2, 2), &wsaData) != 0)
    {
        std::cerr << "Failed to initialize Winsock." << std::endl;
        return 1;
    }

    // Create a socket
    SOCKET clientSocket = socket(AF_INET, SOCK_STREAM, 0);
    if (clientSocket == INVALID_SOCKET)
    {
        std::cerr << "Failed to create socket." << std::endl;
        WSACleanup();
        return 1;
    }

    // Connect to the server (localhost in this example)
    sockaddr_in serverAddr;
    serverAddr.sin_family = AF_INET;
    serverAddr.sin_addr.s_addr = inet_addr("127.0.0.1"); // Server IP (localhost)
    serverAddr.sin_port = htons(12345);                  // Server port

    if (connect(clientSocket, (struct sockaddr *)&serverAddr, sizeof(serverAddr)) == SOCKET_ERROR)
    {
        std::cerr << "Failed to connect to the server." << std::endl;
        closesocket(clientSocket);
        WSACleanup();
        return 1;
    }

    // Send a message to the server
    const char *message = "Hello, server!";
    if (send(clientSocket, message, strlen(message), 0) == SOCKET_ERROR)
    {
        std::cerr << "Failed to send message." << std::endl;
    }
    else
    {
        std::cout << "Message sent to the server." << std::endl;
    }

    // Clean up
    closesocket(clientSocket);
    WSACleanup();

    return 0;
}