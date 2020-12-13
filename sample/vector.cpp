// --------------------------------       includes       --------------------------------
#include <cinttypes>
#include <cuchar>
#include <string>
#include <memory>
#include <iostream>
#include <vector>

// -------------------------------- forward declarations --------------------------------

namespace console
{
}

namespace collection
{
template <class type>
class vector;
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
template <class type>
class vector
{
public:
    vector();
    virtual int32_t size() = 0;
    virtual void resize(int32_t size) = 0;
    virtual int32_t capacity() = 0;
    virtual bool empty() = 0;
    virtual void reserve(int32_t size) = 0;
    virtual void shrink() = 0;
    virtual std::shared_ptr<type> get(int32_t position) = 0;
    virtual std::shared_ptr<type> set(int32_t position, type value) = 0;
    virtual std::shared_ptr<type> front() = 0;
    virtual std::shared_ptr<type> back() = 0;
    virtual void fill(int32_t size, type value) = 0;
    virtual void push(type value) = 0;
    virtual std::shared_ptr<type> pop() = 0;
    virtual void insert(int32_t position, type value) = 0;
    virtual void erase(int32_t position) = 0;
    virtual void clear() = 0;
};
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
template <class type>
vector<type>::vector()
{
}
















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

