#include <iostream>
using namespace std;

int main() {

  int first_number, second_number, diff;
    
  cout << "Enter two integers: ";
  cin >> first_number >> second_number;
  diff = first_number - second_number;
  cout << first_number << " - " <<  second_number << " = " << diff;     

  return 0;
}
