// --------------------------------       includes       --------------------------------
#include <cinttypes>
#include <cuchar>
#include <string>
#include <iostream>

// -------------------------------- forward declarations --------------------------------
enum class color;

class printer;

class empty;

class base_class;

class derive_class;

// --------------------------------     declarations     --------------------------------
bool b = true;

char32_t c = U'a';

int8_t sbyte_v = 1;

int16_t short_v = 0;

int32_t int_v = 0;

int64_t long_v = 0;

int8_t int8_v = 0;

int16_t int16_v = 0;

int32_t int32_v = 0;

int64_t int64_v = 0;

uint8_t byte_v = 0;

uint16_t ushort_v = 0;

uint32_t uint_v = 0;

uint64_t ulong_v = 0;

uint8_t uint8_v = 0;

uint16_t uint16_v = 0;

uint32_t uint32_v = 0;

uint64_t uint64_v = 0;

float float_v = 0;

double double_v = 0;

float f32_v = 0;

double f64_v = 0;

std::string string_v = "hello world\n";

std::string string_v_raw = "\"hello world\"";

template <class T>
void print_line(T t);

int32_t add(int32_t a, int32_t b);

enum class color
{
    red,
    green = 10,
    blue
};

class printer
{
public:
    virtual void print() = 0;
};

class empty
{
public:
    empty();
    virtual ~empty();
};

class base_class
{
public:
    base_class();
    int32_t value = 1;
    base_class(int32_t value);
    virtual ~base_class();
};

class derive_class : public base_class, printer
{
public:
    derive_class();
    derive_class(int32_t value);
    virtual ~derive_class();
    virtual void print();
};

int32_t main();

// --------------------------------      implements      --------------------------------
template <class T>
void print_line(T t)
{
    std::cout << t << std::endl;
}

int32_t add(int32_t a, int32_t b)
{
    return a + b;
}

empty::empty()
{
}

empty::~empty()
{
}

base_class::base_class()
{
}

base_class::base_class(int32_t value)
{
    this->value = value;
    print_line("base class constructor");
}

base_class::~base_class()
{
    print_line("base class destructor");
}

derive_class::derive_class()
{
}

derive_class::derive_class(int32_t value)
{
    print_line("derive class constructor");
}

derive_class::~derive_class()
{
    print_line("derive class destructor");
}

void derive_class::print()
{
    print_line(this->value);
}

int32_t main()
{
    print_line(add(1, 1));
    return 0;
}

