// Code generated from cminus.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // cminus

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 60, 381,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 3, 2, 3, 2, 5, 2, 77,
	10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 84, 10, 3, 12, 3, 14, 3, 87,
	11, 3, 3, 4, 3, 4, 3, 4, 5, 4, 92, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4,
	98, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 107, 10, 5,
	12, 5, 14, 5, 110, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 117, 10,
	6, 12, 6, 14, 6, 120, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7,
	128, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 136, 10, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 145, 10, 10, 12, 10, 14,
	10, 148, 11, 10, 3, 11, 3, 11, 5, 11, 152, 10, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 5, 13, 164, 10, 13, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 175,
	10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 5, 16, 189, 10, 16, 3, 16, 5, 16, 192, 10, 16, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 18, 6, 18, 199, 10, 18, 13, 18, 14, 18, 200,
	3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 207, 10, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 217, 10, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 227, 10, 19, 12, 19, 14, 19,
	230, 11, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 5, 22, 242, 10, 22, 3, 23, 3, 23, 5, 23, 246, 10, 23, 3, 24,
	3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 7, 26, 261, 10, 26, 12, 26, 14, 26, 264, 11, 26, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 7, 27, 281, 10, 27, 12, 27, 14, 27, 284, 11, 27, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 295, 10,
	28, 12, 28, 14, 28, 298, 11, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 7, 29, 306, 10, 29, 12, 29, 14, 29, 309, 11, 29, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 7, 30, 317, 10, 30, 12, 30, 14, 30, 320, 11, 30, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 328, 10, 31, 12, 31, 14,
	31, 331, 11, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 7, 32, 339,
	10, 32, 12, 32, 14, 32, 342, 11, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 7, 33, 350, 10, 33, 12, 33, 14, 33, 353, 11, 33, 3, 34, 3, 34, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 7, 36, 370, 10, 36, 12, 36, 14, 36, 373, 11, 36, 3, 37, 5,
	37, 376, 10, 37, 3, 37, 3, 37, 3, 37, 3, 37, 2, 16, 4, 8, 10, 18, 36, 50,
	52, 54, 56, 58, 60, 62, 64, 70, 38, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
	58, 60, 62, 64, 66, 68, 70, 72, 2, 6, 7, 2, 34, 34, 36, 36, 38, 38, 41,
	41, 46, 47, 4, 2, 34, 34, 36, 36, 3, 2, 38, 40, 4, 2, 11, 14, 17, 17, 2,
	388, 2, 74, 3, 2, 2, 2, 4, 78, 3, 2, 2, 2, 6, 88, 3, 2, 2, 2, 8, 101, 3,
	2, 2, 2, 10, 111, 3, 2, 2, 2, 12, 127, 3, 2, 2, 2, 14, 129, 3, 2, 2, 2,
	16, 133, 3, 2, 2, 2, 18, 139, 3, 2, 2, 2, 20, 151, 3, 2, 2, 2, 22, 153,
	3, 2, 2, 2, 24, 163, 3, 2, 2, 2, 26, 167, 3, 2, 2, 2, 28, 176, 3, 2, 2,
	2, 30, 191, 3, 2, 2, 2, 32, 193, 3, 2, 2, 2, 34, 206, 3, 2, 2, 2, 36, 216,
	3, 2, 2, 2, 38, 231, 3, 2, 2, 2, 40, 233, 3, 2, 2, 2, 42, 235, 3, 2, 2,
	2, 44, 245, 3, 2, 2, 2, 46, 247, 3, 2, 2, 2, 48, 249, 3, 2, 2, 2, 50, 251,
	3, 2, 2, 2, 52, 265, 3, 2, 2, 2, 54, 285, 3, 2, 2, 2, 56, 299, 3, 2, 2,
	2, 58, 310, 3, 2, 2, 2, 60, 321, 3, 2, 2, 2, 62, 332, 3, 2, 2, 2, 64, 343,
	3, 2, 2, 2, 66, 354, 3, 2, 2, 2, 68, 356, 3, 2, 2, 2, 70, 364, 3, 2, 2,
	2, 72, 375, 3, 2, 2, 2, 74, 76, 5, 6, 4, 2, 75, 77, 5, 4, 3, 2, 76, 75,
	3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 3, 3, 2, 2, 2, 78, 79, 8, 3, 1, 2,
	79, 80, 5, 68, 35, 2, 80, 85, 3, 2, 2, 2, 81, 82, 12, 3, 2, 2, 82, 84,
	5, 68, 35, 2, 83, 81, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2,
	2, 85, 86, 3, 2, 2, 2, 86, 5, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 88, 89, 7,
	3, 2, 2, 89, 91, 7, 22, 2, 2, 90, 92, 5, 70, 36, 2, 91, 90, 3, 2, 2, 2,
	91, 92, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 94, 7, 23, 2, 2, 94, 95, 7,
	26, 2, 2, 95, 97, 5, 10, 6, 2, 96, 98, 5, 8, 5, 2, 97, 96, 3, 2, 2, 2,
	97, 98, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 100, 7, 27, 2, 2, 100, 7, 3,
	2, 2, 2, 101, 102, 8, 5, 1, 2, 102, 103, 5, 12, 7, 2, 103, 108, 3, 2, 2,
	2, 104, 105, 12, 3, 2, 2, 105, 107, 5, 12, 7, 2, 106, 104, 3, 2, 2, 2,
	107, 110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109,
	9, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 111, 112, 8, 6, 1, 2, 112, 113, 5,
	22, 12, 2, 113, 118, 3, 2, 2, 2, 114, 115, 12, 3, 2, 2, 115, 117, 5, 22,
	12, 2, 116, 114, 3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2,
	118, 119, 3, 2, 2, 2, 119, 11, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 128,
	5, 16, 9, 2, 122, 128, 5, 24, 13, 2, 123, 128, 5, 26, 14, 2, 124, 128,
	5, 28, 15, 2, 125, 128, 5, 30, 16, 2, 126, 128, 5, 14, 8, 2, 127, 121,
	3, 2, 2, 2, 127, 122, 3, 2, 2, 2, 127, 123, 3, 2, 2, 2, 127, 124, 3, 2,
	2, 2, 127, 125, 3, 2, 2, 2, 127, 126, 3, 2, 2, 2, 128, 13, 3, 2, 2, 2,
	129, 130, 7, 53, 2, 2, 130, 131, 7, 52, 2, 2, 131, 132, 5, 24, 13, 2, 132,
	15, 3, 2, 2, 2, 133, 135, 7, 26, 2, 2, 134, 136, 5, 18, 10, 2, 135, 134,
	3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 138, 7, 27,
	2, 2, 138, 17, 3, 2, 2, 2, 139, 140, 8, 10, 1, 2, 140, 141, 5, 20, 11,
	2, 141, 146, 3, 2, 2, 2, 142, 143, 12, 3, 2, 2, 143, 145, 5, 20, 11, 2,
	144, 142, 3, 2, 2, 2, 145, 148, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146,
	147, 3, 2, 2, 2, 147, 19, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 149, 152, 5,
	12, 7, 2, 150, 152, 5, 22, 12, 2, 151, 149, 3, 2, 2, 2, 151, 150, 3, 2,
	2, 2, 152, 21, 3, 2, 2, 2, 153, 154, 7, 4, 2, 2, 154, 155, 7, 53, 2, 2,
	155, 156, 7, 52, 2, 2, 156, 157, 5, 66, 34, 2, 157, 158, 7, 22, 2, 2, 158,
	159, 7, 54, 2, 2, 159, 160, 7, 23, 2, 2, 160, 161, 7, 50, 2, 2, 161, 23,
	3, 2, 2, 2, 162, 164, 5, 36, 19, 2, 163, 162, 3, 2, 2, 2, 163, 164, 3,
	2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 166, 7, 50, 2, 2, 166, 25, 3, 2, 2,
	2, 167, 168, 7, 10, 2, 2, 168, 169, 7, 22, 2, 2, 169, 170, 5, 36, 19, 2,
	170, 171, 7, 23, 2, 2, 171, 174, 5, 12, 7, 2, 172, 173, 7, 9, 2, 2, 173,
	175, 5, 12, 7, 2, 174, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 27,
	3, 2, 2, 2, 176, 177, 7, 18, 2, 2, 177, 178, 7, 22, 2, 2, 178, 179, 5,
	36, 19, 2, 179, 180, 7, 23, 2, 2, 180, 181, 5, 12, 7, 2, 181, 29, 3, 2,
	2, 2, 182, 183, 7, 8, 2, 2, 183, 192, 7, 50, 2, 2, 184, 185, 7, 7, 2, 2,
	185, 192, 7, 50, 2, 2, 186, 188, 7, 16, 2, 2, 187, 189, 5, 36, 19, 2, 188,
	187, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 192,
	7, 50, 2, 2, 191, 182, 3, 2, 2, 2, 191, 184, 3, 2, 2, 2, 191, 186, 3, 2,
	2, 2, 192, 31, 3, 2, 2, 2, 193, 194, 9, 2, 2, 2, 194, 33, 3, 2, 2, 2, 195,
	207, 7, 53, 2, 2, 196, 207, 7, 54, 2, 2, 197, 199, 7, 56, 2, 2, 198, 197,
	3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2,
	2, 2, 201, 207, 3, 2, 2, 2, 202, 203, 7, 22, 2, 2, 203, 204, 5, 36, 19,
	2, 204, 205, 7, 23, 2, 2, 205, 207, 3, 2, 2, 2, 206, 195, 3, 2, 2, 2, 206,
	196, 3, 2, 2, 2, 206, 198, 3, 2, 2, 2, 206, 202, 3, 2, 2, 2, 207, 35, 3,
	2, 2, 2, 208, 209, 8, 19, 1, 2, 209, 210, 7, 36, 2, 2, 210, 217, 5, 36,
	19, 7, 211, 212, 7, 22, 2, 2, 212, 213, 5, 36, 19, 2, 213, 214, 7, 23,
	2, 2, 214, 217, 3, 2, 2, 2, 215, 217, 7, 54, 2, 2, 216, 208, 3, 2, 2, 2,
	216, 211, 3, 2, 2, 2, 216, 215, 3, 2, 2, 2, 217, 228, 3, 2, 2, 2, 218,
	219, 12, 6, 2, 2, 219, 220, 5, 40, 21, 2, 220, 221, 5, 36, 19, 7, 221,
	227, 3, 2, 2, 2, 222, 223, 12, 5, 2, 2, 223, 224, 5, 38, 20, 2, 224, 225,
	5, 36, 19, 6, 225, 227, 3, 2, 2, 2, 226, 218, 3, 2, 2, 2, 226, 222, 3,
	2, 2, 2, 227, 230, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 228, 229, 3, 2, 2,
	2, 229, 37, 3, 2, 2, 2, 230, 228, 3, 2, 2, 2, 231, 232, 9, 3, 2, 2, 232,
	39, 3, 2, 2, 2, 233, 234, 9, 4, 2, 2, 234, 41, 3, 2, 2, 2, 235, 241, 5,
	64, 33, 2, 236, 237, 7, 48, 2, 2, 237, 238, 5, 36, 19, 2, 238, 239, 7,
	49, 2, 2, 239, 240, 5, 42, 22, 2, 240, 242, 3, 2, 2, 2, 241, 236, 3, 2,
	2, 2, 241, 242, 3, 2, 2, 2, 242, 43, 3, 2, 2, 2, 243, 246, 5, 42, 22, 2,
	244, 246, 7, 55, 2, 2, 245, 243, 3, 2, 2, 2, 245, 244, 3, 2, 2, 2, 246,
	45, 3, 2, 2, 2, 247, 248, 7, 52, 2, 2, 248, 47, 3, 2, 2, 2, 249, 250, 5,
	42, 22, 2, 250, 49, 3, 2, 2, 2, 251, 252, 8, 26, 1, 2, 252, 253, 5, 36,
	19, 2, 253, 262, 3, 2, 2, 2, 254, 255, 12, 4, 2, 2, 255, 256, 7, 32, 2,
	2, 256, 261, 5, 36, 19, 2, 257, 258, 12, 3, 2, 2, 258, 259, 7, 33, 2, 2,
	259, 261, 5, 36, 19, 2, 260, 254, 3, 2, 2, 2, 260, 257, 3, 2, 2, 2, 261,
	264, 3, 2, 2, 2, 262, 260, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 51, 3,
	2, 2, 2, 264, 262, 3, 2, 2, 2, 265, 266, 8, 27, 1, 2, 266, 267, 5, 50,
	26, 2, 267, 282, 3, 2, 2, 2, 268, 269, 12, 6, 2, 2, 269, 270, 7, 28, 2,
	2, 270, 281, 5, 50, 26, 2, 271, 272, 12, 5, 2, 2, 272, 273, 7, 30, 2, 2,
	273, 281, 5, 50, 26, 2, 274, 275, 12, 4, 2, 2, 275, 276, 7, 29, 2, 2, 276,
	281, 5, 50, 26, 2, 277, 278, 12, 3, 2, 2, 278, 279, 7, 31, 2, 2, 279, 281,
	5, 50, 26, 2, 280, 268, 3, 2, 2, 2, 280, 271, 3, 2, 2, 2, 280, 274, 3,
	2, 2, 2, 280, 277, 3, 2, 2, 2, 281, 284, 3, 2, 2, 2, 282, 280, 3, 2, 2,
	2, 282, 283, 3, 2, 2, 2, 283, 53, 3, 2, 2, 2, 284, 282, 3, 2, 2, 2, 285,
	286, 8, 28, 1, 2, 286, 287, 5, 52, 27, 2, 287, 296, 3, 2, 2, 2, 288, 289,
	12, 4, 2, 2, 289, 290, 7, 5, 2, 2, 290, 295, 5, 52, 27, 2, 291, 292, 12,
	3, 2, 2, 292, 293, 7, 6, 2, 2, 293, 295, 5, 52, 27, 2, 294, 288, 3, 2,
	2, 2, 294, 291, 3, 2, 2, 2, 295, 298, 3, 2, 2, 2, 296, 294, 3, 2, 2, 2,
	296, 297, 3, 2, 2, 2, 297, 55, 3, 2, 2, 2, 298, 296, 3, 2, 2, 2, 299, 300,
	8, 29, 1, 2, 300, 301, 5, 54, 28, 2, 301, 307, 3, 2, 2, 2, 302, 303, 12,
	3, 2, 2, 303, 304, 7, 41, 2, 2, 304, 306, 5, 54, 28, 2, 305, 302, 3, 2,
	2, 2, 306, 309, 3, 2, 2, 2, 307, 305, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2,
	308, 57, 3, 2, 2, 2, 309, 307, 3, 2, 2, 2, 310, 311, 8, 30, 1, 2, 311,
	312, 5, 56, 29, 2, 312, 318, 3, 2, 2, 2, 313, 314, 12, 3, 2, 2, 314, 315,
	7, 45, 2, 2, 315, 317, 5, 56, 29, 2, 316, 313, 3, 2, 2, 2, 317, 320, 3,
	2, 2, 2, 318, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 59, 3, 2, 2,
	2, 320, 318, 3, 2, 2, 2, 321, 322, 8, 31, 1, 2, 322, 323, 5, 58, 30, 2,
	323, 329, 3, 2, 2, 2, 324, 325, 12, 3, 2, 2, 325, 326, 7, 42, 2, 2, 326,
	328, 5, 58, 30, 2, 327, 324, 3, 2, 2, 2, 328, 331, 3, 2, 2, 2, 329, 327,
	3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 61, 3, 2, 2, 2, 331, 329, 3, 2,
	2, 2, 332, 333, 8, 32, 1, 2, 333, 334, 5, 60, 31, 2, 334, 340, 3, 2, 2,
	2, 335, 336, 12, 3, 2, 2, 336, 337, 7, 43, 2, 2, 337, 339, 5, 60, 31, 2,
	338, 335, 3, 2, 2, 2, 339, 342, 3, 2, 2, 2, 340, 338, 3, 2, 2, 2, 340,
	341, 3, 2, 2, 2, 341, 63, 3, 2, 2, 2, 342, 340, 3, 2, 2, 2, 343, 344, 8,
	33, 1, 2, 344, 345, 5, 62, 32, 2, 345, 351, 3, 2, 2, 2, 346, 347, 12, 3,
	2, 2, 347, 348, 7, 44, 2, 2, 348, 350, 5, 62, 32, 2, 349, 346, 3, 2, 2,
	2, 350, 353, 3, 2, 2, 2, 351, 349, 3, 2, 2, 2, 351, 352, 3, 2, 2, 2, 352,
	65, 3, 2, 2, 2, 353, 351, 3, 2, 2, 2, 354, 355, 9, 5, 2, 2, 355, 67, 3,
	2, 2, 2, 356, 357, 7, 19, 2, 2, 357, 358, 7, 53, 2, 2, 358, 359, 7, 22,
	2, 2, 359, 360, 5, 70, 36, 2, 360, 361, 7, 23, 2, 2, 361, 362, 5, 66, 34,
	2, 362, 363, 5, 16, 9, 2, 363, 69, 3, 2, 2, 2, 364, 365, 8, 36, 1, 2, 365,
	366, 5, 72, 37, 2, 366, 371, 3, 2, 2, 2, 367, 368, 12, 3, 2, 2, 368, 370,
	5, 72, 37, 2, 369, 367, 3, 2, 2, 2, 370, 373, 3, 2, 2, 2, 371, 369, 3,
	2, 2, 2, 371, 372, 3, 2, 2, 2, 372, 71, 3, 2, 2, 2, 373, 371, 3, 2, 2,
	2, 374, 376, 7, 51, 2, 2, 375, 374, 3, 2, 2, 2, 375, 376, 3, 2, 2, 2, 376,
	377, 3, 2, 2, 2, 377, 378, 5, 66, 34, 2, 378, 379, 7, 53, 2, 2, 379, 73,
	3, 2, 2, 2, 36, 76, 85, 91, 97, 108, 118, 127, 135, 146, 151, 163, 174,
	188, 191, 200, 206, 216, 226, 228, 241, 245, 260, 262, 280, 282, 294, 296,
	307, 318, 329, 340, 351, 371, 375,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'main'", "'var'", "'=='", "'!='", "'break'", "'continue'", "'else'",
	"'if'", "'int'", "'boolean'", "'string'", "'rune'", "'array'", "'return'",
	"'void'", "'while'", "'function'", "'true'", "'false'", "'('", "')'", "'['",
	"']'", "'{'", "'}'", "'<'", "'<='", "'>'", "'>='", "'<<'", "'>>'", "'+'",
	"'++'", "'-'", "'--'", "'*'", "'/'", "'%'", "'&'", "'|'", "'&&'", "'||'",
	"'^'", "'!'", "'~'", "'?'", "':'", "';'", "','", "'='",
}
var symbolicNames = []string{
	"", "", "", "", "", "Break", "Continue", "Else", "If", "Int", "Bool", "StringType",
	"Rune", "Array", "Return", "Void", "While", "Func", "True", "False", "Paren",
	"Thesis", "LeftBracket", "RightBracket", "LeftBrace", "RightBrace", "Less",
	"LessEqual", "Greater", "GreaterEqual", "LeftShift", "RightShift", "Plus",
	"PlusPlus", "Minus", "MinusMinus", "Star", "Div", "Mod", "And", "Or", "AndAnd",
	"OrOr", "Caret", "Not", "Tilde", "Question", "Colon", "Semi", "Comma",
	"Assign", "Identifier", "Constant", "DigitSequence", "StringLiteral", "Whitespace",
	"Newline", "BlockComment", "LineComment",
}

