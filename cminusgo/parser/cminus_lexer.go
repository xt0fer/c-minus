// Code generated from cminus.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 60, 567,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28,
	3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3,
	32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3,
	41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45,
	3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3,
	50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 7, 52, 351, 10, 52, 12, 52, 14,
	52, 354, 11, 52, 3, 53, 3, 53, 5, 53, 358, 10, 53, 3, 54, 3, 54, 3, 55,
	3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3,
	56, 5, 56, 374, 10, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58,
	5, 58, 383, 10, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 5, 59, 390, 10,
	59, 3, 60, 3, 60, 3, 60, 6, 60, 395, 10, 60, 13, 60, 14, 60, 396, 3, 61,
	3, 61, 5, 61, 401, 10, 61, 3, 62, 3, 62, 7, 62, 405, 10, 62, 12, 62, 14,
	62, 408, 11, 62, 3, 63, 3, 63, 7, 63, 412, 10, 63, 12, 63, 14, 63, 415,
	11, 63, 3, 64, 3, 64, 6, 64, 419, 10, 64, 13, 64, 14, 64, 420, 3, 65, 3,
	65, 3, 65, 3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 70,
	6, 70, 435, 10, 70, 13, 70, 14, 70, 436, 3, 71, 6, 71, 440, 10, 71, 13,
	71, 14, 71, 441, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72,
	3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3,
	72, 3, 72, 3, 72, 3, 72, 5, 72, 466, 10, 72, 3, 73, 6, 73, 469, 10, 73,
	13, 73, 14, 73, 470, 3, 74, 3, 74, 5, 74, 475, 10, 74, 3, 75, 3, 75, 3,
	75, 3, 75, 5, 75, 481, 10, 75, 3, 76, 3, 76, 3, 76, 3, 77, 3, 77, 3, 77,
	3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 5, 77, 497, 10,
	77, 3, 78, 3, 78, 3, 78, 3, 78, 6, 78, 503, 10, 78, 13, 78, 14, 78, 504,
	3, 79, 3, 79, 5, 79, 509, 10, 79, 3, 79, 3, 79, 3, 80, 6, 80, 514, 10,
	80, 13, 80, 14, 80, 515, 3, 81, 3, 81, 3, 81, 3, 81, 3, 81, 3, 81, 3, 81,
	5, 81, 525, 10, 81, 3, 82, 6, 82, 528, 10, 82, 13, 82, 14, 82, 529, 3,
	82, 3, 82, 3, 83, 3, 83, 5, 83, 536, 10, 83, 3, 83, 5, 83, 539, 10, 83,
	3, 83, 3, 83, 3, 84, 3, 84, 3, 84, 3, 84, 7, 84, 547, 10, 84, 12, 84, 14,
	84, 550, 11, 84, 3, 84, 3, 84, 3, 84, 3, 84, 3, 84, 3, 85, 3, 85, 3, 85,
	3, 85, 7, 85, 561, 10, 85, 12, 85, 14, 85, 564, 11, 85, 3, 85, 3, 85, 3,
	548, 2, 86, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20,
	39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29,
	57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38,
	75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47,
	93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 2, 107, 2, 109,
	2, 111, 2, 113, 2, 115, 54, 117, 2, 119, 2, 121, 2, 123, 2, 125, 2, 127,
	2, 129, 2, 131, 2, 133, 2, 135, 2, 137, 2, 139, 55, 141, 2, 143, 2, 145,
	2, 147, 2, 149, 2, 151, 2, 153, 2, 155, 2, 157, 56, 159, 2, 161, 2, 163,
	57, 165, 58, 167, 59, 169, 60, 3, 2, 16, 5, 2, 67, 92, 97, 97, 99, 124,
	3, 2, 50, 59, 4, 2, 68, 68, 100, 100, 3, 2, 50, 51, 4, 2, 90, 90, 122,
	122, 3, 2, 51, 59, 3, 2, 50, 57, 5, 2, 50, 59, 67, 72, 99, 104, 4, 2, 45,
	45, 47, 47, 6, 2, 12, 12, 15, 15, 41, 41, 94, 94, 12, 2, 36, 36, 41, 41,
	65, 65, 94, 94, 99, 100, 104, 104, 112, 112, 116, 116, 118, 118, 120, 120,
	6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12,
	15, 15, 2, 577, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3,
	2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47,
	3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2,
	55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2,
	2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2,
	2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2,
	2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3,
	2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93,
	3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2,
	101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 139, 3, 2,
	2, 2, 2, 157, 3, 2, 2, 2, 2, 163, 3, 2, 2, 2, 2, 165, 3, 2, 2, 2, 2, 167,
	3, 2, 2, 2, 2, 169, 3, 2, 2, 2, 3, 171, 3, 2, 2, 2, 5, 174, 3, 2, 2, 2,
	7, 177, 3, 2, 2, 2, 9, 183, 3, 2, 2, 2, 11, 192, 3, 2, 2, 2, 13, 197, 3,
	2, 2, 2, 15, 200, 3, 2, 2, 2, 17, 204, 3, 2, 2, 2, 19, 208, 3, 2, 2, 2,
	21, 216, 3, 2, 2, 2, 23, 223, 3, 2, 2, 2, 25, 228, 3, 2, 2, 2, 27, 234,
	3, 2, 2, 2, 29, 241, 3, 2, 2, 2, 31, 246, 3, 2, 2, 2, 33, 252, 3, 2, 2,
	2, 35, 261, 3, 2, 2, 2, 37, 266, 3, 2, 2, 2, 39, 271, 3, 2, 2, 2, 41, 277,
	3, 2, 2, 2, 43, 279, 3, 2, 2, 2, 45, 281, 3, 2, 2, 2, 47, 283, 3, 2, 2,
	2, 49, 285, 3, 2, 2, 2, 51, 287, 3, 2, 2, 2, 53, 289, 3, 2, 2, 2, 55, 291,
	3, 2, 2, 2, 57, 294, 3, 2, 2, 2, 59, 296, 3, 2, 2, 2, 61, 299, 3, 2, 2,
	2, 63, 302, 3, 2, 2, 2, 65, 305, 3, 2, 2, 2, 67, 307, 3, 2, 2, 2, 69, 310,
	3, 2, 2, 2, 71, 312, 3, 2, 2, 2, 73, 315, 3, 2, 2, 2, 75, 317, 3, 2, 2,
	2, 77, 319, 3, 2, 2, 2, 79, 321, 3, 2, 2, 2, 81, 323, 3, 2, 2, 2, 83, 325,
	3, 2, 2, 2, 85, 328, 3, 2, 2, 2, 87, 331, 3, 2, 2, 2, 89, 333, 3, 2, 2,
	2, 91, 335, 3, 2, 2, 2, 93, 337, 3, 2, 2, 2, 95, 339, 3, 2, 2, 2, 97, 341,
	3, 2, 2, 2, 99, 343, 3, 2, 2, 2, 101, 345, 3, 2, 2, 2, 103, 347, 3, 2,
	2, 2, 105, 357, 3, 2, 2, 2, 107, 359, 3, 2, 2, 2, 109, 361, 3, 2, 2, 2,
	111, 373, 3, 2, 2, 2, 113, 375, 3, 2, 2, 2, 115, 382, 3, 2, 2, 2, 117,
	389, 3, 2, 2, 2, 119, 391, 3, 2, 2, 2, 121, 400, 3, 2, 2, 2, 123, 402,
	3, 2, 2, 2, 125, 409, 3, 2, 2, 2, 127, 416, 3, 2, 2, 2, 129, 422, 3, 2,
	2, 2, 131, 425, 3, 2, 2, 2, 133, 427, 3, 2, 2, 2, 135, 429, 3, 2, 2, 2,
	137, 431, 3, 2, 2, 2, 139, 434, 3, 2, 2, 2, 141, 439, 3, 2, 2, 2, 143,
	465, 3, 2, 2, 2, 145, 468, 3, 2, 2, 2, 147, 474, 3, 2, 2, 2, 149, 480,
	3, 2, 2, 2, 151, 482, 3, 2, 2, 2, 153, 496, 3, 2, 2, 2, 155, 498, 3, 2,
	2, 2, 157, 506, 3, 2, 2, 2, 159, 513, 3, 2, 2, 2, 161, 524, 3, 2, 2, 2,
	163, 527, 3, 2, 2, 2, 165, 538, 3, 2, 2, 2, 167, 542, 3, 2, 2, 2, 169,
	556, 3, 2, 2, 2, 171, 172, 7, 63, 2, 2, 172, 173, 7, 63, 2, 2, 173, 4,
	3, 2, 2, 2, 174, 175, 7, 35, 2, 2, 175, 176, 7, 63, 2, 2, 176, 6, 3, 2,
	2, 2, 177, 178, 7, 100, 2, 2, 178, 179, 7, 116, 2, 2, 179, 180, 7, 103,
	2, 2, 180, 181, 7, 99, 2, 2, 181, 182, 7, 109, 2, 2, 182, 8, 3, 2, 2, 2,
	183, 184, 7, 101, 2, 2, 184, 185, 7, 113, 2, 2, 185, 186, 7, 112, 2, 2,
	186, 187, 7, 118, 2, 2, 187, 188, 7, 107, 2, 2, 188, 189, 7, 112, 2, 2,
	189, 190, 7, 119, 2, 2, 190, 191, 7, 103, 2, 2, 191, 10, 3, 2, 2, 2, 192,
	193, 7, 103, 2, 2, 193, 194, 7, 110, 2, 2, 194, 195, 7, 117, 2, 2, 195,
	196, 7, 103, 2, 2, 196, 12, 3, 2, 2, 2, 197, 198, 7, 107, 2, 2, 198, 199,
	7, 104, 2, 2, 199, 14, 3, 2, 2, 2, 200, 201, 7, 120, 2, 2, 201, 202, 7,
	99, 2, 2, 202, 203, 7, 116, 2, 2, 203, 16, 3, 2, 2, 2, 204, 205, 7, 107,
	2, 2, 205, 206, 7, 112, 2, 2, 206, 207, 7, 118, 2, 2, 207, 18, 3, 2, 2,
	2, 208, 209, 7, 100, 2, 2, 209, 210, 7, 113, 2, 2, 210, 211, 7, 113, 2,
	2, 211, 212, 7, 110, 2, 2, 212, 213, 7, 103, 2, 2, 213, 214, 7, 99, 2,
	2, 214, 215, 7, 112, 2, 2, 215, 20, 3, 2, 2, 2, 216, 217, 7, 117, 2, 2,
	217, 218, 7, 118, 2, 2, 218, 219, 7, 116, 2, 2, 219, 220, 7, 107, 2, 2,
	220, 221, 7, 112, 2, 2, 221, 222, 7, 105, 2, 2, 222, 22, 3, 2, 2, 2, 223,
	224, 7, 116, 2, 2, 224, 225, 7, 119, 2, 2, 225, 226, 7, 112, 2, 2, 226,
	227, 7, 103, 2, 2, 227, 24, 3, 2, 2, 2, 228, 229, 7, 99, 2, 2, 229, 230,
	7, 116, 2, 2, 230, 231, 7, 116, 2, 2, 231, 232, 7, 99, 2, 2, 232, 233,
	7, 123, 2, 2, 233, 26, 3, 2, 2, 2, 234, 235, 7, 116, 2, 2, 235, 236, 7,
	103, 2, 2, 236, 237, 7, 118, 2, 2, 237, 238, 7, 119, 2, 2, 238, 239, 7,
	116, 2, 2, 239, 240, 7, 112, 2, 2, 240, 28, 3, 2, 2, 2, 241, 242, 7, 120,
	2, 2, 242, 243, 7, 113, 2, 2, 243, 244, 7, 107, 2, 2, 244, 245, 7, 102,
	2, 2, 245, 30, 3, 2, 2, 2, 246, 247, 7, 121, 2, 2, 247, 248, 7, 106, 2,
	2, 248, 249, 7, 107, 2, 2, 249, 250, 7, 110, 2, 2, 250, 251, 7, 103, 2,
	2, 251, 32, 3, 2, 2, 2, 252, 253, 7, 104, 2, 2, 253, 254, 7, 119, 2, 2,
	254, 255, 7, 112, 2, 2, 255, 256, 7, 101, 2, 2, 256, 257, 7, 118, 2, 2,
	257, 258, 7, 107, 2, 2, 258, 259, 7, 113, 2, 2, 259, 260, 7, 112, 2, 2,
	260, 34, 3, 2, 2, 2, 261, 262, 7, 111, 2, 2, 262, 263, 7, 99, 2, 2, 263,
	264, 7, 107, 2, 2, 264, 265, 7, 112, 2, 2, 265, 36, 3, 2, 2, 2, 266, 267,
	7, 118, 2, 2, 267, 268, 7, 116, 2, 2, 268, 269, 7, 119, 2, 2, 269, 270,
	7, 103, 2, 2, 270, 38, 3, 2, 2, 2, 271, 272, 7, 104, 2, 2, 272, 273, 7,
	99, 2, 2, 273, 274, 7, 110, 2, 2, 274, 275, 7, 117, 2, 2, 275, 276, 7,
	103, 2, 2, 276, 40, 3, 2, 2, 2, 277, 278, 7, 42, 2, 2, 278, 42, 3, 2, 2,
	2, 279, 280, 7, 43, 2, 2, 280, 44, 3, 2, 2, 2, 281, 282, 7, 93, 2, 2, 282,
	46, 3, 2, 2, 2, 283, 284, 7, 95, 2, 2, 284, 48, 3, 2, 2, 2, 285, 286, 7,
	125, 2, 2, 286, 50, 3, 2, 2, 2, 287, 288, 7, 127, 2, 2, 288, 52, 3, 2,
	2, 2, 289, 290, 7, 62, 2, 2, 290, 54, 3, 2, 2, 2, 291, 292, 7, 62, 2, 2,
	292, 293, 7, 63, 2, 2, 293, 56, 3, 2, 2, 2, 294, 295, 7, 64, 2, 2, 295,
	58, 3, 2, 2, 2, 296, 297, 7, 64, 2, 2, 297, 298, 7, 63, 2, 2, 298, 60,
	3, 2, 2, 2, 299, 300, 7, 62, 2, 2, 300, 301, 7, 62, 2, 2, 301, 62, 3, 2,
	2, 2, 302, 303, 7, 64, 2, 2, 303, 304, 7, 64, 2, 2, 304, 64, 3, 2, 2, 2,
	305, 306, 7, 45, 2, 2, 306, 66, 3, 2, 2, 2, 307, 308, 7, 45, 2, 2, 308,
	309, 7, 45, 2, 2, 309, 68, 3, 2, 2, 2, 310, 311, 7, 47, 2, 2, 311, 70,
	3, 2, 2, 2, 312, 313, 7, 47, 2, 2, 313, 314, 7, 47, 2, 2, 314, 72, 3, 2,
	2, 2, 315, 316, 7, 44, 2, 2, 316, 74, 3, 2, 2, 2, 317, 318, 7, 49, 2, 2,
	318, 76, 3, 2, 2, 2, 319, 320, 7, 39, 2, 2, 320, 78, 3, 2, 2, 2, 321, 322,
	7, 40, 2, 2, 322, 80, 3, 2, 2, 2, 323, 324, 7, 126, 2, 2, 324, 82, 3, 2,
	2, 2, 325, 326, 7, 40, 2, 2, 326, 327, 7, 40, 2, 2, 327, 84, 3, 2, 2, 2,
	328, 329, 7, 126, 2, 2, 329, 330, 7, 126, 2, 2, 330, 86, 3, 2, 2, 2, 331,
	332, 7, 96, 2, 2, 332, 88, 3, 2, 2, 2, 333, 334, 7, 35, 2, 2, 334, 90,
	3, 2, 2, 2, 335, 336, 7, 128, 2, 2, 336, 92, 3, 2, 2, 2, 337, 338, 7, 65,
	2, 2, 338, 94, 3, 2, 2, 2, 339, 340, 7, 60, 2, 2, 340, 96, 3, 2, 2, 2,
	341, 342, 7, 61, 2, 2, 342, 98, 3, 2, 2, 2, 343, 344, 7, 46, 2, 2, 344,
	100, 3, 2, 2, 2, 345, 346, 7, 63, 2, 2, 346, 102, 3, 2, 2, 2, 347, 352,
	5, 105, 53, 2, 348, 351, 5, 105, 53, 2, 349, 351, 5, 109, 55, 2, 350, 348,
	3, 2, 2, 2, 350, 349, 3, 2, 2, 2, 351, 354, 3, 2, 2, 2, 352, 350, 3, 2,
	2, 2, 352, 353, 3, 2, 2, 2, 353, 104, 3, 2, 2, 2, 354, 352, 3, 2, 2, 2,
	355, 358, 5, 107, 54, 2, 356, 358, 5, 111, 56, 2, 357, 355, 3, 2, 2, 2,
	357, 356, 3, 2, 2, 2, 358, 106, 3, 2, 2, 2, 359, 360, 9, 2, 2, 2, 360,
	108, 3, 2, 2, 2, 361, 362, 9, 3, 2, 2, 362, 110, 3, 2, 2, 2, 363, 364,
	7, 94, 2, 2, 364, 365, 7, 119, 2, 2, 365, 366, 3, 2, 2, 2, 366, 374, 5,
	113, 57, 2, 367, 368, 7, 94, 2, 2, 368, 369, 7, 87, 2, 2, 369, 370, 3,
	2, 2, 2, 370, 371, 5, 113, 57, 2, 371, 372, 5, 113, 57, 2, 372, 374, 3,
	2, 2, 2, 373, 363, 3, 2, 2, 2, 373, 367, 3, 2, 2, 2, 374, 112, 3, 2, 2,
	2, 375, 376, 5, 135, 68, 2, 376, 377, 5, 135, 68, 2, 377, 378, 5, 135,
	68, 2, 378, 379, 5, 135, 68, 2, 379, 114, 3, 2, 2, 2, 380, 383, 5, 117,
	59, 2, 381, 383, 5, 143, 72, 2, 382, 380, 3, 2, 2, 2, 382, 381, 3, 2, 2,
	2, 383, 116, 3, 2, 2, 2, 384, 390, 5, 123, 62, 2, 385, 390, 5, 125, 63,
	2, 386, 390, 5, 127, 64, 2, 387, 390, 5, 119, 60, 2, 388, 390, 5, 121,
	61, 2, 389, 384, 3, 2, 2, 2, 389, 385, 3, 2, 2, 2, 389, 386, 3, 2, 2, 2,
	389, 387, 3, 2, 2, 2, 389, 388, 3, 2, 2, 2, 390, 118, 3, 2, 2, 2, 391,
	392, 7, 50, 2, 2, 392, 394, 9, 4, 2, 2, 393, 395, 9, 5, 2, 2, 394, 393,
	3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396, 394, 3, 2, 2, 2, 396, 397, 3, 2,
	2, 2, 397, 120, 3, 2, 2, 2, 398, 401, 5, 37, 19, 2, 399, 401, 5, 39, 20,
	2, 400, 398, 3, 2, 2, 2, 400, 399, 3, 2, 2, 2, 401, 122, 3, 2, 2, 2, 402,
	406, 5, 131, 66, 2, 403, 405, 5, 109, 55, 2, 404, 403, 3, 2, 2, 2, 405,
	408, 3, 2, 2, 2, 406, 404, 3, 2, 2, 2, 406, 407, 3, 2, 2, 2, 407, 124,
	3, 2, 2, 2, 408, 406, 3, 2, 2, 2, 409, 413, 7, 50, 2, 2, 410, 412, 5, 133,
	67, 2, 411, 410, 3, 2, 2, 2, 412, 415, 3, 2, 2, 2, 413, 411, 3, 2, 2, 2,
	413, 414, 3, 2, 2, 2, 414, 126, 3, 2, 2, 2, 415, 413, 3, 2, 2, 2, 416,
	418, 5, 129, 65, 2, 417, 419, 5, 135, 68, 2, 418, 417, 3, 2, 2, 2, 419,
	420, 3, 2, 2, 2, 420, 418, 3, 2, 2, 2, 420, 421, 3, 2, 2, 2, 421, 128,
	3, 2, 2, 2, 422, 423, 7, 50, 2, 2, 423, 424, 9, 6, 2, 2, 424, 130, 3, 2,
	2, 2, 425, 426, 9, 7, 2, 2, 426, 132, 3, 2, 2, 2, 427, 428, 9, 8, 2, 2,
	428, 134, 3, 2, 2, 2, 429, 430, 9, 9, 2, 2, 430, 136, 3, 2, 2, 2, 431,
	432, 9, 10, 2, 2, 432, 138, 3, 2, 2, 2, 433, 435, 5, 109, 55, 2, 434, 433,
	3, 2, 2, 2, 435, 436, 3, 2, 2, 2, 436, 434, 3, 2, 2, 2, 436, 437, 3, 2,
	2, 2, 437, 140, 3, 2, 2, 2, 438, 440, 5, 135, 68, 2, 439, 438, 3, 2, 2,
	2, 440, 441, 3, 2, 2, 2, 441, 439, 3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442,
	142, 3, 2, 2, 2, 443, 444, 7, 41, 2, 2, 444, 445, 5, 145, 73, 2, 445, 446,
	7, 41, 2, 2, 446, 466, 3, 2, 2, 2, 447, 448, 7, 78, 2, 2, 448, 449, 7,
	41, 2, 2, 449, 450, 3, 2, 2, 2, 450, 451, 5, 145, 73, 2, 451, 452, 7, 41,
	2, 2, 452, 466, 3, 2, 2, 2, 453, 454, 7, 119, 2, 2, 454, 455, 7, 41, 2,
	2, 455, 456, 3, 2, 2, 2, 456, 457, 5, 145, 73, 2, 457, 458, 7, 41, 2, 2,
	458, 466, 3, 2, 2, 2, 459, 460, 7, 87, 2, 2, 460, 461, 7, 41, 2, 2, 461,
	462, 3, 2, 2, 2, 462, 463, 5, 145, 73, 2, 463, 464, 7, 41, 2, 2, 464, 466,
	3, 2, 2, 2, 465, 443, 3, 2, 2, 2, 465, 447, 3, 2, 2, 2, 465, 453, 3, 2,
	2, 2, 465, 459, 3, 2, 2, 2, 466, 144, 3, 2, 2, 2, 467, 469, 5, 147, 74,
	2, 468, 467, 3, 2, 2, 2, 469, 470, 3, 2, 2, 2, 470, 468, 3, 2, 2, 2, 470,
	471, 3, 2, 2, 2, 471, 146, 3, 2, 2, 2, 472, 475, 10, 11, 2, 2, 473, 475,
	5, 149, 75, 2, 474, 472, 3, 2, 2, 2, 474, 473, 3, 2, 2, 2, 475, 148, 3,
	2, 2, 2, 476, 481, 5, 151, 76, 2, 477, 481, 5, 153, 77, 2, 478, 481, 5,
	155, 78, 2, 479, 481, 5, 111, 56, 2, 480, 476, 3, 2, 2, 2, 480, 477, 3,
	2, 2, 2, 480, 478, 3, 2, 2, 2, 480, 479, 3, 2, 2, 2, 481, 150, 3, 2, 2,
	2, 482, 483, 7, 94, 2, 2, 483, 484, 9, 12, 2, 2, 484, 152, 3, 2, 2, 2,
	485, 486, 7, 94, 2, 2, 486, 497, 5, 133, 67, 2, 487, 488, 7, 94, 2, 2,
	488, 489, 5, 133, 67, 2, 489, 490, 5, 133, 67, 2, 490, 497, 3, 2, 2, 2,
	491, 492, 7, 94, 2, 2, 492, 493, 5, 133, 67, 2, 493, 494, 5, 133, 67, 2,
	494, 495, 5, 133, 67, 2, 495, 497, 3, 2, 2, 2, 496, 485, 3, 2, 2, 2, 496,
	487, 3, 2, 2, 2, 496, 491, 3, 2, 2, 2, 497, 154, 3, 2, 2, 2, 498, 499,
	7, 94, 2, 2, 499, 500, 7, 122, 2, 2, 500, 502, 3, 2, 2, 2, 501, 503, 5,
	135, 68, 2, 502, 501, 3, 2, 2, 2, 503, 504, 3, 2, 2, 2, 504, 502, 3, 2,
	2, 2, 504, 505, 3, 2, 2, 2, 505, 156, 3, 2, 2, 2, 506, 508, 7, 36, 2, 2,
	507, 509, 5, 159, 80, 2, 508, 507, 3, 2, 2, 2, 508, 509, 3, 2, 2, 2, 509,
	510, 3, 2, 2, 2, 510, 511, 7, 36, 2, 2, 511, 158, 3, 2, 2, 2, 512, 514,
	5, 161, 81, 2, 513, 512, 3, 2, 2, 2, 514, 515, 3, 2, 2, 2, 515, 513, 3,
	2, 2, 2, 515, 516, 3, 2, 2, 2, 516, 160, 3, 2, 2, 2, 517, 525, 10, 13,
	2, 2, 518, 525, 5, 149, 75, 2, 519, 520, 7, 94, 2, 2, 520, 525, 7, 12,
	2, 2, 521, 522, 7, 94, 2, 2, 522, 523, 7, 15, 2, 2, 523, 525, 7, 12, 2,
	2, 524, 517, 3, 2, 2, 2, 524, 518, 3, 2, 2, 2, 524, 519, 3, 2, 2, 2, 524,
	521, 3, 2, 2, 2, 525, 162, 3, 2, 2, 2, 526, 528, 9, 14, 2, 2, 527, 526,
	3, 2, 2, 2, 528, 529, 3, 2, 2, 2, 529, 527, 3, 2, 2, 2, 529, 530, 3, 2,
	2, 2, 530, 531, 3, 2, 2, 2, 531, 532, 8, 82, 2, 2, 532, 164, 3, 2, 2, 2,
	533, 535, 7, 15, 2, 2, 534, 536, 7, 12, 2, 2, 535, 534, 3, 2, 2, 2, 535,
	536, 3, 2, 2, 2, 536, 539, 3, 2, 2, 2, 537, 539, 7, 12, 2, 2, 538, 533,
	3, 2, 2, 2, 538, 537, 3, 2, 2, 2, 539, 540, 3, 2, 2, 2, 540, 541, 8, 83,
	2, 2, 541, 166, 3, 2, 2, 2, 542, 543, 7, 49, 2, 2, 543, 544, 7, 44, 2,
	2, 544, 548, 3, 2, 2, 2, 545, 547, 11, 2, 2, 2, 546, 545, 3, 2, 2, 2, 547,
	550, 3, 2, 2, 2, 548, 549, 3, 2, 2, 2, 548, 546, 3, 2, 2, 2, 549, 551,
	3, 2, 2, 2, 550, 548, 3, 2, 2, 2, 551, 552, 7, 44, 2, 2, 552, 553, 7, 49,
	2, 2, 553, 554, 3, 2, 2, 2, 554, 555, 8, 84, 2, 2, 555, 168, 3, 2, 2, 2,
	556, 557, 7, 49, 2, 2, 557, 558, 7, 49, 2, 2, 558, 562, 3, 2, 2, 2, 559,
	561, 10, 15, 2, 2, 560, 559, 3, 2, 2, 2, 561, 564, 3, 2, 2, 2, 562, 560,
	3, 2, 2, 2, 562, 563, 3, 2, 2, 2, 563, 565, 3, 2, 2, 2, 564, 562, 3, 2,
	2, 2, 565, 566, 8, 85, 2, 2, 566, 170, 3, 2, 2, 2, 30, 2, 350, 352, 357,
	373, 382, 389, 396, 400, 406, 413, 420, 436, 441, 465, 470, 474, 480, 496,
	504, 508, 515, 524, 529, 535, 538, 548, 562, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'=='", "'!='", "'break'", "'continue'", "'else'", "'if'", "'var'",
	"'int'", "'boolean'", "'string'", "'rune'", "'array'", "'return'", "'void'",
	"'while'", "'function'", "'main'", "'true'", "'false'", "'('", "')'", "'['",
	"']'", "'{'", "'}'", "'<'", "'<='", "'>'", "'>='", "'<<'", "'>>'", "'+'",
	"'++'", "'-'", "'--'", "'*'", "'/'", "'%'", "'&'", "'|'", "'&&'", "'||'",
	"'^'", "'!'", "'~'", "'?'", "':'", "';'", "','", "'='",
}

