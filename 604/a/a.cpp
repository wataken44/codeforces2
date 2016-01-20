
/*
  a.cpp
 */

#include <vector>
#include <iostream>

using namespace std;

int main() {
  vector<int> m(5);
  vector<int> w(5);
  vector<int> h(2);

  for(int i = 0; i < 5; ++i) {
    cin >>  m[i];
  }

  for(int i = 0; i < 5; ++i) {
    cin >>  w[i];
  }

  for(int i = 0; i < 2; ++i) {
    cin >> h[i];
  }

  int r = h[0] * 100 - h[1] * 50;
  for(int i = 0; i < 5; ++i) {
    int x = (i + 1) * 500;
    r += max(75 * x, (250 - m[i]) * x - 250 * 50 * w[i]) / 250;
  }
  cout << r << endl;
}
