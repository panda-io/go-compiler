#include <cinttypes>
#include <cuchar>
#include <iostream>
#include <vector>

class console;

namespace collection
{
class vector;
}

class console
{
public:
    template <class type>
virtual void write( value);
    template <class type>
virtual void write_line( value);
};

void main();

collection::vector<int32_t> create_vector();
namespace collection
{
class vector
{
public:
    virtual int32_t size();
    virtual void resize(int32_t size);
    virtual int32_t capacity();
    virtual bool empty();
    virtual void reserve(int32_t size);
    virtual void shrink();
    virtual  get(int32_t position);
    virtual  set(int32_t position,  value);
    virtual  front();
    virtual  back();
    virtual void fill(int32_t size,  value);
    virtual void push( value);
    virtual  pop();
    virtual void insert(int32_t position,  value);
    virtual void erase(int32_t position);
    virtual void clear();
};
}
template <class type>
void console::write( value)
{
    "std::cout << value;";
}
template <class type>
void console::write_line( value)
{
    "std::cout << value << std::endl;";
}
collection::vector<int32_t> v;
void main()
{
    v = ;
    ;
    ;
    ;
    ;
    ;
    ;
    ;
}
collection::vector<int32_t> create_vector()
{
    return ;
}
namespace collection
{
}
