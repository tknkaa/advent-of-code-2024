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

    //左右
    for (int y = 0; y < height; y++) {
        for (int x = 0; x < width - 3; x++) {
            if (field[y][x] == 'X' && field[y][x + 1] == 'M' && field[y][x + 2] == 'A' && field[y][x + 3] == 'S') {
                count++;
            } else if (field[y][x] == 'S' && field[y][x + 1] == 'A' && field[y][x + 2] == 'M' && field[y][x + 3] == 'X') {
                count++;
            } 
        }
    }

    //上下
    for (int x = 0; x < width; x++) {
        for (int y = 0; y < height - 3; y++) {
            if (field[y][x] == 'X' && field[y + 1][x] == 'M' && field[y + 2][x] == 'A' && field[y + 3][x] == 'S') {
                count++;
            } else if (field[y][x] == 'S' && field[y + 1][x] == 'A' && field[y + 2][x] == 'M' && field[y + 3][x] == 'X') {
                count++;
            }
        }
    }

    //左上から右下
    for (int y = 0; y < height - 3; y++) {
        for (int x = 0; x < width - 3; x++) {
            if (field[y][x] == 'X' && field[y + 1][x + 1] == 'M' && field[y + 2][x + 2] == 'A' && field[y + 3][x + 3] == 'S') {
                count++;
            } else if (field[y][x] == 'S' && field[y + 1][x + 1] == 'A' && field[y + 2][x + 2] == 'M' && field[y + 3][x + 3] == 'X') {
                count++;
            }
        }
    }

    //右上から左下
    for (int y = 3; y < height; y++) {
        for (int x = 0; x < width - 3; x++) {
            if (field[y][x] == 'X' && field[y - 1][x + 1] == 'M' && field[y - 2][x + 2] == 'A' && field[y - 3][x + 3] == 'S') {
                count++;
            } else if (field[y][x] == 'S' && field[y - 1][x + 1] == 'A' && field[y - 2][x + 2] == 'M' && field[y - 3][x + 3] == 'X') {
                count++;
            }
        }
    }
    printf("%d\n", count);
}