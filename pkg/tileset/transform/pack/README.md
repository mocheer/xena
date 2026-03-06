# pack

gltfpack.exe 支持 KHR_mesh_quantization、EXT_meshopt_compression、KHR_texture_basisu，既有顶点简化压缩，又有纹理编码转换

命令用法

```shell
gltfpack -i scene.gltf -o scene.glb
```

- 默认KHR_mesh_quantization，做了一些缓存优化和量化针对GPU的消耗
- -c  使用 meshoptimizer 编解码器进一步减少下载大小，生成压缩的gltf/glb文件(需要 EXT_meshopt_compression )
- -cc 更好的压缩，您可以使用 -cc 选项，该选项应用额外的压缩; 需要确保内容分发方法使用deflate (gzip) - meshoptimizer编解码器被设计为可以使用通用压缩器进一步压缩的输出。
- -tc 使用BasisU超压缩将所有纹理转换为KTX2(需要 KHR_texture_basisu ，可能需要 -tp 标记以兼容WebGL 1)
- -tw 纹理转换成webp
- -km 保留材质的命名，禁止材质命名合并
- -ke 保留额外的数据
- -noq 禁用KHR_mesh_quantization、KHR_texture_transform （默认是使用的），禁用所有量化，但会执行网格合并、节点裁剪等其他优化。文件体积减小有限，但完全保留原始数据精度。

## 参考
- https://github.com/zeux/meshoptimizer/blob/master/gltf/README.md
