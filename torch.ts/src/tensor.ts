import util from "util";

function createErrorClass(name: string) {
  return class extends Error {
    constructor(message?: string) {
      super(message || name);
      this.name = name;
    }
  };
}

export const InvalidIndexError = createErrorClass("InvalidIndexError");

type TensorData = number | TensorData[];

function getShape(data: TensorData): number[] {
  if (!Array.isArray(data)) {
    return [];
  }

  const ranks: number[] = [];
  let arr: TensorData = data;
  while (Array.isArray(arr)) {
    ranks.push(arr.length);
    arr = arr[0];
  }
  return ranks;
};

function flattenData(data: TensorData): number[] {
  const result: number[] = [];
  const traverse = (currentData: TensorData): void => {
    if (Array.isArray(currentData)) {
      currentData.forEach(traverse);
    } else {
      result.push(currentData);
    }
  };
  traverse(data);
  return result;
};

function getStrides(shape: number[]): number[] {
  // S_i = ∏(d_k) for k = i+1 to n-1
  const result = new Array(shape.length);
  let stride = 1;
  for (let i = shape.length - 1; i >= 0; i--) {
    result[i] = stride;
    stride *= shape[i];
  }
  return result;
};

function getFlatIndexFromMultiDimensionalIndex(indices: number[], strides: number[]): number {
  // A[i₀, i₁, ..., iₙ₋₁] = A_flat[ ∑(iₖ × Sₖ) for k = 0 to n-1 ]
  return indices.reduce((acc, curr, i) => acc + curr * strides[i], 0);
};

interface TensorI {
  get shape(): number[],
  get(idx: number[]): number,

  data?(): TensorData,
  flatData?(): number[],
  strides?(): number[],
}

class TensorArray implements TensorI {
  _data: number[];
  _shape: number[];
  _strides: number[];
  constructor(
    flatData: number[],
    shape: number[],
  ) {
    if (flatData.length !== shape.reduce((acc, curr) => acc * curr, 1))
      throw new Error(`Invalid data: flat.length=${flatData.length} != shape=${shape}`);

    this._data = flatData;
    this._shape = shape;
    this._strides = getStrides(shape);
  }

  get shape(): number[] {
    return this._shape;
  }

  get rank(): number {
    return this._shape.length;
  }

  strides(): number[] {
    return this._strides;
  }

  flatData(): number[] {
    return this._data;
  }

  data(): TensorData {
    if (this.rank === 0)
      return this._data[0];

    const f = (ii: number, from: number, len: number): TensorData => {
      if (ii === this.rank - 1)
        return this._data.slice(from, from + len);

      const stride = this._strides[ii];
      return [...new Array(this._shape[ii])].map((_, i) => f(ii + 1, from + i * stride, stride));
    };
    return f(0, 0, this._data.length);
  }

  get(indices: number[]): number {
    if (indices.length !== this.rank || indices.find((index, i) => index < 0 || index >= this._shape[i]) !== undefined)
      throw new InvalidIndexError();

    return this.unsafeGet(indices);
  }

  unsafeGet(indices: number[]): number {
    const flatIndex = getFlatIndexFromMultiDimensionalIndex(indices, this._strides);
    return this._data[flatIndex];
  }
}

class TensorBroadcast implements TensorI {
  t: TensorI;
  _shape: number[];
  constructor(
    t: TensorI,
    shape: number[],
  ) {
    if (shape.filter(dim => dim === 0).length !== t.shape.length || shape.some(d => d < 0))
      throw new Error(`Invalid shape for broadcast: ${shape}`);

    this.t = t;
    this._shape = shape;
  }

  get shape(): number[] {
    return (() => {
      let j = 0;
      return this._shape.map(dim => {
        if (dim !== 0)
          return dim;

        const res = this.t.shape[j];
        j++;
        return res;
      });
    })();
  }

