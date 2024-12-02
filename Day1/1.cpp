#include <algorithm>
#include <cstdlib>
#include <iostream>
#include <vector>

int main(void) {
  int size = 1000;
  std::vector<long> list1(size);
  std::vector<long> list2(size);

  for (int i = 0; i < size; i++) {
    std::cin >> list1[i] >> list2[i];
  }

  std::sort(list1.begin(), list1.end());
  std::sort(list2.begin(), list2.end());

  long sum = 0;
  for (int i = 0; i < size; i++) {
    sum += std::abs(list1[i] - list2[i]);
  }
  std::cout << sum;
}
