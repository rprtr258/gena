import {Tensor} from "./tensor";

// Create tensors
const t = Tensor.new([
  [1, 2, 3],
  [4, 5, 6],
]);

// Access tensor properties
console.log(t); // Tensor([[1, 2, 3], [4, 5, 6]], shape=[2, 3])
console.log("rank:", t.rank); // 2
console.log("shape:", t.shape); // [2, 3]
console.log("strides:", t.strides); // [3, 1]
console.log("flat:", t.flatData); // [1, 2, 3, 4, 5, 6]

// Access elements by multi-dimensional index
console.log("t[0,0] =", t.get([0, 0])); // 1
console.log("t[0,2] =", t.get([0, 2])); // 3
console.log("t[1,1] =", t.get([1, 1])); // 5

console.log("[1,2,3].broadcast([4, 0]) =", Tensor.new([1, 2, 3]).broadcast([4, 0]));
console.log("[[1], [2]].transpose([1, 0]) =", Tensor.new([[1], [2]]).transpose([1, 0]));
console.log("t.slice([0, 0], [2, 2]) =", t.slice([0, 0], [2, 2]));
console.log("[[1, 2]].squeeze([0]) =", Tensor.new([[1, 2]]).squeeze([0]));
