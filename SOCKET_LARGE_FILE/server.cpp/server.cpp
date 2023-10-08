#include <iostream>
#include <fstream>
#include <winsock2.h>

int main() {
    // Initialize Winsock
    WSADATA wsaData;
    WSAStartup(MAKEWORD(2, 2), &wsaData);

    // Create a socket
    SOCKET serverSocket = socket(AF_INET, SOCK_STREAM, 0);

    // Server address and port
    sockaddr_in serverAddr;
    serverAddr.sin_family = AF_INET;
    serverAddr.sin_addr.s_addr = INADDR_ANY;
    serverAddr.sin_port = htons(12345);

    // Bind the socket to an address and port
    bind(serverSocket, (struct sockaddr*)&serverAddr, sizeof(serverAddr));
    listen(serverSocket, SOMAXCONN);

    std::cout << "Server is listening on port 12345..." << std::endl;

    // Accept a client connection
    SOCKET clientSocket = accept(serverSocket, nullptr, nullptr);

    // Open and send a large file to the client
    std::ifstream file("large_file.txt", std::ios::binary);
    if (file.is_open()) {
        char buffer[1024];
        while (!file.eof()) {
            file.read(buffer, sizeof(buffer));
            send(clientSocket, buffer, file.gcount(), 0);
        }
        file.close();
    } else {
        std::cerr << "Failed to open the file for sending." << std::endl;
    }

    // Clean up
    closesocket(clientSocket);
    closesocket(serverSocket);
    WSACleanup();

    return 0;
}