var ruleNames = []string{
	"program", "functionList", "mainFunction", "statementList", "declarationList",
	"statement", "assignmentStatement", "compoundStatement", "blockItemList",
	"blockItem", "declaration", "expressionStatement", "selectionStatement",
	"iterationStatement", "jumpStatement", "unaryOperator", "primaryExpression",
	"expression", "addop", "mulop", "conditionalExpression", "assignmentExpression",
	"assignmentOperator", "constantExpression", "shiftExpression", "relationalExpression",
	"equalityExpression", "andExpression", "exclusiveOrExpression", "inclusiveOrExpression",
	"logicalAndExpression", "logicalOrExpression", "typeSpecifier", "functionDefinition",
	"paramList", "param",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type cminusParser struct {
	*antlr.BaseParser
}

func NewcminusParser(input antlr.TokenStream) *cminusParser {
	this := new(cminusParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "cminus.g4"

	return this
}

// cminusParser tokens.
const (
	cminusParserEOF           = antlr.TokenEOF
	cminusParserT__0          = 1
	cminusParserT__1          = 2
	cminusParserT__2          = 3
	cminusParserT__3          = 4
	cminusParserBreak         = 5
	cminusParserContinue      = 6
	cminusParserElse          = 7
	cminusParserIf            = 8
	cminusParserInt           = 9
	cminusParserBool          = 10
	cminusParserStringType    = 11
	cminusParserRune          = 12
	cminusParserArray         = 13
	cminusParserReturn        = 14
	cminusParserVoid          = 15
	cminusParserWhile         = 16
	cminusParserFunc          = 17
	cminusParserTrue          = 18
	cminusParserFalse         = 19
	cminusParserParen         = 20
	cminusParserThesis        = 21
	cminusParserLeftBracket   = 22
	cminusParserRightBracket  = 23
	cminusParserLeftBrace     = 24
	cminusParserRightBrace    = 25
	cminusParserLess          = 26
	cminusParserLessEqual     = 27
	cminusParserGreater       = 28
	cminusParserGreaterEqual  = 29
	cminusParserLeftShift     = 30
	cminusParserRightShift    = 31
	cminusParserPlus          = 32
	cminusParserPlusPlus      = 33
	cminusParserMinus         = 34
	cminusParserMinusMinus    = 35
	cminusParserStar          = 36
	cminusParserDiv           = 37
	cminusParserMod           = 38
	cminusParserAnd           = 39
	cminusParserOr            = 40
	cminusParserAndAnd        = 41
	cminusParserOrOr          = 42
	cminusParserCaret         = 43
	cminusParserNot           = 44
	cminusParserTilde         = 45
	cminusParserQuestion      = 46
	cminusParserColon         = 47
	cminusParserSemi          = 48
	cminusParserComma         = 49
	cminusParserAssign        = 50
	cminusParserIdentifier    = 51
	cminusParserConstant      = 52
	cminusParserDigitSequence = 53
	cminusParserStringLiteral = 54
	cminusParserWhitespace    = 55
	cminusParserNewline       = 56
	cminusParserBlockComment  = 57
	cminusParserLineComment   = 58
)

// cminusParser rules.
const (
	cminusParserRULE_program               = 0
	cminusParserRULE_functionList          = 1
	cminusParserRULE_mainFunction          = 2
	cminusParserRULE_statementList         = 3
	cminusParserRULE_declarationList       = 4
	cminusParserRULE_statement             = 5
	cminusParserRULE_assignmentStatement   = 6
	cminusParserRULE_compoundStatement     = 7
	cminusParserRULE_blockItemList         = 8
	cminusParserRULE_blockItem             = 9
	cminusParserRULE_declaration           = 10
	cminusParserRULE_expressionStatement   = 11
	cminusParserRULE_selectionStatement    = 12
	cminusParserRULE_iterationStatement    = 13
	cminusParserRULE_jumpStatement         = 14
	cminusParserRULE_unaryOperator         = 15
	cminusParserRULE_primaryExpression     = 16
	cminusParserRULE_expression            = 17
	cminusParserRULE_addop                 = 18
	cminusParserRULE_mulop                 = 19
	cminusParserRULE_conditionalExpression = 20
	cminusParserRULE_assignmentExpression  = 21
	cminusParserRULE_assignmentOperator    = 22
	cminusParserRULE_constantExpression    = 23
	cminusParserRULE_shiftExpression       = 24
	cminusParserRULE_relationalExpression  = 25
	cminusParserRULE_equalityExpression    = 26
	cminusParserRULE_andExpression         = 27
	cminusParserRULE_exclusiveOrExpression = 28
	cminusParserRULE_inclusiveOrExpression = 29
	cminusParserRULE_logicalAndExpression  = 30
	cminusParserRULE_logicalOrExpression   = 31
	cminusParserRULE_typeSpecifier         = 32
	cminusParserRULE_functionDefinition    = 33
	cminusParserRULE_paramList             = 34
	cminusParserRULE_param                 = 35
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) MainFunction() IMainFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMainFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMainFunctionContext)
}