  get(indices: number[]): number {
    return this.t.get(indices.filter((_, i) => this._shape[i] === 0));
  }
}

class TensorSqueeze implements TensorI {
  t: TensorI;
  axis: number[];

  constructor(
    t: TensorI,
    axis: number[],
  ) {
    const normalized: number[] = axis.map(a => (a < 0 ? a + t.shape.length : a));

    // validate and check duplicates
    const seen = new Set<number>();
    for (const a of normalized) {
      if (a < 0 || a >= t.shape.length)
        throw new Error(`Axis ${a} out of range`);
      if (seen.has(a))
        throw new Error(`Duplicate axis ${a}`);
      if (t.shape[a] !== 1)
        throw new Error(`Cannot squeeze axis ${a} with size ${t.shape[a]}`);
      seen.add(a);
    }

    this.t = t;
    this.axis = normalized;
  }

  get shape(): number[] {
    return this.t.shape.filter((_, i) => !this.axis.includes(i));
  }

  get(idx: number[]): number {
    const res = new Array(this.t.shape.length).fill(0);
    let j = 0;
    for (let i = 0; i < this.t.shape.length; i++) {
      if (!this.axis.includes(i)) {
        res[i] = idx[j];
        j++;
      }
    }

    return this.t.get(res);
  }
}

class TensorTranspose implements TensorI {
  t: TensorI;
  order: number[];

  constructor(
    t: TensorI,
    order: number[],
  ) {
    if (order.length !== t.shape.length)
      throw new Error("Invalid transpose order");

    const seen = new Array(t.shape.length);
    for (const i of order) {
      if (i < 0 || i >= t.shape.length || seen[i])
        throw new Error("Invalid transpose order");
      seen[i] = true;
    }

    this.t = t;
    this.order = order;
  }

  get shape(): number[] {
    return this.order.map(i => this.t.shape[i]);
  }

  get(idx: number[]): number {
    const rev = new Array(this.order.length);
    for (const [i, o] of this.order.entries())
      rev[o] = i;

    return this.t.get(rev.map(i => idx[i]));
  }
}

class TensorSlice implements TensorI {
  t: TensorI;
  start: number[];
  end: number[];

  constructor(
    t: TensorI,
    start: number[],
    end: number[],
  ) {
    if (start.length !== t.shape.length || end.length !== t.shape.length)
      throw new Error("start/end must have same rank as tensor");

    const newShape = new Array(t.shape.length);
    for (let i = 0; i < t.shape.length; i++) {
      const s = start[i];
      const e = end[i];
      if (s < 0 || e < 0 || s > e || e > t.shape[i])
        throw new Error(`Invalid slice for dimension ${i}: [${s}, ${e})`);
      newShape[i] = e - s;
    }
    if (newShape.some(dim => dim === 0))
      throw new Error(`Empty slice of shape ${newShape}`);

    this.t = t;
    this.start = start;
    this.end = end;
  }

  get shape(): number[] {
    return this.t.shape.map((_, i) => this.end[i] - this.start[i]);
  }

  get(idx: number[]): number {
    return this.t.get(idx.map((x, i) => x + this.start[i]));
  }
}

class TensorUnary implements TensorI {
  t: TensorI;
  op: (a: number) => number;

  constructor(
    t: TensorI,
    op: (a: number) => number,
  ) {
    this.t = t;
    this.op = op;
  }

  get shape(): number[] {
    return this.t.shape;
  }

  get(idx: number[]): number {
    return this.op(this.t.get(idx));
  }
}

class TensorBinary implements TensorI {
  l: TensorI;
  r: TensorI;
  op: (a: number, b: number) => number;

  constructor(
    l: TensorI,
    r: TensorI,
    op: (a: number, b: number) => number,
  ) {
    if (l.shape.length !== r.shape.length || !l.shape.every((dim, i) => dim === r.shape[i]))
      throw new Error(`Incompatible shapes: ${JSON.stringify(l.shape)} != ${JSON.stringify(r.shape)}`);

    this.l = l;
    this.r = r;
    this.op = op;
  }

