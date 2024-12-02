#include <iostream>
#include <vector>

int main(void) {
  const int size = 1000;
  std::vector<long> list1(size);
  std::vector<long> list2(size);

  for (int i = 0; i < size; i++) {
    std::cin >> list1[i] >> list2[i];
  }

  long simScore = 0;

  for (int i = 0; i < size; i++) {
    int count = 0;
    for (int j = 0; j < size; j++) {
      if (list1[i] == list2[j]) {
        count++;
      }
    }
    simScore += list1[i] * count;
  }
  std::cout << simScore;
}
