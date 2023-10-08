#include <iostream>
#include <fstream>
#include <winsock2.h>

int main() {
    // Initialize Winsock
    WSADATA wsaData;
    WSAStartup(MAKEWORD(2, 2), &wsaData);

    // Create a socket
    SOCKET clientSocket = socket(AF_INET, SOCK_STREAM, 0);

    // Server address and port
    sockaddr_in serverAddr;
    serverAddr.sin_family = AF_INET;
    serverAddr.sin_addr.s_addr = inet_addr("127.0.0.1");
    serverAddr.sin_port = htons(12345);

    // Connect to the server
    connect(clientSocket, (struct sockaddr*)&serverAddr, sizeof(serverAddr));

    // Receive and save the file from the server
    std::ofstream file("received_file.txt", std::ios::binary);
    if (file.is_open()) {
        char buffer[1024];
        int bytesRead;
        while ((bytesRead = recv(clientSocket, buffer, sizeof(buffer), 0)) > 0) {
            file.write(buffer, bytesRead);
        }
        file.close();
        std::cout << "File received successfully." << std::endl;
    } else {
        std::cerr << "Failed to open the file for receiving." << std::endl;
    }

    // Clean up
    closesocket(clientSocket);
    WSACleanup();

    return 0;
}
