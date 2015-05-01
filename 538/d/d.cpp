
/*
  d.cpp
 */

#include <algorithm>
#include <iostream>
#include <list>
#include <set>
#include <string>
#include <utility>
#include <vector>

using namespace std;

int main() {
  int n;
  cin >> n;

  vector< string > offered(n);
  for(int i = 0; i < n; ++i) {
    cin >> offered[i];
  }

  list< pair<int, int> > safe;
  for(int y = 0; y < n; ++y) {
    for(int x = 0; x < n; ++x) {
      if(offered[y][x] != 'o') {
        continue;
      }

      for(int yy = 0; yy < n; ++yy) {
        for(int xx = 0; xx < n; ++xx) {
          if(offered[yy][xx] != '.') {
            continue;
          }
          safe.push_back( make_pair(yy-y, xx-x) );
        }
      }
      
    }
  }

  safe.sort();
  list< pair<int, int> >::iterator it = unique(safe.begin(), safe.end());
  safe.erase(it, safe.end());

  for(list< pair<int, int> >::iterator it = safe.begin(); it != safe.end();) {
    bool rem = false;
    for(int y = 0; y < n; ++y) {
      for(int x = 0; x < n; ++x) {
        if(offered[y][x] != 'o') {
          continue;
        }
        int dy = it->first;
        int dx = it->second;
        if(y + dy < 0 || y + dy >= n) {
          continue;
        }
        if(x + dx < 0 || x + dx >= n) {
          continue;
        }
        if(offered[y + dy][x + dx] == 'x') {
          rem = true;
        }
      }
    }
    if(rem) {
      it = safe.erase(it);
    } else {
      ++it;
    }
  }

  vector< string > ret(2 * n - 1, string(2 * n - 1, 'x') );
  ret[n-1][n-1] = 'o';
  for(list< pair<int, int> >::iterator it = safe.begin(); it != safe.end(); ++it) {
    int dy = it->first;
    int dx = it->second;
    ret[n-1 + dy][n-1 + dx] = '.';
  }

  bool ok = false;
  for(int y = 0; y < 2*n - 1; ++y) {
    for(int x = 0; x < 2*n - 1; ++x) {
      ok = ok | (ret[y][x] == 'x');
    }
  }

  if(ok) {
    cout << "YES" << endl;
    for(int y = 0; y < 2*n - 1; ++y) {
      cout << ret[y] << endl;
    }    
  } else {
    cout << "NO" << endl;
  }
  
  return 0;
}
