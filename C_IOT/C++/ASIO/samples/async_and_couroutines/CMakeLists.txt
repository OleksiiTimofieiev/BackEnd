cmake_minimum_required(VERSION 3.15)
project(async_echo_server_lambdas)

set(Boost_USE_STATIC_LIBS OFF)
set(Boost_USE_MULTITHREADED ON)
set(Boost_USE_STATIC_RUNTIME OFF)
# set(CMAKE_CXX_FLAGS "-fcoroutines")

find_package(Boost 1.69.0 REQUIRED)
find_package(Threads)

message(STATUS "Boost version: ${Boost_VERSION}")

# This is needed if your Boost version is newer than your CMake version
# or if you have an old version of CMake (<3.5)

add_executable(async_echo_server_lambdas async_echo_server_lambdas.cpp)

target_link_libraries(async_echo_server_lambdas pthread)

set_target_properties(async_echo_server_lambdas PROPERTIES
            CXX_STANDARD 17
            # CXX_EXTENSIONS OFF
            )

            # add_executable(async_echo_server_coroutines async_echo_server_coroutines.cpp)

            # target_link_libraries(async_echo_server_coroutines pthread)
            
            # set_target_properties(async_echo_server_coroutines PROPERTIES
            #             CXX_STANDARD 17
            #             # CXX_EXTENSIONS OFF
                        # )
# target_compile_features(${TARGET_NAME}  cxx_std_17)

# target_link_libraries(echo_server PUBLIC Boost::asio)