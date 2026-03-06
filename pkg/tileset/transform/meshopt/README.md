# meshopt

Meshopt 是社区的产物，可以针对 Mesh 的 Geometry 和 Animation 等数据进行压缩，包括 Vertex、Index、Morph Target 以及 eeyframe 的 Times 和 Values 等。

社区内有 C++ 版本的开源项目：meshoptimizer，可以对 glTF 进行 Meshopt 压缩，并提供了 WebAssembly 版本的 Encoder 和 Decoder；另外还提供了 WebAssembly 版本的 Simplifier，可以对模型进行简化（减面/减点），但要注意 Simplifier 是有损的。

优化效果稍差于Draco，但解码速度快于draco，综合来说，比draco优秀

## 无损模式
meshopt支持无损模式，可以多次压缩而不损失精度，而draco只有有损模式

## gltfpack
gltfpack 的功能包括各种顶点优化，以及 KHR_mesh_quantization、EXT_meshopt_compression、ktx2拓展，它默认是有损的

- 默认为基础优化
- -c 增加 EXT_meshopt_compression
- -cc 在-c基础上增加

```shell
gltfpack -i input.gltf -o output.glb
gltfpack -i input.glb -o output.glb -cc
gltfpack -i input.glb -o output.glb -cc -tc
```

## 参考
- https://github.com/zeux/meshoptimizer