  get shape(): number[] {
    return this.l.shape;
  }

  get(idx: number[]): number {
    return this.op(this.l.get(idx), this.r.get(idx));
  }
}

export class Tensor {
  t: TensorI;

  private constructor(t: TensorI) {
    this.t = t;
  }

  static new(data: TensorData) {
    const flatData = flattenData(data);
    const shape = getShape(data);
    return new Tensor(new TensorArray(flatData, shape));
  }

  get shape(): number[] {return this.t.shape;}
  get rank(): number {return this.shape.length;}
  get strides(): number[] {
    if (this.t.strides !== undefined)
      return this.t.strides();
    return getStrides(this.shape);
  }

  get data(): TensorData {
    if (this.t.data !== undefined)
      return this.t.data();

    if (this.rank === 0)
      return this.t.get([]);

    const f = (ii: number, idx: number[]): TensorData => {
      const res: TensorData = new Array(this.t.shape[ii]);
      if (ii === this.t.shape.length - 1) {
        for (let i = 0; i < this.t.shape[ii]; i++) {
          res[i] = this.t.get([...idx, i]);
        }
      } else {
        for (let i = 0; i < this.t.shape[ii]; i++) {
          res[i] = f(ii + 1, [...idx, i]);
        }
      }
      return res;
    };
    return f(0, []);
  }
  get flatData(): number[] {
    if (this.t.flatData)
      return this.t.flatData();

    return flattenData(this.data);
  }

  // Given a multi-dimensional index like [1, 2, 3] return the value at that index
  get(idx: number[]): number {
    return this.t.get(idx);
  }

  // broadcasts tensor to given shape with zero replaced with current shape dimensions
  broadcast(shape: number[]): Tensor {
    return new Tensor(new TensorBroadcast(this.t, shape));
  }

  transpose(order: number[]): Tensor {
    return new Tensor(new TensorTranspose(this.t, order));
  }

  /**
   * Extract subtensor defined by two multi-indices.
   * Semantics:
   *  - `start` and `end` are arrays of length rank.
   *  - `start[i]` is inclusive, `end[i]` is exclusive (like JS slice).
   *  - `0 <= start[i] <= end[i] <= this._shape[i]`.
   * Returns a new Tensor containing copied data.
   */
  slice(start: number[], end: number[]): Tensor {
    return new Tensor(new TensorSlice(this.t, start, end));
  }

  /**
   * Remove dimensions of size 1. Supports negative indices.
   * Throws if requested axis does not have size 1 or is out of range.
   */
  squeeze(axis: number[]): Tensor {
    return new Tensor(new TensorSqueeze(this.t, axis));
  }

  private binaryPointwise(other: Tensor, op: (a: number, b: number) => number): Tensor {
    return new Tensor(new TensorBinary(this.t, other.t, op));
  }

  add(other: Tensor): Tensor {return this.binaryPointwise(other, (a, b) => a + b);}
  sub(other: Tensor): Tensor {return this.binaryPointwise(other, (a, b) => a - b);}
  div(other: Tensor): Tensor {return this.binaryPointwise(other, (a, b) => a / b);}
  mul(other: Tensor): Tensor {return this.binaryPointwise(other, (a, b) => a * b);}

  private unary(op: (a: number) => number): Tensor {
    return new Tensor(new TensorUnary(this.t, op));
  }

  neg(): Tensor {return this.unary(value => -value);}
  exp(): Tensor {return this.unary(value => Math.exp(value));}
  log(): Tensor {return this.unary(value => Math.log(value));}
  sqrt(): Tensor {return this.unary(value => Math.sqrt(value));}

