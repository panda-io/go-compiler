// --------------------------------       includes       --------------------------------
#include <cuchar>
#include <string>
#include <iostream>
#include <cinttypes>

// -------------------------------- forward declarations --------------------------------
namespace console
{
}


// --------------------------------     declarations     --------------------------------
namespace console
{
template <class type>
void write(type value);

template <class type>
void write_line(type value);
}

int32_t main();

// --------------------------------      implements      --------------------------------
namespace console
{
template <class type>
void write(type value)
{
    std::cout << value;;
}

template <class type>
void write_line(type value)
{
    std::cout << value << std::endl;;
}
}

int32_t main()
{
    console::write_line("hello world");
    return 0;
}