func (s *ProgramContext) FunctionList() IFunctionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionListContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *cminusParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, cminusParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.MainFunction()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == cminusParserFunc {
		{
			p.SetState(73)
			p.functionList(0)
		}

	}

	return localctx
}

// IFunctionListContext is an interface to support dynamic dispatch.
type IFunctionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionListContext differentiates from other interfaces.
	IsFunctionListContext()
}

type FunctionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionListContext() *FunctionListContext {
	var p = new(FunctionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_functionList
	return p
}

func (*FunctionListContext) IsFunctionListContext() {}

func NewFunctionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionListContext {
	var p = new(FunctionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_functionList

	return p
}

func (s *FunctionListContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionListContext) FunctionDefinition() IFunctionDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *FunctionListContext) FunctionList() IFunctionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionListContext)
}

func (s *FunctionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterFunctionList(s)
	}
}

func (s *FunctionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitFunctionList(s)
	}
}

func (p *cminusParser) FunctionList() (localctx IFunctionListContext) {
	return p.functionList(0)
}

func (p *cminusParser) functionList(_p int) (localctx IFunctionListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFunctionListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFunctionListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, cminusParserRULE_functionList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.FunctionDefinition()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFunctionListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_functionList)
			p.SetState(79)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(80)
				p.FunctionDefinition()
			}

		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IMainFunctionContext is an interface to support dynamic dispatch.
type IMainFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMainFunctionContext differentiates from other interfaces.
	IsMainFunctionContext()
}

type MainFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainFunctionContext() *MainFunctionContext {
	var p = new(MainFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_mainFunction
	return p
}

func (*MainFunctionContext) IsMainFunctionContext() {}

func NewMainFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainFunctionContext {
	var p = new(MainFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_mainFunction

	return p
}

func (s *MainFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *MainFunctionContext) Paren() antlr.TerminalNode {
	return s.GetToken(cminusParserParen, 0)
}

func (s *MainFunctionContext) Thesis() antlr.TerminalNode {
	return s.GetToken(cminusParserThesis, 0)
}

func (s *MainFunctionContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(cminusParserLeftBrace, 0)
}

func (s *MainFunctionContext) DeclarationList() IDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationListContext)
}

func (s *MainFunctionContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(cminusParserRightBrace, 0)
}

func (s *MainFunctionContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *MainFunctionContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *MainFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterMainFunction(s)
	}
}

func (s *MainFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitMainFunction(s)
	}
}

