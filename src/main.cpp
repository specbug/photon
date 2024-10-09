#include <iostream>

int main(const int argc, char* argv[]) {

    if (argc < 2) {
        std::cerr << "incorrect syntax\nusage: photon <file.hv>" << std::endl;
        return EXIT_FAILURE;
    }

    std::cout << argv[1] << std::endl;
    return 0;
}