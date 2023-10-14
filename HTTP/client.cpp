#include <iostream>
#include <fstream>
#include "httplib.h"

int main() {
    httplib::Client cli("127.0.0.1", 12345); // Create an HTTP client

    // Send an HTTP GET request to download the file
    auto res = cli.Get("/download");

    if (res && res->status == 200) {
        // Save the received file
        std::ofstream file("received_file.txt", std::ios::binary);
        if (file.is_open()) {
            file.write(res->body.data(), res->body.length());
            file.close();
            std::cout << "File received successfully." << std::endl;
        } else {
            std::cerr << "Failed to open the file for saving." << std::endl;
        }
    } else {
        std::cerr << "Failed to download the file." << std::endl;
    }

    return 0;
}