func (p *cminusParser) MainFunction() (localctx IMainFunctionContext) {
	localctx = NewMainFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, cminusParserRULE_mainFunction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(cminusParserT__0)
	}
	{
		p.SetState(87)
		p.Match(cminusParserParen)
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<cminusParserInt)|(1<<cminusParserBool)|(1<<cminusParserStringType)|(1<<cminusParserRune)|(1<<cminusParserVoid))) != 0) || _la == cminusParserComma {
		{
			p.SetState(88)
			p.paramList(0)
		}

	}
	{
		p.SetState(91)
		p.Match(cminusParserThesis)
	}
	{
		p.SetState(92)
		p.Match(cminusParserLeftBrace)
	}
	{
		p.SetState(93)
		p.declarationList(0)
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<cminusParserBreak)|(1<<cminusParserContinue)|(1<<cminusParserIf)|(1<<cminusParserReturn)|(1<<cminusParserWhile)|(1<<cminusParserParen)|(1<<cminusParserLeftBrace))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(cminusParserMinus-34))|(1<<(cminusParserSemi-34))|(1<<(cminusParserIdentifier-34))|(1<<(cminusParserConstant-34)))) != 0) {
		{
			p.SetState(94)
			p.statementList(0)
		}

	}
	{
		p.SetState(97)
		p.Match(cminusParserRightBrace)
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (p *cminusParser) StatementList() (localctx IStatementListContext) {
	return p.statementList(0)
}

func (p *cminusParser) statementList(_p int) (localctx IStatementListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IStatementListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, cminusParserRULE_statementList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Statement()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewStatementListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_statementList)
			p.SetState(102)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(103)
				p.Statement()
			}

		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IDeclarationListContext is an interface to support dynamic dispatch.
type IDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationListContext differentiates from other interfaces.
	IsDeclarationListContext()
}

type DeclarationListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationListContext() *DeclarationListContext {
	var p = new(DeclarationListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_declarationList
	return p
}

func (*DeclarationListContext) IsDeclarationListContext() {}

func NewDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationListContext {
	var p = new(DeclarationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_declarationList

	return p
}

func (s *DeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationListContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *DeclarationListContext) DeclarationList() IDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationListContext)
}

func (s *DeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterDeclarationList(s)
	}
}

func (s *DeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitDeclarationList(s)
	}
}

func (p *cminusParser) DeclarationList() (localctx IDeclarationListContext) {
	return p.declarationList(0)
}

func (p *cminusParser) declarationList(_p int) (localctx IDeclarationListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDeclarationListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDeclarationListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, cminusParserRULE_declarationList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Declaration()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDeclarationListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_declarationList)
			p.SetState(112)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(113)
				p.Declaration()
			}

		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CompoundStatement() ICompoundStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *StatementContext) ExpressionStatement() IExpressionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StatementContext) SelectionStatement() ISelectionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionStatementContext)
}

func (s *StatementContext) IterationStatement() IIterationStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIterationStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIterationStatementContext)
}

func (s *StatementContext) JumpStatement() IJumpStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJumpStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJumpStatementContext)
}

func (s *StatementContext) AssignmentStatement() IAssignmentStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *cminusParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, cminusParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(125)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case cminusParserLeftBrace:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.CompoundStatement()
		}

	case cminusParserParen, cminusParserMinus, cminusParserSemi, cminusParserConstant:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.ExpressionStatement()
		}

	case cminusParserIf:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(121)
			p.SelectionStatement()
		}

	case cminusParserWhile:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(122)
			p.IterationStatement()
		}

	case cminusParserBreak, cminusParserContinue, cminusParserReturn:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(123)
			p.JumpStatement()
		}

	case cminusParserIdentifier:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(124)
			p.AssignmentStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_assignmentStatement
	return p
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(cminusParserIdentifier, 0)
}

func (s *AssignmentStatementContext) Assign() antlr.TerminalNode {
	return s.GetToken(cminusParserAssign, 0)
}

func (s *AssignmentStatementContext) ExpressionStatement() IExpressionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (p *cminusParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, cminusParserRULE_assignmentStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(cminusParserIdentifier)
	}
	{
		p.SetState(128)
		p.Match(cminusParserAssign)
	}
	{
		p.SetState(129)
		p.ExpressionStatement()
	}

	return localctx
}

// ICompoundStatementContext is an interface to support dynamic dispatch.
type ICompoundStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompoundStatementContext differentiates from other interfaces.
	IsCompoundStatementContext()
}

type CompoundStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompoundStatementContext() *CompoundStatementContext {
	var p = new(CompoundStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_compoundStatement
	return p
}

func (*CompoundStatementContext) IsCompoundStatementContext() {}

func NewCompoundStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundStatementContext {
	var p = new(CompoundStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_compoundStatement

	return p
}

func (s *CompoundStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CompoundStatementContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(cminusParserLeftBrace, 0)
}

func (s *CompoundStatementContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(cminusParserRightBrace, 0)
}

func (s *CompoundStatementContext) BlockItemList() IBlockItemListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockItemListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockItemListContext)
}

func (s *CompoundStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompoundStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterCompoundStatement(s)
	}
}

func (s *CompoundStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitCompoundStatement(s)
	}
}

func (p *cminusParser) CompoundStatement() (localctx ICompoundStatementContext) {
	localctx = NewCompoundStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, cminusParserRULE_compoundStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(cminusParserLeftBrace)
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<cminusParserT__1)|(1<<cminusParserBreak)|(1<<cminusParserContinue)|(1<<cminusParserIf)|(1<<cminusParserReturn)|(1<<cminusParserWhile)|(1<<cminusParserParen)|(1<<cminusParserLeftBrace))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(cminusParserMinus-34))|(1<<(cminusParserSemi-34))|(1<<(cminusParserIdentifier-34))|(1<<(cminusParserConstant-34)))) != 0) {
		{
			p.SetState(132)
			p.blockItemList(0)
		}

	}
	{
		p.SetState(135)
		p.Match(cminusParserRightBrace)
	}

	return localctx
}

// IBlockItemListContext is an interface to support dynamic dispatch.
type IBlockItemListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockItemListContext differentiates from other interfaces.
	IsBlockItemListContext()
}

type BlockItemListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockItemListContext() *BlockItemListContext {
	var p = new(BlockItemListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_blockItemList
	return p
}

func (*BlockItemListContext) IsBlockItemListContext() {}

func NewBlockItemListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockItemListContext {
	var p = new(BlockItemListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_blockItemList

	return p
}

func (s *BlockItemListContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockItemListContext) BlockItem() IBlockItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockItemContext)
}

func (s *BlockItemListContext) BlockItemList() IBlockItemListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockItemListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockItemListContext)
}

func (s *BlockItemListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockItemListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockItemListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterBlockItemList(s)
	}
}

func (s *BlockItemListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitBlockItemList(s)
	}
}

func (p *cminusParser) BlockItemList() (localctx IBlockItemListContext) {
	return p.blockItemList(0)
}

func (p *cminusParser) blockItemList(_p int) (localctx IBlockItemListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBlockItemListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBlockItemListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, cminusParserRULE_blockItemList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.BlockItem()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBlockItemListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_blockItemList)
			p.SetState(140)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(141)
				p.BlockItem()
			}

		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IBlockItemContext is an interface to support dynamic dispatch.
type IBlockItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockItemContext differentiates from other interfaces.
	IsBlockItemContext()
}

type BlockItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockItemContext() *BlockItemContext {
	var p = new(BlockItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_blockItem
	return p
}

func (*BlockItemContext) IsBlockItemContext() {}

func NewBlockItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockItemContext {
	var p = new(BlockItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_blockItem

	return p
}

func (s *BlockItemContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockItemContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockItemContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *BlockItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterBlockItem(s)
	}
}

func (s *BlockItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitBlockItem(s)
	}
}

func (p *cminusParser) BlockItem() (localctx IBlockItemContext) {
	localctx = NewBlockItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, cminusParserRULE_blockItem)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case cminusParserBreak, cminusParserContinue, cminusParserIf, cminusParserReturn, cminusParserWhile, cminusParserParen, cminusParserLeftBrace, cminusParserMinus, cminusParserSemi, cminusParserIdentifier, cminusParserConstant:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.Statement()
		}

	case cminusParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Declaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(cminusParserIdentifier, 0)
}

func (s *DeclarationContext) Assign() antlr.TerminalNode {
	return s.GetToken(cminusParserAssign, 0)
}

func (s *DeclarationContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *DeclarationContext) Paren() antlr.TerminalNode {
	return s.GetToken(cminusParserParen, 0)
}

func (s *DeclarationContext) Constant() antlr.TerminalNode {
	return s.GetToken(cminusParserConstant, 0)
}

func (s *DeclarationContext) Thesis() antlr.TerminalNode {
	return s.GetToken(cminusParserThesis, 0)
}

func (s *DeclarationContext) Semi() antlr.TerminalNode {
	return s.GetToken(cminusParserSemi, 0)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *cminusParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, cminusParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(cminusParserT__1)
	}
	{
		p.SetState(152)
		p.Match(cminusParserIdentifier)
	}
	{
		p.SetState(153)
		p.Match(cminusParserAssign)
	}
	{
		p.SetState(154)
		p.TypeSpecifier()
	}
	{
		p.SetState(155)
		p.Match(cminusParserParen)
	}
	{
		p.SetState(156)
		p.Match(cminusParserConstant)
	}
	{
		p.SetState(157)
		p.Match(cminusParserThesis)
	}
	{
		p.SetState(158)
		p.Match(cminusParserSemi)
	}

	return localctx
}

