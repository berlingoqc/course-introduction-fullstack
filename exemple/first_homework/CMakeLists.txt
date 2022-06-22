cmake_minimum_required(VERSION 3.0)

project(homework_1 VERSION 0.0.1)

set(CMAKE_CXX_COMPILER "clang++")

# Configure la version de cxx a 20
set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED True)

add_executable(v1 src/v1.cpp)
add_executable(v2 src/v2.cpp)