  product(other: Tensor, axesA: number[], axesB: number[]): Tensor {
    const normalizeAxes = (axes: number[], rank: number) =>
      axes.map(x => (x < 0 ? x + rank : x));

    if (axesA.length !== axesB.length)
      throw new Error("Axes length mismatch");

    const normA = normalizeAxes(axesA, this.rank);
    const normB = normalizeAxes(axesB, other.rank);

    // Validate axes: in-range and no dups
    const checkAxes = (arr: number[], rank: number, name: string) => {
      const seen = new Set<number>();
      for (const ax of arr) {
        if (ax < 0 || ax >= rank)
          throw new Error(`${name} contains invalid axis ${ax}`);
        if (seen.has(ax))
          throw new Error(`${name} contains duplicate axis ${ax}`);
        seen.add(ax);
      }
    };
    checkAxes(normA, this.rank, "axesA");
    checkAxes(normB, other.rank, "axesB");

    // Check contracted dimensions match
    for (let i = 0; i < normA.length; i++) {
      const da = this.shape[normA[i]];
      const db = other.shape[normB[i]];
      if (da !== db)
        throw new Error(`Dimension mismatch on contracted axis ${axesA[i]} vs ${axesB[i]} (${da} != ${db})`);
    }

    // Build free axes lists (preserve order)
    const isContractedA = new Array(this.rank);
    const isContractedB = new Array(other.rank);
    for (const ax of normA)
      isContractedA[ax] = true;
    for (const ax of normB)
      isContractedB[ax] = true;

    const freeA: number[] = [];
    for (let i = 0; i < this.rank; i++)
      if (!isContractedA[i])
        freeA.push(i);

    const freeB: number[] = [];
    for (let i = 0; i < other.rank; i++)
      if (!isContractedB[i])
        freeB.push(i);

    // Output shape = shapes of freeA then freeB
    const outShape = [
      ...freeA.map(i => this.shape[i]),
      ...freeB.map(i => other.shape[i]),
    ];

    if (outShape.some(d => d === 0))
      throw new Error("Output tensor would have dimension zero");

    const aStrides = this.strides;
    const bStrides = other.strides;
    const outStrides = getStrides(outShape);

    // contracted shape and its strides (for enumerating combinations)
    const contractShape = normA.map(i => this.shape[i]); // same as normB mapped dims
    const contractStrides = getStrides(contractShape);
    const contractSize = contractShape.reduce((p, v) => p * v, 1) || 1;

    const outSize = outShape.reduce((p, v) => p * v, 1);
    const outData = new Array(outSize);

    // Pre-allocate coordinate arrays
    const coordsOut = new Array(outShape.length);
    const coordsA = new Array(this.rank);
    const coordsB = new Array(other.rank);
    const coordsContract = new Array(contractShape.length);

    for (let outIdx = 0; outIdx < outSize; outIdx++) {
      // unravel output index into coordsOut
      let tmp = outIdx;
      for (let d = 0; d < outShape.length; d++) {
        const s = outStrides[d];
        const c = Math.floor(tmp / s);
        coordsOut[d] = c;
        tmp %= s;
      }

      // fill free coords for A from the first part of coordsOut
      for (let i = 0; i < freeA.length; i++) {
        coordsA[freeA[i]] = coordsOut[i];
      }
      // fill free coords for B from the latter part of coordsOut
      for (let j = 0; j < freeB.length; j++) {
        coordsB[freeB[j]] = coordsOut[freeA.length + j];
      }

      // accumulate contraction sum
      let acc = 0;
      for (let cIdx = 0; cIdx < contractSize; cIdx++) {
        // unravel contraction index to coordsContract
        let t = cIdx;
        for (let k = 0; k < contractShape.length; k++) {
          const s = contractStrides[k];
          const cv = Math.floor(t / s);
          coordsContract[k] = cv;
          t %= s;
        }

        // place contracted coords into coordsA and coordsB
        for (let k = 0; k < normA.length; k++) {
          coordsA[normA[k]] = coordsContract[k];
          coordsB[normB[k]] = coordsContract[k];
        }

        // compute flat indices
        let idxA = 0;
        for (let d = 0; d < this.rank; d++) idxA += coordsA[d] * aStrides[d];

        let idxB = 0;
        for (let d = 0; d < other.rank; d++) idxB += coordsB[d] * bStrides[d];

        acc += this.flatData[idxA] * other.flatData[idxB];
      }

      outData[outIdx] = acc;
    }

    return new Tensor(new TensorArray(outData, outShape));
  }

