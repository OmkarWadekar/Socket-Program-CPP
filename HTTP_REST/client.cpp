#include <iostream>
#include <fstream>
#include "httplib.h"

int main() {
    httplib::Client cli("127.0.0.1", 12345); // Create an HTTP client

    // Send an HTTP GET request to download the file
    auto get_res = cli.Get("/download");

    if (get_res && get_res->status == 200) {
        // Save the received file
        std::ofstream file("received_file.txt", std::ios::binary);
        if (file.is_open()) {
            file.write(get_res->body.data(), get_res->body.length());
            file.close();
            std::cout << "File downloaded successfully." << std::endl;
        } else {
            std::cerr << "Failed to open the file for saving." << std::endl;
        }
    } else {
        std::cerr << "Failed to download the file." << std::endl;
    }

    // Send an HTTP POST request to upload the file
    std::ifstream upload_file("upload_file.txt", std::ios::binary);
    if (upload_file.is_open()) {
        std::string upload_content((std::istreambuf_iterator<char>(upload_file)), (std::istreambuf_iterator<char>()));
        auto post_res = cli.Post("/upload", upload_content, "text/plain");

        if (post_res && post_res->status == 200) {
            std::cout << "File uploaded successfully." << std::endl;
        } else {
            std::cerr << "Failed to upload the file." << std::endl;
        }

        upload_file.close();
    } else {
        std::cerr << "Failed to open the file for uploading." << std::endl;
    }

    return 0;
}