// IExpressionStatementContext is an interface to support dynamic dispatch.
type IExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionStatementContext differentiates from other interfaces.
	IsExpressionStatementContext()
}

type ExpressionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStatementContext() *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_expressionStatement
	return p
}

func (*ExpressionStatementContext) IsExpressionStatementContext() {}

func NewExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_expressionStatement

	return p
}

func (s *ExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStatementContext) Semi() antlr.TerminalNode {
	return s.GetToken(cminusParserSemi, 0)
}

func (s *ExpressionStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (p *cminusParser) ExpressionStatement() (localctx IExpressionStatementContext) {
	localctx = NewExpressionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, cminusParserRULE_expressionStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == cminusParserParen || _la == cminusParserMinus || _la == cminusParserConstant {
		{
			p.SetState(160)
			p.expression(0)
		}

	}
	{
		p.SetState(163)
		p.Match(cminusParserSemi)
	}

	return localctx
}

// ISelectionStatementContext is an interface to support dynamic dispatch.
type ISelectionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionStatementContext differentiates from other interfaces.
	IsSelectionStatementContext()
}

type SelectionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionStatementContext() *SelectionStatementContext {
	var p = new(SelectionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_selectionStatement
	return p
}

func (*SelectionStatementContext) IsSelectionStatementContext() {}

func NewSelectionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionStatementContext {
	var p = new(SelectionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_selectionStatement

	return p
}

func (s *SelectionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionStatementContext) If() antlr.TerminalNode {
	return s.GetToken(cminusParserIf, 0)
}

func (s *SelectionStatementContext) Paren() antlr.TerminalNode {
	return s.GetToken(cminusParserParen, 0)
}

func (s *SelectionStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SelectionStatementContext) Thesis() antlr.TerminalNode {
	return s.GetToken(cminusParserThesis, 0)
}

func (s *SelectionStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *SelectionStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SelectionStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(cminusParserElse, 0)
}

func (s *SelectionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterSelectionStatement(s)
	}
}

func (s *SelectionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitSelectionStatement(s)
	}
}

func (p *cminusParser) SelectionStatement() (localctx ISelectionStatementContext) {
	localctx = NewSelectionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, cminusParserRULE_selectionStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(cminusParserIf)
	}
	{
		p.SetState(166)
		p.Match(cminusParserParen)
	}
	{
		p.SetState(167)
		p.expression(0)
	}
	{
		p.SetState(168)
		p.Match(cminusParserThesis)
	}
	{
		p.SetState(169)
		p.Statement()
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(170)
			p.Match(cminusParserElse)
		}
		{
			p.SetState(171)
			p.Statement()
		}

	}

	return localctx
}

// IIterationStatementContext is an interface to support dynamic dispatch.
type IIterationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIterationStatementContext differentiates from other interfaces.
	IsIterationStatementContext()
}

type IterationStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterationStatementContext() *IterationStatementContext {
	var p = new(IterationStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_iterationStatement
	return p
}

func (*IterationStatementContext) IsIterationStatementContext() {}

func NewIterationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterationStatementContext {
	var p = new(IterationStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_iterationStatement

	return p
}

func (s *IterationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IterationStatementContext) While() antlr.TerminalNode {
	return s.GetToken(cminusParserWhile, 0)
}

func (s *IterationStatementContext) Paren() antlr.TerminalNode {
	return s.GetToken(cminusParserParen, 0)
}

func (s *IterationStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IterationStatementContext) Thesis() antlr.TerminalNode {
	return s.GetToken(cminusParserThesis, 0)
}

func (s *IterationStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IterationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IterationStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterIterationStatement(s)
	}
}

func (s *IterationStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitIterationStatement(s)
	}
}

func (p *cminusParser) IterationStatement() (localctx IIterationStatementContext) {
	localctx = NewIterationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, cminusParserRULE_iterationStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(cminusParserWhile)
	}
	{
		p.SetState(175)
		p.Match(cminusParserParen)
	}
	{
		p.SetState(176)
		p.expression(0)
	}
	{
		p.SetState(177)
		p.Match(cminusParserThesis)
	}
	{
		p.SetState(178)
		p.Statement()
	}

	return localctx
}

// IJumpStatementContext is an interface to support dynamic dispatch.
type IJumpStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJumpStatementContext differentiates from other interfaces.
	IsJumpStatementContext()
}

type JumpStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJumpStatementContext() *JumpStatementContext {
	var p = new(JumpStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_jumpStatement
	return p
}

func (*JumpStatementContext) IsJumpStatementContext() {}

func NewJumpStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JumpStatementContext {
	var p = new(JumpStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_jumpStatement

	return p
}

func (s *JumpStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *JumpStatementContext) Continue() antlr.TerminalNode {
	return s.GetToken(cminusParserContinue, 0)
}

func (s *JumpStatementContext) Semi() antlr.TerminalNode {
	return s.GetToken(cminusParserSemi, 0)
}

func (s *JumpStatementContext) Break() antlr.TerminalNode {
	return s.GetToken(cminusParserBreak, 0)
}

func (s *JumpStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(cminusParserReturn, 0)
}

func (s *JumpStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *JumpStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JumpStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JumpStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterJumpStatement(s)
	}
}

func (s *JumpStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitJumpStatement(s)
	}
}

func (p *cminusParser) JumpStatement() (localctx IJumpStatementContext) {
	localctx = NewJumpStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, cminusParserRULE_jumpStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(189)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case cminusParserContinue:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Match(cminusParserContinue)
		}
		{
			p.SetState(181)
			p.Match(cminusParserSemi)
		}

	case cminusParserBreak:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(182)
			p.Match(cminusParserBreak)
		}
		{
			p.SetState(183)
			p.Match(cminusParserSemi)
		}

	case cminusParserReturn:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(184)
			p.Match(cminusParserReturn)
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == cminusParserParen || _la == cminusParserMinus || _la == cminusParserConstant {
			{
				p.SetState(185)
				p.expression(0)
			}

		}
		{
			p.SetState(188)
			p.Match(cminusParserSemi)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUnaryOperatorContext is an interface to support dynamic dispatch.
type IUnaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOperatorContext differentiates from other interfaces.
	IsUnaryOperatorContext()
}

type UnaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorContext() *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_unaryOperator
	return p
}

func (*UnaryOperatorContext) IsUnaryOperatorContext() {}

func NewUnaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_unaryOperator

	return p
}

func (s *UnaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperatorContext) And() antlr.TerminalNode {
	return s.GetToken(cminusParserAnd, 0)
}

func (s *UnaryOperatorContext) Star() antlr.TerminalNode {
	return s.GetToken(cminusParserStar, 0)
}

func (s *UnaryOperatorContext) Plus() antlr.TerminalNode {
	return s.GetToken(cminusParserPlus, 0)
}

func (s *UnaryOperatorContext) Minus() antlr.TerminalNode {
	return s.GetToken(cminusParserMinus, 0)
}

func (s *UnaryOperatorContext) Tilde() antlr.TerminalNode {
	return s.GetToken(cminusParserTilde, 0)
}

func (s *UnaryOperatorContext) Not() antlr.TerminalNode {
	return s.GetToken(cminusParserNot, 0)
}

func (s *UnaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitUnaryOperator(s)
	}
}

func (p *cminusParser) UnaryOperator() (localctx IUnaryOperatorContext) {
	localctx = NewUnaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, cminusParserRULE_unaryOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(cminusParserPlus-32))|(1<<(cminusParserMinus-32))|(1<<(cminusParserStar-32))|(1<<(cminusParserAnd-32))|(1<<(cminusParserNot-32))|(1<<(cminusParserTilde-32)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_primaryExpression
	return p
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(cminusParserIdentifier, 0)
}

func (s *PrimaryExpressionContext) Constant() antlr.TerminalNode {
	return s.GetToken(cminusParserConstant, 0)
}

func (s *PrimaryExpressionContext) AllStringLiteral() []antlr.TerminalNode {
	return s.GetTokens(cminusParserStringLiteral)
}

func (s *PrimaryExpressionContext) StringLiteral(i int) antlr.TerminalNode {
	return s.GetToken(cminusParserStringLiteral, i)
}

func (s *PrimaryExpressionContext) Paren() antlr.TerminalNode {
	return s.GetToken(cminusParserParen, 0)
}

func (s *PrimaryExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryExpressionContext) Thesis() antlr.TerminalNode {
	return s.GetToken(cminusParserThesis, 0)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (p *cminusParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, cminusParserRULE_primaryExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(204)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case cminusParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			p.Match(cminusParserIdentifier)
		}

	case cminusParserConstant:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(194)
			p.Match(cminusParserConstant)
		}

	case cminusParserStringLiteral:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == cminusParserStringLiteral {
			{
				p.SetState(195)
				p.Match(cminusParserStringLiteral)
			}

			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case cminusParserParen:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(200)
			p.Match(cminusParserParen)
		}
		{
			p.SetState(201)
			p.expression(0)
		}
		{
			p.SetState(202)
			p.Match(cminusParserThesis)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Minus() antlr.TerminalNode {
	return s.GetToken(cminusParserMinus, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) Paren() antlr.TerminalNode {
	return s.GetToken(cminusParserParen, 0)
}

func (s *ExpressionContext) Thesis() antlr.TerminalNode {
	return s.GetToken(cminusParserThesis, 0)
}

func (s *ExpressionContext) Constant() antlr.TerminalNode {
	return s.GetToken(cminusParserConstant, 0)
}

func (s *ExpressionContext) Mulop() IMulopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMulopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMulopContext)
}