  /**
   * Extract the diagonal across the provided axes.
   * - `axes` is an array of axis indices (may be negative) that must all have the same length.
   * - Returns a new tensor where the `axes` are removed and replaced by a single axis
   *   that enumerates the diagonal index (dimension = common size).
   *   Output shape = keep axes followed by the diagonal axis
   *
   * Example:
   *  trace(Tensor([[1,2],[3,4]]), [0,1]) -> Tensor([1,4])  // diagonal elements [t[0,0], t[1,1]]
   */
  diag(axes: number[]): Tensor {
    const rank = this.shape.length;
    if (!Array.isArray(axes) || axes.length === 0) throw new Error("axes must be a non-empty array");

    // normalize axes (handle negatives) and validate uniqueness / range
    const norm = axes.map(a => (a < 0 ? a + rank : a));
    const seen = new Set<number>();
    for (const a of norm) {
      if (a < 0 || a >= rank) throw new Error(`Invalid axis ${a}`);
      if (seen.has(a)) throw new Error(`Duplicate axis ${a}`);
      seen.add(a);
    }

    // all traced axes must have equal length
    const diagSize = this.shape[norm[0]];
    for (let i = 1; i < norm.length; i++) {
      if (this.shape[norm[i]] !== diagSize) throw new Error("All traced axes must have equal size");
    }

    // determine which axes remain (keepAxes)
    const isTraced = new Array(rank);
    for (const a of norm) isTraced[a] = true;
    const keepAxes: number[] = [];
    for (let i = 0; i < rank; i++)
      if (!isTraced[i])
        keepAxes.push(i);

    const outShape = [...keepAxes.map(i => this.shape[i]), diagSize];

    const outSize = outShape.reduce((p, v) => p * v, 1);
    if (outSize === 0)
      throw new Error("Empty output: all axes were removed or all traced axes had size 0");

    const outStrides = getStrides(outShape);

    const outData = new Array(outSize);
    const coordsOut = new Array(outShape.length);
    const coordsIn = new Array(rank);

    for (let outIdx = 0; outIdx < outSize; outIdx++) {
      // unravel output flat index to coordinates
      let tmp = outIdx;
      for (let d = 0; d < outShape.length; d++) {
        const s = outStrides[d];
        const c = Math.floor(tmp / s);
        coordsOut[d] = c;
        tmp %= s;
      }

      // last coordinate of out is the diagonal index
      const diag = coordsOut[outShape.length - 1];

      // fill in kept axes coords from coordsOut
      for (let i = 0; i < keepAxes.length; i++) {
        coordsIn[keepAxes[i]] = coordsOut[i];
      }

      // set each traced axis to diag
      for (const a of norm) coordsIn[a] = diag;

      // compute input flat index
      let inFlatIdx = 0;
      for (let d = 0; d < rank; d++) inFlatIdx += coordsIn[d] * this.strides[d];

      outData[outIdx] = this.flatData[inFlatIdx];
    }

    return new Tensor(new TensorArray(outData, outShape));
  }

  toString(): string {
    const stringify = (t: TensorData): string => {
      if (Array.isArray(t))
        return `[${t.map(stringify).join(",")}]`;
      return t.toString();
    };
    return `Tensor(${stringify(this.data)}, shape=${JSON.stringify(this.shape)})`;
  }

  [util.inspect.custom](): string {
    return this.toString();
  }
}
