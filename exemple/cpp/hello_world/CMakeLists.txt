# How to compile
# mkdir build
# cd build
# cmake ..
# Each time you want to rebuild the app
# cmake --build .
cmake_minimum_required(VERSION 3.0)

project(hello_world VERSION 0.0.1)

# Configure la version de cxx a 20
set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED True)

add_executable(hello_world src/main.cpp)