func (s *ExpressionContext) Addop() IAddopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddopContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *cminusParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *cminusParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, cminusParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(214)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case cminusParserMinus:
		{
			p.SetState(207)
			p.Match(cminusParserMinus)
		}
		{
			p.SetState(208)
			p.expression(5)
		}

	case cminusParserParen:
		{
			p.SetState(209)
			p.Match(cminusParserParen)
		}
		{
			p.SetState(210)
			p.expression(0)
		}
		{
			p.SetState(211)
			p.Match(cminusParserThesis)
		}

	case cminusParserConstant:
		{
			p.SetState(213)
			p.Match(cminusParserConstant)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(224)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_expression)
				p.SetState(216)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(217)
					p.Mulop()
				}
				{
					p.SetState(218)
					p.expression(5)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_expression)
				p.SetState(220)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(221)
					p.Addop()
				}
				{
					p.SetState(222)
					p.expression(4)
				}

			}

		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// IAddopContext is an interface to support dynamic dispatch.
type IAddopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddopContext differentiates from other interfaces.
	IsAddopContext()
}

type AddopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddopContext() *AddopContext {
	var p = new(AddopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_addop
	return p
}

func (*AddopContext) IsAddopContext() {}

func NewAddopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddopContext {
	var p = new(AddopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_addop

	return p
}

func (s *AddopContext) GetParser() antlr.Parser { return s.parser }

func (s *AddopContext) Plus() antlr.TerminalNode {
	return s.GetToken(cminusParserPlus, 0)
}

func (s *AddopContext) Minus() antlr.TerminalNode {
	return s.GetToken(cminusParserMinus, 0)
}

func (s *AddopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterAddop(s)
	}
}

func (s *AddopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitAddop(s)
	}
}

func (p *cminusParser) Addop() (localctx IAddopContext) {
	localctx = NewAddopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, cminusParserRULE_addop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		_la = p.GetTokenStream().LA(1)

		if !(_la == cminusParserPlus || _la == cminusParserMinus) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMulopContext is an interface to support dynamic dispatch.
type IMulopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMulopContext differentiates from other interfaces.
	IsMulopContext()
}

type MulopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulopContext() *MulopContext {
	var p = new(MulopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_mulop
	return p
}

func (*MulopContext) IsMulopContext() {}

func NewMulopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulopContext {
	var p = new(MulopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_mulop

	return p
}

func (s *MulopContext) GetParser() antlr.Parser { return s.parser }

func (s *MulopContext) Star() antlr.TerminalNode {
	return s.GetToken(cminusParserStar, 0)
}

func (s *MulopContext) Div() antlr.TerminalNode {
	return s.GetToken(cminusParserDiv, 0)
}

func (s *MulopContext) Mod() antlr.TerminalNode {
	return s.GetToken(cminusParserMod, 0)
}

func (s *MulopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterMulop(s)
	}
}

func (s *MulopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitMulop(s)
	}
}

func (p *cminusParser) Mulop() (localctx IMulopContext) {
	localctx = NewMulopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, cminusParserRULE_mulop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(cminusParserStar-36))|(1<<(cminusParserDiv-36))|(1<<(cminusParserMod-36)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConditionalExpressionContext is an interface to support dynamic dispatch.
type IConditionalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalExpressionContext differentiates from other interfaces.
	IsConditionalExpressionContext()
}

type ConditionalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalExpressionContext() *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_conditionalExpression
	return p
}

func (*ConditionalExpressionContext) IsConditionalExpressionContext() {}

func NewConditionalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalExpressionContext {
	var p = new(ConditionalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_conditionalExpression

	return p
}

func (s *ConditionalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalExpressionContext) LogicalOrExpression() ILogicalOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalOrExpressionContext)
}

func (s *ConditionalExpressionContext) Question() antlr.TerminalNode {
	return s.GetToken(cminusParserQuestion, 0)
}

func (s *ConditionalExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalExpressionContext) Colon() antlr.TerminalNode {
	return s.GetToken(cminusParserColon, 0)
}

func (s *ConditionalExpressionContext) ConditionalExpression() IConditionalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalExpressionContext)
}

func (s *ConditionalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterConditionalExpression(s)
	}
}

func (s *ConditionalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitConditionalExpression(s)
	}
}

func (p *cminusParser) ConditionalExpression() (localctx IConditionalExpressionContext) {
	localctx = NewConditionalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, cminusParserRULE_conditionalExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.logicalOrExpression(0)
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == cminusParserQuestion {
		{
			p.SetState(234)
			p.Match(cminusParserQuestion)
		}
		{
			p.SetState(235)
			p.expression(0)
		}
		{
			p.SetState(236)
			p.Match(cminusParserColon)
		}
		{
			p.SetState(237)
			p.ConditionalExpression()
		}

	}

	return localctx
}

// IAssignmentExpressionContext is an interface to support dynamic dispatch.
type IAssignmentExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentExpressionContext differentiates from other interfaces.
	IsAssignmentExpressionContext()
}

type AssignmentExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentExpressionContext() *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_assignmentExpression
	return p
}

func (*AssignmentExpressionContext) IsAssignmentExpressionContext() {}

func NewAssignmentExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_assignmentExpression

	return p
}

func (s *AssignmentExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentExpressionContext) ConditionalExpression() IConditionalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalExpressionContext)
}

func (s *AssignmentExpressionContext) DigitSequence() antlr.TerminalNode {
	return s.GetToken(cminusParserDigitSequence, 0)
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitAssignmentExpression(s)
	}
}

func (p *cminusParser) AssignmentExpression() (localctx IAssignmentExpressionContext) {
	localctx = NewAssignmentExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, cminusParserRULE_assignmentExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(243)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case cminusParserParen, cminusParserMinus, cminusParserConstant:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.ConditionalExpression()
		}

	case cminusParserDigitSequence:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.Match(cminusParserDigitSequence)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssignmentOperatorContext is an interface to support dynamic dispatch.
type IAssignmentOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentOperatorContext differentiates from other interfaces.
	IsAssignmentOperatorContext()
}

type AssignmentOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOperatorContext() *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_assignmentOperator
	return p
}

func (*AssignmentOperatorContext) IsAssignmentOperatorContext() {}

func NewAssignmentOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_assignmentOperator

	return p
}

func (s *AssignmentOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentOperatorContext) Assign() antlr.TerminalNode {
	return s.GetToken(cminusParserAssign, 0)
}

func (s *AssignmentOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitAssignmentOperator(s)
	}
}

func (p *cminusParser) AssignmentOperator() (localctx IAssignmentOperatorContext) {
	localctx = NewAssignmentOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, cminusParserRULE_assignmentOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(cminusParserAssign)
	}

	return localctx
}

// IConstantExpressionContext is an interface to support dynamic dispatch.
type IConstantExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantExpressionContext differentiates from other interfaces.
	IsConstantExpressionContext()
}

type ConstantExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantExpressionContext() *ConstantExpressionContext {
	var p = new(ConstantExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_constantExpression
	return p
}

func (*ConstantExpressionContext) IsConstantExpressionContext() {}

func NewConstantExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantExpressionContext {
	var p = new(ConstantExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_constantExpression

	return p
}

func (s *ConstantExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantExpressionContext) ConditionalExpression() IConditionalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalExpressionContext)
}

func (s *ConstantExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterConstantExpression(s)
	}
}

func (s *ConstantExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitConstantExpression(s)
	}
}

func (p *cminusParser) ConstantExpression() (localctx IConstantExpressionContext) {
	localctx = NewConstantExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, cminusParserRULE_constantExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.ConditionalExpression()
	}

	return localctx
}

// IShiftExpressionContext is an interface to support dynamic dispatch.
type IShiftExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShiftExpressionContext differentiates from other interfaces.
	IsShiftExpressionContext()
}

type ShiftExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShiftExpressionContext() *ShiftExpressionContext {
	var p = new(ShiftExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_shiftExpression
	return p
}

func (*ShiftExpressionContext) IsShiftExpressionContext() {}

func NewShiftExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShiftExpressionContext {
	var p = new(ShiftExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_shiftExpression

	return p
}

func (s *ShiftExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ShiftExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ShiftExpressionContext) ShiftExpression() IShiftExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShiftExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShiftExpressionContext)
}

func (s *ShiftExpressionContext) LeftShift() antlr.TerminalNode {
	return s.GetToken(cminusParserLeftShift, 0)
}

func (s *ShiftExpressionContext) RightShift() antlr.TerminalNode {
	return s.GetToken(cminusParserRightShift, 0)
}

func (s *ShiftExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShiftExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterShiftExpression(s)
	}
}

func (s *ShiftExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitShiftExpression(s)
	}
}