var lexerSymbolicNames = []string{
	"", "", "", "Break", "Continue", "Else", "If", "Var", "Int", "Bool", "StringType",
	"Rune", "Array", "Return", "Void", "While", "Func", "Main", "True", "False",
	"Paren", "Thesis", "LeftBracket", "RightBracket", "LeftBrace", "RightBrace",
	"Less", "LessEqual", "Greater", "GreaterEqual", "LeftShift", "RightShift",
	"Plus", "PlusPlus", "Minus", "MinusMinus", "Star", "Div", "Mod", "And",
	"Or", "AndAnd", "OrOr", "Caret", "Not", "Tilde", "Question", "Colon", "Semi",
	"Comma", "Assign", "Identifier", "Constant", "DigitSequence", "StringLiteral",
	"Whitespace", "Newline", "BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "Break", "Continue", "Else", "If", "Var", "Int", "Bool",
	"StringType", "Rune", "Array", "Return", "Void", "While", "Func", "Main",
	"True", "False", "Paren", "Thesis", "LeftBracket", "RightBracket", "LeftBrace",
	"RightBrace", "Less", "LessEqual", "Greater", "GreaterEqual", "LeftShift",
	"RightShift", "Plus", "PlusPlus", "Minus", "MinusMinus", "Star", "Div",
	"Mod", "And", "Or", "AndAnd", "OrOr", "Caret", "Not", "Tilde", "Question",
	"Colon", "Semi", "Comma", "Assign", "Identifier", "IdentifierNondigit",
	"Nondigit", "Digit", "UniversalCharacterName", "HexQuad", "Constant", "IntegerConstant",
	"BinaryConstant", "BooleanConstant", "DecimalConstant", "OctalConstant",
	"HexadecimalConstant", "HexadecimalPrefix", "NonzeroDigit", "OctalDigit",
	"HexadecimalDigit", "Sign", "DigitSequence", "HexadecimalDigitSequence",
	"CharacterConstant", "CCharSequence", "CChar", "EscapeSequence", "SimpleEscapeSequence",
	"OctalEscapeSequence", "HexadecimalEscapeSequence", "StringLiteral", "SCharSequence",
	"SChar", "Whitespace", "Newline", "BlockComment", "LineComment",
}

type cminusLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewcminusLexer(input antlr.CharStream) *cminusLexer {

	l := new(cminusLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "cminus.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// cminusLexer tokens.
const (
	cminusLexerT__0          = 1
	cminusLexerT__1          = 2
	cminusLexerBreak         = 3
	cminusLexerContinue      = 4
	cminusLexerElse          = 5
	cminusLexerIf            = 6
	cminusLexerVar           = 7
	cminusLexerInt           = 8
	cminusLexerBool          = 9
	cminusLexerStringType    = 10
	cminusLexerRune          = 11
	cminusLexerArray         = 12
	cminusLexerReturn        = 13
	cminusLexerVoid          = 14
	cminusLexerWhile         = 15
	cminusLexerFunc          = 16
	cminusLexerMain          = 17
	cminusLexerTrue          = 18
	cminusLexerFalse         = 19
	cminusLexerParen         = 20
	cminusLexerThesis        = 21
	cminusLexerLeftBracket   = 22
	cminusLexerRightBracket  = 23
	cminusLexerLeftBrace     = 24
	cminusLexerRightBrace    = 25
	cminusLexerLess          = 26
	cminusLexerLessEqual     = 27
	cminusLexerGreater       = 28
	cminusLexerGreaterEqual  = 29
	cminusLexerLeftShift     = 30
	cminusLexerRightShift    = 31
	cminusLexerPlus          = 32
	cminusLexerPlusPlus      = 33
	cminusLexerMinus         = 34
	cminusLexerMinusMinus    = 35
	cminusLexerStar          = 36
	cminusLexerDiv           = 37
	cminusLexerMod           = 38
	cminusLexerAnd           = 39
	cminusLexerOr            = 40
	cminusLexerAndAnd        = 41
	cminusLexerOrOr          = 42
	cminusLexerCaret         = 43
	cminusLexerNot           = 44
	cminusLexerTilde         = 45
	cminusLexerQuestion      = 46
	cminusLexerColon         = 47
	cminusLexerSemi          = 48
	cminusLexerComma         = 49
	cminusLexerAssign        = 50
	cminusLexerIdentifier    = 51
	cminusLexerConstant      = 52
	cminusLexerDigitSequence = 53
	cminusLexerStringLiteral = 54
	cminusLexerWhitespace    = 55
	cminusLexerNewline       = 56
	cminusLexerBlockComment  = 57
	cminusLexerLineComment   = 58
)
