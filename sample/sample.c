int g = 10;

void add(int a, int b)
{
    a = 1;
}

struct ss
{
    int value;
};

int main()
{
    int a = 1;
    {
        int a = 2;
    }

    if (a == 1) {
        int b = 1;
    } else {
        int b = 2;
    }

    int c = 3;
}