func (p *cminusParser) ShiftExpression() (localctx IShiftExpressionContext) {
	return p.shiftExpression(0)
}

func (p *cminusParser) shiftExpression(_p int) (localctx IShiftExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewShiftExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IShiftExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 48
	p.EnterRecursionRule(localctx, 48, cminusParserRULE_shiftExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.expression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(258)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
			case 1:
				localctx = NewShiftExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_shiftExpression)
				p.SetState(252)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(253)
					p.Match(cminusParserLeftShift)
				}
				{
					p.SetState(254)
					p.expression(0)
				}

			case 2:
				localctx = NewShiftExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_shiftExpression)
				p.SetState(255)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(256)
					p.Match(cminusParserRightShift)
				}
				{
					p.SetState(257)
					p.expression(0)
				}

			}

		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_relationalExpression
	return p
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) ShiftExpression() IShiftExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShiftExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShiftExpressionContext)
}

func (s *RelationalExpressionContext) RelationalExpression() IRelationalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *RelationalExpressionContext) Less() antlr.TerminalNode {
	return s.GetToken(cminusParserLess, 0)
}

func (s *RelationalExpressionContext) Greater() antlr.TerminalNode {
	return s.GetToken(cminusParserGreater, 0)
}

func (s *RelationalExpressionContext) LessEqual() antlr.TerminalNode {
	return s.GetToken(cminusParserLessEqual, 0)
}

func (s *RelationalExpressionContext) GreaterEqual() antlr.TerminalNode {
	return s.GetToken(cminusParserGreaterEqual, 0)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (p *cminusParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	return p.relationalExpression(0)
}

func (p *cminusParser) relationalExpression(_p int) (localctx IRelationalExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelationalExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, cminusParserRULE_relationalExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.shiftExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(278)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				localctx = NewRelationalExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_relationalExpression)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(267)
					p.Match(cminusParserLess)
				}
				{
					p.SetState(268)
					p.shiftExpression(0)
				}

			case 2:
				localctx = NewRelationalExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_relationalExpression)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(270)
					p.Match(cminusParserGreater)
				}
				{
					p.SetState(271)
					p.shiftExpression(0)
				}

			case 3:
				localctx = NewRelationalExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_relationalExpression)
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(273)
					p.Match(cminusParserLessEqual)
				}
				{
					p.SetState(274)
					p.shiftExpression(0)
				}

			case 4:
				localctx = NewRelationalExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_relationalExpression)
				p.SetState(275)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(276)
					p.Match(cminusParserGreaterEqual)
				}
				{
					p.SetState(277)
					p.shiftExpression(0)
				}

			}

		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
}

// IEqualityExpressionContext is an interface to support dynamic dispatch.
type IEqualityExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualityExpressionContext differentiates from other interfaces.
	IsEqualityExpressionContext()
}

type EqualityExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExpressionContext() *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_equalityExpression
	return p
}

func (*EqualityExpressionContext) IsEqualityExpressionContext() {}

func NewEqualityExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_equalityExpression

	return p
}

func (s *EqualityExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExpressionContext) RelationalExpression() IRelationalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *EqualityExpressionContext) EqualityExpression() IEqualityExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

func (p *cminusParser) EqualityExpression() (localctx IEqualityExpressionContext) {
	return p.equalityExpression(0)
}

func (p *cminusParser) equalityExpression(_p int) (localctx IEqualityExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewEqualityExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEqualityExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, cminusParserRULE_equalityExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.relationalExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(292)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEqualityExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_equalityExpression)
				p.SetState(286)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(287)
					p.Match(cminusParserT__2)
				}
				{
					p.SetState(288)
					p.relationalExpression(0)
				}

			case 2:
				localctx = NewEqualityExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_equalityExpression)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(290)
					p.Match(cminusParserT__3)
				}
				{
					p.SetState(291)
					p.relationalExpression(0)
				}

			}

		}
		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}

	return localctx
}

// IAndExpressionContext is an interface to support dynamic dispatch.
type IAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndExpressionContext differentiates from other interfaces.
	IsAndExpressionContext()
}

type AndExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExpressionContext() *AndExpressionContext {
	var p = new(AndExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_andExpression
	return p
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_andExpression

	return p
}

func (s *AndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExpressionContext) EqualityExpression() IEqualityExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityExpressionContext)
}

func (s *AndExpressionContext) AndExpression() IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *AndExpressionContext) And() antlr.TerminalNode {
	return s.GetToken(cminusParserAnd, 0)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (p *cminusParser) AndExpression() (localctx IAndExpressionContext) {
	return p.andExpression(0)
}

func (p *cminusParser) andExpression(_p int) (localctx IAndExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAndExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, cminusParserRULE_andExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.equalityExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAndExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_andExpression)
			p.SetState(300)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(301)
				p.Match(cminusParserAnd)
			}
			{
				p.SetState(302)
				p.equalityExpression(0)
			}

		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}

	return localctx
}

// IExclusiveOrExpressionContext is an interface to support dynamic dispatch.
type IExclusiveOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExclusiveOrExpressionContext differentiates from other interfaces.
	IsExclusiveOrExpressionContext()
}

type ExclusiveOrExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExclusiveOrExpressionContext() *ExclusiveOrExpressionContext {
	var p = new(ExclusiveOrExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_exclusiveOrExpression
	return p
}

func (*ExclusiveOrExpressionContext) IsExclusiveOrExpressionContext() {}

func NewExclusiveOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExclusiveOrExpressionContext {
	var p = new(ExclusiveOrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_exclusiveOrExpression

	return p
}

func (s *ExclusiveOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExclusiveOrExpressionContext) AndExpression() IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *ExclusiveOrExpressionContext) ExclusiveOrExpression() IExclusiveOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExclusiveOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExclusiveOrExpressionContext)
}

func (s *ExclusiveOrExpressionContext) Caret() antlr.TerminalNode {
	return s.GetToken(cminusParserCaret, 0)
}

func (s *ExclusiveOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExclusiveOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExclusiveOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterExclusiveOrExpression(s)
	}
}

func (s *ExclusiveOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitExclusiveOrExpression(s)
	}
}

func (p *cminusParser) ExclusiveOrExpression() (localctx IExclusiveOrExpressionContext) {
	return p.exclusiveOrExpression(0)
}

func (p *cminusParser) exclusiveOrExpression(_p int) (localctx IExclusiveOrExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExclusiveOrExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExclusiveOrExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, cminusParserRULE_exclusiveOrExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.andExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExclusiveOrExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_exclusiveOrExpression)
			p.SetState(311)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(312)
				p.Match(cminusParserCaret)
			}
			{
				p.SetState(313)
				p.andExpression(0)
			}

		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}

	return localctx
}

// IInclusiveOrExpressionContext is an interface to support dynamic dispatch.
type IInclusiveOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInclusiveOrExpressionContext differentiates from other interfaces.
	IsInclusiveOrExpressionContext()
}

type InclusiveOrExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInclusiveOrExpressionContext() *InclusiveOrExpressionContext {
	var p = new(InclusiveOrExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_inclusiveOrExpression
	return p
}

func (*InclusiveOrExpressionContext) IsInclusiveOrExpressionContext() {}

func NewInclusiveOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InclusiveOrExpressionContext {
	var p = new(InclusiveOrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_inclusiveOrExpression

	return p
}

func (s *InclusiveOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *InclusiveOrExpressionContext) ExclusiveOrExpression() IExclusiveOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExclusiveOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExclusiveOrExpressionContext)
}

func (s *InclusiveOrExpressionContext) InclusiveOrExpression() IInclusiveOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInclusiveOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInclusiveOrExpressionContext)
}

func (s *InclusiveOrExpressionContext) Or() antlr.TerminalNode {
	return s.GetToken(cminusParserOr, 0)
}

func (s *InclusiveOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InclusiveOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InclusiveOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterInclusiveOrExpression(s)
	}
}

func (s *InclusiveOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitInclusiveOrExpression(s)
	}
}

func (p *cminusParser) InclusiveOrExpression() (localctx IInclusiveOrExpressionContext) {
	return p.inclusiveOrExpression(0)
}

func (p *cminusParser) inclusiveOrExpression(_p int) (localctx IInclusiveOrExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewInclusiveOrExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IInclusiveOrExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 58
	p.EnterRecursionRule(localctx, 58, cminusParserRULE_inclusiveOrExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.exclusiveOrExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewInclusiveOrExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_inclusiveOrExpression)
			p.SetState(322)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(323)
				p.Match(cminusParserOr)
			}
			{
				p.SetState(324)
				p.exclusiveOrExpression(0)
			}

		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
	}

	return localctx
}

// ILogicalAndExpressionContext is an interface to support dynamic dispatch.
type ILogicalAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalAndExpressionContext differentiates from other interfaces.
	IsLogicalAndExpressionContext()
}

type LogicalAndExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalAndExpressionContext() *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_logicalAndExpression
	return p
}

func (*LogicalAndExpressionContext) IsLogicalAndExpressionContext() {}

func NewLogicalAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_logicalAndExpression

	return p
}

