import {describe, it, expect} from "bun:test";
import {Tensor, InvalidIndexError} from "./tensor";

describe("Tensor", () => {
  describe("neg", () => {
    const testcases: Record<string, [Tensor, Tensor]> = {
      "scalar": [
        Tensor.new(5),
        Tensor.new(-5),
      ],
      "1D tensor": [
        Tensor.new([1, -2, 3, -4]),
        Tensor.new([-1, 2, -3, 4]),
      ],
      "2D tensor": [
        Tensor.new([
          [1, -2],
          [-3, 4],
        ]),
        Tensor.new([
          [-1, 2],
          [3, -4],
        ]),
      ],
      "zero": [
        Tensor.new([0, -0]),
        Tensor.new([-0, 0]),
      ],
      "Infinity": [
        Tensor.new([Infinity, -Infinity]),
        Tensor.new([-Infinity, Infinity]),
      ],
      "NaN": [
        Tensor.new([1, NaN, 3]),
        Tensor.new([-1, NaN, -3]),
      ],
    }

    for (const [name, [t, expected]] of Object.entries(testcases)) {
      it(name, () => {
        const actual = t.neg();
        expect(actual.data).toEqual(expected.data);
        expect(actual.shape).toEqual(expected.shape);
      });
    }
  });

  describe("mul", () => {
    const testcases: [string, Tensor, (a: Tensor) => Tensor, Tensor][] = [
      ["multiplies two scalars", Tensor.new(3), _ => Tensor.new(5), Tensor.new(15)],
      ["multiplies two 1D tensors of same shape", Tensor.new([1, 2, 3]), _ => Tensor.new([4, 5, 6]), Tensor.new([4, 10, 18])],
      ["multiplies two 2D tensors of same shape", Tensor.new([
        [1, 2],
        [3, 4],
      ]), _ => Tensor.new([
        [10, 20],
        [30, 40],
      ]),
        Tensor.new([
          [10, 40],
          [90, 160],
        ]),
      ],
      ["broadcasts scalar to 1D tensor", Tensor.new([1, 2, 3]), a => Tensor.new(10).broadcast(a.shape), Tensor.new([10, 20, 30])],
      ["broadcasts 1D tensor across rows of 2D tensor", Tensor.new([
        [1, 2, 3],
        [4, 5, 6],
      ]), _ => Tensor.new([10, 20, 30]).broadcast([2, 0]),
        Tensor.new([
          [10, 40, 90],
          [40, 100, 180],
        ]),
      ],
      ["broadcasts column vector to 2D tensor", Tensor.new([
        [1, 2, 3],
        [4, 5, 6],
      ]), _ => Tensor.new([[10], [100]]).squeeze([1]).broadcast([0, 3]),
        Tensor.new([
          [10, 20, 30],
          [400, 500, 600],
        ]),
      ],
    // Edge cases with Infinity and NaN
      ["handles Infinity in multiplication", Tensor.new([Infinity, -Infinity, 2]), _ => Tensor.new([2, 2, Infinity]), Tensor.new([Infinity, -Infinity, Infinity])],
      ["returns NaN when multiplying Infinity by zero", Tensor.new([Infinity, -Infinity, 0]), _ => Tensor.new([0, 0, Infinity]), Tensor.new([NaN, NaN, NaN])],
      ["handles Infinity signs correctly", Tensor.new([Infinity, Infinity, -Infinity, -Infinity]), _ => Tensor.new([1, -1, 1, -1]), Tensor.new([Infinity, -Infinity, -Infinity, Infinity])],
      ["propagates NaN through multiplication", Tensor.new([1, NaN, 3]), _ => Tensor.new([2, 2, NaN]), Tensor.new([2, NaN, NaN])],
    ];

    for (const [name, a, newb, expected] of testcases) {
      it(name, () => {
        const b = newb(a);
        const actual = a.mul(b);
        expect(actual.shape).toEqual(expected.shape);
        expect(actual.data).toEqual(expected.data);
      });
    }

    it("throws error for incompatible shapes", () => {
      const a = Tensor.new([1, 2, 3]);
      const b = Tensor.new([1, 2]);
      expect(() => a.mul(b)).toThrow("Incompatible shapes: [3] != [2]");
    });
  });

  describe("add", () => {
    const testcases: [string, Tensor, (a: Tensor) => Tensor, Tensor][] = [
      ["adds two scalars", Tensor.new(3), _ => Tensor.new(5), Tensor.new(8)],
      ["adds two 1D tensors of same shape", Tensor.new([1, 2, 3]), _ => Tensor.new([4, 5, 6]), Tensor.new([5, 7, 9])],
      ["adds two 2D tensors of same shape", Tensor.new([
          [1, 2],
          [3, 4],
        ]), _ => Tensor.new([
          [10, 20],
          [30, 40],
        ]), Tensor.new([
          [11, 22],
          [33, 44],
        ])],
      ["broadcasts scalar to 1D tensor", Tensor.new([1, 2, 3]), a => Tensor.new(10).broadcast(a.shape), Tensor.new([11, 12, 13])],
      ["broadcasts scalar to 2D tensor", Tensor.new([
          [1, 2],
          [3, 4],
        ]), a => Tensor.new(10).broadcast(a.shape), Tensor.new([
          [11, 12],
          [13, 14],
        ])],
      ["broadcasts 1D tensor across rows of 2D tensor", Tensor.new([
          [1, 2, 3],
          [4, 5, 6],
        ]), _ => Tensor.new([10, 20, 30]).broadcast([2, 0]), Tensor.new([
          [11, 22, 33],
          [14, 25, 36],
        ])],
      ["broadcasts column vector to 2D tensor", Tensor.new([
          [1, 2, 3],
          [4, 5, 6],
        ]), _ => Tensor.new([[10], [20]]).squeeze([1]).broadcast([0, 3]), Tensor.new([
          [11, 12, 13],
          [24, 25, 26],
        ])],
      ["adds 2D tensors where rows > dimensions [3, 2]", Tensor.new([
          [1, 2],
          [3, 4],
          [5, 6],
        ]), _ => Tensor.new([
          [10, 20],
          [30, 40],
          [50, 60],
        ]), Tensor.new([
          [11, 22],
          [33, 44],
          [55, 66],
        ])],
      ["adds 2D tensors with shape [4, 3]", Tensor.new([
          [1, 2, 3],
          [4, 5, 6],
          [7, 8, 9],
          [10, 11, 12],
        ]), _ => Tensor.new([100, 200, 300]).broadcast([4, 0]), Tensor.new([
          [101, 202, 303],
          [104, 205, 306],
          [107, 208, 309],
          [110, 211, 312],
        ])],
      ["adds 3D tensors where first dim > num dimensions [4, 2, 3]", Tensor.new([
          [
            [1, 2, 3],
            [4, 5, 6],
          ],
          [
            [7, 8, 9],
            [10, 11, 12],
          ],
          [
            [13, 14, 15],
            [16, 17, 18],
          ],
          [
            [19, 20, 21],
            [22, 23, 24],
          ],
        ]), a => Tensor.new(1).broadcast(a.shape), // broadcast scalar
        Tensor.new([
          [
            [2, 3, 4],
            [5, 6, 7],
          ],
          [
            [8, 9, 10],
            [11, 12, 13],
          ],
          [
            [14, 15, 16],
            [17, 18, 19],
          ],
          [
            [20, 21, 22],
            [23, 24, 25],
          ],
        ])],
      // Edge cases with Infinity and NaN
      ["handles Infinity in addition", Tensor.new([1, Infinity, -Infinity]), _ => Tensor.new([1, 1, 1]), Tensor.new([2, Infinity, -Infinity])],
      ["returns NaN when adding Infinity and -Infinity", Tensor.new([Infinity]), _ => Tensor.new([-Infinity]), Tensor.new([NaN])],
      ["propagates NaN through addition", Tensor.new([1, 2, NaN]), _ => Tensor.new([1, 1, 1]), Tensor.new([2, 3, NaN])],
    ];

    for (const [name, a, newb, expected] of testcases) {
      it(name, () => {
        const b = newb(a);
        const actual = a.add(b);
        expect(actual.shape).toEqual(expected.shape);
        expect(actual.data).toEqual(expected.data);
      });
    }

    it("throws error for incompatible shapes", () => {
      const a = Tensor.new([1, 2, 3]);
      const b = Tensor.new([1, 2]);
      expect(() => a.add(b)).toThrow("Incompatible shapes: [3] != [2]");
    });
  });

  describe("broadcast", () => {
    const testcases: [string, Tensor, number[], Tensor][] = [
      ["[1, 2].broadcast([0, 3])", Tensor.new([1, 2]), [0, 3], Tensor.new([[1, 1, 1], [2, 2, 2]])],
    ]

    for (const [name, t, shape, expected] of testcases) {
      it(name, () => {
        const result = t.broadcast(shape);
        expect(result.shape).toEqual(expected.shape);
        expect(result.data).toEqual(expected.data);
      });
    }
  });

  describe("div", () => {
    const testcases: [string, Tensor, (a: Tensor) => Tensor, Tensor][] = [
      ["divides two scalars", Tensor.new(15), _ => Tensor.new(3), Tensor.new(5)],
      ["divides two 1D tensors of same shape", Tensor.new([10, 20, 30]), _ => Tensor.new([2, 4, 5]), Tensor.new([5, 5, 6])],
      ["divides two 2D tensors of same shape", Tensor.new([
          [10, 20],
          [30, 40],
        ]), _ => Tensor.new([
          [2, 4],
          [5, 8],
        ]),
        Tensor.new([
          [5, 5],
          [6, 5],
        ])],
      ["broadcasts scalar to 1D tensor", Tensor.new([10, 20, 30]), a => Tensor.new(10).broadcast(a.shape), Tensor.new([1, 2, 3])],
      ["broadcasts 1D tensor across rows of 2D tensor", Tensor.new([
          [10, 20, 30],
          [40, 50, 60],
        ]), _ => Tensor.new([10, 10, 10]).broadcast([2, 0]), Tensor.new([
          [1, 2, 3],
          [4, 5, 6],
        ])],
      ["broadcasts column vector to 2D tensor", Tensor.new([
          [10, 20, 30],
          [40, 50, 60],
        ]), _ => Tensor.new([[10], [10]]).squeeze([1]).broadcast([0, 3]),
        Tensor.new([
          [1, 2, 3],
          [4, 5, 6],
        ])],
      ["handles floating point division", Tensor.new([1, 2, 3]), a => Tensor.new(2).broadcast(a.shape), Tensor.new([0.5, 1, 1.5])],
      ["returns Infinity when dividing positive by zero", Tensor.new([1, 2, 3]), a => Tensor.new(0).broadcast(a.shape), Tensor.new([Infinity, Infinity, Infinity])],
      ["returns -Infinity when dividing negative by zero", Tensor.new([-1, -2, -3]), a => Tensor.new(0).broadcast(a.shape), Tensor.new([-Infinity, -Infinity, -Infinity])],
      ["returns NaN when dividing zero by zero", Tensor.new([0, 0, 0]), a => Tensor.new(0).broadcast(a.shape), Tensor.new([NaN, NaN, NaN])],
      ["handles mixed division by zero cases", Tensor.new([1, -1, 0]), a => Tensor.new(0).broadcast(a.shape), Tensor.new([Infinity, -Infinity, NaN])],
    ]

    for (const [name, a, newb, expected] of testcases) {
      it(name, () => {
        const b = newb(a);
        const actual = a.div(b);
        expect(actual.shape).toEqual(expected.shape);
        expect(actual.data).toEqual(expected.data);
      });
    }

    it("throws error for incompatible shapes", () => {
      const a = Tensor.new([1, 2, 3]);
      const b = Tensor.new([1, 2]);
      expect(() => a.div(b)).toThrow("Incompatible shapes: [3] != [2]");
    });
  });

  describe("exp", () => {
    it("computes exp of a scalar", () => {
      const a = Tensor.new(0);
      const result = a.exp();
      expect(result.data).toEqual(1); // e^0 = 1
      expect(result.shape).toEqual([]);
    });

    it("computes exp of a 1D tensor", () => {
      const a = Tensor.new([0, 1, 2]);
      const result = a.exp();
      const data = result.data as number[];
      expect(data[0]).toBeCloseTo(1); // e^0
      expect(data[1]).toBeCloseTo(Math.E); // e^1
      expect(data[2]).toBeCloseTo(Math.E ** 2); // e^2
    });

    it("computes exp of a 2D tensor", () => {
      const a = Tensor.new([
        [0, 1],
        [2, 3],
      ]);
      const result = a.exp();
      expect((result.data as number[][])[0][0]).toBeCloseTo(1);
      expect((result.data as number[][])[0][1]).toBeCloseTo(Math.E);
      expect((result.data as number[][])[1][0]).toBeCloseTo(Math.E ** 2);
      expect((result.data as number[][])[1][1]).toBeCloseTo(Math.E ** 3);
    });

    it("handles negative values", () => {
      const a = Tensor.new([-1, -2]);
      const result = a.exp();
      const data = result.data as number[];
      expect(data[0]).toBeCloseTo(1 / Math.E);
      expect(data[1]).toBeCloseTo(1 / Math.E ** 2);
    });

    it("returns Infinity for large positive values", () => {
      const a = Tensor.new([1000]);
      const result = a.exp();
      expect(result.data).toEqual([Infinity]);
    });

    it("returns 0 for large negative values", () => {
      const a = Tensor.new([-1000]);
      const result = a.exp();
      expect(result.data).toEqual([0]);
    });

    it("handles Infinity input", () => {
      const a = Tensor.new([Infinity, -Infinity]);
      const result = a.exp();
      expect(result.data).toEqual([Infinity, 0]);
    });

    it("propagates NaN", () => {
      const a = Tensor.new([0, NaN]);
      const result = a.exp();
      const data = result.data as number[];
      expect(data).toEqual([1, NaN]);
    });
  });

  describe("get", () => {
    const testcases: Record<string, [Tensor, [number[], number][]]> = {
      "1D tensor": [Tensor.new([10, 20, 30, 40]), [
        [[0], 10],
        [[1], 20],
        [[3], 40],
      ]],
      "2D tensor": [Tensor.new([
        [1, 2, 3],
        [4, 5, 6],
      ]), [
        [[0, 0], 1],
        [[0, 2], 3],
        [[1, 0], 4],
        [[1, 2], 6],
      ]],
      "3D tensor": [Tensor.new([
        [
          [1, 2],
          [3, 4],
        ],
        [
          [5, 6],
          [7, 8],
        ],
      ]), [
        [[0, 0, 0], 1],
        [[0, 1, 1], 4],
        [[1, 0, 0], 5],
        [[1, 1, 1], 8],
      ]],
    };

    describe("valid indices", () => {
      for (const [name, [t, tests]] of Object.entries(testcases)) {
        it(name, () => {
          for (const [indices, expected] of tests) {
            const actual = t.get(indices);
            expect(actual).toBe(expected);
          }
        });
      }
    });

    describe("InvalidIndexError", () => {
      const testcases: Record<string, [Tensor, number[][]]> = {
        "too few indices provided": [Tensor.new([
          [1, 2, 3],
          [4, 5, 6],
        ]), [
          [0],
        ]],
        "too many indices provided": [Tensor.new([1, 2, 3]), [[0, 0]]],
        "no indices provided for non-scalar": [Tensor.new([1, 2, 3]), [[]]],
        "index is negative": [Tensor.new([1, 2, 3]), [[-1]]],
        "index equals dimension size": [Tensor.new([1, 2, 3]), [[3]]],
        "index exceeds dimension size": [Tensor.new([1, 2, 3]), [[5]]],
        "any index in multi-dim is out of bounds": [Tensor.new([
          [1, 2, 3],
          [4, 5, 6],
        ]), [
          [2, 0],
          [0, 3],
        ]],
      };

      for (const [name, [t, tests]] of Object.entries(testcases)) {
        it(name, () => {
          for (const indices of tests)
            expect(() => t.get(indices)).toThrow(InvalidIndexError);
        });
      }
    });
  });

  describe("log", () => {
    it("computes log of a scalar", () => {
      const a = Tensor.new(1);
      const result = a.log();
      expect(result.data).toEqual(0); // ln(1) = 0
      expect(result.shape).toEqual([]);
    });

    it("computes log of e", () => {
      const a = Tensor.new(Math.E);
      const result = a.log();
      expect(result.data).toBeCloseTo(1); // ln(e) = 1
    });

    it("computes log of a 1D tensor", () => {
      const a = Tensor.new([1, Math.E, Math.E ** 2]);
      const result = a.log();
      const data = result.data as number[];
      expect(data[0]).toBeCloseTo(0);
      expect(data[1]).toBeCloseTo(1);
      expect(data[2]).toBeCloseTo(2);
    });

    it("computes log of a 2D tensor", () => {
      const a = Tensor.new([
        [1, Math.E],
        [Math.E ** 2, Math.E ** 3],
      ]);
      const result = a.log();
      expect((result.data as number[][])[0][0]).toBeCloseTo(0);
      expect((result.data as number[][])[0][1]).toBeCloseTo(1);
      expect((result.data as number[][])[1][0]).toBeCloseTo(2);
      expect((result.data as number[][])[1][1]).toBeCloseTo(3);
    });

    it("returns -Infinity for zero", () => {
      const a = Tensor.new([0]);
      const result = a.log();
      expect(result.data).toEqual([-Infinity]);
    });

    it("returns NaN for negative values", () => {
      const a = Tensor.new([-1, -2, -100]);
      const result = a.log();
      expect(result.data).toEqual([NaN, NaN, NaN]);
    });

    it("returns Infinity for Infinity input", () => {
      const a = Tensor.new([Infinity]);
      const result = a.log();
      expect(result.data).toEqual([Infinity]);
    });

    it("returns NaN for -Infinity input", () => {
      const a = Tensor.new([-Infinity]);
      const result = a.log();
      expect(result.data).toEqual([NaN]);
    });

    it("propagates NaN", () => {
      const a = Tensor.new([1, NaN]);
      const result = a.log();
      expect(result.data).toEqual([0, NaN]);
    });
  });

  describe("matmul", () => {
    describe("2D matrix multiplication", () => {
      it("multiplies two 2D matrices [2,3] @ [3,2]", () => {
        const a = Tensor.new([
          [1, 2, 3],
          [4, 5, 6],
        ]);
        const b = Tensor.new([
          [7, 8],
          [9, 10],
          [11, 12],
        ]);
        const result = a.product(b, [1], [0]);
        // [1*7+2*9+3*11, 1*8+2*10+3*12] = [58, 64]
        // [4*7+5*9+6*11, 4*8+5*10+6*12] = [139, 154]
        expect(result.data).toEqual([
          [58, 64],
          [139, 154],
        ]);
        expect(result.shape).toEqual([2, 2]);
      });

      it("multiplies [3,2] @ [2,4]", () => {
        const a = Tensor.new([
          [1, 2],
          [3, 4],
          [5, 6],
        ]);
        const b = Tensor.new([
          [1, 2, 3, 4],
          [5, 6, 7, 8],
        ]);
        const result = a.product(b, [1], [0]);
        // Row 0: [1*1+2*5, 1*2+2*6, 1*3+2*7, 1*4+2*8] = [11, 14, 17, 20]
        // Row 1: [3*1+4*5, 3*2+4*6, 3*3+4*7, 3*4+4*8] = [23, 30, 37, 44]
        // Row 2: [5*1+6*5, 5*2+6*6, 5*3+6*7, 5*4+6*8] = [35, 46, 57, 68]
        expect(result.data).toEqual([
          [11, 14, 17, 20],
          [23, 30, 37, 44],
          [35, 46, 57, 68],
        ]);
        expect(result.shape).toEqual([3, 4]);
      });

      it("multiplies square matrices [2,2] @ [2,2]", () => {
        const a = Tensor.new([
          [1, 2],
          [3, 4],
        ]);
        const b = Tensor.new([
          [5, 6],
          [7, 8],
        ]);
        const result = a.product(b, [1], [0]);
        // [1*5+2*7, 1*6+2*8] = [19, 22]
        // [3*5+4*7, 3*6+4*8] = [43, 50]
        expect(result.data).toEqual([
          [19, 22],
          [43, 50],
        ]);
        expect(result.shape).toEqual([2, 2]);
      });
    });

    describe("1D vector operations", () => {
      it("computes dot product of two 1D vectors [3] @ [3]", () => {
        const a = Tensor.new([1, 2, 3]);
        const b = Tensor.new([4, 5, 6]);
        const result = a.product(b, [0], [0]);
        // 1*4 + 2*5 + 3*6 = 32
        expect(result.data).toEqual(32);
        expect(result.shape).toEqual([]);
      });

      it("computes matrix-vector product [2,3] @ [3]", () => {
        const a = Tensor.new([
          [1, 2, 3],
          [4, 5, 6],
        ]);
        const b = Tensor.new([1, 2, 3]);
        const result = a.product(b, [1], [0]);
        // [1*1+2*2+3*3, 4*1+5*2+6*3] = [14, 32]
        expect(result.data).toEqual([14, 32]);
      });

      it("computes vector-matrix product [3] @ [3,2]", () => {
        const a = Tensor.new([1, 2, 3]);
        const b = Tensor.new([
          [1, 2],
          [3, 4],
          [5, 6],
        ]);
        const result = a.product(b, [0], [0]);
        // [1*1+2*3+3*5, 1*2+2*4+3*6] = [22, 28]
        expect(result.data).toEqual([22, 28]);
      });
    });

    describe("batched matrix multiplication", () => {
      it("multiplies batched 3D tensors [2,3,4] @ [2,4,5]", () => {
        // Batch of 2 matrices
        const a = Tensor.new([
          [
            [1, 0, 0, 0],
            [0, 1, 0, 0],
            [0, 0, 1, 0],
          ],
          [
            [2, 0, 0, 0],
            [0, 2, 0, 0],
            [0, 0, 2, 0],
          ],
        ]);
        const b = Tensor.new([
          [
            [ 1,  2,  3,  4,  5],
            [ 6,  7,  8,  9, 10],
            [11, 12, 13, 14, 15],
            [16, 17, 18, 19, 20],
          ],
          [
            [1, 1, 1, 1, 1],
            [2, 2, 2, 2, 2],
            [3, 3, 3, 3, 3],
            [4, 4, 4, 4, 4],
          ],
        ]);
        const result = a.product(b, [2], [1]).diag([0, 2]).transpose([2, 0, 1]);
        expect(result.shape).toEqual([2, 3, 5]);
        expect(result.data).toEqual([
          // First batch: identity-ish extracts first 3 rows
          [
            [ 1,  2,  3,  4,  5],
            [ 6,  7,  8,  9, 10],
            [11, 12, 13, 14, 15],
          ],
          // Second batch: 2x identity extracts and doubles first 3 rows
          [
            [2, 2, 2, 2, 2],
            [4, 4, 4, 4, 4],
            [6, 6, 6, 6, 6],
          ],
        ]);
      });

      it("broadcasts 2D matrix against batched 3D tensor [2,3,4] @ [4,5]", () => {
        const a = Tensor.new([
          [
            [1, 0, 0, 0],
            [0, 1, 0, 0],
            [0, 0, 1, 0],
          ],
          [
            [1, 0, 0, 0],
            [0, 1, 0, 0],
            [0, 0, 1, 0],
          ],
        ]);
        const b = Tensor.new([
          [ 1,  2,  3,  4,  5],
          [ 6,  7,  8,  9, 10],
          [11, 12, 13, 14, 15],
          [16, 17, 18, 19, 20],
        ]);
        const result = a.product(b, [2], [0]);
        expect(result.shape).toEqual([2, 3, 5]);
        // Both batches get same result (extracting first 3 rows of b)
        const expected = [
          [ 1,  2,  3,  4,  5],
          [ 6,  7,  8,  9, 10],
          [11, 12, 13, 14, 15],
        ];
        expect((result.data as number[][][])[0]).toEqual(expected);
        expect((result.data as number[][][])[1]).toEqual(expected);
      });

      it("broadcasts 2D matrix against batched 3D tensor [3,4] @ [2,4,5]", () => {
        const a = Tensor.new([
          [1, 0, 0, 0],
          [0, 1, 0, 0],
          [0, 0, 1, 0],
        ]);
        const b = Tensor.new([
          [
            [1, 2, 3, 4, 5],
            [6, 7, 8, 9, 10],
            [11, 12, 13, 14, 15],
            [16, 17, 18, 19, 20],
          ],
          [
            [1, 1, 1, 1, 1],
            [2, 2, 2, 2, 2],
            [3, 3, 3, 3, 3],
            [4, 4, 4, 4, 4],
          ],
        ]);
        const result = a.product(b, [1], [1]).transpose([1, 0, 2]);
        expect(result.shape).toEqual([2, 3, 5]);
      });

      it("multiplies 4D batched tensors [2,3,4,5] @ [2,3,5,6] (multi-head attention style)", () => {
        // Shape: [batch=2, heads=3, seq=4, dim=5] @ [batch=2, heads=3, dim=5, out=6]
        // This is like Q @ K^T in multi-head attention
        // Using simple values: each 2D matrix is filled with 1s
        const a = Tensor.new([
          [
            [
              [1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1],
            ],
            [
              [2, 2, 2, 2, 2],
              [2, 2, 2, 2, 2],
              [2, 2, 2, 2, 2],
              [2, 2, 2, 2, 2],
            ],
            [
              [1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1],
            ],
          ],
          [
            [
              [1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1],
            ],
            [
              [1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1],
            ],
            [
              [3, 3, 3, 3, 3],
              [3, 3, 3, 3, 3],
              [3, 3, 3, 3, 3],
              [3, 3, 3, 3, 3],
            ],
          ],
        ]);
        const b = Tensor.new([
          [
            [
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
            ],
            [
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
            ],
            [
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
            ],
          ],
          [
            [
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
            ],
            [
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
            ],
            [
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
              [1, 1, 1, 1, 1, 1],
            ],
          ],
        ]);
        const result = a.product(b, [3], [2]).diag([0, 3]).diag([0, 2]).transpose([2, 3, 0, 1]);
        expect(result.shape).toEqual([2, 3, 4, 6]);
        // [1,1,1,1,1] @ [[1]*6, [1]*6, ...] = [5,5,5,5,5,5] for each row
        // For a[0][0] (all 1s), each row sums to 5
        expect((result.data as number[][][][])[0][0][0]).toEqual([
          5, 5, 5, 5, 5, 5,
        ]);
        // For a[0][1] (all 2s), each row sums to 10
        expect((result.data as number[][][][])[0][1][0]).toEqual([
          10, 10, 10, 10, 10, 10,
        ]);
        // For a[1][2] (all 3s), each row sums to 15
        expect((result.data as number[][][][])[1][2][0]).toEqual([
          15, 15, 15, 15, 15, 15,
        ]);
      });

      it("complex batch broadcast [2,1,3,4] @ [3,4,5] -> [2,3,3,5]", () => {
        // A has shape [2, 1, 3, 4] - batch dims [2, 1]
        // B has shape [3, 4, 5] - batch dims [3]
        // Broadcast: [2, 1] with [3] -> [2, 3]
        // Result: [2, 3, 3, 5]
        const a = Tensor.new([
          [
            // batch [0, 0] - this will broadcast to [0,0], [0,1], [0,2]
            [
              [1, 0, 0, 0],
              [0, 1, 0, 0],
              [0, 0, 1, 0],
            ],
          ],
          [
            // batch [1, 0] - this will broadcast to [1,0], [1,1], [1,2]
            [
              [2, 0, 0, 0],
              [0, 2, 0, 0],
              [0, 0, 2, 0],
            ],
          ],
        ]);
        const b = Tensor.new([
          // batch [0]
          [
            [1, 1, 1, 1, 1],
            [2, 2, 2, 2, 2],
            [3, 3, 3, 3, 3],
            [4, 4, 4, 4, 4],
          ],
          // batch [1]
          [
            [10, 10, 10, 10, 10],
            [20, 20, 20, 20, 20],
            [30, 30, 30, 30, 30],
            [40, 40, 40, 40, 40],
          ],
          // batch [2]
          [
            [100, 100, 100, 100, 100],
            [200, 200, 200, 200, 200],
            [300, 300, 300, 300, 300],
            [400, 400, 400, 400, 400],
          ],
        ]);
        const result = a.squeeze([1]).product(b, [2], [1]).transpose([0, 2, 1, 3]);
        expect(result.shape).toEqual([2, 3, 3, 5]);

        // result[0, 0] = a[0, 0] @ b[0] = identity-ish @ b[0] = first 3 rows of b[0]
        expect((result.data as number[][][][])[0][0]).toEqual([
          [1, 1, 1, 1, 1],
          [2, 2, 2, 2, 2],
          [3, 3, 3, 3, 3],
        ]);

        // result[0, 1] = a[0, 0] @ b[1] (a broadcasts, uses same slice)
        expect((result.data as number[][][][])[0][1]).toEqual([
          [10, 10, 10, 10, 10],
          [20, 20, 20, 20, 20],
          [30, 30, 30, 30, 30],
        ]);

        // result[1, 2] = a[1, 0] @ b[2] = 2x identity-ish @ b[2] = doubled first 3 rows
        expect((result.data as number[][][][])[1][2]).toEqual([
          [200, 200, 200, 200, 200],
          [400, 400, 400, 400, 400],
          [600, 600, 600, 600, 600],
        ]);
      });
    });

    describe("error cases", () => {
      it("throws error when inner dimensions don't match [2,3] @ [4,5]", () => {
        const a = Tensor.new([
          [1, 2, 3],
          [4, 5, 6],
        ]);
        const b = Tensor.new([
          [1, 2, 3, 4, 5],
          [1, 2, 3, 4, 5],
          [1, 2, 3, 4, 5],
          [1, 2, 3, 4, 5],
        ]);
        expect(() => a.product(b, [1], [0])).toThrow();
      });

      it("throws error when 1D vectors have different lengths", () => {
        const a = Tensor.new([1, 2, 3]);
        const b = Tensor.new([1, 2, 3, 4]);
        expect(() => a.product(b, [1], [0])).toThrow();
      });

      it("throws error when batch dimensions are incompatible [2,3,4] @ [3,4,5]", () => {
        // Batch dim 2 vs 3 - not broadcastable (neither is 1)
        const a = Tensor.new([
          [
            [1, 2, 3, 4],
            [5, 6, 7, 8],
            [9, 10, 11, 12],
          ],
          [
            [1, 2, 3, 4],
            [5, 6, 7, 8],
            [9, 10, 11, 12],
          ],
        ]);
        const b = Tensor.new([
          [
            [1, 2, 3, 4, 5],
            [1, 2, 3, 4, 5],
            [1, 2, 3, 4, 5],
            [1, 2, 3, 4, 5],
          ],
          [
            [1, 2, 3, 4, 5],
            [1, 2, 3, 4, 5],
            [1, 2, 3, 4, 5],
            [1, 2, 3, 4, 5],
          ],
          [
            [1, 2, 3, 4, 5],
            [1, 2, 3, 4, 5],
            [1, 2, 3, 4, 5],
            [1, 2, 3, 4, 5],
          ],
        ]);
        expect(() => a.product(b, [1], [2])).toThrow();
      });
    });

    describe("identity and special cases", () => {
      it("multiplying by identity matrix returns same values", () => {
        const a = Tensor.new([
          [1, 2, 3],
          [4, 5, 6],
        ]);
        const identity = Tensor.new([
          [1, 0, 0],
          [0, 1, 0],
          [0, 0, 1],
        ]);
        const result = a.product(identity, [1], [0]);
        expect(result.data).toEqual([
          [1, 2, 3],
          [4, 5, 6],
        ]);
      });

      it("handles single element matrices", () => {
        const a = Tensor.new([[5]]);
        const b = Tensor.new([[3]]);
        const result = a.product(b, [1], [0]);;
        expect(result.data).toEqual([[15]]);
        expect(result.shape).toEqual([1, 1]);
      });
    });
  });

  describe("shape", () => {
    const testcases: Record<string, [Tensor, number[]]> = {
      "empty array for scalar": [Tensor.new(42), []],
      "[1] for single-element 1D tensor": [Tensor.new([7]), [1]],
      "[n] for 1D tensor": [Tensor.new([1, 2, 3, 4, 5]), [5]],
      "[1, 1] for single-element 2D tensor": [Tensor.new([[7]]), [1, 1]],
      "[rows, cols] for 2D tensor": [
        Tensor.new([
          [1, 2, 3],
          [4, 5, 6],
        ]),
        [2, 3],
      ],
      "[d1, d2, d3] for 3D tensor": [
        Tensor.new([
          [
            [1, 2],
            [3, 4],
            [5, 6],
          ],
          [
            [7, 8],
            [9, 10],
            [11, 12],
          ],
        ]),
        [2, 3, 2],
      ],
      "correct shape for 4D tensor": [
        Tensor.new([
          [
            [
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
            ],
            [
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
            ],
            [
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
            ],
          ],
          [
            [
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
            ],
            [
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
            ],
            [
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
              [1, 2, 3, 4, 5],
            ],
          ],
        ]),
        [2, 3, 4, 5],
      ],
    }

    for (const [name, [t, expected]] of Object.entries(testcases)) {
      it(name, () => {
        expect(t.shape).toEqual(expected);
      });
    }
  });

  describe("sqrt", () => {
    it("computes sqrt of a scalar", () => {
      const a = Tensor.new(4);
      const result = a.sqrt();
      expect(result.data).toEqual(2);
      expect(result.shape).toEqual([]);
    });

    it("computes sqrt of a 1D tensor", () => {
      const a = Tensor.new([0, 1, 4, 9, 16]);
      const result = a.sqrt();
      expect(result.data).toEqual([0, 1, 2, 3, 4]);
    });

    it("computes sqrt of a 2D tensor", () => {
      const a = Tensor.new([
        [1, 4],
        [9, 16],
      ]);
      const result = a.sqrt();
      expect(result.data).toEqual([
        [1, 2],
        [3, 4],
      ]);
    });

    it("handles non-perfect squares", () => {
      const a = Tensor.new([2, 3, 5]);
      const result = a.sqrt();
      const data = result.data as number[];
      expect(data[0]).toBeCloseTo(Math.sqrt(2));
      expect(data[1]).toBeCloseTo(Math.sqrt(3));
      expect(data[2]).toBeCloseTo(Math.sqrt(5));
    });

    it("returns 0 for zero", () => {
      const a = Tensor.new([0]);
      const result = a.sqrt();
      expect(result.data).toEqual([0]);
    });

    it("returns NaN for negative values", () => {
      const a = Tensor.new([-1, -4, -9]);
      const result = a.sqrt();
      const data = result.data as number[];
      expect(data[0]).toBeNaN();
      expect(data[1]).toBeNaN();
      expect(data[2]).toBeNaN();
    });

    it("returns Infinity for Infinity input", () => {
      const a = Tensor.new([Infinity]);
      const result = a.sqrt();
      expect(result.data).toEqual([Infinity]);
    });

    it("returns NaN for -Infinity input", () => {
      const a = Tensor.new([-Infinity]);
      const result = a.sqrt();
      const data = result.data as number[];
      expect(data[0]).toBeNaN();
    });

    it("propagates NaN", () => {
      const a = Tensor.new([4, NaN, 9]);
      const result = a.sqrt();
      const data = result.data as number[];
      expect(data[0]).toEqual(2);
      expect(data[1]).toBeNaN();
      expect(data[2]).toEqual(3);
    });
  });

  describe("sub", () => {
    const testcases: [string, Tensor, (a: Tensor) => Tensor, Tensor][] = [
      ["subtracts two scalars", Tensor.new(10), _ => Tensor.new(3), Tensor.new(7)],
      ["subtracts two 1D tensors of same shape", Tensor.new([10, 20, 30]), _ => Tensor.new([1, 2, 3]), Tensor.new([9, 18, 27])],
      ["subtracts two 2D tensors of same shape", Tensor.new([
          [10, 20],
          [30, 40],
        ]), _ => Tensor.new([
          [1, 2],
          [3, 4],
        ]), Tensor.new([
          [9, 18],
          [27, 36],
        ])],
      ["broadcasts scalar to 1D tensor", Tensor.new([10, 20, 30]), a => Tensor.new(5).broadcast(a.shape), Tensor.new([5, 15, 25])],
      ["broadcasts 1D tensor across rows of 2D tensor", Tensor.new([
          [10, 20, 30],
          [40, 50, 60],
        ]), _ => Tensor.new([1, 2, 3]).broadcast([2, 0]), Tensor.new([
          [9, 18, 27],
          [39, 48, 57],
        ])],
      ["broadcasts column vector to 2D tensor", Tensor.new([
          [10, 20, 30],
          [40, 50, 60],
        ]), _ => Tensor.new([[1], [10]]).squeeze([1]).broadcast([0, 3]), Tensor.new([
          [9, 19, 29],
          [30, 40, 50],
        ])],
      // Edge cases with Infinity and NaN
      ["handles Infinity in subtraction", Tensor.new([Infinity, -Infinity, 1]), _ => Tensor.new([1, 1, Infinity]), Tensor.new([Infinity, -Infinity, -Infinity])],
      ["returns NaN when subtracting Infinity from Infinity", Tensor.new([Infinity, -Infinity]), _ => Tensor.new([Infinity, -Infinity]), Tensor.new([NaN, NaN])],
      ["propagates NaN through subtraction", Tensor.new([1, NaN, 3]), _ => Tensor.new([1, 1, NaN]), Tensor.new([0, NaN, NaN])],
    ]

    for (const [name, a, newb, expected] of testcases) {
      it(name, () => {
        const b = newb(a);
        const actual = a.sub(b);
        expect(actual.shape).toEqual(expected.shape);
        expect(actual.data).toEqual(expected.data);
      });
    }

    it("throws error for incompatible shapes", () => {
      const a = Tensor.new([1, 2, 3]);
      const b = Tensor.new([1, 2]);
      expect(() => a.sub(b)).toThrow("Incompatible shapes: [3] != [2]");
    });
  });

  describe("toString", () => {
    const testcases: [string, Tensor, string][] = [
      ["scalar", Tensor.new(42), "Tensor(42, shape=[])"],
      ["1D tensor", Tensor.new([1, 2, 3]), "Tensor([1,2,3], shape=[3])"],
      ["2D tensor", Tensor.new([[1, 2], [3, 4]]), "Tensor([[1,2],[3,4]], shape=[2,2])"],
      ["3D tensor", Tensor.new([[[1, 2], [3, 4]], [[5, 6], [7, 8]]]), "Tensor([[[1,2],[3,4]],[[5,6],[7,8]]], shape=[2,2,2])"],
      ["single-element 1D tensor", Tensor.new([7]), "Tensor([7], shape=[1])"],
      ["single-element 2D tensor", Tensor.new([[7]]), "Tensor([[7]], shape=[1,1])"],
      ["negative numbers", Tensor.new([-1, -2, -3]), "Tensor([-1,-2,-3], shape=[3])"],
      ["floating point numbers", Tensor.new([1.5, 2.7, 3.14]), "Tensor([1.5,2.7,3.14], shape=[3])"],
      ["Infinity", Tensor.new([1, Infinity, -Infinity]), "Tensor([1,Infinity,-Infinity], shape=[3])"],
      ["NaN", Tensor.new([1, 2, NaN]), "Tensor([1,2,NaN], shape=[3])"],
      ["empty 1D tensor", Tensor.new([]), "Tensor([], shape=[0])"],
    ];

    for (const [name, t, expected] of testcases) {
      it(name, () => {
        expect(t.toString()).toEqual(expected);
      });
    }
  });

  describe("new", () => {
    it("throws on 2D tensor with different row lengths", () => {
      expect(() => Tensor.new([
        [1, 2, 3],
        [4, 5],
      ])).toThrow("Invalid data: flat.length=5 != shape=2,3");
    });
  });
});
