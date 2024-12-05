#include <stdio.h>
#include <string.h>

int main() {
  char field[1000][1000];
  int height = 0;
  int width = 0;
  int count = 0;
  while (fgets(field[height], sizeof(field[height]), stdin) != NULL) {
    height++;
  }
  width = strlen(field[0]);
  for (int y = 1; y < height - 1; y++) {
    for (int x = 1; x < width - 1; x++) {
      if (field[y][x] == 'A') {
        char ul = field[y - 1][x - 1];
        char ur = field[y - 1][x + 1];
        char ll = field[y + 1][x - 1];
        char lr = field[y + 1][x + 1];
        if (ul == 'M' && lr == 'S' && ur == 'M' && ll == 'S') {
          count++;
        } else if (ul == 'M' && lr == 'S' && ur == 'S' && ll == 'M') {
          count++;
        } else if (ul == 'S' && lr == 'M' && ur == 'M' && ll == 'S') {
          count++;
        } else if (ul == 'S' && lr == 'M' && ur == 'S' && ll == 'M') {
          count++;
        }
      }
    }
  }
  printf("%d\n", count);
}