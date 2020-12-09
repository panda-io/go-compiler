// --------------------------------       includes       --------------------------------
#include <cinttypes>
#include <cuchar>
#include <string>

// -------------------------------- forward declarations --------------------------------
enum class color;

class must_do;

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

std::string string_v_raw = "hello \\n\n    world\\n";

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
    int32_t value = 100;
    empty();
    virtual ~empty();
};

class base_class
{
public:
    int32_t value = 1;
    base_class(int32_t value);
    virtual void destroy();
};

class derive_class : public base_class, must_do
{
public:
    virtual void print();
};

int32_t main();

// --------------------------------      implements      --------------------------------
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
