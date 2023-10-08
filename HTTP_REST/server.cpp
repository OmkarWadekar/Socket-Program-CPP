#include <iostream>
#include <fstream>
#include "httplib.h"

int main() {
    httplib::Server svr;

    // Define an HTTP GET handler for downloading the file
    svr.Get("/download", [](const httplib::Request& req, httplib::Response& res) {
        // Open and send a large file to the client
        std::ifstream file("large_file.txt", std::ios::binary);
        if (file.is_open()) {
            res.set_content([&file](size_t offset, size_t length, httplib::DataSink& sink) {
                char buffer[1024];
                file.seekg(offset);
                size_t bytesRead = 0;
                while (length > 0 && !file.eof()) {
                    file.read(buffer, std::min(sizeof(buffer), length));
                    size_t readCount = file.gcount();
                    sink.write(buffer, readCount);
                    length -= readCount;
                    bytesRead += readCount;
                }
                return bytesRead;
            });
            file.close();
        } else {
            res.status = 500;
            res.set_content("Failed to open the file for sending.", "text/plain");
        }
    });

    // Define an HTTP POST handler for uploading a file
    svr.Post("/upload", [](const httplib::Request& req, httplib::Response& res) {
        std::ofstream file("received_file.txt", std::ios::binary);
        if (file.is_open()) {
            file.write(req.body.c_str(), req.body.length());
            file.close();
            res.status = 200;
            res.set_content("File uploaded successfully.", "text/plain");
        } else {
            res.status = 500;
            res.set_content("Failed to save the uploaded file.", "text/plain");
        }
    });

    std::cout << "Server is listening on port 12345..." << std::endl;
    svr.listen("127.0.0.1", 12345); // Start the server on localhost:12345

    return 0;
}
