int g = 10;

void add(int a, int b)
{
    a = 1;
    int c = a + b;
}

struct ss
{
    int value;
};

int main()
{
    struct ss s;
    s.value = 10;
    add(1, 2);
}