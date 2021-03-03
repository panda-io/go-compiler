#include <stdio.h>
#include <stdbool.h>

int main()
{
   float a = 0.1;
   float b = -a;
   int c = 1;
   int d = -c;
   int e = +c;
   return 0;
}

bool test(int* a)
{
   if (a == NULL)
   {
      return true;
   }
   return false;
}