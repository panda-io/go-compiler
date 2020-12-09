// --------------------------------       includes       --------------------------------
#include <string>
#include <iostream>
#include <cinttypes>
#include <cuchar>

// -------------------------------- forward declarations --------------------------------
namespace console
{
}

// --------------------------------     declarations     --------------------------------
int32_t main();

namespace console
{
template <class type>
void write(type value);

template <class type>
void write_line(type value);
}

// --------------------------------      implements      --------------------------------
int32_t main()
{
    console::write_line("hello world");
    return 0;
}

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

