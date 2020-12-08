#include <cinttypes>
#include <cuchar>
#include <string>

enum class color;

class must_do;

class empty;

class base_class;

class derive_class;

template <class T>
void print(T t);

template <class T>
void print_line(T t);

int32_t add(int32_t a, int32_t b);

enum class color
{
    red,
    green = 10,
    blue
};

class must_do
{
public:
    virtual void print();
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
        base_class(int32_t value);
    virtual void destroy();
};

class derive_class : public base_class, must_do
{
public:
    virtual void print();
};

int32_t main();
bool b = true;
char32_t c = 'a';
int8_t sbyte_v = 1;
int16_t short_v = 0;
int32_t int_v;
int64_t long_v;
int8_t int8_v;
int16_t int16_v;
int32_t int32_v;
int64_t int64_v;
uint8_t byte_v;
uint16_t ushort_v;
uint32_t uint_v;
uint64_t ulong_v;
uint8_t uint8_v;
uint16_t uint16_v;
uint32_t uint32_v;
uint64_t uint64_v;
float float_v;
double double_v;
float f32_v;
double f64_v;
std::string string_v = "hello world\n";
std::string string_v_raw = "hello \\n\n    world\\n";
template <class T>
void print(T t)
{
    "std::cout << t";
}
template <class T>
void print_line(T t)
{
    "std::cout << t << std::endl";
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
base_class::base_class(int32_t value)
{
    this.value = value;
    ;
}
void base_class::destroy()
{
    ;
}
void derive_class::print()
{
    ;
}
int32_t main()
{
    ;
    ;
    return 0;
}
