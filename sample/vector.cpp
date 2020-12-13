// --------------------------------       includes       --------------------------------
#include <cuchar>
#include <string>
#include <memory>
#include <iostream>
#include <vector>
#include <cinttypes>

// -------------------------------- forward declarations --------------------------------
namespace console
{
}

namespace collection
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

namespace collection
{
}

std::shared_ptr<collection::vector<int32_t>> v;

int32_t main();

std::shared_ptr<collection::vector<int32_t>> create_vector();

// --------------------------------      implements      --------------------------------
namespace console
{
template <class type>
void write(type value)
{
    std::cout << value;
}

template <class type>
void write_line(type value)
{
    std::cout << value << std::endl;
}
}

namespace collection
{
}

int32_t main()
{
    console::write_line("hello world");
    v = create_vector();
    console::write_line(v->size());
    v->push(1);
    v->push(2);
    console::write_line(v->size());
    auto x = v->pop();
    console::write_line(x);
    console::write_line(v->size());
    return 0;
}

std::shared_ptr<collection::vector<int32_t>> create_vector()
{
    return std::make_shared<collection::vector<int32_t>>();
}