func (s *LogicalAndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndExpressionContext) InclusiveOrExpression() IInclusiveOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInclusiveOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInclusiveOrExpressionContext)
}

func (s *LogicalAndExpressionContext) LogicalAndExpression() ILogicalAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalAndExpressionContext)
}

func (s *LogicalAndExpressionContext) AndAnd() antlr.TerminalNode {
	return s.GetToken(cminusParserAndAnd, 0)
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitLogicalAndExpression(s)
	}
}

func (p *cminusParser) LogicalAndExpression() (localctx ILogicalAndExpressionContext) {
	return p.logicalAndExpression(0)
}

func (p *cminusParser) logicalAndExpression(_p int) (localctx ILogicalAndExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewLogicalAndExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogicalAndExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 60
	p.EnterRecursionRule(localctx, 60, cminusParserRULE_logicalAndExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(331)
		p.inclusiveOrExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalAndExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_logicalAndExpression)
			p.SetState(333)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(334)
				p.Match(cminusParserAndAnd)
			}
			{
				p.SetState(335)
				p.inclusiveOrExpression(0)
			}

		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}

	return localctx
}

// ILogicalOrExpressionContext is an interface to support dynamic dispatch.
type ILogicalOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalOrExpressionContext differentiates from other interfaces.
	IsLogicalOrExpressionContext()
}

type LogicalOrExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOrExpressionContext() *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_logicalOrExpression
	return p
}

func (*LogicalOrExpressionContext) IsLogicalOrExpressionContext() {}

func NewLogicalOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_logicalOrExpression

	return p
}

func (s *LogicalOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrExpressionContext) LogicalAndExpression() ILogicalAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalAndExpressionContext)
}

func (s *LogicalOrExpressionContext) LogicalOrExpression() ILogicalOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalOrExpressionContext)
}

func (s *LogicalOrExpressionContext) OrOr() antlr.TerminalNode {
	return s.GetToken(cminusParserOrOr, 0)
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitLogicalOrExpression(s)
	}
}

func (p *cminusParser) LogicalOrExpression() (localctx ILogicalOrExpressionContext) {
	return p.logicalOrExpression(0)
}

func (p *cminusParser) logicalOrExpression(_p int) (localctx ILogicalOrExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewLogicalOrExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogicalOrExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, cminusParserRULE_logicalOrExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.logicalAndExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalOrExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_logicalOrExpression)
			p.SetState(344)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(345)
				p.Match(cminusParserOrOr)
			}
			{
				p.SetState(346)
				p.logicalAndExpression(0)
			}

		}
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeSpecifierContext is an interface to support dynamic dispatch.
type ITypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecifierContext differentiates from other interfaces.
	IsTypeSpecifierContext()
}

type TypeSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierContext() *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_typeSpecifier
	return p
}

func (*TypeSpecifierContext) IsTypeSpecifierContext() {}

func NewTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_typeSpecifier

	return p
}

func (s *TypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierContext) Void() antlr.TerminalNode {
	return s.GetToken(cminusParserVoid, 0)
}

func (s *TypeSpecifierContext) Rune() antlr.TerminalNode {
	return s.GetToken(cminusParserRune, 0)
}

func (s *TypeSpecifierContext) Int() antlr.TerminalNode {
	return s.GetToken(cminusParserInt, 0)
}

func (s *TypeSpecifierContext) StringType() antlr.TerminalNode {
	return s.GetToken(cminusParserStringType, 0)
}

func (s *TypeSpecifierContext) Bool() antlr.TerminalNode {
	return s.GetToken(cminusParserBool, 0)
}

func (s *TypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterTypeSpecifier(s)
	}
}

func (s *TypeSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitTypeSpecifier(s)
	}
}

func (p *cminusParser) TypeSpecifier() (localctx ITypeSpecifierContext) {
	localctx = NewTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, cminusParserRULE_typeSpecifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<cminusParserInt)|(1<<cminusParserBool)|(1<<cminusParserStringType)|(1<<cminusParserRune)|(1<<cminusParserVoid))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_functionDefinition
	return p
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) Func() antlr.TerminalNode {
	return s.GetToken(cminusParserFunc, 0)
}

func (s *FunctionDefinitionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(cminusParserIdentifier, 0)
}

func (s *FunctionDefinitionContext) Paren() antlr.TerminalNode {
	return s.GetToken(cminusParserParen, 0)
}

func (s *FunctionDefinitionContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *FunctionDefinitionContext) Thesis() antlr.TerminalNode {
	return s.GetToken(cminusParserThesis, 0)
}

func (s *FunctionDefinitionContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *FunctionDefinitionContext) CompoundStatement() ICompoundStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundStatementContext)
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitFunctionDefinition(s)
	}
}

func (p *cminusParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, cminusParserRULE_functionDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.Match(cminusParserFunc)
	}
	{
		p.SetState(355)
		p.Match(cminusParserIdentifier)
	}
	{
		p.SetState(356)
		p.Match(cminusParserParen)
	}
	{
		p.SetState(357)
		p.paramList(0)
	}
	{
		p.SetState(358)
		p.Match(cminusParserThesis)
	}
	{
		p.SetState(359)
		p.TypeSpecifier()
	}
	{
		p.SetState(360)
		p.CompoundStatement()
	}

	return localctx
}

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_paramList
	return p
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) Param() IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamListContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (p *cminusParser) ParamList() (localctx IParamListContext) {
	return p.paramList(0)
}

func (p *cminusParser) paramList(_p int) (localctx IParamListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewParamListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IParamListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 68
	p.EnterRecursionRule(localctx, 68, cminusParserRULE_paramList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Param()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParamListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, cminusParserRULE_paramList)
			p.SetState(365)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(366)
				p.Param()
			}

		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = cminusParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = cminusParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *ParamContext) Identifier() antlr.TerminalNode {
	return s.GetToken(cminusParserIdentifier, 0)
}

func (s *ParamContext) Comma() antlr.TerminalNode {
	return s.GetToken(cminusParserComma, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(cminusListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *cminusParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, cminusParserRULE_param)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == cminusParserComma {
		{
			p.SetState(372)
			p.Match(cminusParserComma)
		}

	}
	{
		p.SetState(375)
		p.TypeSpecifier()
	}
	{
		p.SetState(376)
		p.Match(cminusParserIdentifier)
	}

	return localctx
}

func (p *cminusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *FunctionListContext = nil
		if localctx != nil {
			t = localctx.(*FunctionListContext)
		}
		return p.FunctionList_Sempred(t, predIndex)

	case 3:
		var t *StatementListContext = nil
		if localctx != nil {
			t = localctx.(*StatementListContext)
		}
		return p.StatementList_Sempred(t, predIndex)

	case 4:
		var t *DeclarationListContext = nil
		if localctx != nil {
			t = localctx.(*DeclarationListContext)
		}
		return p.DeclarationList_Sempred(t, predIndex)

	case 8:
		var t *BlockItemListContext = nil
		if localctx != nil {
			t = localctx.(*BlockItemListContext)
		}
		return p.BlockItemList_Sempred(t, predIndex)

	case 17:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 24:
		var t *ShiftExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ShiftExpressionContext)
		}
		return p.ShiftExpression_Sempred(t, predIndex)

	case 25:
		var t *RelationalExpressionContext = nil
		if localctx != nil {
			t = localctx.(*RelationalExpressionContext)
		}
		return p.RelationalExpression_Sempred(t, predIndex)

	case 26:
		var t *EqualityExpressionContext = nil
		if localctx != nil {
			t = localctx.(*EqualityExpressionContext)
		}
		return p.EqualityExpression_Sempred(t, predIndex)

	case 27:
		var t *AndExpressionContext = nil
		if localctx != nil {
			t = localctx.(*AndExpressionContext)
		}
		return p.AndExpression_Sempred(t, predIndex)

	case 28:
		var t *ExclusiveOrExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExclusiveOrExpressionContext)
		}
		return p.ExclusiveOrExpression_Sempred(t, predIndex)

	case 29:
		var t *InclusiveOrExpressionContext = nil
		if localctx != nil {
			t = localctx.(*InclusiveOrExpressionContext)
		}
		return p.InclusiveOrExpression_Sempred(t, predIndex)

	case 30:
		var t *LogicalAndExpressionContext = nil
		if localctx != nil {
			t = localctx.(*LogicalAndExpressionContext)
		}
		return p.LogicalAndExpression_Sempred(t, predIndex)

	case 31:
		var t *LogicalOrExpressionContext = nil
		if localctx != nil {
			t = localctx.(*LogicalOrExpressionContext)
		}
		return p.LogicalOrExpression_Sempred(t, predIndex)

	case 34:
		var t *ParamListContext = nil
		if localctx != nil {
			t = localctx.(*ParamListContext)
		}
		return p.ParamList_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *cminusParser) FunctionList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) StatementList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) DeclarationList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) BlockItemList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) ShiftExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) RelationalExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) EqualityExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) AndExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) ExclusiveOrExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 15:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) InclusiveOrExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 16:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) LogicalAndExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 17:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) LogicalOrExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 18:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *cminusParser) ParamList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 19:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
