# graph


## cesium加载glTF:
1. https://github.com/CesiumGS/cesium/issues/11280 : 应该是glTF文件中某个mesh的primitives使用了带有纹理的material，但该primitive的几何数据（attributes）中缺少TEXCOORD_0属性

建议：可以校验下是否发生这种情况，并生成默认的uv坐标


## gltf库
- https://github.com/EliCDavis/polyform/tree/main/formats/gltf

## 参考
- https://github.com/donmccurdy/property-graph
- https://gltf-transform.dev/modules/core/classes/Accessor